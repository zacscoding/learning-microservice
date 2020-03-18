package data

var data = []Kitten{
	{
		Id:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	{
		Id:     "3",
		Name:   "Garfield",
		Weight: 35.0,
	},
}

// simple in memory data store that implements Store
type MemoryStore struct {
}

func (m *MemoryStore) Search(name string) []Kitten {
	var kittens []Kitten

	for _, k := range data {
		if k.Name == name {
			kittens = append(kittens, k)
		}
	}

	return kittens
}
