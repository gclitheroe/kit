package quake_test

import (
	"github.com/GeoNet/kit/quake"
	"testing"
	"time"
)

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
