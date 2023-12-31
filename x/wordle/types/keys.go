package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "wordle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wordle"
)

const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventPlayer    = "player"           // Is it relevant to me?
)

const (
	MovePlayedEventType       = "move-played"
	MovePlayedEventCreator    = "creator"
	MovePlayedEventGameIndex  = "game-index"
	MovePlayedEventGuessState = "guess-state"
	MovePlayedEventWinner     = "winner"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	NoFifoIndex = "-1"
)
