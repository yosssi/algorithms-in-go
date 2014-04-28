package bags

type Bag struct {
	items []interface{}
}

func (b *Bag) Add(item interface{}) {
	b.items = append(b.items, item)
}

func (b *Bag) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *Bag) Size() int {
	return len(b.items)
}

func NewBag() *Bag {
	return &Bag{}
}
