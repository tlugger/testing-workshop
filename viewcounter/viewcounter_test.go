package viewcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddView(t *testing.T) {
	// first view should set 1
	video := &VideoMetadata{
		Title: "my first video",
		Views: 0,
	}
	expectCount := 1

	video.AddView()

	assert.Equal(t, expectCount, int(video.Views))

	// milliionth view should set 1000000
	video2 := &VideoMetadata{
		Title: "my best video",
		Views: 999999,
	}
	expectCount2 := 1000000

	video2.AddView()

	assert.Equal(t, expectCount2, int(video2.Views))
}

func TestAddViewTableTest(t *testing.T) {
	var testcases = []struct {
		description   string
		video         *VideoMetadata
		expectedCount int
	}{
		{
			description: "first view",
			video: &VideoMetadata{
				Title: "my first video",
				Views: 0,
			},
			expectedCount: 1,
		},
		{
			description: "millionth view",
			video: &VideoMetadata{
				Title: "my best video",
				Views: 999999,
			},
			expectedCount: 1000000,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			tc.video.AddView()
			assert.Equal(t, tc.expectedCount, int(tc.video.Views))
		})
	}
}
