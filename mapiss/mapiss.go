package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"net/http"
	"strconv"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
)

type Location struct {
	Position *ISSPosition `json:"iss_position"`
}

type ISSPosition struct {
	Lat  string `json:"latitude"`
	Long string `json:"longitude"`
}

type mapper struct {
	ImageName string
}

type Mapper interface {
	CallGetISSNow() ([]byte, error)
	RenderMap(ctx *sm.Context) error
}

func NewMapper() Mapper {
	return &mapper{
		ImageName: "space_station_locaiton.png",
	}
}

// CallGetISSNow calls the iss-now API for current position of the
// International Space Station
func (m *mapper) CallGetISSNow() ([]byte, error) {
	response, err := http.Get("http://api.open-notify.org/iss-now.json")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetLocation parses json as bytes into a Location struct
func GetLocation(m Mapper) (*Location, error) {
	data, err := m.CallGetISSNow()
	if err != nil {
		return nil, err
	}

	var loc Location
	if err = json.Unmarshal([]byte(data), &loc); err != nil {
		return nil, err
	}

	return &loc, err

}

// FormatLatLong converts latitude/longitude string values into float64
func FormatLatLong(latitude, longitude string) (float64, float64, error) {
	var lat, long float64
	var err error

	lat, err = strconv.ParseFloat(latitude, 64)
	if err != nil {
		return lat, long, err
	}

	long, err = strconv.ParseFloat(longitude, 64)
	if err != nil {
		return lat, long, err
	}

	return lat, long, nil

}

// CreateMapContext generates the static maps Context object used to render a marked map
func CreateMapContext(width, height, zoom int, latitude, longitude string) (*sm.Context, error) {
	lat, long, err := FormatLatLong(latitude, longitude)
	if err != nil {
		return nil, err
	}

	ctx := sm.NewContext()
	ctx.SetSize(width, height)
	ctx.AddObject(
		sm.NewMarker(
			s2.LatLngFromDegrees(lat, long),
			color.RGBA{0xff, 0, 0, 0xff},
			16.0,
		),
	)
	ctx.SetZoom(zoom)

	return ctx, nil

}

// RenderMap renders a static marked map as a PNG
func (m *mapper) RenderMap(ctx *sm.Context) error {
	img, err := ctx.Render()
	if err != nil {
		return err
	}

	if err := gg.SavePNG(m.ImageName, img); err != nil {
		return err
	}

	return nil
}

// MapISSLocation generates a static map with the current location of the
// International Space Station marked
func MapISSLocation(m Mapper) error {
	loc, err := GetLocation(m)
	if err != nil {
		return err
	}

	imgWidth := 910
	imgHeight := 512
	imgZoom := 5

	lat := loc.Position.Lat
	long := loc.Position.Long

	ctx, err := CreateMapContext(imgWidth, imgHeight, imgZoom, lat, long)
	if err != nil {
		return err
	}

	err = m.RenderMap(ctx)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	m := NewMapper()
	err := MapISSLocation(m)
	if err != nil {
		panic(err)
	}
}
