package request

type Request struct {
	Method string `json:"method"`
	Host   string `json:"host"`
}

type Response struct {
}

func NewRequest(method, host string) *Request {
	return &Request{
		Method: method,
		Host:   host,
	}
}
