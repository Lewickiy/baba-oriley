package instruments

import "math"

// ArpSynth is a simple synthesizer that generates a sine wave for a given MIDI note
// Each instance has a unique ID for tracking purposes
type ArpSynth struct {
	id int
}

// Play generates audio samples for a given note.
// Parameters:
//   - note: MIDI note number (integer, e.g., 69 for A4).
//   - velocity: amplitude of the note (0-127, affects volume).
//   - duration: duration of the note in seconds.
//   - sampleRate: samples per second (e.g., 44100).
//
// Returns a slice of int16 samples representing a monophonic sine wave at the
// corresponding frequency.
func (s *ArpSynth) Play(note int, velocity int, duration float64, sampleRate int) []int16 {
	numSamples := int(duration * float64(sampleRate))
	buf := make([]int16, numSamples)
	freq := 440.0 * math.Pow(2, float64(note-69)/12.0)

	phase := 0.0
	for i := 0; i < numSamples; i++ {
		sample := math.Sin(2 * math.Pi * freq * phase / float64(sampleRate))
		buf[i] = int16(sample * float64(velocity))
		phase++
	}

	return buf
}
