package model

type Position string
type Navigation string

const (
	Single Position = "SINGLE"
	First  Position = "FIRST"
	Middle Position = "MIDDLE"
	Last   Position = "LAST"

	Next     Navigation = "NEXT"
	Previous Navigation = "PREVIOUS"

	DefaultPageNumber       = 1
	DefaultRemindCardNumber = 1
)

type UserState struct {
	UserID         int64
	IsDeleteMode   bool
	DictionaryPage *DictionaryPage
	RemindSession  *RemindSession
	LastMessageID  int
}

type DictionaryPage struct {
	CurrentPage int64
	TotalPages  int64
	Position    Position
	MessageID   int
	Words       []Word
}

type RemindSession struct {
	UserID      int64
	CurrentCard int
	TotalCards  int
	Position    Position
	MessageID   int
	Words       []Word
}

func NewUserState(userID int64) *UserState {
	return &UserState{
		UserID:       userID,
		IsDeleteMode: false,
	}
}

func NewDictionaryPage() *DictionaryPage {
	page := &DictionaryPage{
		CurrentPage: DefaultPageNumber,
	}
	page.DeterminePosition()
	return page
}

func NewRemindSession(words []Word) *RemindSession {
	rs := &RemindSession{
		CurrentCard: DefaultRemindCardNumber,
		TotalCards:  len(words),
		Words:       words,
	}
	return rs
}

func (p *DictionaryPage) ChangeCurrenctPage(navigation Navigation) {
	if navigation == Next {
		p.CurrentPage++
	} else {
		p.CurrentPage--
	}
	p.DeterminePosition()
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

func (rs *RemindSession) GoToNextCard() {
	rs.CurrentCard++
	rs.DeterminePosition()
}

func (rs *RemindSession) DeterminePosition() {
	switch {
	case rs.CurrentCard == 1 && rs.TotalCards == 1:
		rs.Position = Single
	case rs.CurrentCard == 1 && rs.TotalCards > 1:
		rs.Position = First
	case rs.CurrentCard != 1 && rs.CurrentCard == rs.TotalCards:
		rs.Position = Last
	default:
		rs.Position = Middle
	}
}
