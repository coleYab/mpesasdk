package mpesaclient 

func payBillOnlineUSSDPush() {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/simulatetransaction/v1/request"
	method := "POST"
	payload := strings.NewReader({
			"BusinessShortCode": "174379",
			"Password": "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
			"Timestamp": "20241219091721",
			"TransactionType": "CustomerPayBillOnline",
			"Amount": 100,
			"PartyA": 600991,
			"PartyB": 600000,
			"PartyB": 600000,
			"PhoneNumber": "251708374149",
			"TransactionDesc": "Good Transaction",
			"CallBackURL": "http://172.29.65.59:13345",
			"AccountReference": "sjdhfjksdhfjk",
			"MerchantRequestID": "12",
			"ReferenceData": [{"Key":"BundleName","Value":"Monthly Unlimited Bundle"},{"Key":"BundleType","Value":"Self"},{"Key":"TINNumber","Value":"89234093223"}],
		  })
	client := &http.Client {
	  }
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	}
	  
	req.Header.Add("Content-Type", "application/json")
	// the acess token retrived by the first function
	req.Header.Add("Authorization", "Bearer OJJBlbe8y3GYJXzpXzFHokrAZ65K")
  
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}


// Customer initiated
func payBillOnline() {
	url := "https://apisandbox.safaricom.et/mpesa/b2c/simulatetransaction/v1/request"
	method := "POST"
	payload := strings.NewReader({
	  "ShortCode": 600998,
	  "CommandID": "CustomerPayBillOnline",
	  "Amount": "100",
	  "Msisdn": "251712653434",
	  "BillRefNumber": "01293091209382190781",
	})
	client := &http.Client {
	  }
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	}
	  
	req.Header.Add("Content-Type", "application/json")
	// the acess token retrived by the first function
	req.Header.Add("Authorization", "Bearer OJJBlbe8y3GYJXzpXzFHokrAZ65K")
  
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}