package quake_test

import (
	"github.com/GeoNet/kit/quake"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	testCases := []struct {
		id     string
		q      quake.Quake
		status string
	}{
		{id: loc(), q: quake.Quake{}, status: "automatic"},
		{id: loc(), q: quake.Quake{Type: "not existing"}, status: "deleted"},
		{id: loc(), q: quake.Quake{Type: "duplicate"}, status: "duplicate"},
		{id: loc(), q: quake.Quake{EvaluationMode: "manual"}, status: "reviewed"},
		{id: loc(), q: quake.Quake{EvaluationStatus: "confirmed"}, status: "reviewed"},
	}

	for _, v := range testCases {
		if v.status != v.q.Status() {
			t.Errorf("%s expected status %s got %s", v.id, v.status, v.q.Status())
		}
	}
}

func TestQuality(t *testing.T) {
	testCases := []struct {
		id      string
		q       quake.Quake
		quality string
	}{
		{id: loc(), quality: "caution", q: quake.Quake{}},
		{id: loc(), quality: "best", q: quake.Quake{EvaluationMode: "manual"}},
		{id: loc(), quality: "deleted", q: quake.Quake{Type: "not existing"}},
		{id: loc(), quality: "caution", q: quake.Quake{UsedPhaseCount: 19, MagnitudeStationCount: 9}},
		{id: loc(), quality: "good", q: quake.Quake{UsedPhaseCount: 20, MagnitudeStationCount: 10}},
		{id: loc(), quality: "deleted", q: quake.Quake{Type: "not existing", UsedPhaseCount: 20, MagnitudeStationCount: 10}},
	}
	for _, v := range testCases {
		if v.quality != v.q.Quality() {
			t.Errorf("%s expected quality %s got %s", v.id, v.quality, v.q.Quality())
		}
	}
}

func TestAlertQuality(t *testing.T) {
	testCases := []struct {
		id    string
		q     quake.Quake
		alert bool
	}{
		{id: loc(), alert: false, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 8, MagnitudeStationCount: 8}},
		{id: loc(), alert: false, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 8}},
		{id: loc(), alert: true, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12}},
		{id: loc(), alert: false, q: quake.Quake{Time: time.Now().UTC(), Type: "not existing", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12}},
		{id: loc(), alert: false, q: quake.Quake{Time: time.Now().UTC(), Type: "duplicate", EvaluationMode: "automatic", UsedPhaseCount: 22, MagnitudeStationCount: 12}},
		{id: loc(), alert: true, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 8, MagnitudeStationCount: 8}},
		{id: loc(), alert: true, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 8, MagnitudeStationCount: 8}},
		{id: loc(), alert: true, q: quake.Quake{Time: time.Now().UTC(), Type: "earthquake", EvaluationMode: "", EvaluationStatus: "confirmed", UsedPhaseCount: 28, MagnitudeStationCount: 28}},
		{id: loc(), alert: false, q: quake.Quake{Time: time.Now().UTC().Add(time.Minute * -61), Type: "earthquake", EvaluationMode: "manual", UsedPhaseCount: 8, MagnitudeStationCount: 8}},
	}

	for _, v := range testCases {
		if v.q.AlertQuality() != v.alert {
			t.Errorf("%s got expected alert quality %t got %t", v.id, v.alert, v.q.AlertQuality())
		}
	}
}

func TestPublish(t *testing.T) {
	testCases := []struct {
		id      string
		q       quake.Quake
		publish bool
	}{
		{id: loc(), publish: false, q: quake.Quake{}},
		{id: loc(), publish: false, q: quake.Quake{Site: "backup"}},
		{id: loc(), publish: false, q: quake.Quake{Site: "backup", EvaluationStatus: "automatic"}},
		{id: loc(), publish: false, q: quake.Quake{Site: "backup", EvaluationStatus: "automatic"}},
		{id: loc(), publish: false, q: quake.Quake{Site: "primary", EvaluationStatus: "automatic"}},
		{id: loc(), publish: true, q: quake.Quake{Site: "primary", EvaluationStatus: "confirmed"}},
		{id: loc(), publish: true, q: quake.Quake{Site: "backup", EvaluationStatus: "confirmed"}},
		{id: loc(), publish: false, q: quake.Quake{Site: "primary", EvaluationStatus: "automatic"}},
		{id: loc(), publish: false, q: quake.Quake{Site: "primary", EvaluationStatus: "automatic", Depth: 0.01, AzimuthalGap: 321.0, MinimumDistance: 3.0}},
		{id: loc(), publish: false, q: quake.Quake{Site: "primary", EvaluationStatus: "automatic", Depth: 0.01, AzimuthalGap: 321.0, MinimumDistance: 3.0}},
		{id: loc(), publish: false, q: quake.Quake{Site: "backup", EvaluationStatus: "automatic", Depth: 0.2, AzimuthalGap: 319.0, MinimumDistance: 2.4}},
		{id: loc(), publish: true, q: quake.Quake{Site: "primary", EvaluationStatus: "automatic", Depth: 0.2, AzimuthalGap: 319.0, MinimumDistance: 2.4}},
		{id: loc(), publish: true, q: quake.Quake{Site: "backup", EvaluationStatus: "confirmed", Depth: 0.2, AzimuthalGap: 319.0, MinimumDistance: 2.4}},
	}

	for _, v := range testCases {
		if v.publish != v.q.Publish() {
			t.Errorf("%s Publish expected %t got %t", v.id, v.publish, v.q.Publish())
		}
	}
}
