package instruments

import "math"

// PolySynth is a simple polyphonic-style lead pad sound.
type PolySynth struct {
	id int
}

// Play generates a bright, slightly detuned waveform.
func (p *PolySynth) Play(note int, velocity int, duration float64, sampleRate int) []int16 {
	numSamples := int(duration * float64(sampleRate))
	buf := make([]int16, numSamples)
	baseFreq := 440.0 * math.Pow(2, float64(note-69)/12.0)
	amp := float64(velocity)

	attack := math.Min(duration*0.2, 0.03)
	release := math.Min(duration*0.25, 0.08)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)

		env := 1.0
		if attack > 0 && t < attack {
			env = t / attack
		}
		if release > 0 && t > duration-release {
			env *= math.Max(0, (duration-t)/release)
		}

		osc1 := math.Sin(2 * math.Pi * baseFreq * t)
		osc2 := math.Sin(2 * math.Pi * baseFreq * 1.005 * t)
		osc3 := math.Sin(2 * math.Pi * baseFreq * 0.997 * t)
		sample := (osc1 + 0.6*osc2 + 0.6*osc3) * 0.5 * env * amp
		buf[i] = int16(sample)
	}

	return buf
}
