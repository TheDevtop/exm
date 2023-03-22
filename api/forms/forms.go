package forms

type SearchForm struct {
	Object string
	Regex  string
}

type ReplaceForm struct {
	Object  string
	Regex   string
	Mapping string
}

type ResultForm struct {
	Route   string
	Error   string
	Results []string
}
