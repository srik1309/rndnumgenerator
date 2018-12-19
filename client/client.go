package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/srik1309/rndnumgenerator/rndnumgenerator"

	"net/http"

	"google.golang.org/grpc"
)

const (
	address          = "localhost:50052"
	defaultMaxNumber = 12345
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRandomNumberGeneratorClient(conn)

	// Contact the server and print out its response.
	var maxNumber int32 = defaultMaxNumber
	if len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 32) //Is third parameter not honored??
		if err != nil {
			panic(err)
		}
		maxNumber = int32(i)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour*25)
	defer cancel()
	getRandomNumber(c, ctx, maxNumber)

}

func getRandomNumber(c pb.RandomNumberGeneratorClient, ctx context.Context, maxNumber int32) {
	for i := 0; i < 300; i++ {
		r, err := c.SendRandomNumber(ctx, &pb.RandomNumberRequest{MaxNumber: maxNumber})
		if err != nil {
			log.Fatalf("could not get random number: %v", err)
		}
		log.Printf("Random number from server: %s", r.Message)
		postToRestService(r.Message)
		time.Sleep(10 * time.Minute)
	}
}

func postToRestService(randomnumber string) {
	message := map[string]interface{}{
		"name":        randomnumber,
		"description": time.Now(),
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	/*var server string
	if len(os.Args) > 2 {
		server = os.Args[2]
	} else {
		server = "localhost"
	}

	resp, err := http.Post("http://"+server+":9013/", "application/json", bytes.NewBuffer(bytesRepresentation))*/
	resp, err := http.Post("http://appinventory-go", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	//log.Println(result)
	//log.Println(result["data"])
}
