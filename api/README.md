# Go Chess API

A web backend for real-time online chess, built with Go's standard library and PostgreSQL. Pairs with a standalone chess engine (`../engine`) that handles all game logic, move validation, and board state.

## Architecture

The API follows a **layered architecture** (Handler → Service → Repository) with no external web framework — just `net/http` with Go 1.22+ routing patterns. Dependency injection is done manually in `cmd/main.go` by wiring concrete implementations into each layer.

```
                        ┌──────────────────┐
                        │  HTTP + WS Client │
                        └─────────┬────────┘
                                  │  REST (JSON) / WebSocket
                                  ▼
                ┌─────────────────────────────────┐
                │        net/http ServeMux        │
                │  Go 1.22+ pattern routing        │
                └────────────────┬────────────────┘
                                 │
                                 ▼
                ┌─────────────────────────────────┐
                │  Middleware (middleware/auth.go) │
                │  JWT validation + userID in ctx  │
                └────────────────┬────────────────┘
                                 │
        ┌────────────────────────┴────────────────────────┐
        ▼                                                 ▼
┌───────────────────┐                        ┌─────────────────────┐
│   auth Handler    │                        │    match Handler     │
│   user Handler    │                        │  REST + WebSocket    │
└─────────┬─────────┘                        └──────────┬──────────┘
          ▼                                            ▼
┌───────────────────┐                        ┌─────────────────────┐
│   Service Layer   │                        │    match.Service     │
│  auth / user /    │                        │  owns a MatchManager │
│  match            │                        └──────────┬──────────┘
└─────────┬─────────┘                                   │
          │                                    ┌────────┴────────┐
          ▼                                    ▼                 ▼
┌───────────────────┐              ┌─────────────────┐  ┌────────────────┐
│   Repository      │              │  MatchManager   │  │ Reference to   │
│   (Postgres)      │              │  (in-memory,    │  │ Match object   │
└─────────┬─────────┘              │  per-game lock) │  │ w/ own game    │
          ▼                        └────────┬────────┘  │ loop goroutine │
┌───────────────────┐                       │          └───────┬────────┘
│ sqlc generated    │                       │                  │
│ pgx driver / DB   │                       ▼                  ▼
└─────────┬─────────┘              ┌────────────────────────────────┐
          ▼                        │         Chess Engine             │
┌───────────────────┐              │  gochess.Game (own goroutine     │
│    PostgreSQL     │              │  reading CommandsCh, fanout)     │
└───────────────────┘              └────────────────────────────────┘
```

*(Note: `MatchManager` is a memory-`map[string]*Match` registry guarded by a `sync.RWMutex`; each active `Match` additionally carries its own `sync.RWMutex` protecting `Status` and its `listeners` channels.)*

### Request flows

- **Stateless REST** (auth, users, match CRUD): request hits the handler, calls a service, the service talks to a repository, which runs a sqlc-generated SQL query against PostgreSQL, then the response is written back. No per-request goroutines beyond what `net/http` spawns.
- **Real-time WebSocket** (`/ws/match/{id}`): the handler upgrades the connection, then spawns a **broadcast goroutine** (writing responses from the match's per-user channel to the socket) and a **blocking read loop** (forwarding client messages into the match's shared `CommandsCh`). The actual game loop runs independently in a goroutine started by the service.

## Directory Structure

```
api/
├── cmd/main.go                          # Entry point: loads .env, connects DB, wires DI, registers routes
├── sqlc.yaml                            # sqlc codegen configuration
├── Makefile                             # migrate-up/down, generate (sqlc), dev-db
├── internal/
│   ├── auth/                            # Registration & login
│   │   ├── dto.go                       # Request validation (UserRegister, UserCredentials)
│   │   ├── service.go                   # Register / login logic
│   │   └── handler.go                   # POST /auth/register, POST /auth/login
│   ├── token/                           # JWT token issuance & validation
│   │   └── token.go                     # JWTTokenProvider (HS256, access + refresh)
│   ├── user/                            # User domain
│   │   ├── user.go                      # User model + Repository interface
│   │   ├── service.go                   # Business logic
│   │   ├── handler.go                   # GET /users
│   │   ├── postgres_repository.go       # PostgreSQL implementation
│   │   └── inmem_repository.go          # In-memory (testing)
│   ├── match/                           # Chess match domain
│   │   ├── match.go                     # Match model, game loop (Start), listener fanout
│   │   ├── dto.go                       # Commands, statuses, request/response types
│   │   ├── service.go                   # Create match, matchmaking handshake, start game
│   │   ├── handler.go                   # POST /matches, GET /matches/live, WS handler
│   │   ├── match_manager.go             # In-memory registry of active matches
│   │   ├── postgres_repository.go       # PostgreSQL implementation
│   │   └── inmem_repository.go          # In-memory (testing)
│   ├── middleware/
│   │   ├── middleware.go                # Middleware chain composition (Use)
│   │   └── auth.go                      # JWT auth middleware (extracts userID)
│   ├── websocket/
│   │   └── upgrader.go                  # Gorilla WebSocket upgrader
│   ├── database/
│   │   ├── migrations/                  # Goose SQL migrations (schema source)
│   │   │   ├── 00001_create_table_users.sql
│   │   │   ├── 00002_create_table_matches.sql
│   │   │   └── 00003_matches_add_fen_col.sql
│   │   ├── queries/                     # sqlc query source
│   │   │   ├── users.sql
│   │   │   └── matches.sql
│   │   └── generated/                   # sqlc output (DO NOT EDIT)
│   │       ├── db.go
│   │       ├── models.go
│   │       ├── users.sql.go
│   │       └── matches.sql.go
│   └── database/                        # (see above)
└── utils/json.go                        # HTTPError, JSON request/response helpers
```

## Data Persistence: PostgreSQL + sqlc

Persistence uses **PostgreSQL** accessed through **sqlc-generated Go code**, so there are no hand-written SQL strings in the application code.

### Migration workflow (goose)

Schema lives in `internal/database/migrations/` as numbered SQL files managed by [goose](https://github.com/pressly/goose):

| Migration | Contents |
|-----------|----------|
| `00001_create_table_users.sql` | `users` table |
| `00002_create_table_matches.sql` | `matches` table + two indexes |
| `00003_matches_add_fen_col.sql` | adds `fen` column to `matches` |

```bash
make migrate-up     # apply all pending migrations
make migrate-down   # revert the latest migration
make migrate-create NAME=add_foo   # scaffold a new migration
```

### Schema

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(20) NOT NULL UNIQUE,
    hashed_password VARCHAR(60),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    opponent_id UUID REFERENCES users(id) ON DELETE CASCADE,   -- nullable until opponent joins
    status VARCHAR(32) NOT NULL DEFAULT 'PENDING',
    owner_white BOOLEAN NOT NULL,
    move_history TEXT[] NOT NULL DEFAULT '{}',                 -- algebraic moves
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    fen TEXT                                                    -- current board FEN
);
CREATE INDEX idx_matches_owner_id ON matches(owner_id);
CREATE INDEX idx_matches_opponent_id ON matches(opponent_id);
```

### sqlc workflow

[sqlc](https://sqlc.dev) reads the migrations (schema) plus hand-written `.sql` query files and generates type-safe Go functions.

**Queries** (`internal/database/queries/`):

- `users.sql`: `CreateUser`, `GetUsers`, `GetUserByID`, `ExistsUserByID`
- `matches.sql`: `CreateMatch`, `GetMatchByID`, `GetMatchesByUser`, `UpdateGameFinalState`, `SetMatchStatus`, `SetMatchStatusAndOpponent`

**Configuration** (`sqlc.yaml`) — PostgreSQL engine, `pgx/v5` SQL package, JSON tags emitted, and a type override mapping the `uuid` DB type to `github.com/google/uuid.UUID`:

```yaml
version: "2"
sql:
  - engine: "postgresql"
    schema: "internal/database/migrations"
    queries: "internal/database/queries"
    gen:
      go:
        package: "generated"
        out: "internal/database/generated"
        sql_package: "pgx/v5"
        emit_json_tags: true
        overrides:
          - db_type: "uuid"
            go_type: "github.com/google/uuid.UUID"
```

**Generated output** (`internal/database/generated/`):
- `db.go` — `DBTX` interface (Exec/Query/QueryRow) and the `Queries` struct; `New(conn)` binds it to a `pgx` connection
- `models.go` — DB row structs (`User`, `Match`) with `pgtype` wrappers for nullable fields (`OpponentID pgtype.UUID`, `Fen pgtype.Text`, `HashedPassword pgtype.Text`)
- `users.sql.go` / `matches.sql.go` — one typed Go function + params struct per query

**Runtime flow:** `cmd/main.go` opens a `pgx` connection → `generated.New(conn)` → passed into `NewPostgresRepository(queries)`. Repositories map generated DB rows ↔ domain structs (e.g. `ParseMatchDB` converts `generated.Match` → `match.Match`, converting `uuid.UUID`/`pgtype.UUID` back to plain strings).

```bash
make generate   # regenerate code after editing queries/*.sql
```

## Models

### User (domain)

| Field            | Type     | DB Column         | Notes                     |
|------------------|----------|-------------------|---------------------------|
| `ID`             | `string` | `id`              | UUID, auto-generated      |
| `Username`       | `string` | `username`        | Unique, min 6 chars       |
| `HashedPassword` | `string` | `hashed_password` | Stored as-is (see limitations) |

### Match (domain)

| Field         | Type       | DB Column       | Notes                                   |
|---------------|------------|-----------------|-----------------------------------------|
| `ID`          | `string`   | `id`            | UUID, auto-generated                    |
| `OwnerID`     | `string`   | `owner_id`      | FK → users, ON DELETE CASCADE           |
| `OpponentID`  | `string`   | `opponent_id`   | FK → users, nullable (set on connect)   |
| `Status`      | `string`   | `status`        | PENDING → MATCHMAKING → PLAYING → result|
| `OwnerWhite`  | `bool`     | `owner_white`   | Whether creator plays white             |
| `MoveHistory` | `[]string` | `move_history`  | Algebraic notation per move             |
| `FEN`         | `string`   | `fen`           | Current board state (FEN)               |

**Status lifecycle:** `PENDING` → `MATCHMAKING` → `PLAYING` → `WHITE_WINS` / `BLACK_WINS` / `DRAW` / `ABORTED`

### In-Memory Match Object

A `Match` also holds runtime (non-persisted) fields used during active play:

- `game` — a `gochess.Game` instance (the engine)
- `listeners` — `map[string]chan GameResponse` (one buffered channel per connected player)
- `CommandsCh` — `chan GameCommand` (buffered, shared by both players; the game loop reads from it)
- `MoveHistory` — accumulated during play, persisted on finalization
- `mu` — a `sync.RWMutex` protecting `Status` and `listeners`

### Game Command / Response (WebSocket)

```jsonc
// Client → Server
{ "cmd": "match.cmd.move", "move": { "from": "e2", "to": "e4" } }

// Server → Client (broadcast to both players)
{ "valid": true, "cmd": "match.cmd.move", "fen": "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1" }

// Server → Client (game over)
{ "cmd": "match.cmd.end", "status": "WHITE_WINS", "fen": "..." }
```

## The MatchManager

The `MatchManager` (`match_manager.go`) is an **in-memory registry of active matches**, used to coordinate real-time play without hitting the database on every move.

**Structure:**

```go
type MatchManager struct {
    matches map[string]*Match
    mu      sync.RWMutex          // guards map access
}
```

It does **two levels of locking**:
- The manager's own `mu (RWMutex)` protects the `matches` map itself (`AddMatch`, `GetMatch`, `GetMatches`).
- Each individual `Match` carries its **own** `sync.RWMutex` (`m.mu`) that protects its mutable `Status` and `listeners` map. This keeps the hot per-game resources isolated so one busy match doesn't block map lookups for all others.

**Key methods:**

| Method | What it does |
|--------|--------------|
| `AddMatch(match)` | Inserts a new match; returns `ErrMatchAlreadyExists` on duplicate |
| `GetMatch(id)` | Lookup by ID; returns `ErrMatchNotFound` |
| `GetStatus(id)` / `SetStatus(id, status)` | Read/write status under the match's lock |
| `SetOpponentID(id, userID)` | Assigns the second player during the handshake |
| `AddListener(matchID, userID)` | Creates a buffered `chan GameResponse` (size 1) for a player; returns `ErrUserAlreadyConnected` if they already have one |
| `GetListener(matchID, userID)` | Retrieves a player's channel |
| `RemoveListener(matchID, userID)` | Closes and removes a player's channel (on disconnect) |
| `GetMatches()` | Returns all live matches (backed by `slices.Collect(maps.Values(...))`) |

The service uses it for the **matchmaking handshake** in `AddUserToMatch`: it reads the current status and, based on it, assigns the owner (`PENDING` → `MATCHMAKING`), assigns the opponent (`MATCHMAKING` → `PLAYING`), then registers the connecting player as a listener. When the second player joins, `StartMatch` fires `go match.Start()` to launch the game loop.

Because each WS connection adds/removes a listener, `RemoveListener` (which `close()`s the channel) is what signals the broadcast goroutine to stop when a player disconnects.

## Concurrency

The API is heavily concurrent; here is how each layer handles it:

### Per-connection goroutines (handler)

Each WebSocket connection spawns **two goroutines** in `HandleGameWebSocket`:
- A **broadcast goroutine** that ranges over the player's `responseCh` and writes each `GameResponse` to the socket; it exits on write error or when it sees `MatchEndCmd` (also finalizing the match in the DB).
- The **main handler goroutine** runs a blocking `conn.ReadJSON` loop, stamps each inbound `GameCommand` with the authenticated `UserID`, and pushes it into the match's shared `CommandsCh`.

### Per-match game loop goroutine

`Match.Start()` runs in its own goroutine per match. It:
1. Builds the `gochess.Game` and broadcasts the begin signal via `sendMessage` → `fanout`.
2. Loops on `select` reading from the shared `CommandsCh`.
3. For each `MovePieceCmd`, calls `game.MovePlayer(...)`, then broadcasts the updated FEN; on error, echoes the failure **only to the offending user** (`response.UserID` set).
4. On game end, upgrades the reply to `MatchEndCmd`, records `MoveHistory`, returns, and `defer CloseChannels()` closes every listener channel (unblocking the broadcast goroutines).

### Message fanout / channels

- Each player has a **buffered channel (size 1)** — the game loop does a non-blocking-ish `ch <- msg` push per listener; a consumer blocking on socket writes won't stall the whole game because buffers absorb latency between broadcasts.
- `sendMessage` distinguishes *global* broadcasts (`UserID == ""` → `fanout` to all listeners) from *targeted* replies (only the sender's channel).

### Locking summary

| Shared resource        | Guard                              | Accessed by                          |
|------------------------|------------------------------------|--------------------------------------|
| `matches` map          | `MatchManager.mu` (`RWMutex`)      | Multiple WS handlers / REST handlers |
| Per-match `Status`     | `Match.mu` (`RWMutex`)             | Game loop, WS handlers               |
| Per-match `listeners`  | `Match.mu`                         | Game loop (fanout), handlers (add/remove) |
| SQL connection (`pgx`) | Handled by pgx connection pool     | All repository calls                 |

`net/http` itself also runs each request in its own goroutine, so REST handlers are already safe to call concurrently; their shared repository (`*generated.Queries`) is stateless, and `pgx` handles connection-pool synchronization.

## Chess Engine Integration

The API imports the engine as `gochess "github.com/tomascaceres14/go-chess/engine"`. Key points:

- **`gochess.Game`** — owns the board state, validates moves, detects check/checkmate/stalemate
- **`Game.MovePlayer(from, to, playerName)`** — accepts algebraic square strings (`"e2"`, `"e4"`), returns move result
- **`Game.GetFENString()`** — serializes current state as FEN for client sync
- **Board representation:** `[8][8]Movable` grid, a1 = `[0][0]`, h8 = `[7][7]`
- **Pieces:** King, Queen, Rook, Bishop, Knight, Pawn — each implements the `Movable` interface (`visibleSquares`, `legalMoves`, `move`)
- **Move validation:** legal moves → king-safety check via board cloning and simulation (handles pins)
- **Special moves:** castling, en passant, pawn promotion (auto-queen)

## API Endpoints

### Public

| Method | Path             | Description                        |
|--------|------------------|------------------------------------|
| `GET`  | `/`              | Health check ("Server running!")   |
| `POST` | `/auth/register` | Create account                     |
| `POST` | `/auth/login`    | Get JWT tokens                     |
| `GET`  | `/users`         | List all users                     |
| `GET`  | `/matches/live`  | List all in-progress matches (from `MatchManager`) |

### Protected (JWT Bearer token required)

| Method | Path               | Description                                   |
|--------|--------------------|-----------------------------------------------|
| `GET`  | `/users/matches`   | Get matches for authenticated user (from DB)  |
| `POST` | `/matches`         | Create a new match (`{"whites": true/false}`) |
| `GET`  | `/ws/match/{id}`   | Connect to match via WebSocket                |

## Authentication Flow

1. **Register:** `POST /auth/register` with `{username, password, repeat_password}`. Returns `{token, refresh_token}`.
2. **Login:** `POST /auth/login` with `{username, password}`. Returns `{token, refresh_token}`.
3. **Token:** HS256 JWT, 1h access / 24h refresh. Claims: `iss`, `sub` (user UUID), `iat`, `exp`, `type`.
4. **Middleware:** `Authorization: Bearer <token>` → validates signature + expiry → checks user exists → injects `userID` into request context (`context.WithValue(ctx, "userID", id)`).

## Match Lifecycle

1. **Create:** Authenticated user `POST /matches {"whites": true}` → DB row (`PENDING`) via repository → added to `MatchManager`
2. **Connect Owner:** Owner hits `GET /ws/match/{id}` → status `PENDING` → `AssignOwner` → `MATCHMAKING` → receives `{"cmd":"match.status.waiting"}`
3. **Connect Opponent:** Opponent hits same URL → status `MATCHMAKING` → `AssignOpponent` (DB + manager both set `PLAYING`) → `go match.Start()` → both receive `{"cmd":"match.status.begin", "fen":"..."}`
4. **Play:** Clients send `match.cmd.move` → engine validates → broadcast updated FEN to both players via per-user channels
5. **End:** Checkmate/stalemate → `match.cmd.end` with final status → handler calls `FinalizeMatch` (DB updated with status, FEN, move history) → channels closed

## Dependencies

| Package                           | Purpose                              |
|-----------------------------------|--------------------------------------|
| `net/http`                        | HTTP server & routing (stdlib)       |
| `github.com/gorilla/websocket`    | WebSocket upgrades                   |
| `github.com/golang-jwt/jwt/v5`    | JWT token creation/validation        |
| `github.com/jackc/pgx/v5`         | PostgreSQL driver                    |
| `github.com/google/uuid`          | UUID generation for DB keys          |
| `github.com/joho/godotenv`        | `.env` file loading (autoload)       |
| `github.com/pressly/goose`        | Database migrations (via Makefile)   |
| `sqlc` (tool)                     | Generate type-safe DB code           |
| `gochess` (../engine)             | Chess game logic engine              |

## Environment Variables

| Variable          | Description                               |
|-------------------|-------------------------------------------|
| `JWT_SIGNING_KEY` | Secret key for HS256 JWT signing          |
| `PORT_HTTP`       | Server listen address (default `:80`)     |
| `DATABASE_URL`    | PostgreSQL connection string              |

## Getting Started

```bash
# Set up database (create schema + apply migrations)
make dev-db

# Alternatively, apply migrations only
make migrate-up

# Regenerate sqlc code (only if queries changed)
make generate

# Run the server
go run cmd/main.go
```

## Known Limitations

- Passwords are not hashed on storage or verified on login
- Match `Start()` uses a hardcoded FEN instead of the classic opening position (temporary)
- `GetByUsername` performs a full table scan
- WebSocket `CheckOrigin` allows all origins
- No refresh token rotation or revocation
- Engine lacks: PGN export, undo/redo, 50-move rule, threefold repetition, insufficient material draws
