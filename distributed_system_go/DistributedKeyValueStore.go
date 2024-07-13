
package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

// Store holds the key-value data
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// Node represents a server node
type Node struct {
	address string
	store   *Store
}

// NewStore creates a new Store
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// NewNode creates a new Node
func NewNode(address string) *Node {
	return &Node{
		address: address,
		store:   NewStore(),
	}
}

// Set a key-value pair
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get a value by key
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	return value, exists
}

// Delete a key
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// HandleConnection handles incoming connections
func (n *Node) HandleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println("Received:", string(buffer[:length]))
	}
}

// Start starts the node server
func (n *Node) Start() {
	ln, err := net.Listen("tcp", n.address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Listening on", n.address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		go n.HandleConnection(conn)
	}
}

func main() {
	node := NewNode(":3000")
	node.Start()
}
