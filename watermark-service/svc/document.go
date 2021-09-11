package svc

type Repository interface {
	add(d Document) error
	get(key string) (Document, error)
	delete(key string) error
}

type Document struct {
	// Document type definition
}
