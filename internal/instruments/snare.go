package instruments

import "math"

// Snare is a lightweight synthesized snare drum (noise + short tone).
type Snare struct {
	id int
}

// Play generates a percussive snare-like sound.
func (s *Snare) Play(note int, velocity int, duration float64, sampleRate int) []int16 {
	numSamples := int(duration * float64(sampleRate))
	if numSamples < sampleRate/20 {
		numSamples = sampleRate / 20
	}

	buf := make([]int16, numSamples)
	amp := float64(velocity)
	toneFreq := 180.0 * math.Pow(2, float64(note-60)/24.0)

	// Deterministic pseudo-noise so rendering is reproducible across runs.
	var noiseState uint32 = uint32(note*131 + velocity*17 + 1)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		noiseEnv := math.Exp(-24.0 * t)
		toneEnv := math.Exp(-12.0 * t)

		noiseState = noiseState*1664525 + 1013904223
		noise := (float64(noiseState>>16)/32767.5 - 1.0) * noiseEnv

		tone := math.Sin(2*math.Pi*toneFreq*t) * toneEnv * 0.35
		sample := (noise*0.65 + tone) * amp
		buf[i] = int16(sample)
	}

	return buf
}
