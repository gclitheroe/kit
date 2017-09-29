package quake

import (
	"fmt"
	"time"
)

// Quake for earthquakes.  Members (particularly strings) should follow SC3ML / QuakeML typing.
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

// Status returns a simplified status.
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

// Quality returns a simplified quality.
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
func (q *Quake) AlertQuality() (bool, string) {
	switch {
	case q.Status() == "deleted":
		return false, fmt.Sprintf("%s status deleted not suitable for alerting.", q.PublicID)
	case q.Status() == "duplicate":
		return false, fmt.Sprintf("%s status duplicate not suitable for alerting.", q.PublicID)
	case q.Status() == "automatic" && (q.UsedPhaseCount < 20 || q.MagnitudeStationCount < 10):
		return false, fmt.Sprintf("%s unreviewed with %d phases and %d magnitudes not suitable for alerting.", q.PublicID, q.UsedPhaseCount, q.MagnitudeStationCount)
	case q.Status() == "automatic" && !(q.Depth >= 0.1 && q.AzimuthalGap <= 320.0 && q.MinimumDistance <= 2.5):
		return false, fmt.Sprintf("%s automatic with poor location criteria", q.PublicID)
	default:
		return true, ""
	}
}

// manual returns true if the quake has been manually reviewed in some way.
func (q *Quake) Manual() bool {
	switch {
	case q.Type == "not existing":
		return true
	case q.Type == "duplicate":
		return true
	case q.EvaluationMode == "manual":
		return true
	case q.EvaluationStatus == "confirmed":
		return true
	default:
		return false
	}
}
