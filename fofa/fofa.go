package fofa

const BaseURL = "https://fofa.info"

type Client struct {
	email  string
	apiKey string
}

func New(email string, apiKey string) *Client {
	return &Client{email: email, apiKey: apiKey}
}
