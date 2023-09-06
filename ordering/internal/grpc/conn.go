package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(ctx context.Context, addr string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				return
			}
		}
		go func() {
			<-ctx.Done()
			if err = conn.Close(); err != nil {
				return
			}
		}()
	}()

	return
}