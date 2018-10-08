package handlers

type RequestEntity interface {
	UnmarshalJSON([]byte) error
}

type ResponseEntity interface {
	MarshalJSON() ([]byte, error)
}
