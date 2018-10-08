package handlers

type Entity interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

type RequestEntity interface {
	Entity
}

type ResponseEntity interface {
	Entity
}
