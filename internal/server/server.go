package server

import (
	"context"
	"net"
	"steppingDisk/internal/config"
	"sync"
)

type Server struct {
	cfg      *config.Config
	listener net.Listener
	connMgr  *connection.Manager
	wg       sync.WaitGroup
}

func New(cfg *config.Config) (*Server, error) {
	listener, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return nil, err
	}

	connMgr, err := connection.NewManager()
	if err != nil {
		return nil, err
	}

	return &Server{
		cfg:      cfg,
		listener: listener,
		connMgr:  connMgr,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				// Handle temporary errors differently
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					continue
				}
				return err
			}

			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				s.handleConnection(ctx, conn)
			}()
		}
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	// Close listener to stop accepting new connections
	if err := s.listener.Close(); err != nil {
		return err
	}

	// Wait for all connections to finish
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}
func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {func (s *Server) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	// This is where we'll implement the protocol detection and handling
	// For now, we'll just log the connection
}}
