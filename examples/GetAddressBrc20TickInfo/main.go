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
	resp, err := unisat.GetAddressBrc20TickInfo(ctx, unisat.DefaultMainnet, bear, "bc1pxgpyskz606rlpfjsg3jsluuqedxflyrnfdscftcc7slesqlgqxfsa2mvcf", "3518")
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "  ")
	log.Printf("%s", string(data))
}
