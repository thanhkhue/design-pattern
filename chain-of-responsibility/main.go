package main

func main() {
	objects := &Response{}
	codeFields := map[string][]string{
		"author":    []string{"name", "age", "phone"},
		"chapter":   []string{"name", "chapter_number"},
		"character": []string{"name", "bio"},
	}
	retriever(objects, codeFields)
}
