package b2c

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func MakeRefund(transactionId string) {
		url := "https://apisandbox.safaricom.et/mpesa/reversal/v2/request"
		method := "POST"
		data, err := json.Marshal(map[string]interface{}{
			"initiator": "testapi",
			"SecurityCredential": "e7RuwD8vFKw/boiDASY7PdBvLz2Mq4lcfTplTqRxgGXMrM/zJS80L9G8D41RxfJL3but+G4qkyWPP+bG9c/1TcUc9sT2UkP3kwb51IKB+mNtlKRG1yfM0JMQe8fbLrPM6GGyy9O5VWnvLNQgTu6xk9dInoGcPpPGpZCU1f516kutSH4qAAwynPNsXO9SmAbpyzVfr80zYi2OCtNkuhegUkiNudNOI12tnJjmZiwNQX2p8w80ZrXj0mueimQWy55LyqXuZh9saASjOlbSf4bXX/n6XID/7r98uScH5oovAd0EEAOaKoXgnIpNjyNjg5fE1WR02WkiGNUoQoJNAE0nJQ==",
			"CommandID": "TransactionReversal",
			"TransactionID": transactionId,
			"Amount": 100,
			"ReceiverParty": "600991",
			"RecieverIdentifierType": "11",
			"ResultURL": "https://mydomain.com/Reversal/result/",
			"QueueTimeOutURL": "https://mydomain.com/Reversal/queue/",
			"remarks": "Good Reversal",
			"occassion": "Make reversal here",
		  	"OriginatorConversationID": "767062273001823",
		 })

		payload := strings.NewReader(string(data))
		client := &http.Client {}
		req, err := http.NewRequest(method, url, payload)
	  
		if err != nil {
		  	fmt.Println(err)
		}
		  
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer squAEenc6HTGthHw9OYi08I8R7xA")
	  
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		fmt.Println(string(body))	
}