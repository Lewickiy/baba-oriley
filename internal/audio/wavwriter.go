package audio

import (
	"encoding/binary"
	"os"
)

// WAVWriter represents an object for writing WAV files
type WAVWriter struct {
	f          *os.File // file descriptor
	numSamples int      // total number of samples
}

// writeOrFail writes data to a file and returns an error if it occurs
func writeOrFail(f *os.File, data any) error {
	switch v := data.(type) {
	case []byte:
		_, err := f.Write(v)
		return err
	default:
		return binary.Write(f, binary.LittleEndian, v)
	}
}

// NewWAV creates a WAV file with 16-bit PCM, 1 channel (mono), at the given sample rate.
// filename - output file name (without extension, stored in "out/" folder)
// sampleRate - sampling rate in Hz
// durationSec - duration of the file in seconds
func NewWAV(filename string, sampleRate int, durationSec int) (*WAVWriter, error) {
	pesPath := "out/" + filename + ".wav"
	f, err := os.Create(pesPath)
	if err != nil {
		return nil, err
	}

	numSamples := sampleRate * durationSec
	dataSize := numSamples * 2 // 16-bit mono
	fileSize := 36 + dataSize

	// RIFF chunk
	if err := writeOrFail(f, []byte("RIFF")); err != nil {
		return nil, err
	}
	if err := writeOrFail(f, uint32(fileSize)); err != nil {
		return nil, err
	}
	if err := writeOrFail(f, []byte("WAVE")); err != nil {
		return nil, err
	}

	// fmt subchunk
	if err := writeOrFail(f, []byte("fmt ")); err != nil {
		return nil, err
	}
	if err := writeOrFail(f, uint32(16)); err != nil { // Subchunk1Size
		return nil, err
	}
	if err := writeOrFail(f, uint16(1)); err != nil { // PCM
		return nil, err
	}
	if err := writeOrFail(f, uint16(1)); err != nil { // Mono
		return nil, err
	}
	if err := writeOrFail(f, uint32(sampleRate)); err != nil {
		return nil, err
	}
	byteRate := sampleRate * 2
	if err := writeOrFail(f, uint32(byteRate)); err != nil {
		return nil, err
	}
	if err := writeOrFail(f, uint16(2)); err != nil { // BlockAlign
		return nil, err
	}
	if err := writeOrFail(f, uint16(16)); err != nil { // BitsPerSample
		return nil, err
	}

	// data subchunk
	if err := writeOrFail(f, []byte("data")); err != nil {
		return nil, err
	}
	if err := writeOrFail(f, uint32(dataSize)); err != nil {
		return nil, err
	}

	return &WAVWriter{
		f:          f,
		numSamples: numSamples,
	}, nil
}

// WriteSamples writes an array of int16 samples to the WAV file
// The samples are written in little-endian byte order
func (w *WAVWriter) WriteSamples(samples []int16) error {
	buf := make([]byte, len(samples)*2)
	for i, s := range samples {
		buf[2*i] = byte(s)
		buf[2*i+1] = byte(s >> 8) // little-endian
	}
	return writeOrFail(w.f, buf)
}

// Close closes the WAV file
func (w *WAVWriter) Close() error {
	return w.f.Close()
}
