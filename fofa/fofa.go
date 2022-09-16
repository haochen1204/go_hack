package github.com/haochen1204/go_hack

const BaseURL = "https://fofa.info"

type FoFa_Client struct {
	email  string
	apiKey string
}

func New_FoFa_Client(email string, apiKey string) *Fofa_Client {
	return &FoFa_Client{email: email, apiKey: apiKey}
}
