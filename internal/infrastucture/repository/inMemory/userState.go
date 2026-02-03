package inmemory

import (
	apperrors "langbrv/internal/app_errors"
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

func (r *UserStateRepo) Save(state *model.UserState) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.states[state.UserID] = state
	return nil
}

func (r *UserStateRepo) Get(userID int64) (*model.UserState, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	state := r.states[userID]
	if state == nil {
		return nil, apperrors.ErrUserStateNotFound
	}
	return state, nil
}
