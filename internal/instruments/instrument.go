package instruments

// Instrument defines a musical instrument that can generate audio samples.
// The Play method receives a MIDI note number, velocity, duration in seconds,
// and sample rate, and returns a slice of int16 PCM samples.
type Instrument interface {
	Play(note int, velocity int, duration float64, sampleRate int) []int16
}
