package item

type Item struct {
	Name string
}

func NewItem(name string) *Item {
	return &Item{name}
}
