package models

//
type MQ struct {
	Value   string  `json:"Value"`
	Time    string  `json:"Time"`
	CPUINFO float64 `json:"CPUINFO"`
	MEMINFO float64 `json:"MEMINFO"`
}
