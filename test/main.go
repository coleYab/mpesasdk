package main

import (
	"fmt"

	"github.com/coleYab/mpesasdk"
	"github.com/coleYab/mpesasdk/c2b"
)

const consumerKey string = "<Your Consumer Key>"
const consumerSecret string = "<Your Consumer Secret>"
const passkey string = "<passkey>"
const SecurityCredential string = "<Your security credentials>"

func main() {
    mpesaClient := mpesasdk.NewMpesaClient(consumerKey, consumerSecret)
    outputCount := 0
    for true {
        go mpesaClient.RegisterNewURL( c2b.RegisterC2BURLRequest{
            ShortCode: 123456,
            ResponseType: "Completed",
            CommandID: "RegisterURL",
            ConfirmationURL: "https://gg.tt/sfsd",
            ValidationURL: "https://gg.tt/fsd",
        })
        go mpesaClient.RegisterNewURL( c2b.RegisterC2BURLRequest{
            ShortCode: 123456,
            ResponseType: "Completed",
            CommandID: "RegisterURL",
            ConfirmationURL: "https://gg.tt/sfsd",
            ValidationURL: "https://gg.tt/fsd",
        })
        go mpesaClient.RegisterNewURL( c2b.RegisterC2BURLRequest{
            ShortCode: 123456,
            ResponseType: "Completed",
            CommandID: "RegisterURL",
            ConfirmationURL: "https://gg.tt/sfsd",
            ValidationURL: "https://gg.tt/fsd",
        })
        outputCount += 3
        fmt.Printf("-- Output count: [%v] --- \n", outputCount)

        // mpesaClient.SimulateCustomerInititatedPayment(mpesasdk.SimulateCustomerInititatedPayment{
        //     ShortCode: "123456",
        //     BillRefNumber: "ET343434",
        //     CommandID: "CustomerPayBillOnline",
        //     Amount: 54,
        //     Msisdn: "251789898989",
        // })


        // mpesaClient.MakeB2CPaymentRequest(mpesasdk.B2CRequest{
        //     InitiatorName:       "testapi", // Updated InitiatorName
        //     SecurityCredential: SecurityCredential, // Updated SecurityCredential
        //     CommandID:          "BusinessPayment", // CommandID remains the same
        //     Occasion:           "skjdfksdj", // Updated Occasion
        //     PartyA:             101010, // Updated PartyA
        //     PartyB:             251700100100, // Updated PartyB
        //     Amount:             100, // Amount remains the same
        //     Remarks:            "kjdsfksdj", // Updated Remarks
        //     QueueTimeOutURL:    "https://webhook.site/d9f5ca00-51af-43a6-9a64-7cb9fdf51b2c", // Updated QueueTimeOutURL
        //     ResultURL:          "https://webhook.site/d9f5ca00-51af-43a6-9a64-7cb9fdf51b2c", // Updated ResultURL
        //     OriginatorConversationID: "543168755246895", // Updated OriginatorConversationID
        // })

        // mpesaClient.CheckTransactionStatus(mpesasdk.TransactionStatusRequest{
        //     Initiator: "testApi",
        //     SecurityCredential: SecurityCredential,
        //     CommandID: "TransactionStatusQuery",
        //     TransactionID: "RHJ4BTOYS8",
        //     OriginatorConversationID: "AG_20190826_0000777ab7d848b9e721",
        //     PartyA: "101010",
        //     IdentifierType: "4",
        //     ResultURL: "https://mydomain.com/api/transaction-status/result",
        //     QueueTimeOutURL: "https://mydomain.com/transaction-status/timeout",
        //     Remarks: "OK",
        //     Occasion: "OK",
        // })

        // mpesaClient.AccountBalance(mpesasdk.AccountBalanceRequest{
        //     CommandID:               mpesasdk.AccountBalanceCommandID,
        //     Initiator:               "testapi",
        //     SecurityCredential:      SecurityCredential,
        //     PartyA:                  600984,
        //     IdentifierType:          4,
        //     Remarks:                 "sfsdf",
        //     ResultURL:               "https://webhook.site/d9f5ca00-51af-43a6-9a64-7cb9fdf51b2c",
        //     QueueTimeOutURL:         "https://webhook.site/d9f5ca00-51af-43a6-9a64-7cb9fdf51b2c",
        //     OriginatorConversationID: "277406453938994",
        // })

        // mpesaClient.ReverseTransaction(mpesasdk.TransactionReversalRequest{
        //     Initiator:"testapiuser",
        //     SecurityCredential:"dsam==",
        //     CommandID:"TransactionReversal",
        //     TransactionID:"A644545RED",
        //     Amount:2000,
        //     ReceiverParty:"600610",
        //     RecieverIdentifierType:"4",
        //     ResultURL:"https://darajambili.herokuapp.com/b2c/result",
        //     QueueTimeOutURL:"https://darajambili.herokuapp.com/b2c/timeout",
        //     Remarks:"please",
        //     Occasion:"work",
        //     OriginatorConversationID: "543168754246895",
        // })

        // mpesaClient.STKPushPaymentRequest(passkey, mpesasdk.USSDPushRequest{
        //     MerchantRequestID: "SFC-Testing-9146-4216-9455-e3947ac570fc",
        //     BusinessShortCode: 554433,
        //     Password: "123",
        //     Timestamp: "20160216165627",
        //     TransactionType: mpesasdk.CustomerPayBillOnlineCommandID,
        //     Amount: 10,
        //     PartyA: "251700404789",
        //     PartyB: "554433",
        //     PhoneNumber: "251700404789",
        //     TransactionDesc: "Monthly Unlimited Package via Chatbot",
        //     CallBackURL: "https://apigee-listener.oat.mpesa.safaricomet.net/api/ussd-push/result",
        //     AccountReference: "DATA",
        //     ReferenceData: []mpesasdk.ReferenceDataRequest{},
        // })

    }
}
