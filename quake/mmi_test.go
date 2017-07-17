package quake_test

import (
	"github.com/GeoNet/kit/quake"
	"math"
	"runtime"
	"strconv"
	"testing"
)

func TestMMI(t *testing.T) {
	testCases := []struct {
		id                    string
		magnitude, depth, mmi float64
	}{
		{id: "Christchurch 2011", depth: 5.0, magnitude: 6.3, mmi: 8.86},
		{id: "Gisbourne 2007", depth: 40.0, magnitude: 6.8, mmi: 8.19},
		{id: "Darfield 2010", depth: 11.0, magnitude: 7.1, mmi: 9.96},
		{id: "small deep event", depth: 150.0, magnitude: 1.5, mmi: -1.0},
		{id: "large deep event", depth: 150.0, magnitude: 6.5, mmi: 6.23},
		{id: "moderate shallow event", depth: 7.0, magnitude: 4.4, mmi: 6.41},
		{id: "zero values", mmi: -1.0},
	}

	for _, v := range testCases {
		mmi := quake.MMI(v.depth, v.magnitude)

		if math.Abs(v.mmi-mmi) > 0.005 {
			t.Errorf("%s expected mmi %f got %f", v.id, v.mmi, mmi)
		}
	}
}

func TestMMIDistance(t *testing.T) {
	testCases := []struct {
		id                               string
		distance, depth, magnitude, mmid float64
	}{
		{id: loc(), distance: 110.0, depth: 27.4, magnitude: 3.9, mmid: 2.65},
		{id: loc(), distance: 5.0, depth: 22.2, magnitude: 4.2, mmid: 5.27},
		{id: loc(), distance: 0.0, depth: 22.2, magnitude: 4.2, mmid: 5.27},
	}

	for _, v := range testCases {
		mmid := quake.MMIDistance(v.distance, v.depth, quake.MMI(v.depth, v.magnitude))

		if math.Abs(v.mmid-mmid) > 0.1 {
			t.Errorf("%s expected mmid %f got %f", v.id, v.mmid, mmid)
		}
	}
}

func loc() string {
	_, _, l, _ := runtime.Caller(1)
	return "L" + strconv.Itoa(l)
}
