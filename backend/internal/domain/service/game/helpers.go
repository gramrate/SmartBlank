package game

import "backend/internal/domain/entity"

func findPlayerIndex(players []entity.PlayerState, position int) int {
	for i, p := range players {
		if p.Position == position {
			return i
		}
	}
	return -1
}

func findRegistrationIndex(players []entity.RegistrationPlayerState, position int) int {
	for i, p := range players {
		if p.Position == position {
			return i
		}
	}
	return -1
}

func hasRegistrationPosition(players []entity.RegistrationPlayerState, position int) bool {
	return findRegistrationIndex(players, position) != -1
}

func nextAvailablePosition(players []entity.RegistrationPlayerState) int {
	pos := 1
	for {
		if !hasRegistrationPosition(players, pos) {
			return pos
		}
		pos++
	}
}

func recountAlive(game *entity.GameState) {
	count := 0
	for i := range game.Players.Players {
		if game.Players.Players[i].IsAlive {
			count++
		}
	}
	game.Players.AlivePlayers = count
}

func lastDayIndex(game *entity.GameState) int {
	if len(game.Days) == 0 {
		return -1
	}
	return len(game.Days) - 1
}

func lastNightIndex(game *entity.GameState) int {
	if len(game.Nights) == 0 {
		return -1
	}
	return len(game.Nights) - 1
}

func initVoteCounts(suggested []int) []entity.VoteCountState {
	voting := make([]entity.VoteCountState, 0, len(suggested))
	for _, pos := range suggested {
		voting = append(voting, entity.VoteCountState{Position: pos, VotesCount: 0})
	}
	return voting
}

func setVoteCount(vote *entity.VoteStateState, position int, votes int) {
	for i := range vote.Voting {
		if vote.Voting[i].Position == position {
			vote.Voting[i].VotesCount = votes
			return
		}
	}
	vote.Voting = append(vote.Voting, entity.VoteCountState{Position: position, VotesCount: votes})
}

func resolveVoteTies(vote *entity.VoteStateState) (int, []int) {
	votesMap := make(map[int]int)
	for _, v := range vote.Voting {
		votesMap[v.Position] = v.VotesCount
	}
	maxVotes := 0
	for _, pos := range vote.SuggestedPlayers {
		if votesMap[pos] > maxVotes {
			maxVotes = votesMap[pos]
		}
	}
	if maxVotes == 0 {
		return 0, nil
	}
	var tied []int
	for _, pos := range vote.SuggestedPlayers {
		if votesMap[pos] == maxVotes {
			tied = append(tied, pos)
		}
	}
	return maxVotes, tied
}

func kickPlayers(game *entity.GameState, positions []int) {
	for _, pos := range positions {
		idx := findPlayerIndex(game.Players.Players, pos)
		if idx == -1 {
			continue
		}
		game.Players.Players[idx].IsAlive = false
		game.Players.Players[idx].PendingDeath = false
	}
	recountAlive(game)
}

func upsertForbiddenRole(deal *entity.DealState, position int, roles []entity.Role) {
	for i := range deal.ForbiddenRoles {
		if deal.ForbiddenRoles[i].Position == position {
			deal.ForbiddenRoles[i].Roles = roles
			return
		}
	}
	deal.ForbiddenRoles = append(deal.ForbiddenRoles, entity.RoleRestriction{Position: position, Roles: roles})
}

func upsertForcedRole(deal *entity.DealState, position int, role entity.Role) {
	for i := range deal.ForcedRoles {
		if deal.ForcedRoles[i].Position == position {
			deal.ForcedRoles[i].Role = role
			return
		}
	}
	deal.ForcedRoles = append(deal.ForcedRoles, entity.RoleAssignment{Position: position, Role: role})
}

func buildRoles(mafiaCount int, includeSheriff bool, includeDon bool) []entity.Role {
	roles := make([]entity.Role, 0, mafiaCount+2)
	for i := 0; i < mafiaCount; i++ {
		roles = append(roles, entity.RoleMafia)
	}
	if includeSheriff {
		roles = append(roles, entity.RoleSheriff)
	}
	if includeDon {
		roles = append(roles, entity.RoleDon)
	}
	return roles
}

func removeForcedRoles(roles []entity.Role, forced []entity.RoleAssignment) []entity.Role {
	result := append([]entity.Role{}, roles...)
	for _, f := range forced {
		if f.Role == entity.RoleCivilian {
			continue
		}
		for i, r := range result {
			if r == f.Role {
				result = append(result[:i], result[i+1:]...)
				break
			}
		}
	}
	return result
}

func isRoleForbidden(forbidden []entity.RoleRestriction, position int, role entity.Role) bool {
	for _, fr := range forbidden {
		if fr.Position != position {
			continue
		}
		for _, r := range fr.Roles {
			if r == role {
				return true
			}
		}
	}
	return false
}

func availableForRole(players []entity.PlayerState, forbidden []entity.RoleRestriction, role entity.Role) []int {
	var candidates []int
	for i, p := range players {
		if p.Role != entity.RoleCivilian {
			continue
		}
		if isRoleForbidden(forbidden, p.Position, role) {
			continue
		}
		candidates = append(candidates, i)
	}
	return candidates
}

func containsInt(items []int, value int) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}
