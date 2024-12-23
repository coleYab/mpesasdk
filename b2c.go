package mpesaclient

func makeBuisnessPayment() {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/v2/paymentrequest"
	method := "POST"
  	payload := strings.NewReader({
    	"InitiatorName": "testapi",
    	"SecurityCredential": "IYK44XLqal3Ijr2qIgGq+EWZjd8+XfYlbaMX7SsK5Djw6gKami5xC2+bt+14BfqwzzsWvavQuRGoqB9O0arjO1AEzfaoFKK/Uilbrcz/WPcc+qJKo/yF5X0uHt2w7MrDjmpQaY6F5cBKhuemnU5lRHNi9LvsPwiiOK8yuMxCtD0rXweR4y9N6G2bfJgu4RO2gn6nelHH0cZZaI+NsEfprwQT62miMmxBeD3VB913bT3SpaBHvgeC0FMfaameHnPe7utdMCAThmeTbx4JN0K3D/VwiUjbwyXqCly30nmXqkvs9aXIvHi54mF74ofXfO/gKx7IZj6wy+lByx1vIvvbJw==",
    	"CommandID": "BusinessPayment",
    	"PartyA": 600998,
    	"PartyB": 251789454545,
    	"Amount": 100,
    	"ResultURL": "https://mydomain.com/b2c/result/",
    	"QueueTimeOutURL": "https://mydomain.com/b2c/timeout",
    	"Remarks": "Good Payment",
    	"Occassion": "Bad payment",
  		"OriginatorConversationID": "1234",
 	})
  	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
	
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer TlHZh21JAxrG9D6j1dp0SdBqynxu")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}

func reverseTransaction() {
	url := "https://apisandbox.safaricom.et/mpesa/reversal/v2/request"
	method := "POST"
	payload := strings.NewReader({
	  "initiator": "testapi",
	  "SecurityCredential": "e7RuwD8vFKw/boiDASY7PdBvLz2Mq4lcfTplTqRxgGXMrM/zJS80L9G8D41RxfJL3but+G4qkyWPP+bG9c/1TcUc9sT2UkP3kwb51IKB+mNtlKRG1yfM0JMQe8fbLrPM6GGyy9O5VWnvLNQgTu6xk9dInoGcPpPGpZCU1f516kutSH4qAAwynPNsXO9SmAbpyzVfr80zYi2OCtNkuhegUkiNudNOI12tnJjmZiwNQX2p8w80ZrXj0mueimQWy55LyqXuZh9saASjOlbSf4bXX/n6XID/7r98uScH5oovAd0EEAOaKoXgnIpNjyNjg5fE1WR02WkiGNUoQoJNAE0nJQ==",
	  "CommandID": "TransactionReversal",
	  "TransactionID": jdfhj8372843,
	  "Amount": 100,
	  "ReceiverParty": "600991",
	  "RecieverIdentifierType": "11",
	  "ResultURL": "https://mydomain.com/Reversal/result/",
	  "QueueTimeOutURL": "https://mydomain.com/Reversal/queue/",
	  "remarks": "Good Reversal",
	  "occassion": "Make reversal here",
	"OriginatorConversationID": "767062273001823",
   })
	client := &http.Client {
	  }
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

func makeSalaryPayment()  {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/v2/paymentrequest"
	method := "POST"
	payload := strings.NewReader({
	  "InitiatorName": "testapi",
	  "SecurityCredential": "IYK44XLqal3Ijr2qIgGq+EWZjd8+XfYlbaMX7SsK5Djw6gKami5xC2+bt+14BfqwzzsWvavQuRGoqB9O0arjO1AEzfaoFKK/Uilbrcz/WPcc+qJKo/yF5X0uHt2w7MrDjmpQaY6F5cBKhuemnU5lRHNi9LvsPwiiOK8yuMxCtD0rXweR4y9N6G2bfJgu4RO2gn6nelHH0cZZaI+NsEfprwQT62miMmxBeD3VB913bT3SpaBHvgeC0FMfaameHnPe7utdMCAThmeTbx4JN0K3D/VwiUjbwyXqCly30nmXqkvs9aXIvHi54mF74ofXfO/gKx7IZj6wy+lByx1vIvvbJw==",
	  "CommandID": "SalaryPayment",
	  "PartyA": 600998,
	  "PartyB": 251789454545,
	  "Amount": 100,
	  "ResultURL": "https://mydomain.com/b2c/result/",
	  "QueueTimeOutURL": "https://mydomain.com/b2c/timeout",
	  "Remarks": "Good Payment",
	  "Occassion": "Bad payment",
	"OriginatorConversationID": "112233",
   })
	client := &http.Client {
	  }
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	}
	  
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer HYzHKWVSaqByH98cfHPDMiMG6Acu")
  
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func makePromtionPayment() {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/v2/paymentrequest"
  method := "POST"
  payload := strings.NewReader({
    "InitiatorName": "testapi",
    "SecurityCredential": "Hbe/vB/dOatkEZOpwhiygdHoykgrqJXKjoh9iCWliWAPqOVQZMMuQce/P7QB4vGsug59bBibshi8PunvN0kD4wGShTsRpysr0me6uwClzR03x5FLyGDGgzJYjAI4yuxjm93cyZ3WDvja1HwtXqWMkS0WKxwNUHXGJlcn+yKIBxhl9EyqcKKNkFNg0jQrGy+y/ZP/FxFa1eazUdHzKv9x+olcfH+H8FH9o3FL+ekuCoAAFEq/1JUUMLQR6IQi9IibleitkYtFup0Gl0tgbGtLgmdvprCh8cgAsQUrh0aniVD1nv8BCAV06FL7PXb4R7wnt/YtNW2iA4+hLpjfaBHAAQ==",
    "CommandID": "PromotionPayment",
    "PartyA": 600981,
    "PartyB": 251723232323,
    "Amount": 100,
    "ResultURL": "https://mydomain.com/b2c/result/",
    "QueueTimeOutURL": "https://mydomain.com/b2c/timeout",
    "Remarks": "Wow this is good",
    "Occassion": "Wow this was bad",
  "OriginatorConversationID": "680325814515894",
 })
  client := &http.Client {
	}
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
	
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer pWpJWigeoi9scITCD9XU7CyMZbjm")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}