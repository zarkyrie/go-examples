package programmingmode

import (
	"fmt"
	"testing"
)

type Config struct {
	Timeout int
	MaxConn int
}

type Server struct {
	addr     string
	port     int
	protocol string
	config   *Config
}

type Option func(*Server)

func NewTimeout(timeout int) Option {
	return func(s *Server) {
		s.config.Timeout = timeout
	}
}

func NewMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.config.MaxConn = maxConn
	}
}

func NewServer(addr string, port int, protocol string, options ...Option) (*Server, error) {
	s := Server{"localhost", 8080, "tcp", &Config{10, 20}}

	for _, op := range options {
		op(&s)
	}

	return &s, nil
}

func TestFunctionalOptions(t *testing.T) {
	s, _ := NewServer("localhost", 8080, "tcp")
	fmt.Println(s)
}
