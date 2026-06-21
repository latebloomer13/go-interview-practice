// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"sync"
	// Add any other necessary imports
)

//hints here are inconsistent

// Client represents a connected chat client
type Client struct {
	// TODO: Implement this struct
	// Hint: username, message channel, mutex, disconnected flag
	Username string
	Messages chan string
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	select {
	case c.Messages <- message:
	default:
		// Channel is full, handle gracefully
		// return errors.New("recipient's message queue is full")
	}
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	// Hint: read from channel, handle closed channel
	message, ok := <-c.Messages
	if !ok {
		return "Channel is closed"
	}
	return message
}

type BroadcastMessage struct {
	Content string
	Sender  *Client
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// TODO: Implement this struct
	// Hint: clients map, mutex
	clients    map[string]*Client
	broadcast  chan BroadcastMessage
	disconnect chan *Client
	connect    chan *Client
	mu         sync.RWMutex
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	// TODO: Implement this function
	chatServer := &ChatServer{clients: map[string]*Client{}}
	chatServer.run()
	return chatServer
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// TODO: Implement this method
	// Hint: check username, create client, add to map
	s.mu.Lock()
	_, ok := s.clients[username]
	s.mu.Unlock()
	if ok {
		return nil, ErrUsernameAlreadyTaken
	}
	client := &Client{
		Username: username,
		Messages: make(chan string, 100),
	}
	s.connect <- client

	return client, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// Hint: remove from map, close channels
	s.mu.Lock()
	if _, ok := s.clients[client.Username]; ok {
		s.disconnect <- client
	}
	s.mu.Unlock()
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	s.broadcast <- BroadcastMessage{
		Sender:  sender,
		Content: message,
	}
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	s.mu.RLock()
	client, exists := s.clients[recipient]
	s.mu.RUnlock()
	if !exists {
		return errors.New("recipient not found")
	}
	select {
	case client.Messages <- message:
		return nil
	default:
		return errors.New("recipient's message queue is full")
	}
}

func (s *ChatServer) run() {
	for {
		select {
		case client := <-s.connect:
			s.mu.Lock()
			s.clients[client.Username] = client
			s.mu.Unlock()
		case client := <-s.disconnect:
			// Handle disconnection
			s.mu.Lock()
			close(client.Messages)
			delete(s.clients, client.Username)
			s.mu.Unlock()
		case msg := <-s.broadcast:
			s.mu.Lock()
			for _, client := range s.clients {
				client.Send(msg.Content)
			}
			s.mu.Unlock()
		}
	}
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
