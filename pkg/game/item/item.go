package item

type Item interface {
	Name() string
}

type BaseItem struct {
	name string
}

func NewBaseItem(name string) *BaseItem {
	return &BaseItem{name}
}

func (i *BaseItem) Name() string {
	return i.name
}
