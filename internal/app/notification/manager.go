package notification

import "sync"

type Manager interface {
	AddClient(key string, client chan string)
	RemoveClient(key string, client chan string)
	Notify(key, message string)
}

type InMemoryManager struct {
	clients map[string]map[chan string]bool
	mu      sync.RWMutex
}

func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		clients: make(map[string]map[chan string]bool),
	}
}

func (n *InMemoryManager) AddClient(key string, client chan string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.clients[key] == nil {
		n.clients[key] = make(map[chan string]bool)
	}
	n.clients[key][client] = true
}

func (n *InMemoryManager) RemoveClient(key string, client chan string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if clients := n.clients[key]; clients != nil {
		delete(clients, client)
		if len(clients) == 0 {
			delete(n.clients, key)
		}
	}
	close(client)
}

func (n *InMemoryManager) Notify(key, message string) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	for client := range n.clients[key] {
		select {
		case client <- message:
		default:
		}
	}
}
