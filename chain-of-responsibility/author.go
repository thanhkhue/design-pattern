package main

const (
	fieldAuthorName  = "name"
	fieldAuthorAge   = "age"
	fieldAuthorPhone = "phone"
)

func init() {
	registerEntryHandlerFunc("author", RetrieveAuthorInfo, []string{fieldAuthorName, fieldAuthorAge, fieldAuthorPhone})
}

// RetrieveAuthorInfo retrieve author info
func RetrieveAuthorInfo(objects *Response, fields *FieldFlag) *Response {

	retrieveFields := make([]string, 0)

	if (*fields)[fieldAuthorName] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldAuthorName)
	}

	if (*fields)[fieldAuthorAge] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldAuthorAge)
	}

	if (*fields)[fieldAuthorPhone] == WaitingFlag {
		retrieveFields = append(retrieveFields, fieldAuthorPhone)
	}

	if len(retrieveFields) == 0 {
		return objects
	}

	queryAuthorInfo(objects, retrieveFields)
	// assign queried data

	// mark as done
	for _, field := range retrieveFields {
		(*fields)[field] = DoneFlag
	}

	return objects
}

func queryAuthorInfo(objects *Response, fields []string) *Response {
	// query author fields
	return objects
}
