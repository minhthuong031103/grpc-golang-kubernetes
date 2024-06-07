package orderclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	serviceAddr string
	conn        *grpc.ClientConn
}

// NewOrderClient creates a new client for the Order service
func NewOrderClient(serviceAddr string) *OrderClient {
	return &OrderClient{
		serviceAddr: serviceAddr,
	}
}

func (oc *OrderClient) Connect() error {
	var err error
	oc.conn, err = grpc.Dial(oc.serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return err
}

func (oc *OrderClient) Disconnect() error {
	if oc.conn != nil {
		return oc.conn.Close()
	}
	return nil
}

func (oc *OrderClient) GetConnection() *grpc.ClientConn {
	return oc.conn
}
