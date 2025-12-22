package inmemory

import (
	"fmt"
	"langbrv/internal/core/model"
	"sync"
)

type RemindCardRepo struct {
	card map[int64]*model.RemindCard
	mu   sync.RWMutex
}

func NewRemindCardRepo() *RemindCardRepo {
	return &RemindCardRepo{
		card: make(map[int64]*model.RemindCard, DefaultInMemorySize),
	}
}

func (r *RemindCardRepo) Save(page *model.RemindCard) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.card[page.UserID] = page
	return nil
}

func (r *RemindCardRepo) Get(userID int64) (*model.RemindCard, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	page := r.card[userID]
	if page == nil {
		return nil, fmt.Errorf("failed to find remind card with such userID")
	}
	return page, nil
}
