package quake

import (
	"fmt"
	"math"
	"time"
)

var alertAge = time.Duration(-60) * time.Minute

type Quake struct {
	PublicID              string
	Type                  string
	AgencyID              string
	ModificationTime      time.Time
	Time                  time.Time
	Latitude              float64
	Longitude             float64
	Depth                 float64
	DepthType             string
	MethodID              string
	EarthModelID          string
	EvaluationMode        string
	EvaluationStatus      string
	UsedPhaseCount        int
	UsedStationCount      int
	StandardError         float64
	AzimuthalGap          float64
	MinimumDistance       float64
	Magnitude             float64
	MagnitudeUncertainty  float64
	MagnitudeType         string
	MagnitudeStationCount int
	Site                  string
}

// Status returns the public status for the Quake referred to by q.
func (q *Quake) Status() string {
	switch {
	case q.Type == "not existing":
		return "deleted"
	case q.Type == "duplicate":
		return "duplicate"
	case q.EvaluationMode == "manual":
		return "reviewed"
	case q.EvaluationStatus == "confirmed":
		return "reviewed"
	default:
		return "automatic"
	}
}

func (q *Quake) Quality() string {
	status := q.Status()

	switch {
	case status == "reviewed":
		return "best"
	case status == "deleted":
		return "deleted"
	case q.UsedPhaseCount >= 20 && q.MagnitudeStationCount >= 10:
		return "good"
	default:
		return "caution"
	}
}

// Returns true of the Quake is of high enough quality to consider for alerting.
//  false if not.
func (q *Quake) AlertQuality() bool {
	switch {
	case q.Status() == "deleted":
		return false
	case q.Status() == "duplicate":
		return false
	case q.Status() == "automatic" && (q.UsedPhaseCount < 20 || q.MagnitudeStationCount < 10):
		return false
	case q.Time.Before(time.Now().UTC().Add(alertAge)):
		return false
	}

	return true
}

func Distance(km float64) string {
	s := "Within 5 km of"

	d := math.Floor(km / 5.0)
	if d > 0 {
		s = fmt.Sprintf("%.f km", d*5)
	}
	return s
}

// Publish returns true if the quake is suitable for publishing.
// site is either 'primary' or 'backup'.
func (q *Quake) Publish() bool {
	switch q.Site {
	case "primary", "":
		if q.Status() == "automatic" && !(q.Depth >= 0.1 && q.AzimuthalGap <= 320.0 && q.MinimumDistance <= 2.5) {
			return false
		}
	case "backup":
		if q.Status() == "automatic" {
			return false
		}
	}

	return true
}
