package data

var data = []Kitten{
	Kitten{
		Id:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	Kitten{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	Kitten{
		Id:     "3",
		Name:   "Garfield",
		Weight: 35.0,
	},
}

// MemoryStore is a simple in memory datastore that implements Store
type MemoryStore struct {
}

//Search returns a slice of Kitten which have a name matching the name in the parameters
func (m *MemoryStore) Search(name string) []Kitten {
	var kittens []Kitten

	for _, k := range data {
		if k.Name == name {
			kittens = append(kittens, k)
		}
	}

	return kittens
}
