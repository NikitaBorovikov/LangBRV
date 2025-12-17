package model

type UserState struct {
	UserID     int64
	DeleteMode bool
	LastMsgID  int
}

func NewUserState(userID int64, deleteMode bool, msgID int) *UserState {
	return &UserState{
		UserID:     userID,
		DeleteMode: deleteMode,
		LastMsgID:  msgID,
	}
}
