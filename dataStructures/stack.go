package dataStructures

type stack[T any] struct {
	storage   []T
	pointer   int
	isDynamic bool
}

func CreateStack[T any](initialLen int, isDynamic bool) *stack[T] {
	st := stack[T]{
		storage:   make([]T, initialLen),
		pointer:   1,
		isDynamic: isDynamic,
	}
	return &st
}

func (st *stack[T]) Push(value T) {
	if cap(st.storage) > st.pointer && st.isDynamic {
		st.storage = append(st.storage, value)
		st.pointer++
		return
	}

	if cap(st.storage) > st.pointer {
		panic("Stack Overflow!")
	}

	st.pointer++
	st.storage[st.pointer] = value

}

func (st *stack[T]) Pop() (value T) {
	if st.pointer <= 1 {
		panic("Stack is empty!")
	}

	value = st.storage[st.pointer]
	st.pointer--
	return
}
