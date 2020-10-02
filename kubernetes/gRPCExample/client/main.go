package main

import (
	"context"
	"log"
	"os"
	"time"

	echo "github.com/WeiWeiWesley/ithome12/kubernetes/gRPCExample"
	"google.golang.org/grpc"
)

func main() {
	currentPod := os.Getenv("POD_IP")
	log.Println("Current pod ip:", currentPod)

	conn, err := grpc.Dial("grpc-server:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("grpc conn fail: %v", err)
	}
	defer conn.Close()

	echoClient := echo.NewEchoServiceClient(conn)

	for {
		time.Sleep(time.Second)
		resp, err := echoClient.Echo(context.Background(), &echo.EchoRequest{ClientAddress: currentPod})
		if err != nil {
			log.Println("Request err", err.Error())
		}

		log.Println("Resp:", resp.ServerAddress)
	}

}
