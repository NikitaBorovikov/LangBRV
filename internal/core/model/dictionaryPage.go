package model

type DictionaryPageStatus string

const (
	SinglePage DictionaryPageStatus = "SINGLE"
	FirstPage  DictionaryPageStatus = "FIRST"
	MiddlePage DictionaryPageStatus = "MIDDLE"
	LastPage   DictionaryPageStatus = "LAST"
)

const (
	DefaultPageNumber = 1
)

type DictionaryPage struct {
	UserID          int64
	CurrentPage     int64
	TotalPages      int64
	Status          DictionaryPageStatus
	DictionaryMsgID int
	Words           []Word
}

func NewDictionaryPage(userID int64) *DictionaryPage {
	return &DictionaryPage{
		UserID:      userID,
		CurrentPage: DefaultPageNumber,
	}
}

func (p *DictionaryPage) DetermineStatus() {
	switch {
	case p.CurrentPage == 1 && p.TotalPages == 1:
		p.Status = SinglePage
	case p.CurrentPage == 1 && p.TotalPages > 1:
		p.Status = FirstPage
	case p.CurrentPage != 1 && p.CurrentPage == p.TotalPages:
		p.Status = LastPage
	default:
		p.Status = MiddlePage
	}
}
