package ipapi

const (
    APIbase = "http://ip-api.com/json/"
    fieldsStr = "status,message,continent,continentCode,country,countryCode,region,regionName,city,lat,lon,timezone"
    fieldsNum = "3195359"
)

type GeoLocation struct {
    IP              string  `json:"ip"`
    Status          string  `json:"status"`
    Message         string  `json:"message"`
    Continent       string  `json:"continent"`
    ContinentCode   string  `json:"continentCode"`
    Country         string  `json:"country"`
    CountryCode     string  `json:"countryCode"`
    Region          string  `json:"region"`
    RegionName      string  `json:"regionName"`
    City            string  `json:"city"`
    Lat             float64 `json:"lat"`
    Lon             float64 `json:"lon"`
    Timezone        string  `json:"timezone"`
    Offset          int     `json:"offset"`
    Currency        string  `json:"currency"`
    Mobile          bool    `json:"mobile"`
    Proxy           bool    `json:"proxy"`
    Hosting         bool    `json:"hosting"`
}
