package main

import "sync"

const (
	// WaitingFlag waiting flag status
	WaitingFlag byte = 3
	// DoneFlag done flag status
	DoneFlag byte = 1
	// DeniedFlag denied flag status
	DeniedFlag byte = 2
)

// ObjectResponse object response
type ObjectResponse map[string][]byte

// Response data response
type Response map[string]*ObjectResponse

// FieldFlag field flag
type FieldFlag map[string]byte

// RetrieverFunc handler function
type RetrieverFunc func(objects *Response, fields *FieldFlag) *Response

// RetrieverEntry handler and supported fields
type RetrieverEntry struct {
	Handler         RetrieverFunc
	SupportedFields map[string]bool
}

var registries = make(map[string]*RetrieverEntry)
var lock sync.RWMutex

func registerEntryHandlerFunc(code string, handler RetrieverFunc, fields []string) {
	supportedFields := make(map[string]bool)
	for _, field := range fields {
		supportedFields[field] = true
	}
	retrieverEntry := &RetrieverEntry{
		Handler:         handler,
		SupportedFields: supportedFields,
	}
	lock.Lock()
	registries[code] = retrieverEntry
	lock.Unlock()
}

func getEntryHandlerFunc(code string) *RetrieverEntry {
	lock.RLock()
	entryHandler := registries[code]
	lock.RUnlock()
	return entryHandler
}
