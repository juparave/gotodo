package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

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
