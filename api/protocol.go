package api

type AllCasesResponse struct {
	Cases int `json:"cases"`
	Deaths int `json:"deaths"`
	Recovered int `json:"recovered"`
}

type CountryResponse struct {
	Name string `json:"country"`
	Cases int `json:"cases"`
	Deaths int `json:"deaths"`
	Recovered int `json:"recovered"`
	TodayCases int `json:"todayCases"`
	TodayDeaths int `json:"todayDeaths"`
}
