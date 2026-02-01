# SmartBlank

SmartBlank is an interactive score sheet for sports mafia hosts. The frontend is already in development.

## What it does
- Generates seating layouts.
- Provides a convenient role deal flow.
- Reduces mistakes during Don and Sheriff checks.
- Keeps timer, music, and participant list on one screen.
- Supports stats tracking and digitization of score sheets after the game.

## Tech stack
- Backend: Go
- Database: MongoDB
- Transport: HTTP + WebSocket
- Frontend: in development

## Architecture (brief)
- `backend/cmd` — entry points.
- `backend/internal/adapter` — delivery (HTTP/WebSocket), configuration, repositories.
- `backend/internal/domain` — entities, DTOs, business logic.
- `backend/pkg` — shared utilities.

## Status
Active development.

## Run
Instructions will be added later.
