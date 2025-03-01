package protocol

import (
	"fmt"
	"github.com/j3ro3n/onionscan/config"
	"github.com/j3ro3n/onionscan/report"
	"github.com/j3ro3n/onionscan/utils"
)

type RicochetProtocolScanner struct {
}

func (rps *RicochetProtocolScanner) ScanProtocol(hiddenService string, osc *config.OnionScanConfig, report *report.OnionScanReport) {
	// Ricochet
	osc.LogInfo(fmt.Sprintf("Checking %s ricochet(9878)\n", hiddenService))
	conn, err := utils.GetNetworkConnection(hiddenService, 9878, osc.TorProxyAddress, osc.Timeout)
	if err != nil {
		osc.LogInfo("Failed to connect to service on port 9878\n")
		report.RicochetDetected = false
	} else {
		osc.LogInfo("Detected possible ricochet instance\n")
		// TODO: Actual Analysis
		report.RicochetDetected = true
	}
	if conn != nil {
		conn.Close()
	}
}
