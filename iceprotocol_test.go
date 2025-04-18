// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewICEProtocol(t *testing.T) {
	testCases := []struct {
		protoString   string
		shouldFail    bool
		expectedProto ICEProtocol
	}{
		{ErrUnknownType.Error(), true, ICEProtocolUnknown},
		{"udp", false, ICEProtocolUDP},
		{"tcp", false, ICEProtocolTCP},
		{"UDP", false, ICEProtocolUDP},
		{"TCP", false, ICEProtocolTCP},
	}

	for i, testCase := range testCases {
		actual, err := NewICEProtocol(testCase.protoString)
		if testCase.shouldFail {
			assert.Error(t, err, "testCase: %d %v", i, testCase)
		} else {
			assert.NoError(t, err, "testCase: %d %v", i, testCase)
		}
		assert.Equal(t,
			testCase.expectedProto,
			actual,
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestICEProtocol_String(t *testing.T) {
	testCases := []struct {
		proto          ICEProtocol
		expectedString string
	}{
		{ICEProtocolUnknown, ErrUnknownType.Error()},
		{ICEProtocolUDP, "udp"},
		{ICEProtocolTCP, "tcp"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.expectedString,
			testCase.proto.String(),
			"testCase: %d %v", i, testCase,
		)
	}
}
