package quake

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"time"
)

// types etc for unmarshaling SC3ML
// these are private.  There is the kit/sc3ml pkg with public types.

var sc3ml07 = []byte(`<seiscomp xmlns="http://geofon.gfz-potsdam.de/ns/seiscomp3-schema/0.7" version="0.7">`)
var sc3ml08 = []byte(`<seiscomp xmlns="http://geofon.gfz-potsdam.de/ns/seiscomp3-schema/0.8" version="0.8">`)
var sc3ml09 = []byte(`<seiscomp xmlns="http://geofon.gfz-potsdam.de/ns/seiscomp3-schema/0.9" version="0.9">`)

type seiscomp struct {
	EventParameters eventParameters `xml:"EventParameters"`
}

type eventParameters struct {
	Events     []event     `xml:"event"`
	Picks      []pick      `xml:"pick"`
	Amplitudes []amplitude `xml:"amplitude"`
	Origins    []origin    `xml:"origin"`
}

type event struct {
	PublicID             string `xml:"publicID,attr"`
	PreferredOriginID    string `xml:"preferredOriginID"`
	PreferredMagnitudeID string `xml:"preferredMagnitudeID"`
	Type                 string `xml:"type"`
	PreferredOrigin      origin
	PreferredMagnitude   magnitude
	ModificationTime     time.Time    `xml:"-"` // most recent modification time for all objects in the event.  Not in the XML.
	CreationInfo         creationInfo `xml:"creationInfo"`
}

type creationInfo struct {
	AgencyID         string    `xml:"agencyID"`
	CreationTime     time.Time `xml:"creationTime"`
	ModificationTime time.Time `xml:"modificationTime"`
}

type origin struct {
	PublicID          string             `xml:"publicID,attr"`
	Time              timeValue          `xml:"time"`
	Latitude          realQuantity       `xml:"latitude"`
	Longitude         realQuantity       `xml:"longitude"`
	Depth             realQuantity       `xml:"depth"`
	DepthType         string             `xml:"depthType"`
	MethodID          string             `xml:"methodID"`
	EarthModelID      string             `xml:"earthModelID"`
	Quality           quality            `xml:"quality"`
	EvaluationMode    string             `xml:"evaluationMode"`
	EvaluationStatus  string             `xml:"evaluationStatus"`
	Arrivals          []arrival          `xml:"arrival"`
	StationMagnitudes []stationMagnitude `xml:"stationMagnitude"`
	Magnitudes        []magnitude        `xml:"magnitude"`
}

type quality struct {
	UsedPhaseCount   int64   `xml:"usedPhaseCount"`
	UsedStationCount int64   `xml:"usedStationCount"`
	StandardError    float64 `xml:"standardError"`
	AzimuthalGap     float64 `xml:"azimuthalGap"`
	MinimumDistance  float64 `xml:"minimumDistance"`
}

type arrival struct {
	PickID       string  `xml:"pickID"`
	Phase        string  `xml:"phase"`
	Azimuth      float64 `xml:"azimuth"`
	Distance     float64 `xml:"distance"`
	TimeResidual float64 `xml:"timeResidual"`
	Weight       float64 `xml:"weight"`
	Pick         pick
}

type pick struct {
	PublicID         string     `xml:"publicID,attr"`
	Time             timeValue  `xml:"time"`
	WaveformID       waveformID `xml:"waveformID"`
	EvaluationMode   string     `xml:"evaluationMode"`
	EvaluationStatus string     `xml:"evaluationStatus"`
}

type waveformID struct {
	NetworkCode  string `xml:"networkCode,attr"`
	StationCode  string `xml:"stationCode,attr"`
	LocationCode string `xml:"locationCode,attr"`
	ChannelCode  string `xml:"channelCode,attr"`
}

type realQuantity struct {
	Value       float64 `xml:"value"`
	Uncertainty float64 `xml:"uncertainty"`
}

type timeValue struct {
	Value time.Time `xml:"value"`
}

type magnitude struct {
	PublicID                      string                         `xml:"publicID,attr"`
	Magnitude                     realQuantity                   `xml:"magnitude"`
	Type                          string                         `xml:"type"`
	MethodID                      string                         `xml:"methodID"`
	StationCount                  int64                          `xml:"stationCount"`
	StationMagnitudeContributions []stationMagnitudeContribution `xml:"stationMagnitudeContribution"`
}

type stationMagnitudeContribution struct {
	StationMagnitudeID string  `xml:"stationMagnitudeID"`
	Weight             float64 `xml:"weight"`
	Residual           float64 `xml:"residual"`
	StationMagnitude   stationMagnitude
}

type stationMagnitude struct {
	PublicID    string       `xml:"publicID,attr"`
	Magnitude   realQuantity `xml:"magnitude"`
	Type        string       `xml:"type"`
	AmplitudeID string       `xml:"amplitudeID"`
	WaveformID  waveformID   `xml:"waveformID"`
	Amplitude   amplitude
}

type amplitude struct {
	PublicID  string       `xml:"publicID,attr"`
	Amplitude realQuantity `xml:"amplitude"`
	PickID    string       `xml:"pickID"`
	Azimuth   float64      // not in the SC3ML - will be mapped from arrival using PickID
	Distance  float64      // not in the SC3ML - will be mapped from arrival using PickID
}

// unmarshal unmarshals the SeisComPML in r and initialises all
// the objects referenced by ID in the SeisComPML e.g., PreferredOrigin,
// PreferredMagnitude etc.
//
// Supports SC3ML versions are 0.7, 0.8, 0.9
// Any other versions will result in a error.
// It is an error if there is more than 1 Event in the SeisComPML.
func unmarshal(r io.Reader, s *seiscomp) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	switch {
	case bytes.Contains(b, sc3ml07):
	case bytes.Contains(b, sc3ml08):
	case bytes.Contains(b, sc3ml09):
	default:
		return errors.New("unsupported SC3ML version.")
	}

	if err := xml.Unmarshal(b, s); err != nil {
		return err
	}

	var picks = make(map[string]pick)
	for k, v := range s.EventParameters.Picks {
		picks[v.PublicID] = s.EventParameters.Picks[k]
	}

	var arrivals = make(map[string]arrival)
	for i := range s.EventParameters.Origins {
		for _, v := range s.EventParameters.Origins[i].Arrivals {
			arrivals[v.PickID] = v
		}
	}

	var amplitudes = make(map[string]amplitude)
	for k, v := range s.EventParameters.Amplitudes {
		a := s.EventParameters.Amplitudes[k]

		// add distance and azimuth from the arrival with the matching PickID.
		pk := arrivals[v.PickID]

		a.Distance = pk.Distance
		a.Azimuth = pk.Azimuth

		amplitudes[v.PublicID] = a
	}

	for i := range s.EventParameters.Origins {
		for k, v := range s.EventParameters.Origins[i].Arrivals {
			s.EventParameters.Origins[i].Arrivals[k].Pick = picks[v.PickID]
		}

		var stationMagnitudes = make(map[string]stationMagnitude)

		for k, v := range s.EventParameters.Origins[i].StationMagnitudes {
			s.EventParameters.Origins[i].StationMagnitudes[k].Amplitude = amplitudes[v.AmplitudeID]
			stationMagnitudes[v.PublicID] = s.EventParameters.Origins[i].StationMagnitudes[k]
		}

		for j := range s.EventParameters.Origins[i].Magnitudes {
			for k, v := range s.EventParameters.Origins[i].Magnitudes[j].StationMagnitudeContributions {
				s.EventParameters.Origins[i].Magnitudes[j].StationMagnitudeContributions[k].StationMagnitude = stationMagnitudes[v.StationMagnitudeID]
			}
		}
	}

	// set the preferred origin.
	// set the preferred mag which can come from any origin
	for i := range s.EventParameters.Events {
		for k, v := range s.EventParameters.Origins {
			if v.PublicID == s.EventParameters.Events[i].PreferredOriginID {
				s.EventParameters.Events[i].PreferredOrigin = s.EventParameters.Origins[k]
			}
			for _, mag := range v.Magnitudes {
				if mag.PublicID == s.EventParameters.Events[i].PreferredMagnitudeID {
					s.EventParameters.Events[i].PreferredMagnitude = mag
				}
			}
		}
	}

	// set the most recent modified time as long as there is only one event.
	// XML token parse the entire SC3ML, looking for creationInfo.
	// assumes all objects in the SC3ML are associated, or have been associated,
	// with the event somehow.
	if len(s.EventParameters.Events) != 1 {
		return errors.New("found more than 1 Event")
	}

	var by bytes.Buffer
	by.Write(b)
	d := xml.NewDecoder(&by)
	var tk xml.Token

	for {
		// Read tokens from the XML document in a stream.
		tk, err = d.Token()
		if tk == nil {
			break
		}
		if err != nil {
			return err
		}

		switch se := tk.(type) {
		case xml.StartElement:
			if se.Name.Local == "creationInfo" {
				var c creationInfo
				err = d.DecodeElement(&c, &se)
				if err != nil {
					return err
				}
				if c.ModificationTime.After(s.EventParameters.Events[0].ModificationTime) {
					s.EventParameters.Events[0].ModificationTime = c.ModificationTime
				}
				if c.CreationTime.After(s.EventParameters.Events[0].ModificationTime) {
					s.EventParameters.Events[0].ModificationTime = c.CreationTime
				}
			}
		}
	}

	return nil
}
