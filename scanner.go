package portscanner

import (
	"net"
	"strconv"
	"time"
)

// ScanResult holds the result of a port scan.
type ScanResult struct {
	// Protocal is the protocol of the url used to scan the port
	// e.g udp, tcp.
	Protocol string
	// Port is the integer representation of the port being scanned.
	Port int
	// Status holds the port availability status which can either be
	// Open or Closed.
	Status string
}

// ScanPort scans a single port and outputs the result using the ScanResult
// format.
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{
		Protocol: protocol,
		Port:     port,
		Status:   "Open",
	}
	scanURL := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, scanURL, 5*time.Second)
	if err != nil {
		result.Status = "Closed"
		return result
	}
	conn.Close()
	return result
}

// ScanPorts scans multiple ports starting from startRange to endRange and then
// returns the ScanResults.
func ScanPorts(protocol, hostname string, startRange, endRange int) []ScanResult {
	var results []ScanResult
	for startRange <= endRange {
		scanResult := ScanPort(protocol, hostname, startRange)
		results = append(results, scanResult)
		startRange++
	}
	return results
}
