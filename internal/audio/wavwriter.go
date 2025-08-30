package audio

import (
	"encoding/binary"
	"os"
)

// WAVWriter представляет объект для записи WAV файлов.
// Хранит дескриптор файла и количество сэмплов
type WAVWriter struct {
	f          *os.File
	numSamples int
}

// CreateSilenceWav создаёт WAV файл с тишиной указанной длительности.
// filename - имя файла.
// sampleRate - частота дискретизации (например, 44100).
// durationSec - длительность в секундах.
// Возвращает ошибку, если не удалось создать файл.
func CreateSilenceWav(filename string, sampleRate int, durationSec int) error {
	writer, err := newWAV(filename, sampleRate, durationSec)
	if err != nil {
		return err
	}

	err = writer.writeSilence()
	if err != nil {
		return err
	}
	return writer.close()
}

// newWAV создаёт файл WAV с 16-bit PCM, 1 канал, sampleRate Гц
func newWAV(filename string, sampleRate int, durationSec int) (*WAVWriter, error) {
	pesPath := "out/" + filename + ".wav"
	f, err := os.Create(pesPath)
	if err != nil {
		return nil, err
	}

	numSamples := sampleRate * durationSec

	// заголовок WAV (RIFF)
	// 44 байта стандартного заголовка
	dataSize := numSamples * 2 // 16-bit mono
	fileSize := 36 + dataSize

	// RIFF chunk
	f.Write([]byte("RIFF"))
	binary.Write(f, binary.LittleEndian, uint32(fileSize))
	f.Write([]byte("WAVE"))

	// fmt subchunk
	f.Write([]byte("fmt "))
	binary.Write(f, binary.LittleEndian, uint32(16)) // Subchunk1Size
	binary.Write(f, binary.LittleEndian, uint16(1))  // PCM
	binary.Write(f, binary.LittleEndian, uint16(1))  // Mono
	binary.Write(f, binary.LittleEndian, uint32(sampleRate))
	byteRate := sampleRate * 2
	binary.Write(f, binary.LittleEndian, uint32(byteRate))
	binary.Write(f, binary.LittleEndian, uint16(2))  // BlockAlign
	binary.Write(f, binary.LittleEndian, uint16(16)) // BitsPerSample

	// data subchunk
	f.Write([]byte("data"))
	binary.Write(f, binary.LittleEndian, uint32(dataSize))

	return &WAVWriter{
		f:          f,
		numSamples: numSamples,
	}, nil
}

// writeSilence записывает тишину
func (w *WAVWriter) writeSilence() error {
	buf := make([]byte, w.numSamples*2) // 2 байта на сэмпл
	_, err := w.f.Write(buf)
	return err
}

// close закрывает файл
func (w *WAVWriter) close() error {
	return w.f.Close()
}
