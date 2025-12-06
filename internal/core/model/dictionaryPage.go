package model

const (
	DefaultPageNumber = 1
)

type DictionaryPage struct {
	UserID      int64
	CurrentPage int
	TotalPages  int
	Words       []Word
}

func NewDictionaryPage(userID int64) *DictionaryPage {
	return &DictionaryPage{
		UserID:      userID,
		CurrentPage: DefaultPageNumber,
	}
}
