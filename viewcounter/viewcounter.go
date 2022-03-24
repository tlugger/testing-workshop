package viewcounter

type VideoMetadata struct {
	Title string
	Views int32
}

func (v *VideoMetadata) AddView() {
	v.Views++
}
