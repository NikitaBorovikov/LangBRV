package model

type State string

const (
	AddWord       State = "ADD_WORD"
	DelWord       State = "DEL_WORD"
	GetDictionary State = "GET_DICT"
)

type UserState struct {
	UserID    int64
	State     State
	LastMsgID int
}

func NewUserState(userID int64, state State, msgID int) *UserState {
	return &UserState{
		UserID:    userID,
		State:     state,
		LastMsgID: msgID,
	}
}
