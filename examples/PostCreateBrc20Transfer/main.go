package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ninja0404/go-unisat"
)

func main() {
	bear := ""
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := unisat.CreateBrc20Transfer(ctx, unisat.FractalTest, bear, "bc1p0crww6qy20v64z5urkn4vmza8eqw8l5gsquaa9mhlewxm326289q3kyjdf", 2, 546, "satsuni", "58", "", 0)
	if err != nil {
		log.Fatalf("CreateBrc20Transfer error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
