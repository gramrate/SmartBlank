package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardColor string

const (
	CardNone  CardColor = "none"
	CardYellow CardColor = "yellow"
	CardGray  CardColor = "gray"
	CardRed   CardColor = "red"
)

// GameState represents the full game document stored in MongoDB.
type GameState struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LobbyID             string             `bson:"lobby_id" json:"lobby_id"`
	LobbyName           string             `bson:"lobby_name" json:"lobby_name"`
	IsActive            bool               `bson:"is_active" json:"is_active"`
	Players             PlayersState       `bson:"players" json:"players"`
	StageType           StageType          `bson:"stage_type" json:"stage_type"`
	Registration        RegistrationState  `bson:"registration" json:"registration"`
	Days                []DayState         `bson:"days" json:"days"`
	DayNumber           int                `bson:"day_number" json:"day_number"`
	Nights              []NightState       `bson:"nights" json:"nights"`
	NightNumber         int                `bson:"night_number" json:"night_number"`
	End                 *EndState          `bson:"end,omitempty" json:"end,omitempty"`
	Deal                DealState          `bson:"deal,omitempty" json:"deal,omitempty"`
	Music               NightMusicState    `bson:"music,omitempty" json:"music,omitempty"`
	FirstNightKillPos   *int               `bson:"first_night_kill_pos,omitempty" json:"first_night_kill_pos,omitempty"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updated_at"`
}

type PlayersState struct {
	AlivePlayers int           `bson:"alive_players" json:"alive_players"`
	Players      []PlayerState `bson:"players" json:"players"`
}

type PlayerState struct {
	Name          string    `bson:"name" json:"name"`
	Position      int       `bson:"position" json:"position"`
	Fouls         int       `bson:"fouls" json:"fouls"`
	IsAlive       bool      `bson:"is_alive" json:"is_alive"`
	IsDisqualified bool     `bson:"is_disqualified" json:"is_disqualified"`
	Role          Role      `bson:"role" json:"role"`
	BestTurn      []int     `bson:"best_turn,omitempty" json:"best_turn,omitempty"`
	Card          CardColor `bson:"card,omitempty" json:"card,omitempty"`
	PendingDeath  bool      `bson:"pending_death,omitempty" json:"pending_death,omitempty"`
}

type RegistrationState struct {
	Players []RegistrationPlayerState `bson:"players" json:"players"`
}

type RegistrationPlayerState struct {
	Name     string `bson:"name" json:"name"`
	Position int    `bson:"position" json:"position"`
}

type DealState struct {
	ForbiddenRoles []RoleRestriction `bson:"forbidden_roles,omitempty" json:"forbidden_roles,omitempty"`
	ForcedRoles    []RoleAssignment  `bson:"forced_roles,omitempty" json:"forced_roles,omitempty"`
}

type RoleRestriction struct {
	Position int    `bson:"position" json:"position"`
	Roles    []Role `bson:"roles" json:"roles"`
}

type RoleAssignment struct {
	Position int  `bson:"position" json:"position"`
	Role     Role `bson:"role" json:"role"`
}

type DayState struct {
	Number           int       `bson:"number" json:"number"`
	TodayHaveVote    bool      `bson:"today_have_vote" json:"today_have_vote"`
	Vote             VoteStateState `bson:"vote,omitempty" json:"vote,omitempty"`
	ReVote           VoteStateState `bson:"re_vote,omitempty" json:"re_vote,omitempty"`
	KickAllVotes     int       `bson:"kick_all_votes" json:"kick_all_votes"`
	KickAllCandidates []int    `bson:"kick_all_candidates,omitempty" json:"kick_all_candidates,omitempty"`
}

type VoteStateState struct {
	SuggestedPlayers []int            `bson:"suggested_players" json:"suggested_players"`
	Voting           []VoteCountState `bson:"voting" json:"voting"`
	KickedPlayers    []int            `bson:"kicked_players,omitempty" json:"kicked_players,omitempty"`
}

type VoteCountState struct {
	Position   int `bson:"position" json:"position"`
	VotesCount int `bson:"votes_count" json:"votes_count"`
}

type NightState struct {
	Number             int   `bson:"number" json:"number"`
	MafiaTarget        *int  `bson:"mafia_target,omitempty" json:"mafia_target,omitempty"`
	MafiaMiss          bool  `bson:"mafia_miss" json:"mafia_miss"`
	SheriffCheck       *int  `bson:"sheriff_check,omitempty" json:"sheriff_check,omitempty"`
	SheriffCheckResult *bool `bson:"sheriff_check_result,omitempty" json:"sheriff_check_result,omitempty"`
	DonCheck           *int  `bson:"don_check,omitempty" json:"don_check,omitempty"`
	DonCheckResult     *bool `bson:"don_check_result,omitempty" json:"don_check_result,omitempty"`
}

type EndState struct {
	Winner  Team             `bson:"winner" json:"winner"`
	Players []PlayerEndState `bson:"players" json:"players"`
}

type PlayerEndState struct {
	Name          string   `bson:"name" json:"name"`
	Role          Role     `bson:"role" json:"role"`
	Position      int      `bson:"position" json:"position"`
	Fouls         int      `bson:"fouls" json:"fouls"`
	IsDisqualified bool    `bson:"is_disqualified" json:"is_disqualified"`
	IsWin         bool     `bson:"is_win" json:"is_win"`
	ExtraPoints   float64  `bson:"extra_points" json:"extra_points"`
	Compensation  int      `bson:"compensation" json:"compensation"`
	Result        int      `bson:"result" json:"result"`
	BestTurn      []int    `bson:"best_turn,omitempty" json:"best_turn,omitempty"`
}

type NightMusicState struct {
	Paused bool    `bson:"paused" json:"paused"`
	Volume float64 `bson:"volume" json:"volume"`
}
