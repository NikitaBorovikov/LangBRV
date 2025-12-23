package model

type RemindCardPosition string

const (
	SingleCard RemindCardPosition = "SINGLE"
	FirstCard  RemindCardPosition = "FIRST"
	MiddleCard RemindCardPosition = "MIDDLE"
	LastCard   RemindCardPosition = "LAST"

	DefaultCardNumber int = 1
)

type RemindCard struct {
	UserID      int64
	CurrentCard int
	TotalCards  int
	Position    RemindCardPosition
	MessageID   int
	Words       []Word
}

func NewRemindCard(userID int64, words []Word) *RemindCard {
	return &RemindCard{
		UserID:      userID,
		CurrentCard: DefaultCardNumber,
		TotalCards:  len(words),
		Words:       words,
	}
}

func (c *RemindCard) DeterminePosition() {
	switch {
	case c.CurrentCard == 1 && c.TotalCards == 1:
		c.Position = SingleCard
	case c.CurrentCard == 1 && c.TotalCards > 1:
		c.Position = FirstCard
	case c.CurrentCard != 1 && c.CurrentCard == c.TotalCards:
		c.Position = LastCard
	default:
		c.Position = MiddleCard
	}
}
