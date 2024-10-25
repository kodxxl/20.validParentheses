package slice

func IsValid(source string) bool {
	var stack Stack
	open := map[rune]rune{
		'{': '{',
		'[': '[',
		'(': '(',
	}
	close := map[rune]rune{
		'}': open['{'],
		']': open['['],
		')': open['('],
	}
	for _, value := range source {
		if _, ok := open[value]; ok {
			stack.push(value)
		}
		if current, ok := close[value]; ok {
			if stack.top() == current {
				stack.pop()
				continue
			}
			return false
		}
	}
	return (stack.size() <= 0)
}

// /-----------------
type Stack struct {
	state []rune
}

func (s *Stack) push(value rune) {
	s.state = append(s.state, value)
}

func (s *Stack) pop() rune {
	if s.size() <= 0 {
		return rune(0)
	}
	top := s.top()
	s.state = s.state[:len(s.state)-1]

	return top
}

func (s Stack) top() rune {
	if s.size() <= 0 {
		return rune(0)
	}
	return s.state[len(s.state)-1]
}

func (s Stack) size() int {
	return len(s.state)
}
