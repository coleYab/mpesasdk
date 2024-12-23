package b2c

import (
	"encoding/json"
	"fmt"
	"io"
	mpesaclient "mymodule"
	"net/http"
	"strings"
)

type Payout struct {
	ApplicationCtx mpesaclient.MpesaAppliation
	Sender int
	Reciver int
	Amount float64
	ResultURL string
	QueueTimeoutURL string
	Remarks string
	PayoutType string // will be taken as the command id
	Occassion string
	OriginatorConversationID string
	Status string // COMPLETED, FAILED, INPROGRESS, 
}

func (p Payout) GeneratePayload() (*strings.Reader, error) {
	// do the required things here and return the data
	data := map[string]interface{} {
		"InitiatorName": p.ApplicationCtx.Name,
		"SecurityCredential": p.ApplicationCtx.Password,
		"CommandID": p.PayoutType,
		"PartyA": p.Sender,
		"PartyB": p.Reciver,
		"Amount": p.Amount,
		"ResultURL": p.ResultURL,
		"QueueTimeOutURL": p.QueueTimeoutURL,
		"Remarks": p.Remarks,
		"Occassion": p.Occassion,
	  	"OriginatorConversationID": p.OriginatorConversationID,
   }

   json_data, err := json.Marshal(data)
   if err != nil {
		return nil, err
   }

   return strings.NewReader(string(json_data)), nil
}

func (p *Payout) makePayout() {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/v2/paymentrequest"
	method := "POST"
	payload, err := p.GeneratePayload()
	if err != nil {
		return
	}
	
	client := &http.Client {}
	
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	}
	  
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer HYzHKWVSaqByH98cfHPDMiMG6Acu")
  
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
}


func HandlePayout() {
	
}

func NewPayout() *Payout {
	return &Payout{}
}