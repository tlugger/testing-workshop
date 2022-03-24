package main

import (
	"os"
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
	var testcases = []struct {
		description string
		inputLat    string
		inputLong   string
		expectLat   float64
		expectLong  float64
		expectErr   bool
	}{
		{
			description: "lat/long success",
			inputLat:    "12.3456",
			inputLong:   "-12.3456",
			expectLat:   float64(12.3456),
			expectLong:  float64(-12.3456),
			expectErr:   false,
		},
		{
			description: "errorr lat",
			inputLat:    "not a float",
			inputLong:   "-12.3456",
			expectErr:   true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			lat, long, err := FormatLatLong(tc.inputLat, tc.inputLong)
			assert.Equal(t, tc.expectLat, lat)
			assert.Equal(t, tc.expectLong, long)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMapISSLocationIntegration(t *testing.T) {
	var testcases = []struct {
		description  string
		expectErr    bool
		validateFunc func(t *testing.T)
		cleanupFunc  func(t *testing.T)
	}{
		{
			description: "success",
			expectErr:   false,
			validateFunc: func(t *testing.T) {
				_, err := os.Stat("space_station_locaiton.png")
				assert.NoError(t, err)
			},
			cleanupFunc: func(t *testing.T) {
				err := os.Remove("space_station_locaiton.png")
				assert.NoError(t, err)
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			m := NewMapper()
			err := MapISSLocation(m)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			tc.validateFunc(t)
			tc.cleanupFunc(t)
		})
	}
}
