package main

import (
	"testing"

	sm "github.com/flopp/go-staticmaps"
	"github.com/stretchr/testify/assert"
)

var testData = "{\"iss_position\": {\"latitude\": \"25.9866\", \"longitude\": \"-100.7594\"}}"

type MockMapper struct {
	CallGetISSNowData []byte
	CallGetISSNowErr  error

	RenderMapErr error
}

func (m *MockMapper) CallGetISSNow() ([]byte, error) {
	return m.CallGetISSNowData, m.CallGetISSNowErr
}

func (m *MockMapper) RenderMap(ctx *sm.Context) error {
	return m.RenderMapErr
}

func TestGetLocation(t *testing.T) {
	var testcases = []struct {
		description string
		mapper      *MockMapper
		expectLoc   *Location
		expectErr   bool
	}{
		{
			description: "Get location success",
			mapper: &MockMapper{
				CallGetISSNowData: []byte(testData),
			},
			expectLoc: &Location{
				Position: &ISSPosition{
					Lat:  "25.9866",
					Long: "-100.7594",
				},
			},
			expectErr: false,
		},
		{
			description: "Get location error",
			mapper: &MockMapper{
				CallGetISSNowErr: assert.AnError,
			},
			expectLoc: nil,
			expectErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			actualLoc, actualErr := GetLocation(tc.mapper)
			assert.Equal(t, tc.expectLoc, actualLoc)
			if tc.expectErr {
				assert.Error(t, actualErr)
			} else {
				assert.NoError(t, actualErr)
			}
		})
	}
}

func TestFormatLatLong(t *testing.T) {
	// formatLatLong()
}

func TestGetMapContext(t *testing.T) {
	// createMapContext()
}

func TestMapISSLocation(t *testing.T) {
	// MapISSLocation()
}
