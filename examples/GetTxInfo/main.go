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
	resp, err := unisat.GetTxInfo(ctx, unisat.DefaultMainnet, bear, "4306046f45e39dd1736c9f543eee4d22df053f1fdc3eacac734319a9fa2e5a9b")
	if err != nil {
		log.Fatalf("GetTxInfo error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
