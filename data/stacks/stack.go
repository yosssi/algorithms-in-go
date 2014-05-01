package stacks

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(*s) - 1
	item := []interface{}(*s)[lastIndex]
	*s = []interface{}(*s)[:lastIndex]
	return item
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func NewStack() *Stack {
	return &Stack{}
}
