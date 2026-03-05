package instruments

import "math"

// BassSynth generates a warm bass waveform with a light harmonic mix.
type BassSynth struct {
	id int
}

// Play generates bass samples for the provided note.
func (b *BassSynth) Play(note int, velocity int, duration float64, sampleRate int) []int16 {
	numSamples := int(duration * float64(sampleRate))
	buf := make([]int16, numSamples)
	freq := 55.0 * math.Pow(2, float64(note-33)/12.0)
	amp := float64(velocity)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		env := 1.0
		releaseStart := math.Max(0.0, duration-0.05)
		if t > releaseStart {
			env = math.Exp(-40.0 * (t - releaseStart))
		}

		fundamental := math.Sin(2 * math.Pi * freq * t)
		harmonic := 0.35 * math.Sin(2*math.Pi*freq*2*t)
		sample := (fundamental + harmonic) * env * amp
		buf[i] = int16(sample)
	}

	return buf
}
