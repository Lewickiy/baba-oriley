package player

// NoteEvent represents a single musical note event in a track.
// Each event contains information about when it starts, how long it lasts,
// its pitch, intensity, and which instrument should play it.
type NoteEvent struct {
	Start      float64 `json:"start"`      // Start time in seconds
	Duration   float64 `json:"duration"`   // Duration of the note in seconds
	Note       int     `json:"note"`       // MIDI note number
	Velocity   int     `json:"velocity"`   // Note intensity (volume)
	Instrument string  `json:"instrument"` // Name of the instrument
}
