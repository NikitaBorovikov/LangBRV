package model

type RemindCardStatus string

const (
	SingleCard RemindCardStatus = "SINGLE"
	FirstCard  RemindCardStatus = "FIRST"
	MiddleCard RemindCardStatus = "MIDDLE"
	LastCard   RemindCardStatus = "LAST"

	DefaultCardNumber int = 1
)

type RemindCard struct {
	UserID      int64
	Status      RemindCardStatus
	CurrentCard int
	TotalCards  int
	RemindMsgID int
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

func (c *RemindCard) DetermineStatus() {
	switch {
	case c.CurrentCard == 1 && c.TotalCards == 1:
		c.Status = SingleCard
	case c.CurrentCard == 1 && c.TotalCards > 1:
		c.Status = FirstCard
	case c.CurrentCard != 1 && c.CurrentCard == c.TotalCards:
		c.Status = LastCard
	default:
		c.Status = MiddleCard
	}
}
