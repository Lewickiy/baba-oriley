package instruments

import (
	"log"
	"strings"
	"sync"
	"sync/atomic"
)

// nextInstrumentID is a global counter for unique instrument IDs
var nextInstrumentID uint64 = 0

// InstrumentInstances stores created instrument instances by name
var InstrumentInstances = map[string]Instrument{}

// getNextID returns a unique integer ID for a new instrument
func getNextID() int {
	return int(atomic.AddUint64(&nextInstrumentID, 1))
}

// protects access to InstrumentInstances
var instMu sync.Mutex

// GetInstrument returns an instrument object by name
// If an instrument with this name already exists, it returns the existing instance
// Otherwise, it creates a new instrument (Kick, ArpSynth, or default ArpSynth)
// assigns a unique ID, and saves it in InstrumentInstances
// Logs the creation of the instrument and its object address
func GetInstrument(name string) Instrument {
	instMu.Lock()
	defer instMu.Unlock()

	if inst, ok := InstrumentInstances[name]; ok {
		return inst
	}

	var inst Instrument

	switch {
	case strings.HasPrefix(name, "kick"):
		inst = &Kick{id: getNextID()}
		log.Printf("[Factory] Created Kick instrument for name '%s', object address: %p\n", name, inst)
	case strings.HasPrefix(name, "arp"):
		inst = &ArpSynth{id: getNextID()}
		log.Printf("[Factory] Created ArpSynth instrument for name '%s', object address: %p\n", name, inst)
	default:
		inst = &ArpSynth{id: getNextID()}
		log.Printf("[Factory] Created default SineSynth instrument for name '%s', object address: %p\n", name, inst)
	}

	InstrumentInstances[name] = inst
	return inst
}
