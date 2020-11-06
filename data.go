package main

const (
	dateFormat = "20060102150405" // equivalent to yyyymmddhhmmss
)

type Response struct {
	IP    string
	Date  string `json:"UTC"`
	Zinfo Zdata  `json:"ATZ"`
	Where string `json:"lctn"` // country/region/city
	Lat   string `json:"lat"`
	Lon   string `json:"lng"`
	Fw    FwData `json:"FOTA"`
}

type Zdata struct {
	Timezone string `json:"n"`
	Current  Zone   `json:"cEra"`
	Next     Zone   `json:"nEra"`
}

type Zone struct {
	Start  string `json:"s"`
	End    string `json:"e"`
	Name   string `json:"n"`
	IsDST  string `json:"d"`
	Offset string `json:"o"`
}

type FwData struct {
	Status string `json:"s"` //NO_FOTA_AVAIL
}
