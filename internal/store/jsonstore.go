package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/juparave/gotodo/internal/model"
)

type JSONFileStore struct {
	path  string
	mutex sync.Mutex
	Todos []*model.Todo `json:"todos"`
}

func NewJSONFileStore(path string) *JSONFileStore {
	return &JSONFileStore{path: path}
}

func (s *JSONFileStore) Init() error {
	// Create initial empty store and persist it. Don't hold the mutex while
	// calling Save (which also locks) to avoid deadlock.
	s.mutex.Lock()
	if _, err := os.Stat(s.path); err == nil {
		s.mutex.Unlock()
		return errors.New("file already exists")
	}
	s.Todos = []*model.Todo{}
	s.mutex.Unlock()
	return s.Save()
}

func (s *JSONFileStore) Load() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	b, err := ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}

func (s *JSONFileStore) Save() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	tmp := s.path + ".tmp"
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

func (s *JSONFileStore) Add(t *model.Todo) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Todos = append(s.Todos, t)
}

func (s *JSONFileStore) All() []*model.Todo {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.Todos
}

// OpenIndexes returns a slice of global indexes corresponding to non-done
// todos (open items) in the store. The returned slice should be treated as
// immutable by callers.
func (s *JSONFileStore) OpenIndexes() []int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var idxes []int
	for i, t := range s.Todos {
		if !t.Done {
			idxes = append(idxes, i)
		}
	}
	return idxes
}

// MarkDoneByIndex marks the todo at the given 0-based index as done.
func (s *JSONFileStore) MarkDoneByIndex(i int) error {
	s.mutex.Lock()
	if i < 0 || i >= len(s.Todos) {
		s.mutex.Unlock()
		return errors.New("index out of range")
	}
	s.Todos[i].Done = true
	s.Todos[i].DoneAt = time.Now().UTC()
	s.mutex.Unlock()
	return s.Save()
}

// MarkDoneByID finds a todo by ID and marks it done.
func (s *JSONFileStore) MarkDoneByID(id string) error {
	s.mutex.Lock()
	var found bool
	for _, t := range s.Todos {
		if t.ID == id {
			t.Done = true
			t.DoneAt = time.Now().UTC()
			found = true
			break
		}
	}
	s.mutex.Unlock()
	if !found {
		return errors.New("id not found")
	}
	return s.Save()
}

// RemoveByIndex removes the todo at the given global 0-based index.
func (s *JSONFileStore) RemoveByIndex(i int) error {
	s.mutex.Lock()
	if i < 0 || i >= len(s.Todos) {
		s.mutex.Unlock()
		return errors.New("index out of range")
	}
	s.Todos = append(s.Todos[:i], s.Todos[i+1:]...)
	s.mutex.Unlock()
	return s.Save()
}

// RemoveByID removes the todo with the matching ID.
func (s *JSONFileStore) RemoveByID(id string) error {
	s.mutex.Lock()
	var found bool
	var idx int
	for i, t := range s.Todos {
		if t.ID == id {
			found = true
			idx = i
			break
		}
	}
	if !found {
		s.mutex.Unlock()
		return errors.New("id not found")
	}
	s.Todos = append(s.Todos[:idx], s.Todos[idx+1:]...)
	s.mutex.Unlock()
	return s.Save()
}
