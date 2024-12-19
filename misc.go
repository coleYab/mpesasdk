package mpesaclient

import "io"

func getTransactionStatus() {
	url := "https://apisandbox.safaricom.et/mpesa/transactionstatus/v1/query"
  method := "POST"
  payload := strings.NewReader({
    "initiator": "testapi",
    "SecurityCredential": "Hbe/vB/dOatkEZOpwhiygdHoykgrqJXKjoh9iCWliWAPqOVQZMMuQce/P7QB4vGsug59bBibshi8PunvN0kD4wGShTsRpysr0me6uwClzR03x5FLyGDGgzJYjAI4yuxjm93cyZ3WDvja1HwtXqWMkS0WKxwNUHXGJlcn+yKIBxhl9EyqcKKNkFNg0jQrGy+y/ZP/FxFa1eazUdHzKv9x+olcfH+H8FH9o3FL+ekuCoAAFEq/1JUUMLQR6IQi9IibleitkYtFup0Gl0tgbGtLgmdvprCh8cgAsQUrh0aniVD1nv8BCAV06FL7PXb4R7wnt/YtNW2iA4+hLpjfaBHAAQ==",
    "CommandID": "TransactionStatusQuery",
    "TransactionID": "OEI2AK4Q16",
    "PartyA": 600981,
    "IdentifierType": "1",
    "ResultURL": "https://mydomain.com/TransactionStatus/result/",
    "QueueTimeOutURL": "https://mydomain.com/TransactionStatus/queue/",
    "Remarks": "jkshjkfhjksdhfj",
    "Occassion": "jshkfjkdshkjfhsdjk",
  })

  client := &http.Client {
	}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
	
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer 5XSQlMPT87AG7KGdfbSO5c4EMa40")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}


func getAccountBalance() {
  url := "https://apisandbox.safaricom.et/mpesa/accountbalance/v2/query"
  method := "POST"
  payload := strings.NewReader({
    "initiator": "testapi",
    "SecurityCredential": "e7RuwD8vFKw/boiDASY7PdBvLz2Mq4lcfTplTqRxgGXMrM/zJS80L9G8D41RxfJL3but+G4qkyWPP+bG9c/1TcUc9sT2UkP3kwb51IKB+mNtlKRG1yfM0JMQe8fbLrPM6GGyy9O5VWnvLNQgTu6xk9dInoGcPpPGpZCU1f516kutSH4qAAwynPNsXO9SmAbpyzVfr80zYi2OCtNkuhegUkiNudNOI12tnJjmZiwNQX2p8w80ZrXj0mueimQWy55LyqXuZh9saASjOlbSf4bXX/n6XID/7r98uScH5oovAd0EEAOaKoXgnIpNjyNjg5fE1WR02WkiGNUoQoJNAE0nJQ==",
    "CommandID": "AccountBalance",
    "PartyA": 600991,
    "IdentifierType": "4",
    "Remarks": "Account Balance",
    "QueueTimeOutURL": "https://95ae-167-172-63-139.eu.ngrok.io/b2b/result",
    "ResultURL": "https://95ae-167-172-63-139.eu.ngrok.io/b2b/result",
  "OriginatorConversationID": "537736249139743",
  })
  client := &http.Client {
	}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
	
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer x4FmGlK2EPMYqiqLDq0bm7FyiKUS")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}