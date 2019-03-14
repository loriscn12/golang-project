package main

import (
	"google.golang.org/grpc"
	
	"log"
	"fmt"
	"net"
	"time"
	"context"
	"flag"
	"golang-project/db/service"

	dpb "golang-project/db/proto"
)

var (
	port = flag.Int("port", 10000, "The server port")
	mongoAddr = flag.String("mongoAddress", "mongodb://localhost:27017", "The address of the MongoDB")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel();
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	dbServer, err := service.New(ctx, &service.Config{MongoAddress: *mongoAddr})
	if err != nil {
		log.Fatalf("failed to instanciate a new DBService client: %s", err)
	}
	defer dbServer.Close(ctx)
	grpcServer := grpc.NewServer()
	dpb.RegisterDatabaseServer(grpcServer, dbServer)
	grpcServer.Serve(lis)
}