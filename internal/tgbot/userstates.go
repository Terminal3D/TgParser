package tgbot

import "sync"

// AwaitingUrlState User's states
const (
	AwaitingUrlState            = "awaiting_url"
	AwaitingFilterModeState     = "awaiting_filter_mode"
	AwaitingNameFilter          = "awaiting_name_filter"
	AwaitingBrandFilter         = "awaiting_brand_filter"
	AwaitingPriceFilter         = "awaiting_price_filter"
	AwaitingBrandAndPriceFilter = "awaiting_brand_and_price_filter"
)

type UserStates struct {
	mu     sync.RWMutex
	states map[int64]string
}

func NewUserStates() *UserStates {
	return &UserStates{
		states: make(map[int64]string),
	}
}

func (us *UserStates) Set(userID int64, state string) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.states[userID] = state
}

func (us *UserStates) Get(userID int64) (string, bool) {
	us.mu.Lock()
	defer us.mu.Unlock()
	state, exists := us.states[userID]
	return state, exists
}

func (us *UserStates) Delete(userID int64) {
	us.mu.Lock()
	defer us.mu.Unlock()
	delete(us.states, userID)
}
