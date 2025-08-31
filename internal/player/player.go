package player

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"baba-oriley/internal/audio"
	"baba-oriley/internal/instruments"
)

// LoadEvents reads a JSON file containing note events
// jsonFilename — the filename without extension
// Returns a slice of NoteEvent and an error if the file does not exist or the JSON is invalid
func LoadEvents(jsonFilename string) ([]NoteEvent, error) {
	data, err := os.ReadFile("assets/" + jsonFilename + ".json")
	if err != nil {
		return nil, err
	}

	var events []NoteEvent
	if err := json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// PlayEvents plays events from a slice of NoteEvent
// Each instrument is played in a separate goroutine for parallel generation
// events — slice of events
// sampleRate — WAV sample rate
// outputFile — name of the output WAV file
// speedCoefficient — playback speed multiplier (1.0 = normal speed, 0.7 = slow speed, 1.7 = fast speed)
func PlayEvents(events []NoteEvent, sampleRate int, outputFile string, speedCoefficient float64) error {
	// группируем события по инструменту
	instrEvents := map[string][]NoteEvent{}
	for _, e := range events {
		instrEvents[e.Instrument] = append(instrEvents[e.Instrument], e)
	}

	// вычисляем общую длительность трека
	durationSec := calcDuration(events, speedCoefficient)
	totalSamples := sampleRate * durationSec
	finalBuf := make([]int16, totalSamples)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for name, evts := range instrEvents {
		wg.Add(1)
		go func(name string, evts []NoteEvent) {
			defer wg.Done()

			inst := instruments.GetInstrument(name)

			log.Printf("Goroutine %d: instrument %s, object address %p\n", getGID(), name, inst)

			localBuf := make([]int16, totalSamples)
			for _, e := range evts {
				start := e.Start / speedCoefficient
				dur := e.Duration / speedCoefficient
				samples := inst.Play(e.Note, e.Velocity, dur, sampleRate)
				startSample := int(start * float64(sampleRate))

				for i := 0; i < len(samples) && startSample+i < len(localBuf); i++ {
					localBuf[startSample+i] += samples[i]
				}
			}

			mu.Lock()
			for i := range finalBuf {
				finalBuf[i] += localBuf[i]
			}
			mu.Unlock()
		}(name, evts)
	}

	wg.Wait()

	writer, err := audio.NewWAV(outputFile, sampleRate, durationSec)
	if err != nil {
		return err
	}

	if err := writer.WriteSamples(finalBuf); err != nil {
		return err
	}

	return writer.Close()
}

// calcDuration calculates the total duration of the track considering all events and playback speed
// Returns the duration in seconds (int)
func calcDuration(events []NoteEvent, speed float64) int {
	var maxEnd float64
	for _, e := range events {
		end := (e.Start + e.Duration) / speed
		if end > maxEnd {
			maxEnd = end
		}
	}
	return int(maxEnd) + 1
}

// getGID returns the ID of the current goroutine
// Used for logging to track which instrument is being processed in which goroutine
func getGID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var gid uint64
	_, err := fmt.Sscanf(string(buf[:n]), "goroutine %d ", &gid)
	if err != nil {
		log.Printf("getGID: failed to read goroutine ID: %v", err)
		return 0
	}
	return gid
}
