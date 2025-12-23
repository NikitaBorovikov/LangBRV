package model

type UserState struct {
	UserID        int64
	DeleteMode    bool
	LastMessageID int
}

func NewUserState(userID int64, deleteMode bool) *UserState {
	return &UserState{
		UserID:     userID,
		DeleteMode: deleteMode,
	}
}
