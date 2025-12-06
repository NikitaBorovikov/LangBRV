package model

type DictionaryPage struct {
	UserID      int64
	CurrentPage int
	TotalPages  int
	Words       []Word
}
