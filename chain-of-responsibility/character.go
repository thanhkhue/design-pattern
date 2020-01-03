package main

const (
	fieldCharacterName   = "name"
	fieldCharacterNumber = "bio"
)

func init() {
	registerEntryHandlerFunc("character", RetrieveCharacterInfo, []string{fieldCharacterName, fieldCharacterNumber})
}

// RetrieveCharacterInfo retrieve character info
func RetrieveCharacterInfo(objects *Response, fields *FieldFlag) *Response {

	retrieveFields := make([]string, 0)

	if (*fields)[fieldCharacterName] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldCharacterName)
	}

	if (*fields)[fieldCharacterNumber] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldCharacterNumber)
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

func queryCharacterInfo(objects *Response, fields []string) *Response {
	// query character fields
	return objects
}
