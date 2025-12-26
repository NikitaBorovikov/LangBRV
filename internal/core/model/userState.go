package model

type Mode string
type Position string

const (
	DeleteMode         Mode = "DELETE"
	AddMode            Mode = "ADD"
	ViewDictionaryMode Mode = "VIEW_DICTIONARY"
	RemidMode          Mode = "REMIND"

	Single Position = "SINGLE"
	First  Position = "FIRST"
	Middle Position = "MIDDLE"
	Last   Position = "LAST"

	DefaultPageNumber = 1
	DefaultCardNumber = 1
)

type UserState struct {
	UserID         int64
	Mode           Mode
	DictionaryPage *DictionaryPage
	RemindCard     *RemindCard
	LastMessageID  int
}

type DictionaryPage struct {
	CurrentPage int64
	TotalPages  int64
	Position    Position
	MessageID   int
	Words       []Word
}

type RemindCard struct {
	UserID      int64
	CurrentCard int
	TotalCards  int
	Position    Position
	MessageID   int
	Words       []Word
}

func NewUserState(userID int64, mode Mode) *UserState {
	return &UserState{
		UserID: userID,
		Mode:   mode,
	}
}

func NewDictionaryPage() *DictionaryPage {
	return &DictionaryPage{
		CurrentPage: DefaultPageNumber,
	}
}

func NewRemindCard(words []Word) *RemindCard {
	return &RemindCard{
		CurrentCard: DefaultCardNumber,
		TotalCards:  len(words),
		Words:       words,
	}
}

func (p *DictionaryPage) DeterminePosition() {
	switch {
	case p.CurrentPage == 1 && p.TotalPages == 1:
		p.Position = Single
	case p.CurrentPage == 1 && p.TotalPages > 1:
		p.Position = First
	case p.CurrentPage != 1 && p.CurrentPage == p.TotalPages:
		p.Position = Last
	default:
		p.Position = Middle
	}
}

func (c *RemindCard) DeterminePosition() {
	switch {
	case c.CurrentCard == 1 && c.TotalCards == 1:
		c.Position = Single
	case c.CurrentCard == 1 && c.TotalCards > 1:
		c.Position = First
	case c.CurrentCard != 1 && c.CurrentCard == c.TotalCards:
		c.Position = Last
	default:
		c.Position = Middle
	}
}
