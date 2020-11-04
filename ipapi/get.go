package ipapi

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
)

func GetLocation (ip string) (*GeoLocation, error) {
    url := strings.Join([]string{APIbase, ip, "?fields=", fieldsNum}, "")

    // prepare request object
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    // execute request
    resp, err := http.DefaultClient.Do(req)
    defer resp.Body.Close()
    if err != nil {
        return nil, err
    }

    // check HTTP response code
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed: %s", resp.Status)
    }

    // decode received JSON
    var geoLocation GeoLocation
    if err := json.NewDecoder(resp.Body).Decode(&geoLocation); err != nil {
        return nil, err
    }
    return &geoLocation, nil
}
