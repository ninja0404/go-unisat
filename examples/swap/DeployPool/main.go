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
	resp, err := swap.PreDeployPool(ctx, unisat.DefaultMainnet, bear, "bc1pxgpyskz606rlpfjsg3jsluuqedxflyrnfdscftcc7slesqlgqxfsa2mvcf", "ordi", "sats", time.Now().Unix())
	if err != nil {
		log.Fatalf("request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
