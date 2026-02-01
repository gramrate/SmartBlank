package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game - основная сущность игры
type Game struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	LobbyID   string             `bson:"lobby_id,omitempty"`
	LobbyName string             `bson:"lobby_name"`
	IsActive  bool               `bson:"is_active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`

	// Статистика для быстрого доступа
	Stats GameStats `bson:"stats"`

	// Текущее состояние
	Current CurrentState `bson:"current"`

	// История (отдельные коллекции через ссылки)
	Days    []primitive.ObjectID `bson:"days,omitempty"`    // ссылки на Day
	Nights  []primitive.ObjectID `bson:"nights,omitempty"`  // ссылки на Night
	Players []primitive.ObjectID `bson:"players,omitempty"` // ссылки на Player

	// Результат игры
	Result *GameResult `bson:"result,omitempty"`
}

// GameStats - статистика для быстрого поиска
type GameStats struct {
	TotalPlayers int       `bson:"total_players"`
	AlivePlayers int       `bson:"alive_players"`
	DayNumber    int       `bson:"day_number"`
	NightNumber  int       `bson:"night_number"`
	StageType    StageType `bson:"stage_type"`
	HasEnded     bool      `bson:"has_ended"`
	Winner       *Team     `bson:"winner,omitempty"`
}

// CurrentState - текущее состояние (меняется часто)
type CurrentState struct {
	Stage        StageType            `bson:"stage"`
	DayNumber    int                  `bson:"day_number"`
	NightNumber  int                  `bson:"night_number"`
	Players      []PlayerInGame       `bson:"players"` // легковесная копия
	Vote         *VoteState           `bson:"vote,omitempty"`
	ReVote       *VoteState           `bson:"re_vote,omitempty"`
	KickAllVotes int                  `bson:"kick_all_votes"`
	Registration []RegistrationPlayer `bson:"registration,omitempty"`
}

// Player - отдельная коллекция (можно переиспользовать между играми)
type Player struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	GlobalStats PlayerStats        `bson:"global_stats"`
	GamesPlayed []GameReference    `bson:"games_played,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
}

// PlayerInGame - легковесная копия игрока в конкретной игре
type PlayerInGame struct {
	PlayerID       primitive.ObjectID `bson:"player_id"` // ссылка на Player
	Position       int                `bson:"position"`
	Role           Role               `bson:"role"`
	Fouls          int                `bson:"fouls"`
	IsAlive        bool               `bson:"is_alive"`
	IsDisqualified bool               `bson:"is_disqualified"`
	BestTurn       []int              `bson:"best_turn,omitempty"`
}

// Day - день игры (отдельная коллекция)
type Day struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	GameID       primitive.ObjectID `bson:"game_id"`
	Number       int                `bson:"number"`
	HasVote      bool               `bson:"has_vote"`
	Vote         *VoteResult        `bson:"vote,omitempty"`
	ReVote       *VoteResult        `bson:"re_vote,omitempty"`
	KickAllVotes int                `bson:"kick_all_votes"`
	CreatedAt    time.Time          `bson:"created_at"`
}

// Night - ночь игры (отдельная коллекция)
type Night struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	GameID       primitive.ObjectID `bson:"game_id"`
	Number       int                `bson:"number"`
	MafiaTarget  *int               `bson:"mafia_target,omitempty"`
	SheriffCheck *int               `bson:"sheriff_check,omitempty"`
	DonCheck     *int               `bson:"don_check,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
}

// GameResult - результат игры
type GameResult struct {
	Winner  Team           `bson:"winner"`
	Players []PlayerResult `bson:"players"`
	EndedAt time.Time      `bson:"ended_at"`
}

// Дополнительные структуры
type VoteState struct {
	SuggestedPlayers []int       `bson:"suggested_players"`
	Voting           []VoteCount `bson:"voting"`
}

type VoteResult struct {
	SuggestedPlayers []int       `bson:"suggested_players"`
	Voting           []VoteCount `bson:"voting"`
	KickedPlayer     *int        `bson:"kicked_player,omitempty"`
}

type VoteCount struct {
	Position   int `bson:"position"`
	VotesCount int `bson:"votes_count"`
}

type RegistrationPlayer struct {
	Name     string `bson:"name"`
	Position int    `bson:"position"`
}

type PlayerResult struct {
	PlayerID       primitive.ObjectID `bson:"player_id"`
	Role           Role               `bson:"role"`
	Position       int                `bson:"position"`
	Fouls          int                `bson:"fouls"`
	IsDisqualified bool               `bson:"is_disqualified"`
	IsWin          bool               `bson:"is_win"`
	ExtraPoints    float64            `bson:"extra_points"`
	Compensation   int                `bson:"compensation"`
	Result         int                `bson:"result"`
	BestTurn       []int              `bson:"best_turn,omitempty"`
}

type PlayerStats struct {
	TotalGames    int     `bson:"total_games"`
	Wins          int     `bson:"wins"`
	Losses        int     `bson:"losses"`
	TotalFouls    int     `bson:"total_fouls"`
	AverageResult float64 `bson:"average_result"`
}

type GameReference struct {
	GameID   primitive.ObjectID `bson:"game_id"`
	Role     Role               `bson:"role"`
	Position int                `bson:"position"`
	Result   int                `bson:"result"`
	IsWin    bool               `bson:"is_win"`
	PlayedAt time.Time          `bson:"played_at"`
}

// Enums
type StageType int

const (
	StageStart     StageType = 0 // Начало
	StageDeal      StageType = 1 // Раздача
	StageNightZero StageType = 2 // Нулевая ночь
	StageDay       StageType = 3 // День
	StageNight     StageType = 4 // Ночь
	StageEnd       StageType = 5 // Конец
)

type Role int

const (
	RoleCivilian Role = 0 // Мирный
	RoleMafia    Role = 1 // Мафия
	RoleSheriff  Role = 2 // Шериф
	RoleDon      Role = 3 // Дон
)

type Team int

const (
	TeamRed   Team = 0 // Мирные
	TeamBlack Team = 1 // Мафия
)
