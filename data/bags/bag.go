package bags

type bag struct {
	items []interface{}
}

func (b *bag) Add(item interface{}) {
	b.items = append(b.items, item)
}

func (b *bag) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *bag) Size() int {
	return len(b.items)
}

func NewBag() *bag {
	return &bag{}
}
