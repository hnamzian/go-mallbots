package domain

type BotStatus string

const (
	BotStatusUnknown BotStatus = "unknown"
	BotStatusIdle    BotStatus = "idle"
	BotStatusActive  BotStatus = "active"
)

type Bot struct {
	ID    string
	Name  string
	Staus BotStatus
}

func (s BotStatus) String() string {
	switch s {
	case BotStatusIdle, BotStatusActive:
		return string(s)
	default:
		return string(BotStatusUnknown)
	}
}

func ToBotStatus(s string) BotStatus {
	switch s{
		case BotStatusActive.String():
			return BotStatusActive
		case BotStatusIdle.String():
			return BotStatusIdle
		default:
			return BotStatusUnknown
	}
}