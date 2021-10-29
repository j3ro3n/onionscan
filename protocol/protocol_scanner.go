package protocol

import (
	"github.com/j3ro3n/onionscan/config"
	"github.com/j3ro3n/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
