package quake

import (
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestManual(t *testing.T) {
	in := []struct {
		id     string
		q      Quake
		manual bool
	}{
		{id: loc(), q: Quake{}, manual: false},
		{id: loc(), q: Quake{Type: "not existing"}, manual: true},
		{id: loc(), q: Quake{Type: "duplicate"}, manual: true},
		{id: loc(), q: Quake{EvaluationMode: "manual"}, manual: true},
		{id: loc(), q: Quake{EvaluationStatus: "confirmed"}, manual: true},
	}

	for _, v := range in {
		if v.q.Manual() != v.manual {
			t.Errorf("%s expected manual %t got %t", v.id, v.manual, v.q.Manual())
		}
	}
}

func TestStatus(t *testing.T) {
	in := []struct {
		id     string
		q      Quake
		status string
	}{
		{id: loc(), q: Quake{}, status: "automatic"},
		{id: loc(), q: Quake{Type: "not existing"}, status: "deleted"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationMode: "manual"}, status: "deleted"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationStatus: "confirmed"}, status: "deleted"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationMode: "manual", EvaluationStatus: "confirmed"}, status: "deleted"},
		{id: loc(), q: Quake{Type: "duplicate"}, status: "duplicate"},
		{id: loc(), q: Quake{Type: "duplicate", EvaluationMode: "manual"}, status: "duplicate"},
		{id: loc(), q: Quake{Type: "duplicate", EvaluationStatus: "confirmed"}, status: "duplicate"},
		{id: loc(), q: Quake{Type: "duplicate", EvaluationMode: "manual", EvaluationStatus: "confirmed"}, status: "duplicate"},
		{id: loc(), q: Quake{EvaluationMode: "manual"}, status: "reviewed"},
		{id: loc(), q: Quake{EvaluationStatus: "confirmed"}, status: "reviewed"},
	}

	for _, v := range in {
		if v.q.Status() != v.status {
			t.Errorf("%s expected status %s got %s", v.id, v.status, v.q.Status())
		}
	}
}

func TestQuality(t *testing.T) {
	in := []struct {
		id      string
		q       Quake
		quality string
	}{
		{id: loc(), q: Quake{}, quality: "caution"},
		{id: loc(), q: Quake{EvaluationMode: "manual"}, quality: "best"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationMode: "manual"}, quality: "deleted"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationStatus: "confirmed"}, quality: "deleted"},
		{id: loc(), q: Quake{Type: "not existing", EvaluationMode: "manual", EvaluationStatus: "confirmed"}, quality: "deleted"},
		{id: loc(), q: Quake{UsedPhaseCount: 10, MagnitudeStationCount: 4}, quality: "caution"},
		{id: loc(), q: Quake{UsedPhaseCount: 19, MagnitudeStationCount: 10}, quality: "caution"},
		{id: loc(), q: Quake{UsedPhaseCount: 20, MagnitudeStationCount: 9}, quality: "caution"},
		{id: loc(), q: Quake{UsedPhaseCount: 20, MagnitudeStationCount: 10}, quality: "good"},
	}

	for _, v := range in {
		if v.q.Quality() != v.quality {
			t.Errorf("%s expected quality %s got %s", v.id, v.quality, v.q.Quality())
		}
	}
}

func TestAlertQuality(t *testing.T) {
	in := []struct {
		id    string
		q     Quake
		alert bool
	}{
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 8, MagnitudeStationCount: 8, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 8, MagnitudeStationCount: 8, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 28, MagnitudeStationCount: 28, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 8, MagnitudeStationCount: 8, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 8, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "not existing", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "duplicate", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: false, // shallow automatic shouldn't alert
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: false, // high azimuthal gap automatic shouldn't alert
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 330, MinimumDistance: 1.0}},
		{id: loc(), alert: false, // far outside network automatic shouldn't alert
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 3.0}},
		//	 combinations of bad quality parameters shouldn't alert
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 320, MinimumDistance: 3.0}},
		{id: loc(), alert: false,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 320, MinimumDistance: 3.0}},
		//	bad quality parameters but confirmed.  Should all alert.
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 330, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 3.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 320, MinimumDistance: 3.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 320, MinimumDistance: 3.0}},
		//	bad quality parameters but manual.  Should all alert.
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 200, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 330, MinimumDistance: 1.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 200, MinimumDistance: 3.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 10.0, AzimuthalGap: 320, MinimumDistance: 3.0}},
		{id: loc(), alert: true,
			q: Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 22, MagnitudeStationCount: 12, Depth: 0.01, AzimuthalGap: 320, MinimumDistance: 3.0}},
	}

	for _, v := range in {
		alert, _ := v.q.AlertQuality()

		if alert != v.alert {
			t.Errorf("%s incorrect alert quality got %t expected %t", v.id, alert, v.alert)
		}
	}
}

func loc() string {
	_, _, l, _ := runtime.Caller(1)
	return "L" + strconv.Itoa(l)
}
