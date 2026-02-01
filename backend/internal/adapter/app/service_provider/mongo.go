package service_provider

import (
	"backend/pkg/closer"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB возвращает MongoDB клиент с ленивой инициализацией
func (s *ServiceProvider) MongoDB() *mongo.Database {
	if s.mongoDB == nil {
		s.Logger().Debugf("Connecting to MongoDB (uri=%s)", s.MongoConfig().UriAddr())

		// Создаем контекст с таймаутом для подключения
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Настраиваем опции клиента
		clientOptions := options.Client().
			ApplyURI(s.MongoConfig().UriAddr()).
			SetConnectTimeout(10 * time.Second).
			SetSocketTimeout(30 * time.Second).
			SetServerSelectionTimeout(10 * time.Second).
			SetMaxPoolSize(100).
			SetMinPoolSize(5).
			SetRetryWrites(true)

		// Подключаемся к MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			s.Logger().Panicf("Failed to connect to MongoDB: %v", err)
		}

		// Проверяем подключение
		err = client.Ping(ctx, nil)
		if err != nil {
			s.Logger().Panicf("Failed to ping MongoDB: %v", err)
		}

		// Регистрируем закрытие соединения
		closer.Add(func() error {
			s.Logger().Debug("Closing MongoDB connection...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := client.Disconnect(ctx); err != nil {
				s.Logger().Errorf("Failed to disconnect from MongoDB: %v", err)
				return err
			}
			return nil
		})

		s.mongoDB = client.Database(s.MongoConfig().Database())
		s.Logger().Infof("Successfully connected to MongoDB database '%s'",
			s.MongoConfig().Database())
	}

	return s.mongoDB
}

// runMigrations выполняет миграции схемы
func (s *ServiceProvider) runMigrations(client *mongo.Client) {
	s.Logger().Debug("Running MongoDB schema migrations...")
	ctx := context.Background()
	mongoDB := client.Database(s.MongoConfig().Database())

	// 1. Создаем коллекцию для отслеживания миграций
	migrationsCol := mongoDB.Collection("_migrations")

	// Создаем индекс для версий (игнорируем ошибку если уже существует)
	_, _ = migrationsCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"version": 1},
		Options: options.Index().SetUnique(true),
	})

	// 2. Создаем коллекции и индексы
	migrations := []struct {
		Version     int
		Description string
		Migrate     func(context.Context, *mongo.Database) error
	}{
		{
			Version:     1,
			Description: "Create users collection with indexes",
			Migrate: func(ctx context.Context, mongoDB *mongo.Database) error {
				// Создаем коллекцию users
				err := mongoDB.CreateCollection(ctx, "users")
				if err != nil && !isCollectionExistsError(err) {
					return err
				}

				// Создаем индексы для users
				usersCol := mongoDB.Collection("users")
				_, err = usersCol.Indexes().CreateMany(ctx, []mongo.IndexModel{
					{
						Keys:    bson.M{"email": 1},
						Options: options.Index().SetUnique(true).SetName("idx_users_email"),
					},
					{
						Keys:    bson.M{"created_at": -1},
						Options: options.Index().SetName("idx_users_created_at"),
					},
				})

				return err
			},
		},
		{
			Version:     2,
			Description: "Create sessions collection",
			Migrate: func(ctx context.Context, mongoDB *mongo.Database) error {
				err := mongoDB.CreateCollection(ctx, "sessions")
				if err != nil && !isCollectionExistsError(err) {
					return err
				}

				sessionsCol := mongoDB.Collection("sessions")
				_, err = sessionsCol.Indexes().CreateOne(ctx, mongo.IndexModel{
					Keys:    bson.M{"expires_at": 1},
					Options: options.Index().SetExpireAfterSeconds(0).SetName("idx_sessions_expires"),
				})

				return err
			},
		},
	}

	// 3. Применяем миграции
	for _, migration := range migrations {
		// Проверяем, применена ли уже миграция
		count, err := migrationsCol.CountDocuments(ctx, bson.M{"version": migration.Version})
		if err != nil {
			s.Logger().Errorf("Failed to check migration %d: %v", migration.Version, err)
			continue
		}

		if count > 0 {
			s.Logger().Debugf("Migration %d already applied: %s",
				migration.Version, migration.Description)
			continue
		}

		// Выполняем миграцию
		s.Logger().Debugf("Applying migration %d: %s", migration.Version, migration.Description)
		if err := migration.Migrate(ctx, mongoDB); err != nil {
			s.Logger().Errorf("Failed to apply migration %d: %v", migration.Version, err)
			continue
		}

		// Записываем в историю
		_, err = migrationsCol.InsertOne(ctx, bson.M{
			"version":     migration.Version,
			"description": migration.Description,
			"applied_at":  time.Now(),
		})
		if err != nil {
			s.Logger().Errorf("Failed to record migration %d: %v", migration.Version, err)
		}
	}
}

// isCollectionExistsError проверяет, что ошибка - "коллекция уже существует"
func isCollectionExistsError(err error) bool {
	return err != nil &&
		(err.Error() == "collection already exists" ||
			mongo.IsDuplicateKeyError(err))
}
