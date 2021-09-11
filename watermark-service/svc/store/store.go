// store would implement the Database operations

package store

type dbStore struct {
	// DB object pointer

}

// DB interface
func (s dbStore) add(key string, value string) error {
	panic("not implemented") // TODO: Implement
}

func (s dbStore) get(key string) error {
	panic("not implemented") // TODO: Implement
}

func (s dbStore) delete(key string) error {
	panic("not implemented") // TODO: Implement
}
