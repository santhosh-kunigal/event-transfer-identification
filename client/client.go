package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type IClient interface {
	GetTransactionReceiptByHash(txHash string) (Receipts, error)
}

type Client struct {
	client http.Client
}

func NewClient(client http.Client) IClient {
	return &Client{
		client: client,
	}
}

func (c *Client) GetTransactionReceiptByHash(txHash string) (Receipts, error) {
	var receipts Receipts
	url := "https://docs-demo.quiknode.pro/"
	method := "POST"
	payload := strings.NewReader(`{"method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":1,"jsonrpc":"2.0"}`)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return Receipts{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return Receipts{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		defer res.Body.Close()
	}
	err = json.Unmarshal(body, &receipts)
	if err != nil {
		return Receipts{}, err
	}
	return receipts, nil
}
