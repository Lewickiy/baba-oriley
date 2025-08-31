package instruments

import "math"

// Kick — bass drum synthesizer
// Each instance has a unique ID for tracking purposes
type Kick struct {
	id int
}

// Play generates audio samples for the bass drum
// Parameters:
//   - note: MIDI note number, determines the pitch of the bass
//   - velocity: note amplitude (volume) from 0 to 127
//   - duration: duration of the note in seconds
//   - sampleRate: sampling rate in Hz
//
// Returns a slice of int16 samples representing a mono decaying sine wave
// that simulates a bass drum
func (k *Kick) Play(note int, velocity int, duration float64, sampleRate int) []int16 {
	numSamples := int(duration * float64(sampleRate))
	buf := make([]int16, numSamples)
	freq := 40.0 * math.Pow(2, float64(note)/12.0) // низкая бочка в зависимости от ноты

	for i := 0; i < numSamples; i++ {
		env := math.Exp(-8.0 * float64(i) / float64(numSamples))
		sample := math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)) * env
		buf[i] = int16(sample * float64(velocity))
	}

	return buf
}
