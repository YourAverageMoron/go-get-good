package main

type MinStack struct {
	l     int
	stack []int
	min   []int
}

func MinStackConstructor() MinStack {
	return MinStack{
		l:     0,
		stack: []int{},
		min:   []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if m.l == 0 || m.GetMin() > val {
		m.min = append(m.min, val)
	} else {
		m.min = append(m.min, m.GetMin())
	}
	m.l++
}

func (m *MinStack) Pop() {
	m.l--
	m.stack = m.stack[:m.l]
	m.min = m.min[:m.l]
}

func (m *MinStack) Top() int {
	return m.stack[m.l-1]
}

func (m *MinStack) GetMin() int {
	return m.min[m.l-1]
}
