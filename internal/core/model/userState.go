package model

type State string

const (
	AddWord State = "ADD_WORD"
)

type UserState struct {
	ChatID int64
	State  State
}

func NewUserState(chatID int64, state State) *UserState {
	return &UserState{
		ChatID: chatID,
		State:  state,
	}
}
