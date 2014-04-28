package bags

type Bag []interface{}

func (b *Bag) Add(item interface{}) {
	*b = append(*b, item)
}

func (b *Bag) IsEmpty() bool {
	return len(*b) == 0
}

func (b *Bag) Size() int {
	return len(*b)
}

func NewBag() *Bag {
	return &Bag{}
}
