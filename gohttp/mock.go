package gohttp


type Mock struct {
	Method             string
	Url                string
	RequestBody        string

	// Response
	Error error
	ResponseBody       string
	ResponseStatusCode int


}

