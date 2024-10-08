package main

import (
	"context"
	"encoding/json"
	"github.com/ninja0404/go-unisat"
	"github.com/ninja0404/go-unisat/swap"
	"log"
	"os"
	"time"
)

func main() {
	bear := os.Getenv("BEAR")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := swap.PoolList(ctx, unisat.DefaultMainnet, bear, "ordi", 0, 100)
	if err != nil {
		log.Fatalf("request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
