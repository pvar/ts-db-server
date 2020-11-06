package main

import (
    "fmt"
    tzdb "github.com/pvar/ts-db-generator/tzdb"
)

func findZoneIndex(instant int64, zoneSet []tzdb.Zone) (int, error) {

    // abort if no zones at all!
    if len(zoneSet) == 0 {
        return 0, fmt.Errorf("No zone available!")
    }

    // abort if all zones are in the future
    if zoneSet[0].Start > instant {
        return 0, fmt.Errorf("No suitable zone found!")
    }

    // get last one if all zones are in the past
    if zoneSet[len(zoneSet) - 1].Start < instant {
        return len(zoneSet) - 1, nil
    }

    lo := 0
    hi := len(zoneSet) - 1

    for hi - lo > 1 {
        med := (lo + hi) / 2 // from: lo + (hi - lo) / 2
        if instant < zoneSet[med].Start {
            hi = med
        } else {
            lo = med
        }
    }

    // abort if no zone found or the remaining is in the past
    if lo == len(zoneSet) || zoneSet[lo].Start > instant {
        return 0, fmt.Errorf("No suitable zone found!")
    }

    return lo, nil
}
