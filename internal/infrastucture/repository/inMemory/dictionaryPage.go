package inmemory

import (
	"fmt"
	"langbrv/internal/core/model"
	"sync"
)

type DictionaryPageRepo struct {
	page map[int64]*model.DictionaryPage
	mu   sync.RWMutex
}

func NewDictionaryPageRepo() *DictionaryPageRepo {
	return &DictionaryPageRepo{
		page: make(map[int64]*model.DictionaryPage, DefaultInMemorySize),
	}
}

func (r *DictionaryPageRepo) Save(page *model.DictionaryPage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.page[1] = page //REMOVE
	return nil
}

func (r *DictionaryPageRepo) Get(userID int64) (*model.DictionaryPage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	page := r.page[userID]
	if page == nil {
		return nil, fmt.Errorf("failed to find dictionary page with such chatID")
	}
	return page, nil
}
