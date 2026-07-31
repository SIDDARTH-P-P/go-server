package store

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var ErrItemNotFound = errors.New("item not found")

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Store struct {
	mu     sync.RWMutex
	items  map[int]Item
	nextID int
}

func NewStore() *Store {
	return &Store{
		items:  make(map[int]Item),
		nextID: 1,
	}
}

func (s *Store) Create(title string) (Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item := Item{ID: s.nextID, Title: title}
	s.items[item.ID] = item
	s.nextID++
	return item, nil
}

func (s *Store) Get(id int) (Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[id]
	if !ok {
		return Item{}, ErrItemNotFound
	}
	return item, nil
}

func (s *Store) Update(id int, title string) (Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[id]
	if !ok {
		return Item{}, ErrItemNotFound
	}
	item.Title = title
	s.items[id] = item
	return item, nil
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return ErrItemNotFound
	}
	delete(s.items, id)
	return nil
}

func (s *Store) List() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	return items
}

func (s *Store) String() string {
	return fmt.Sprintf("store with %d items", len(s.items))
}
