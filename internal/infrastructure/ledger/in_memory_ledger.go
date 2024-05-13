package ledger

import (
	"fmt"
	"github.com/asaphin/all-databases-go/internal/app"
	"strings"
	"sync"
)

const keySrparator = "#"

func NewInMemoryEntitiesLedger() *InMemoryEntitiesLedger {
	return &InMemoryEntitiesLedger{
		entities: make(map[string]map[string]*app.Entity),
		mu:       sync.RWMutex{},
	}
}

type InMemoryEntitiesLedger struct {
	entities map[string]map[string]*app.Entity
	mu       sync.RWMutex
}

func (l *InMemoryEntitiesLedger) Add(entity *app.Entity) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.entities[entity.Resource]; ok {
		l.entities[entity.Resource][keyToString(entity.Key)] = entity
		return nil
	}

	l.entities[entity.Resource] = make(map[string]*app.Entity)
	l.entities[entity.Resource][keyToString(entity.Key)] = entity

	return nil
}

func (l *InMemoryEntitiesLedger) Remove(entity *app.Entity) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.entities[entity.Resource], keyToString(entity.Key))

	return nil
}

func (l *InMemoryEntitiesLedger) GetByResource(resource string) ([]*app.Entity, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if len(l.entities[resource]) == 0 {
		return nil, nil
	}

	entities := make([]*app.Entity, 0, len(l.entities[resource]))
	for _, entity := range l.entities[resource] {
		entities = append(entities, entity)
	}

	return entities, nil
}

func (l *InMemoryEntitiesLedger) GetAll() ([]*app.Entity, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	entities := make([]*app.Entity, 0, len(l.entities)*16)

	for k := range l.entities {
		for s := range l.entities[k] {
			entities = append(entities, l.entities[k][s])
		}
	}

	return entities, nil
}

func (l *InMemoryEntitiesLedger) ClearByResource(resource string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.entities, resource)

	return nil
}

func (l *InMemoryEntitiesLedger) ClearAll() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.entities = make(map[string]map[string]*app.Entity)

	return nil
}

func keyToString(k app.EntityKey) string {
	s := make([]string, len(k))

	for _, v := range k {
		switch v.(type) {
		case string:
			s = append(s, v.(string))
		default:
			s = append(s, fmt.Sprintf("%v", v))
		}
	}

	return strings.Join(s, keySrparator)
}
