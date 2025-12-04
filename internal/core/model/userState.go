package model

type State string

const (
	AddWord State = "ADD_WORD"
	DelWord State = "DEL_WORD"
)

type UserState struct {
	UserID int64
	State  State
}

func NewUserState(userID int64, state State) *UserState {
	return &UserState{
		UserID: userID,
		State:  state,
	}
}
