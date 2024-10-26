package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	thumbnailgrpc "thumbnail-proxy/internal/grpc/thumbnail"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	thumbnailService thumbnailgrpc.Thumbnail,
	port int,
) *Server {

	gRPCServer := grpc.NewServer()

	thumbnailgrpc.Register(gRPCServer, thumbnailService)
	return &Server{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *Server) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *Server) Run() error {
	const op = "grpcServer.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *Server) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
