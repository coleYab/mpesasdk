package mpesasdk;


type CommandID string
const (
    TransactionReversalCommandID CommandID = "TransactionReversal"
    AccountBalanceCommandID CommandID = "AccountBalance"
    TransactionStatusCommandID CommandID = "TransactionStatusQuery"
    SalaryPaymentCommandID CommandID = "SalaryPayment"
    BusinessPaymentCommandID CommandID = "BusinessPayment"
    PromotionPaymentCommandID CommandID = "PromotionPayment"
    RegisterURLCommandID CommandID = "RegisterURL"
    CustomerPayBillOnlineCommandID CommandID = "CustomerPayBillOnline"
    CustomerBuyGoodsOnlineCommandID CommandID = "CustomerBuyGoodsOnline"
)
