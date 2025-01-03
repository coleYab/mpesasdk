package account;

// AccountBalanceRequest represents the parameters for querying the account balance for a shortcode.
// This is used to check the balance of a business shortcode.
type AccountBalanceRequest struct {
    // CommandID specifies the request type, typically "AccountBalanceCommandID".
    CommandID string `json:"CommandID"`

    // IdentifierType defines the type of identifier used for PartyA (usually "Shortcode").
    IdentifierType int `json:"IdentifierType"` // Changed to int

    // Initiator is the username used to authenticate the balance query request.
    Initiator string `json:"Initiator"`

    // PartyA is the shortcode querying for the balance.
    PartyA int `json:"PartyA"` // Changed to int to match the example

    // QueueTimeOutURL is the URL for notifications if the request times out.
    QueueTimeOutURL string `json:"QueueTimeOutURL"`

    // Remarks are optional comments that can be sent with the request.
    Remarks string `json:"Remarks"`

    // ResultURL is the URL to receive the result of the balance query.
    ResultURL string `json:"ResultURL"`

    // SecurityCredential is the encrypted password for the initiator.
    SecurityCredential string `json:"SecurityCredential"`

    // OriginatorConversationID is a unique identifier for the originator of the transaction.
    OriginatorConversationID string `json:"OriginatorConversationID"`
}

type Response struct {

}
