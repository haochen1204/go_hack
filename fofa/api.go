package fofa

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	Error      bool   `json:"error"`
	Email      string `json:"email"`
	UserName   string `json:"username"`
	Fcoin      int    `json:"fcoin"`
	Isvip      bool   `json:"isvip"`
	VipLevel   int    `json:"vip_level"`
	IsVerfied  bool   `json:"is_verified"`
	Avatar     string `json:"avatar"`
	Message    string `json:"message"`
	FofaCliVer string `json:"fofacli_ver"`
	FofaServer bool   `json:"fofa_server"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api/v1/info/my?email=%s&key=%s", BaseURL, s.email, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
