package client

// Response contains the JSON body
type Response struct {
	Headers map[string]string
	Status  int
	Body    []byte
	Data    string
}
