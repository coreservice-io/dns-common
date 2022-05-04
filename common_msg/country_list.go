package common_msg

type Msg_Resp_CountryList struct {
	API_META_STATUS
	Country_list map[string]CountryData `json:"country_list"`
}

type CountryData struct {
	Country_code   string `json:"country_code"`
	Country_name   string `json:"country_name"`
	Continent_code string `json:"continent_code"`
	Continent_name string `json:"continent_name"`
}
