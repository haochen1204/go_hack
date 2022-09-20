package go_hack

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://api.fofa.info"

type FoFa_Client struct {
	email  string
	apiKey string
}

type FoFa_APIInfo struct {
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

type FoFa_Host struct {
	Ip              int    `json:"ip"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	Country         int    `json:"country"`
	Country_Name    string `json:"country_name"`
	Region          string `json:"region"`
	City            string `json:"city"`
	Longitude       string `json:"logitude"`
	As_Number       string `json:"as_number"`
	As_OrganizAtion string `json:"as_organization"`
	Host            string `json:"host"`
	Domain          string `json:"domain"`
	Os              string `json:"os"`
	Server          string `json:"server"`
	Icp             string `json:"icp"`
	Title           string `json:"title"`
	Jarm            string `json:"jarm"`
	Header          string `json:"header"`
	Banner          string `json:"banner"`
	Cert            string `json:"cert"`
}

type FoFa_MsgSearch struct {
	Error   bool        `json:"error"`
	Size    int         `json:"size"`
	Page    int         `json:"page"`
	Mode    string      `json:"mode"`
	Query   string      `json:"query"`
	Results []FoFa_Host `json:"results"`
}

type FoFa_InfoSearch struct {
	Qbase64 string
	Fields  string
	Page    int
	Size    int
	Full    bool
}

func New_FoFa_Client(email string, apiKey string) *FoFa_Client {
	return &FoFa_Client{email: email, apiKey: apiKey}
}

func New_FoFa_InfoSearch(q string) *FoFa_InfoSearch {
	q = base64.StdEncoding.EncodeToString([]byte(q))
	p := FoFa_InfoSearch{Qbase64: q, Fields: "no", Page: 0, Size: 0, Full: false}
	return &p
}

func (s *FoFa_Client) HostSearch(q *FoFa_InfoSearch) (*FoFa_MsgSearch, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/api/v1/search/all?email=%s&key=%s&qbase64=%s", BaseURL, s.email, s.apiKey, q.Qbase64),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret FoFa_MsgSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, err
}

func (s *FoFa_Client) APIInfo() (*FoFa_APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api/v1/info/my?email=%s&key=%s", BaseURL, s.email, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret FoFa_APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
