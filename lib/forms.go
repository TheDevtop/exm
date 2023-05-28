package lib

type FormRequest struct {
	Object string
	Regex  string
}

type FormResult struct {
	Error   string
	Count   int
	Results []string
}

type FormMetadata struct {
	Error        string
	Object       string
	Type         string
	Size         int64
	Source       string
	LastModified string
}
