package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type Client struct {
	Conn     net.Conn
	Username string
	Outgoing chan string
}

func (c *Client) ReadMessages(scanner *bufio.Scanner, chat *Chat) {
	defer func() {
		chat.leave <- c
	}()

	for scanner.Scan() {
		message := scanner.Text()
		if message == "/quit" {
			break
		}

		chat.broadcast <- fmt.Sprintf("%s: %s", c.Username, message)
	}
}

// WriteMessages sends messages from the chat to the client
func (c *Client) WriteMessages() {
	for message := range c.Outgoing {
		fmt.Fprintln(c.Conn, message)
	}
}

type Chat struct {
	clients   map[*Client]bool
	broadcast chan string
	join      chan *Client
	leave     chan *Client
	mu        sync.Mutex
}

func NewChat() *Chat {
	return &Chat{
		clients:   make(map[*Client]bool),
		broadcast: make(chan string),
		join:      make(chan *Client),
		leave:     make(chan *Client),
	}
}

func (c *Chat) Run() {
	for {
		select {
		case client := <-c.join:
			c.mu.Lock()
			c.clients[client] = true
			c.mu.Unlock()
			c.broadcast <- fmt.Sprintf("%s has joined the chat", client.Username)

		case client := <-c.leave:
			c.mu.Lock()
			delete(c.clients, client)
			c.mu.Unlock()
			close(client.Outgoing)
			c.broadcast <- fmt.Sprintf("%s has left the chat", client.Username)

		case message := <-c.broadcast:
			c.mu.Lock()
			for client := range c.clients {
				select {
				case client.Outgoing <- message:
					// Message sent successfully
				default:
					// Client buffer is full, remove them
					delete(c.clients, client)
					close(client.Outgoing)
				}
			}
			c.mu.Unlock()
		}
	}
}

func handleconnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	var username string
	if scanner.Scan() {
		username = scanner.Text()
	}

	client := &Client{
		Conn:     conn,
		Username: username,
		Outgoing: make(chan string),
	}

	go client.ReadMessages(scanner)
	client.WriteMessages()

}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	chat := NewChat()
	chat.Run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connections", err)
			continue
		}

		go handleconnection(conn)
	}
}
