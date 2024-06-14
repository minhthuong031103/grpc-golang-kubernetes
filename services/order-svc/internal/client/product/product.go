package productclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	serviceAddr string
	conn        *grpc.ClientConn
}

func NewProductClient(serviceAddr string) *ProductClient {
	return &ProductClient{
		serviceAddr: serviceAddr,
	}
}

func (oc *ProductClient) Connect() error {
	var err error
	oc.conn, err = grpc.Dial(oc.serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return err
}

func (oc *ProductClient) Disconnect() error {
	if oc.conn != nil {
		return oc.conn.Close()
	}
	return nil
}

func (oc *ProductClient) GetConnection() *grpc.ClientConn {
	return oc.conn
}
