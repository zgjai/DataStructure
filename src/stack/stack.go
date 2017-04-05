package stack

type Stack struct {
	items     []interface{}
	lastIndex int
}

func NewStack(c int) *Stack {
	return &Stack{items: make([]interface{}, c), lastIndex: 0}
}

func (s *Stack) Push(item interface{}) {
	if s.lastIndex < len(s.items) {
		s.items[s.lastIndex] = item
	} else {
		s.items = append(s.items, item)
	}
	s.lastIndex++
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.lastIndex--
	return s.items[s.lastIndex]
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[s.lastIndex-1]
}

func (s *Stack) Depth() int {
	return s.lastIndex
}

func (s *Stack) IsEmpty() bool {
	return s.Depth() == 0;
}