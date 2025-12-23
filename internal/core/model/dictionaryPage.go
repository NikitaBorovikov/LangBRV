package model

type DictionaryPagePosition string

const (
	SinglePage DictionaryPagePosition = "SINGLE"
	FirstPage  DictionaryPagePosition = "FIRST"
	MiddlePage DictionaryPagePosition = "MIDDLE"
	LastPage   DictionaryPagePosition = "LAST"

	DefaultPageNumber = 1
)

type DictionaryPage struct {
	UserID      int64
	CurrentPage int64
	TotalPages  int64
	Position    DictionaryPagePosition
	MessageID   int
	Words       []Word
}

func NewDictionaryPage(userID int64) *DictionaryPage {
	return &DictionaryPage{
		UserID:      userID,
		CurrentPage: DefaultPageNumber,
	}
}

func (p *DictionaryPage) DeterminePosition() {
	switch {
	case p.CurrentPage == 1 && p.TotalPages == 1:
		p.Position = SinglePage
	case p.CurrentPage == 1 && p.TotalPages > 1:
		p.Position = FirstPage
	case p.CurrentPage != 1 && p.CurrentPage == p.TotalPages:
		p.Position = LastPage
	default:
		p.Position = MiddlePage
	}
}
