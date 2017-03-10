package configurationstruct

import "time"

type Configuration struct {
	Interval int64		`json:"Interval"`
	Timespec time.Duration	`json:"Timespec"`
        UpdateUsingAPI  int64	`json:"UpdateUsingAPI"`
	DynamicInterval int64	`json:"DynamicInterval"`
        DynamicTimespec time.Duration	`json:"DynamicTimespec"`
}
