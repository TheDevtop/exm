package forms

type ObjectForm struct {
	Object string
}

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
