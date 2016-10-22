package data

// Store is an interface used for interacting with the backend datastore
type Store interface {
	Search(name string) []Kitten
}
