package inmemory

import (
	"fmt"
	"langbrv/internal/core/model"
	"sync"
)

const (
	DefaultInMemorySize = 100
)

type UserStateRepo struct {
	states map[int64]*model.UserState
	mu     sync.RWMutex
}

func NewUserStateRepo() *UserStateRepo {
	return &UserStateRepo{
		states: make(map[int64]*model.UserState, DefaultInMemorySize),
	}
}

func (r *UserStateRepo) Set(state *model.UserState) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.states[state.ChatID] = state
	return nil
}

func (r *UserStateRepo) Get(chatID int64) (*model.UserState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	state := r.states[chatID]
	if state == nil {
		return nil, fmt.Errorf("failed to find state with such chatID")
	}
	return state, nil
}
