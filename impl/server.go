package impl

import (
	"gitlab.okta-solutions.com/mashroom/backend/common/health"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"gitlab.okta-solutions.com/mashroom/backend/rightmove"
	"gitlab.okta-solutions.com/mashroom/backend/rightmove/version"
	"google.golang.org/grpc"
	"net"
)

type Server interface {
	rightmove.RightmoveServiceServer
	Serve(addr string)
	Background()
}

type serverImpl struct {
}

func (server *serverImpl) Example(context.Context, *rightmove.RightmoveExampleRequest) (*rightmove.RightmoveExampleResponse, error) {
	panic("implement me")
}

func (server *serverImpl) Background() {
	// background processes
}

func (server *serverImpl) Serve(addr string) {
	if listener, err := net.Listen("tcp", addr); err != nil {
		panic(err)
	} else {
		log.SetHost("rightmove")
		grpcServer := grpc.NewServer()

		rightmove.RegisterRightmoveServiceServer(grpcServer, server)

		healthServer := version.NewHealthServer()
		health.RegisterHealthServiceServer(grpcServer, healthServer)

		log.Infoln("rightmove started")
		if err := grpcServer.Serve(listener); err != nil {
			log.Errorln("gRPC error", err)
		}
	}
}

func NewServer() Server {
	server := &serverImpl{}
	go server.Background()
	return server
}
