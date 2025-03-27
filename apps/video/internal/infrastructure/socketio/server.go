package socketio

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type Server struct {
	server    *socketio.Server
	mutex     sync.RWMutex
	rooms     map[string]map[string]bool // roomID -> map[sessionID]bool
	sessions  map[string]*Session
	onConnect func(s *Session)
}

type Session struct {
	ID        string
	UserID    string // Optional, can be empty for non-authenticated users
	Socket    socketio.Conn
	CreatedAt int64
	Rooms     map[string]bool
}

type VideoEvent struct {
	VideoID string      `json:"video_id"`
	UserID  string      `json:"user_id"`
	Type    string      `json:"type"`
	Data    interface{} `json:"data,omitempty"`
}

func NewServer() (*Server, error) {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{},
			&websocket.Transport{},
		},
	})

	s := &Server{
		server:   server,
		rooms:    make(map[string]map[string]bool),
		sessions: make(map[string]*Session),
	}

	// Handle connection
	server.OnConnect("/", func(socket socketio.Conn) error {
		session := &Session{
			ID:        socket.ID(),
			Socket:    socket,
			CreatedAt: time.Now().Unix(),
			Rooms:     make(map[string]bool),
		}
		fmt.Println("session", session)
		s.mutex.Lock()
		s.sessions[session.ID] = session
		s.mutex.Unlock()

		if s.onConnect != nil {
			s.onConnect(session)
		}

		log.Printf("Client connected: %s", session.ID)
		return nil
	})

	// Handle disconnection
	server.OnDisconnect("/", func(socket socketio.Conn, reason string) {
		s.mutex.Lock()
		if session, ok := s.sessions[socket.ID()]; ok {
			// Remove from all rooms
			for room := range session.Rooms {
				delete(s.rooms[room], session.ID)
				if len(s.rooms[room]) == 0 {
					delete(s.rooms, room)
				}
			}
			delete(s.sessions, socket.ID())
		}
		s.mutex.Unlock()
		log.Printf("Client disconnected: %s, reason: %s", socket.ID(), reason)
	})

	// Handle video events
	server.OnEvent("/", "video:view", s.handleVideoView)
	server.OnEvent("/", "video:like", s.handleVideoLike)
	server.OnEvent("/", "video:comment", s.handleVideoComment)

	return s, nil
}

func (s *Server) handleVideoView(socket socketio.Conn, msg string) {
	var event VideoEvent
	if err := json.Unmarshal([]byte(msg), &event); err != nil {
		log.Printf("Error unmarshaling video view event: %v", err)
		return
	}

	// Join video room
	roomID := fmt.Sprintf("video:%s", event.VideoID)
	s.JoinRoom(socket.ID(), roomID)

	// Broadcast to room
	s.BroadcastToRoom(roomID, "video:view", event)
}

func (s *Server) handleVideoLike(socket socketio.Conn, msg string) {
	var event VideoEvent
	if err := json.Unmarshal([]byte(msg), &event); err != nil {
		log.Printf("Error unmarshaling video like event: %v", err)
		return
	}

	roomID := fmt.Sprintf("video:%s", event.VideoID)
	s.BroadcastToRoom(roomID, "video:like", event)
}

func (s *Server) handleVideoComment(socket socketio.Conn, msg string) {
	var event VideoEvent
	if err := json.Unmarshal([]byte(msg), &event); err != nil {
		log.Printf("Error unmarshaling video comment event: %v", err)
		return
	}

	roomID := fmt.Sprintf("video:%s", event.VideoID)
	s.BroadcastToRoom(roomID, "video:comment", event)
}

func (s *Server) JoinRoom(sessionID string, roomID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	session, ok := s.sessions[sessionID]
	if !ok {
		return
	}

	if _, ok := s.rooms[roomID]; !ok {
		s.rooms[roomID] = make(map[string]bool)
	}

	s.rooms[roomID][sessionID] = true
	session.Rooms[roomID] = true
	session.Socket.Join(roomID)
}

func (s *Server) LeaveRoom(sessionID string, roomID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	session, ok := s.sessions[sessionID]
	if !ok {
		return
	}

	if _, ok := s.rooms[roomID]; ok {
		delete(s.rooms[roomID], sessionID)
		if len(s.rooms[roomID]) == 0 {
			delete(s.rooms, roomID)
		}
	}

	delete(session.Rooms, roomID)
	session.Socket.Leave(roomID)
}

func (s *Server) BroadcastToRoom(roomID string, event string, data interface{}) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if room, ok := s.rooms[roomID]; ok {
		for sessionID := range room {
			if session, ok := s.sessions[sessionID]; ok {
				session.Socket.Emit(event, data)
			}
		}
	}
}

func (s *Server) SetOnConnect(fn func(s *Session)) {
	s.onConnect = fn
}

func (s *Server) Serve() error {
	go s.server.Serve()
	return nil
}

func (s *Server) Close() error {
	return s.server.Close()
}
