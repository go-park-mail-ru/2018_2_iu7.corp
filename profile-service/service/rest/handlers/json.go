package handlers

type requestEntity interface {
	UnmarshalJSON([]byte) error
}

type responseEntity interface {
	MarshalJSON() ([]byte, error)
}
