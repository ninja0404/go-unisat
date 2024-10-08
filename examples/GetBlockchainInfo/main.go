package main

import (
	"context"
	"encoding/json"
	"github.com/ninja0404/go-unisat"
	"log"
	"os"
	"time"
)

func main() {
	bear := os.Getenv("BEAR")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := unisat.GetBlockchainInfo(ctx, unisat.DefaultMainnet, bear)
	if err != nil {
		log.Fatalf("GetBlockchainInfo error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
