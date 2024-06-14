package grpcclientconn

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	serviceAddr string
	conn        *grpc.ClientConn
}

func NewGRPCClient(serviceAddr string) *GRPCClient {
	return &GRPCClient{
		serviceAddr: serviceAddr,
	}
}

func (oc *GRPCClient) Connect() error {
	var err error
	oc.conn, err = grpc.Dial(oc.serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return err
}

func (oc *GRPCClient) Disconnect() error {
	if oc.conn != nil {
		return oc.conn.Close()
	}
	return nil
}

func (oc *GRPCClient) GetConnection() *grpc.ClientConn {
	return oc.conn
}
