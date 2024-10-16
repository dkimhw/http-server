package main

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	store map[string]int
}

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}
