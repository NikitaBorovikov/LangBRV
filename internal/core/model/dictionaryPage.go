package model

const (
	DefaultPageNumber = 1
)

type DictionaryPage struct {
	UserID          int64
	CurrentPage     int64
	TotalPages      int64
	DictionaryMsgID int
	Words           []Word
}

func NewDictionaryPage(userID int64) *DictionaryPage {
	return &DictionaryPage{
		UserID:      userID,
		CurrentPage: DefaultPageNumber,
	}
}
