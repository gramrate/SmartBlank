package game

import (
	"backend/internal/domain/dto"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleWebSocket(c echo.Context) error {
	ws, err := h.wsUtils.Upgrader().Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	defer h.clearLobbyID(ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("ws read error: %v", err)
			break
		}

		var wsMsg dto.WSMessage
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			h.wsUtils.SendError(ws, "invalid message format")
			continue
		}

		if wsMsg.Type != "game.create" && wsMsg.Type != "game.bind" {
			lobbyID, ok := h.getLobbyID(ws)
			if !ok || lobbyID == "" {
				h.wsUtils.SendError(ws, "lobby_id is not bound to connection")
				continue
			}
			dataMap, ok := wsMsg.Data.(map[string]interface{})
			if !ok || dataMap == nil {
				dataMap = map[string]interface{}{}
			}
			idRaw, has := dataMap["lobby_id"]
			idStr, _ := idRaw.(string)
			if !has || idStr == "" {
				dataMap["lobby_id"] = lobbyID
			}
			wsMsg.Data = dataMap
		}

		switch wsMsg.Type {
		case "game.create":
			h.handleCreate(ws, &wsMsg)
		case "game.bind":
			h.handleBind(ws, &wsMsg)
		case "game.get":
			h.handleGet(ws, &wsMsg)
		case "game.set_stage":
			h.handleSetStage(ws, &wsMsg)
		case "registration.add":
			h.handleAddRegistrationPlayer(ws, &wsMsg)
		case "registration.update":
			h.handleUpdateRegistrationPlayer(ws, &wsMsg)
		case "registration.remove":
			h.handleRemoveRegistrationPlayer(ws, &wsMsg)
		case "registration.swap":
			h.handleSwapRegistrationPositions(ws, &wsMsg)
		case "registration.generate":
			h.handleGenerateSeating(ws, &wsMsg)
		case "deal.assign_role":
			h.handleAssignRole(ws, &wsMsg)
		case "deal.forbid_role":
			h.handleSetForbiddenRole(ws, &wsMsg)
		case "deal.auto":
			h.handleAutoDeal(ws, &wsMsg)
		case "day.start":
			h.handleStartDay(ws, &wsMsg)
		case "day.foul.add":
			h.handleAddFoul(ws, &wsMsg)
		case "day.card.set":
			h.handleSetCard(ws, &wsMsg)
		case "player.remove":
			h.handleRemovePlayer(ws, &wsMsg)
		case "vote.start":
			h.handleStartVote(ws, &wsMsg)
		case "vote.set":
			h.handleSetVote(ws, &wsMsg)
		case "vote.resolve":
			h.handleResolveVote(ws, &wsMsg)
		case "vote.kick_all":
			h.handleKickAllVote(ws, &wsMsg)
		case "night.start":
			h.handleStartNight(ws, &wsMsg)
		case "night.mafia.target":
			h.handleMafiaTarget(ws, &wsMsg)
		case "night.mafia.miss":
			h.handleMafiaMiss(ws, &wsMsg)
		case "night.sheriff.check":
			h.handleSheriffCheck(ws, &wsMsg)
		case "night.don.check":
			h.handleDonCheck(ws, &wsMsg)
		case "night.apply_results":
			h.handleApplyNightResults(ws, &wsMsg)
		case "night.best_turn.set":
			h.handleSetBestTurn(ws, &wsMsg)
		case "night.music.update":
			h.handleUpdateMusic(ws, &wsMsg)
		case "game.end":
			h.handleEndGame(ws, &wsMsg)
		default:
			h.wsUtils.SendError(ws, "unknown message type")
		}
	}

	return nil
}
