# Baba O`Riley... Sorry.
Experimental synthesizer/sequencer in Go that plays events (notes) from JSON and saves the output to a WAV file. The name is inspired by the song **Baba O'Riley by The Who**.

## Disclaimer
While developing this application, I in no way intended to offend the artistic legacy of **The Who**.
If the quality of the performance of **Baba O’Riley** seemed inappropriate — I am truly sorry that it turned out this way...

## Features
- Playback of events (notes) from `JSON`.
- Support for different instruments (e.g., `Kick`).
- Rendering to mono `WAV` (16-bit PCM).
- Multithreaded generation: each instrument is processed in its own goroutine.

## Project Structure
```text
baba-oriley/
├── assets/           # JSON files with notes
├── out/              # Output folder for generated .wav files
├── internal/
│   ├── audio/        # WAVWriter and audio handling
│   ├── instruments/  # Instruments (Kick, etc.)
│   └── player/       # Logic for loading and playing events
└── cmd/
    └── main.go       # Entry point
```
## Example JSON (assets/demo.json)
```json
[
  {"start": 17.70, "duration": 0.59, "note": 72, "velocity": 105, "instrument": "arp3"},
  {"start": 9.60, "duration": 0.28, "note": 72, "velocity": 105, "instrument": "arp1"}
]
```

## Running
Build and run:
```bash
go run ./cmd
```
After that, a `baba.wav` file will appear in the `out/` folder.

## Adding Instruments

A new instrument must implement the method:

```go
Play(note int, velocity int, duration float64, sampleRate int) []int16
```
and be registered in `instruments`.
Example: the Kick (bass drum) is synthesized as a decaying sine wave.

## TODO
- Add more instruments (snare, bass, synth).
- Implement stereo support.
- MIDI import/export.
- Real-time player.


