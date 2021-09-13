package server

import "context"

type Server struct {
	ctx context.Context
}

func NewServer(ctx context.Context) *Server {
	return &Server{
		ctx: ctx,
	}
}

func (s *Server) StartStart() {

}
