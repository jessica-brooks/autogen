package request

//Config Configuration for HTTP Request
type Config struct {
	APIURL  string
	Timeout Timeout
	Headers Headers
}

//Timeout struct for connection timeout
type Timeout struct {
	ConnectionTimeout int
}

//Headers for http request
type Headers struct {
	RequestHeaders map[string]string
}
