package models

type Data struct {
	MinimumSeverity  string
	Active           bool
	Verifed          bool
	ScanType         int64
	Engagement       string
	CloseOldFindings bool
	PushToJira       bool
	Token            string
	URI              string
	File             string
	ListScan         bool
}
