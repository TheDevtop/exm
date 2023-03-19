package forms

type SearchForm struct {
	Object string
	Regex  string
}

type ResultForm struct {
	Route   string
	Error   string
	Results []string
}
