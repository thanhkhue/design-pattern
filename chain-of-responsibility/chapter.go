package main

const (
	fieldChapterName   = "name"
	fieldChapterNumber = "chapter_number"
)

func init() {
	registerEntryHandlerFunc("chapter", RetrieveChapterInfo, []string{fieldChapterName, fieldChapterNumber})
}

// RetrieveChapterInfo retrieve chapter info
func RetrieveChapterInfo(objects *Response, fields *FieldFlag) *Response {

	retrieveFields := make([]string, 0)

	if (*fields)[fieldChapterName] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldChapterName)
	}

	if (*fields)[fieldChapterName] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldChapterName)
	}

	if len(retrieveFields) == 0 {
		return objects
	}

	queryChapterInfo(objects, retrieveFields)
	// assign queried data

	// mark as done
	for _, field := range retrieveFields {
		(*fields)[field] = DoneFlag
	}

	return objects
}

func queryChapterInfo(objects *Response, fields []string) *Response {
	// query chapter fields
	return objects
}
