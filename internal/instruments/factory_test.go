package instruments

import "testing"

func TestGetInstrumentByPrefix(t *testing.T) {
	InstrumentInstances = map[string]Instrument{}
	nextInstrumentID = 0

	cases := []struct {
		name string
		want any
	}{
		{name: "kick1", want: &Kick{}},
		{name: "snare1", want: &Snare{}},
		{name: "bass1", want: &BassSynth{}},
		{name: "synth1", want: &PolySynth{}},
		{name: "arp1", want: &ArpSynth{}},
		{name: "unknown", want: &ArpSynth{}},
	}

	for _, tc := range cases {
		got := GetInstrument(tc.name)
		switch tc.want.(type) {
		case *Kick:
			if _, ok := got.(*Kick); !ok {
				t.Fatalf("%s: expected *Kick, got %T", tc.name, got)
			}
		case *Snare:
			if _, ok := got.(*Snare); !ok {
				t.Fatalf("%s: expected *Snare, got %T", tc.name, got)
			}
		case *BassSynth:
			if _, ok := got.(*BassSynth); !ok {
				t.Fatalf("%s: expected *BassSynth, got %T", tc.name, got)
			}
		case *PolySynth:
			if _, ok := got.(*PolySynth); !ok {
				t.Fatalf("%s: expected *PolySynth, got %T", tc.name, got)
			}
		case *ArpSynth:
			if _, ok := got.(*ArpSynth); !ok {
				t.Fatalf("%s: expected *ArpSynth, got %T", tc.name, got)
			}
		}
	}
}
