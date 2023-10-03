package main

import (
	"fmt"
	"time"
)

func addDaysToEpoch(epochSeconds int64, daysToAdd int) int64 {
	t := time.Unix(epochSeconds, 0)
	t = t.Add(time.Duration(daysToAdd*24) * time.Hour)
	return t.Unix()
}

func main() {
	epochTimestamp := int64(1672473600)
	daysToAdd := 10
	newEpochTimestamp := addDaysToEpoch(epochTimestamp, daysToAdd)
	fmt.Println("New Epoch Timestamp:", newEpochTimestamp)
}
