package unisat

import (
	"context"
	"github.com/ninja0404/go-unisat/common"
)

func GetBestBlockHeight(ctx context.Context, server, bear string) (ResponseBestHeight, error) {
	var resp ResponseBestHeight
	url := BestBlockHeight(server)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20List(ctx context.Context, server, bear string, offset, limit int64) (ResponseBrc20List, error) {
	var resp ResponseBrc20List
	url := Brc20List(server, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20Info(ctx context.Context, server, bear, ticker string) (ResponseBrc20Info, error) {
	var resp ResponseBrc20Info
	url := Brc20Info(server, ticker)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20Holders(ctx context.Context, server, bear, ticker string, offset, limit int64) (ResponseBrc20Holders, error) {
	var resp ResponseBrc20Holders
	url := Brc20Holders(server, ticker, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20History(ctx context.Context, server, bear, ticker, eventType string, height, offset, limit int64) (ResponseBrc20History, error) {
	var resp ResponseBrc20History
	url := Brc20History(server, ticker, eventType, height, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20TxHistory(ctx context.Context, server, bear, ticker, tx, eventType string, offset, limit int64) (ResponseBrc20TxHistory, error) {
	var resp ResponseBrc20TxHistory
	url := Brc20TxHistory(server, ticker, tx, eventType, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBrc20HistoryByHeight(ctx context.Context, server, bear string, height, offset, limit int64) (ResponseBrc20HistoryByHeight, error) {
	var resp ResponseBrc20HistoryByHeight
	url := Brc20HistoryByHeight(server, height, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetAddressBrc20Summary(ctx context.Context, server, bear, address string, offset, limit int64) (ResponseBrc20Summary, error) {
	var resp ResponseBrc20Summary
	url := AddressBrc20Summary(server, address, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetAddressBrc20TickInfo(ctx context.Context, server, bear, address, ticker string) (ResponseAddressBrc20TickInfo, error) {
	var resp ResponseAddressBrc20TickInfo
	url := AddressBrc20TickInfo(server, address, ticker)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetAddressBrc20History(ctx context.Context, server, bear, address, ticker, eventType string, offset, limit int64) (ResponseAddressBrc20History, error) {
	var resp ResponseAddressBrc20History
	url := AddressBrc20History(server, address, ticker, eventType, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetTransferableInscriptions(ctx context.Context, server, bear, address, ticker string, offset, limit int64) (ResponseTransferableInscriptions, error) {
	var resp ResponseTransferableInscriptions
	url := TransferableInscriptions(server, address, ticker, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func CreateBrc20Transfer(ctx context.Context, server, bear, receiveAddress string, feeRate int64, outputValue int64, brc20Ticker string, brc20Amount string, devAddress string, devFee int64) (ResponseCreateBrc20Transfer, error) {
	var resp ResponseCreateBrc20Transfer
	url := Brc20Transfer(server)
	data := map[string]interface{}{
		"receiveAddress": receiveAddress,
		"feeRate":        feeRate,
		"outputValue":    outputValue,
		"devAddress":     devAddress,
		"devFee":         devFee,
		"brc20Ticker":    brc20Ticker,
		"brc20Amount":    brc20Amount,
	}
	err := common.PostWithBear(ctx, url, bear, data, &resp)
	return resp, err
}
