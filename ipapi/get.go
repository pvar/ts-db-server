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
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // check HTTP response code
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed: %s", resp.Status)
    }

    // decode received JSON
    var geoLocation GeoLocation
    if err := json.NewDecoder(resp.Body).Decode(&geoLocation); err != nil {
        return nil, err
    }

    if geoLocation.Status != "success" {
        return nil, fmt.Errorf("%s: %s", geoLocation.Status, geoLocation.Message)
    }

    return &geoLocation, nil
}
