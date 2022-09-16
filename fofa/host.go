package github.com/haochen1204/go_hack

import (
	base64 "encoding/base64"
	"net/http"
	"fmt"
)

type FoFa_Host struct {
	Ip				int 	`json:"ip"`
	Port 			int 	`json:"port"`
	Protocol 		string 	`json:"protocol"`
	Country 		int 	`json:"country"`
	Country_Name 	string 	`json:"country_name"`
	Region 			string 	`json:"region"`
	City 			string 	`json:"city"`
	Longitude 		string 	`json:"logitude"`
	As_Number		string 	`json:"as_number"å`
	As_OrganizAtion string 	`json:"as_organization"`
	Host 			string 	`json:"host"`
	Domain 			string 	`json:"domain"`
	Os 				string 	`json:"os"`
	Server 			string 	`json:"server"`
	Icp 			string 	`json:"icp"`
	Title 			string 	`json:"title"`
	Jarm 			string 	`json:"jarm"`
	Header 			string 	`json:"header"`
	Banner 			string 	`json:"banner"`
	Cert 			string 	`json:"cert"`
}

type FoFa_MsgSearch struct {
	Error 	bool	 	`json:"error"`
	Size 	int 		`json:"size"`
	Page 	int 		`json:"page"`
	Mode 	string 		`json:"mode"`
	Query   string 		`json:"query"`
	Results	[]FoFa_Host `json:"results"`
}

type Fofa_InfoSearch struct {
	Qbase64 string
	Fields 	string
	Page 	int
	Size	int
	Full	bool
}

func New_Fofa_InfoSearch(q string) *Fofa_InfoSearch{
	q := base64.StdEncoding.EncodeToString([]byte(q))
	return Fofa_InfoSearch{Qbase64: q, Fields: "no", Page: 0, Size: 0, Full: false}
}

func (s *FoFa_Client) HostSearch(q *Fofa_InfoSearch) (*FoFa_HostSearch, error){
	res, err := http.Get(
		fmt.Sprintf("%s/api/v1/search/all?email=%s&key=%s&qbase64=%s", BaseURL, s.email, s.apiKey, q.Qbase64),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostSearch
	if err := json.NewDecoder(re.Body).Decode(&ret); err != nil {
		return nil ,err
	}
	return &ret, err
}
