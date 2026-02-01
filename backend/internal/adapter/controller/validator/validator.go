package validator

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	minPosition = 1
	maxPosition = 10
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateCreateLobbyRequest(_ *dto.CreateLobbyRequest) error {
	return nil
}

func (v *Validator) ValidateCloseLobbyRequest(req *dto.CloseLobbyRequest) error {
	if req.LobbyID == uuid.Nil {
		return errors.New("lobby_id is required")
	}
	return nil
}

func (v *Validator) ValidateLobbyIDRequest(req *dto.LobbyIDRequest) error {
	if strings.TrimSpace(req.LobbyID) == "" {
		return errors.New("lobby_id is required")
	}
	return nil
}

func (v *Validator) ValidateCreateGameRequest(req *dto.CreateGameRequest) error {
	if strings.TrimSpace(req.LobbyID) == "" && strings.TrimSpace(req.LobbyName) == "" {
		return nil
	}
	if req.LobbyID != "" && strings.TrimSpace(req.LobbyID) == "" {
		return errors.New("lobby_id is invalid")
	}
	return nil
}

func (v *Validator) ValidateAddRegistrationPlayer(req *dto.AddRegistrationPlayerRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}
	if req.Position != 0 && !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateUpdateRegistrationPlayer(req *dto.UpdateRegistrationPlayerRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if req.NewPosition != 0 && !validPosition(req.NewPosition) {
		return fmt.Errorf("new_position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateRemoveRegistrationPlayer(req *dto.RemoveRegistrationPlayerRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateSwapRegistrationPositions(req *dto.SwapRegistrationPositionsRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.PositionA) || !validPosition(req.PositionB) {
		return fmt.Errorf("positions must be %d-%d", minPosition, maxPosition)
	}
	if req.PositionA == req.PositionB {
		return errors.New("positions must be different")
	}
	return nil
}

func (v *Validator) ValidateGenerateSeatingRequest(req *dto.GenerateSeatingRequest) error {
	return validateLobbyID(req.LobbyID)
}

func (v *Validator) ValidateSetStageRequest(req *dto.SetStageRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if req.StageType < entity.StageStart || req.StageType > entity.StageEnd {
		return errors.New("invalid stage_type")
	}
	return nil
}

func (v *Validator) ValidateAssignRoleRequest(req *dto.AssignRoleRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if req.Role < entity.RoleCivilian || req.Role > entity.RoleDon {
		return errors.New("invalid role")
	}
	return nil
}

func (v *Validator) ValidateForbiddenRoleRequest(req *dto.ForbiddenRoleRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if len(req.Roles) == 0 {
		return errors.New("roles are required")
	}
	for _, role := range req.Roles {
		if role < entity.RoleCivilian || role > entity.RoleDon {
			return errors.New("invalid role")
		}
	}
	return nil
}

func (v *Validator) ValidateAutoDealRequest(req *dto.AutoDealRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if req.MafiaCount < 0 {
		return errors.New("mafia_count must be >= 0")
	}
	return nil
}

func (v *Validator) ValidateAddFoulRequest(req *dto.AddFoulRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if req.Delta == 0 {
		return errors.New("delta must be non-zero")
	}
	return nil
}

func (v *Validator) ValidateSetCardRequest(req *dto.SetCardRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	switch req.Card {
	case entity.CardNone, entity.CardYellow, entity.CardGray, entity.CardRed:
		return nil
	default:
		return errors.New("invalid card color")
	}
}

func (v *Validator) ValidateRemovePlayerRequest(req *dto.RemovePlayerRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateStartDayRequest(req *dto.StartDayRequest) error {
	return validateLobbyID(req.LobbyID)
}

func (v *Validator) ValidateStartNightRequest(req *dto.StartNightRequest) error {
	return validateLobbyID(req.LobbyID)
}

func (v *Validator) ValidateStartVoteRequest(req *dto.StartVoteRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if len(req.SuggestedPlayers) == 0 {
		return errors.New("suggested_players must not be empty")
	}
	seen := map[int]bool{}
	for _, pos := range req.SuggestedPlayers {
		if !validPosition(pos) {
			return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
		}
		if seen[pos] {
			return errors.New("duplicate position")
		}
		seen[pos] = true
	}
	return nil
}

func (v *Validator) ValidateSetVoteRequest(req *dto.SetVoteRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if req.VotesCount < 0 {
		return errors.New("votes_count must be >= 0")
	}
	return nil
}

func (v *Validator) ValidateResolveVoteRequest(_ *dto.ResolveVoteRequest) error {
	return nil
}

func (v *Validator) ValidateKickAllVoteRequest(req *dto.KickAllVoteRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if req.Votes < 0 {
		return errors.New("votes must be >= 0")
	}
	return nil
}

func (v *Validator) ValidateMafiaTargetRequest(req *dto.MafiaTargetRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateMafiaMissRequest(req *dto.MafiaMissRequest) error {
	return validateLobbyID(req.LobbyID)
}

func (v *Validator) ValidateSheriffCheckRequest(req *dto.SheriffCheckRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateDonCheckRequest(req *dto.DonCheckRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	return nil
}

func (v *Validator) ValidateApplyNightResultsRequest(req *dto.ApplyNightResultsRequest) error {
	return validateLobbyID(req.LobbyID)
}

func (v *Validator) ValidateSetBestTurnRequest(req *dto.SetBestTurnRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if !validPosition(req.Position) {
		return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
	}
	if len(req.BestTurn) != 3 {
		return errors.New("best_turn must contain 3 numbers")
	}
	seen := map[int]bool{}
	for _, value := range req.BestTurn {
		if value < minPosition || value > maxPosition {
			return fmt.Errorf("best_turn values must be %d-%d", minPosition, maxPosition)
		}
		if seen[value] {
			return errors.New("best_turn values must be unique")
		}
		seen[value] = true
	}
	return nil
}

func (v *Validator) ValidateUpdateMusicRequest(req *dto.UpdateMusicRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if req.Volume < 0 || req.Volume > 1 {
		return errors.New("volume must be between 0 and 1")
	}
	return nil
}

func (v *Validator) ValidateEndGameRequest(req *dto.EndGameRequest) error {
	if err := validateLobbyID(req.LobbyID); err != nil {
		return err
	}
	if req.Winner < entity.TeamRed || req.Winner > entity.TeamBlack {
		return errors.New("winner must be 0 or 1")
	}
	for _, player := range req.Players {
		if !validPosition(player.Position) {
			return fmt.Errorf("position must be %d-%d", minPosition, maxPosition)
		}
		if player.Role < entity.RoleCivilian || player.Role > entity.RoleDon {
			return errors.New("invalid role")
		}
	}
	return nil
}

func validPosition(position int) bool {
	return position >= minPosition && position <= maxPosition
}

func validateLobbyID(lobbyID string) error {
	if strings.TrimSpace(lobbyID) == "" {
		return errors.New("lobby_id is required")
	}
	return nil
}
