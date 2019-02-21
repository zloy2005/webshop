package signedpay_api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"

	"gopkg.in/resty.v1"
)

const baseUri = `https://pay.sp-stage.us/api/v1/`

type Api struct {
	privateKey string
	merchantID string
	client     *resty.Client
}

func New(merchantID, privateKey string, debug bool) *Api {
	client := resty.New().SetDebug(debug)
	client.Header.Add("Content-Type", "application/json")
	client.Header.Add("Accept", "application/json")
	client.Header.Add("Merchant", merchantID)
	return &Api{privateKey: privateKey, merchantID: merchantID, client: client}
}

func (a *Api) Charge(req *ChargeRequest) (*PaymentResponse, error) {
	resp := &chargeResponseWithApiError{}
	if err := a.postRequest("charge", req, resp); err != nil {
		return nil, err
	}
	return &resp.PaymentResponse, nil
}

func (a *Api) Recurring(req *RecurringRequest) (*PaymentResponse, error) {
	resp := &recurringResponseWithApiError{}
	if err := a.postRequest("recurring", req, resp); err != nil {
		return nil, err
	}
	return &resp.PaymentResponse, nil
}

func (a *Api) Refund(req *RefundRequest) (*PaymentResponse, error) {
	resp := &recurringResponseWithApiError{}
	if err := a.postRequest("refund", req, resp); err != nil {
		return nil, err
	}
	return &resp.PaymentResponse, nil
}

func (a *Api) Status(req *StatusRequest) (*PaymentResponse, error) {
	resp := &statusResponseWithApiError{}
	if err := a.postRequest("status", req, resp); err != nil {
		return nil, err
	}
	return &resp.PaymentResponse, nil
}

func (a *Api) signature(data string) string {
	h := hmac.New(sha512.New, []byte(a.privateKey))
	h.Write([]byte(a.merchantID + data + a.merchantID))
	return base64.URLEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}

func (a *Api) postRequest(method string, req, result interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	data := string(b)
	response, err := a.client.R().SetHeader("Signature", a.signature(data)).SetBody(data).
		Post(baseUri + method)
	if err != nil {
		return err
	}
	if response.StatusCode() != 200 {
		return errors.New(response.Status())
	}
	if err := json.Unmarshal(response.Body(), result); err != nil {
		return err
	}
	if err := apiError(result.(Error)); err != nil {
		return err
	}
	return nil
}
