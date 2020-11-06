package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tzdb "github.com/pvar/ts-db-generator/tzdb"
	ipapi "github.com/pvar/ts-db-server/ipapi"
	"net/http"
	"time"
)

const dbfile = "./tsdb.sqlite"

var router *gin.Engine

func main() {
	// All times should be UTC times!
	utclocation, _ := time.LoadLocation("UTC")
	time.Local = utclocation

	// Open db with parsed timezone data
	tzdb.OpenRO(dbfile)
	defer tzdb.Close()

	// Set the router as the default one provided by Gin.
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	// Define the route for the server root.
	router.GET("/", handleRootGet)

	// Start serving the application.
	router.Run(":8086")
}

func handleRootGet(c *gin.Context) {
	//ip := c.ClientIP()
	ip := "94.66.58.66"

	// Use IP-API to get location information
	location, err := ipapi.GetLocation(ip)
	if err != nil {
		// No location data!
		// Panic would be a good option too..
		fmt.Printf("No location data received!\n")
		return
	}

	timezone := location.Timezone

	var czone, nzone Zone

	// Use tzdb to get zones in location
	lookupZone := true
	zones, err := tzdb.GetZones(timezone)
	if err != nil {
		// No zones found!
		// Use default values for timezone...
		fmt.Printf("No zones found!\n")
		czone, nzone = getDefaultZones(timezone)
		// Should not even attempt to search in Zones.
		lookupZone = false
	}

	// Find active zone at this time
	if lookupZone {
		now := time.Now().Unix()
		zi, err := findZoneIndex(now, zones)
		if err != nil {
			// No active zone found!
			// Use default values for timezone...
			fmt.Printf("No active zone found!\n")
			czone, nzone = getDefaultZones(timezone)
		} else {
			czone, nzone = getCorrectZones(zi, zones)
		}
	}

	var zdata Zdata
	zdata.Timezone = location.Timezone
	zdata.Current = czone
	zdata.Next = nzone

	var fwdata FwData
	fwdata.Status = "NO_FOTA_AVAIL"

	var respData Response
	respData.IP = ip
	respData.Date = time.Now().Format(dateFormat)
	respData.Where = location.Country + "/" + location.RegionName + "/" + location.City
	respData.Lat = fmt.Sprintf("%f", location.Lat)
	respData.Lon = fmt.Sprintf("%f", location.Lon)
	respData.Zinfo = zdata
	respData.Fw = fwdata

	c.IndentedJSON(http.StatusOK, respData)
}

func getDefaultZones(timezone string) (current, next Zone) {
	var name string
	var offset int64

	original, err := tzdb.GetOriginalByName(timezone)
	if err != nil {
		name = ""
		offset = -1
	} else {
		name = original.DZone
		offset = original.DOffset
	}

	current.Name = name
	current.Start = "?"
	current.End = "gnabgib"
	current.IsDST = "0"
	current.Offset = fmt.Sprintf("%d", offset)

	next.Name = ""
	next.Start = ""
	next.End = ""
	next.IsDST = ""
	next.Offset = ""

	return current, next
}

func getCorrectZones(zi int, zones []tzdb.Zone) (current, next Zone) {

	var boolStr = map[bool]string{true: "1", false: "0"}

	current.Name = zones[zi].Name
	current.Start = time.Unix(zones[zi].Start, 0).Format(dateFormat)
	if zones[zi].End == -1 {
		current.End = "gnabgib"
	} else {
		current.End = time.Unix(zones[zi].End, 0).Format(dateFormat)
	}
	current.IsDST = boolStr[zones[zi].IsDST]
	current.Offset = fmt.Sprintf("%d", zones[zi].Offset)

	if zi < (len(zones) - 1) {
		next.Name = zones[zi+1].Name
		next.Start = time.Unix(zones[zi+1].Start, 0).Format(dateFormat)
		if zones[zi+1].End == -1 {
			current.End = "gnabgib"
		} else {
			next.End = time.Unix(zones[zi+1].End, 0).Format(dateFormat)
		}
		next.IsDST = boolStr[zones[zi+1].IsDST]
		next.Offset = fmt.Sprintf("%d", zones[zi+1].Offset)
	} else {
		next.Name = ""
		next.Start = ""
		next.End = ""
		next.IsDST = ""
		next.Offset = ""
	}

	return current, next
}
