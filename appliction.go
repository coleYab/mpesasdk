package mpesaclient



type MpesaAppliation struct {
	Name string
	ClientKey string
	ClientSecret string
	Password string
	ApplicationState string // either production or test mode
}

