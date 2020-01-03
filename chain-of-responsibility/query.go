package main

func retriever(objects *Response, codeFields map[string][]string) {
	entries := make([]*RetrieverEntry, 0)
	entryFields := make(map[*RetrieverEntry]FieldFlag)
	for code, fields := range codeFields {
		entry := getEntryHandlerFunc(code)
		if entry == nil {
			continue
		}
		entries = append(entries, entry)
		entryFields[entry] = markFieldFlag(fields)
	}

	if len(entries) == 0 || len(entryFields) == 0 {
		return
	}

	for i := 0; i < len(entries); i++ {
		process := false
		entry := entries[i]
		for field := range entry.SupportedFields {
			if entryFields[entry][field] == WaitingFlag {
				process = true
			}
		}
		if process {
			fields := entryFields[entry]
			entry.Handler(objects, &fields)
		}
	}
}

func markFieldFlag(fields []string) FieldFlag {
	fieldsFlag := make(FieldFlag)
	for _, field := range fields {
		fieldsFlag[field] = WaitingFlag
	}
	return fieldsFlag
}
