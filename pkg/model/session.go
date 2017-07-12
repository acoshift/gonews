package model

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session struct {
	ID     string
	UserID string
	Flash  url.Values
}

var sessionStore struct {
	sync.RWMutex
	data map[string]*Session
}

func generateID() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return base64.URLEncoding.EncodeToString(buf)
}

func CreateSession() *Session {
	return &Session{
		ID:    generateID(),
		Flash: make(url.Values),
	}
}

func GetSession(r *http.Request) *Session {
	id, err := r.Cookie("session")
	if err != nil {
		return CreateSession()
	}
	sessionStore.RLock()
	defer sessionStore.RUnlock()
	if sessionStore.data == nil {
		return CreateSession()
	}
	s := sessionStore.data[id.Value]
	if s == nil {
		return CreateSession()
	}
	return s
}

func (s *Session) Save(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    s.ID,
		MaxAge:   int(10 * time.Minute / time.Second),
		HttpOnly: true,
		Path:     "/",
	})
	sessionStore.Lock()
	defer sessionStore.Unlock()
	if sessionStore.data == nil {
		sessionStore.data = make(map[string]*Session)
	}
	sessionStore.data[s.ID] = s
}
