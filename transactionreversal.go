package mpesasdk;

// TransactionReversalRequest represents the parameters for reversing a transaction.
// This is used to cancel or refund a transaction that has already been processed.
type TransactionReversalRequest struct {
    // Initiator is the username of the API operator initiating the reversal.
    Initiator string `json:"Initiator"`

    // SecurityCredential is the encrypted password for the operator initiating the reversal.
    SecurityCredential string `json:"SecurityCredential"`

    // CommandID must be "TransactionReversal" for a reversal request.
    CommandID CommandID `json:"CommandID"`

    // TransactionID is the ID of the transaction to be reversed.
    TransactionID string `json:"TransactionID"`

    // Amount is the amount to be reversed.
    Amount uint64 `json:"Amount"`

    // ReceiverParty is the organization receiving the funds.
    ReceiverParty string `json:"ReceiverParty"`

    // RecieverIdentifierType defines the type of organization receiving the transaction.
    RecieverIdentifierType string `json:"RecieverIdentifierType"`

    // QueueTimeOutURL is the URL for timeout notifications.
    QueueTimeOutURL string `json:"QueueTimeOutURL"`

    // ResultURL is the URL to receive the result of the reversal.
    ResultURL string `json:"ResultURL"`

    // Remarks are optional comments to include with the reversal.
    Remarks string `json:"Remarks"`

    // Occasion is an optional field for additional details.
    Occasion string `json:"Occasion"`

    OriginatorConversationID string `json:"OriginatorConversationID"`
}

