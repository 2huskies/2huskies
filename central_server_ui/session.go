package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

func initCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		log.Fatalf("init cache: %s", err)
	}
	// Assign the connection to the package level `cache` variable
	cache = conn
}

type session struct {
	token string
}

func getSession(req *http.Request) (*session, error) {
	c, err := req.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		}
		return nil, err
	}
	sessionToken := c.Value

	return &session{sessionToken}, nil
}

func (s *session) getAbiturientID() int64 {
	key := s.token + "-abiturient-id"
	str, err := redis.String(cache.Do("GET", key))
	if err != nil {
		panic(err)
	}
	//	log.Printf("sess.resp: %s", str)
	id, _ := strconv.ParseInt(str, 10, 64)
	return id
}

func (s *session) getRole() string {
	resp, _ := cache.Do("GET", s.token+"-role")
	str, _ := resp.(string)
	return str
}
