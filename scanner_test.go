package portscanner

import (
	"testing"
)

func TestScanPort(t *testing.T) {
	testCases := map[string]struct {
		port                     int
		protocol, expectedStatus string
	}{
		"using an open udp port": {
			port:           80,
			protocol:       "udp",
			expectedStatus: "Open",
		},
		"using a closed udp port": {
			port:           51999999,
			protocol:       "udp",
			expectedStatus: "Closed",
		},
		"using a closed tcp port": {
			port:           519999999,
			protocol:       "tcp",
			expectedStatus: "Closed",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			result := ScanPort(testCase.protocol, "localhost", testCase.port)
			if testCase.expectedStatus != result.Status {
				t.Errorf("Expected: %s, got %s", testCase.expectedStatus, result.Status)
			}
		})
	}
}

func TestScanPorts(t *testing.T) {
	testCases := map[string]struct {
		protocol                       string
		startRange, endRange, expected int
	}{
		"udp ports from 80 - 100": {
			protocol:   "udp",
			startRange: 80,
			endRange:   100,
			expected:   21,
		},
		"udp ports from 499999978 - 499999999": {
			protocol:   "udp",
			startRange: 499999978,
			endRange:   499999999,
			expected:   22,
		},
		"tcp ports from 1000 - 1030": {
			protocol:   "tcp",
			startRange: 1000,
			endRange:   1030,
			expected:   31,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			results := ScanPorts(testCase.protocol, "localhost", testCase.startRange, testCase.endRange)
			if testCase.expected != len(results) {
				t.Errorf("Expected: %v, got %v", testCase.expected, len(results))
			}
		})
	}
}
