package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	formatter "github.com/x-cray/logrus-prefixed-formatter"
	"google.golang.org/grpc"
)

const (
	TCP  = "tcp"
	TIME = "2006-01-02 15:04:05"
	PORT = 9000
)

func main() {
	logger := logrus.New()
	logger.Formatter = &formatter.TextFormatter{
		ForceFormatting:  true,
		FullTimestamp:    true,
		TimestampFormat:  TIME,
	}
	server, err := net.Listen(TCP, fmt.Sprintf(":%d", PORT))
	if err != nil {
		logger.Fatalf("Unable to start server: %s\n", err.Error())
	}

	grpcServer := grpc.NewServer()
	logger.Infof("Running server on :%d", PORT)
	if err = grpcServer.Serve(server); err != nil {
		logger.Fatalf("Unable to start server: %s\n", err.Error())
	}
}
