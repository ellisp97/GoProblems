package Utility

type Stack struct {
	data []interface{}
}

func (s Stack) Len() int { return len(s.data) }

func (s Stack) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() interface{} {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *Stack) Push(el interface{}) {
	s.data = append(s.data, el)
}
