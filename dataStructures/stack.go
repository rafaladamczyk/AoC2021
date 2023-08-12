package datastructures

type Stack []byte

func (s Stack) Push(v byte) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Peek() byte {
	return s[len(s)-1]
}
