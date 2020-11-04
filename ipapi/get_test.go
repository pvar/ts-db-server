package ipapi

import (
    "fmt"
    "testing"
)

func TestGetLocation (t *testing.T) {
    ipz := []string{"94.66.58.66",  // Greece
                    "5.23.80.100",  // Iceland
                    "27.116.57.120",// Afghanistan
                    "27.148.0.100", // China
                    "1.132.10.100", // Australia
                    "202.65.32.10", // Cook Islands
                    "8.242.100.10", // United States
                    "0.0.0.1" }     // invalid IP address

    for _, ip := range ipz {
        gLoc, err := GetLocation(ip)
        if err != nil {
            t.Errorf("%s", err)
            continue
        }
        fmt.Printf("\nQuerying : %q\n", ip)
        fmt.Printf("Continent: %q\n", gLoc.Continent)
        fmt.Printf("Country  : %q\n", gLoc.Country)
        fmt.Printf("Region   : %q\n", gLoc.Region)
        fmt.Printf("City     : %q\n", gLoc.City)
        fmt.Printf("Timezone : %q\n", gLoc.Timezone)
    }
}
