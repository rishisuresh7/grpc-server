package main

import (
	"fmt"
	"log"
	"net"

	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	formatter "github.com/x-cray/logrus-prefixed-formatter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpc-server/config"
	"grpc-server/factory"
	"grpc-server/interceptor"
	"grpc-server/proto"
	"grpc-server/wrapper"
)

const (
	TCP  = "tcp"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	conf, err := config.NewConfig(wrapper.NewWrapper())
	if err != nil {
		log.Fatalf("Invalid configurations: %s", err.Error())
	}

	logger := logrus.New()
	logger.Formatter = &formatter.TextFormatter{
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: TIME,
	}
	logger.Out = conf.LogFile
	server, err := net.Listen(TCP, fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		logger.Fatalf("Unable to start server: %s\n", err.Error())
	}

	f := factory.NewFactory(conf, logger)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcLogrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
		interceptor.NewInterceptor(conf.Token),
	))
	proto.RegisterBookServiceServer(grpcServer, f.NewBook())
	reflection.Register(grpcServer)
	logger.Infof("Running server on :%d", conf.Port)
	if err = grpcServer.Serve(server); err != nil {
		logger.Fatalf("Unable to start server: %s\n", err.Error())
	}
}
