package chapa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	chapaAcceptPaymentV1APIURL = "https://api.chapa.co/v1/transaction/initialize"
	chapaVerifyPaymentV1APIURL = "https://api.chapa.co/v1/transaction/verify/%v"
)

type (
	ChapaAPI interface {
		PaymentRequest(request *ChapaPaymentRequest) (*ChapaPaymentResponse, error)
		Verify(txnRef string) error
	}

	Chapa struct {
		apiKey string
		client *http.Client
	}
)

func New(apiKey string) *Chapa {
	return &Chapa{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}

func (c *Chapa) PaymentRequest(request *ChapaPaymentRequest) (*ChapaPaymentResponse, error) {

	data, err := json.Marshal(request)
	if err != nil {
		return &ChapaPaymentResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, chapaAcceptPaymentV1APIURL, bytes.NewBuffer(data))
	if err != nil {
		return &ChapaPaymentResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Close = true

	resp, err := c.client.Do(req)
	if err != nil {
		return &ChapaPaymentResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ChapaPaymentResponse{}, err
	}

	var chapaPaymentResponse ChapaPaymentResponse

	err = json.Unmarshal(body, &chapaPaymentResponse)
	if err != nil {
		return &ChapaPaymentResponse{}, err
	}

	return &chapaPaymentResponse, nil
}

func (c *Chapa) Verify(txnRef string) (*ChapaVerifyResponse, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(chapaVerifyPaymentV1APIURL, txnRef), nil)
	if err != nil {
		return &ChapaVerifyResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Close = true

	resp, err := c.client.Do(req)
	if err != nil {
		return &ChapaVerifyResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ChapaVerifyResponse{}, err
	}

	var chapaVerifyResponse ChapaVerifyResponse

	err = json.Unmarshal(body, &chapaVerifyResponse)
	if err != nil {
		return &ChapaVerifyResponse{}, err
	}

	return &chapaVerifyResponse, nil
}
