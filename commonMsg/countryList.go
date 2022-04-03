package commonMsg

type Msg_Resp_CountryList struct {
	API_META_STATUS
	Country_list map[string]CountryData
}

type CountryData struct {
	Country_code   string
	Country_name   string
	Continent_code string
	Continent_name string
}
