package leetcode

import "errors"

type Stack[T any] interface {
  Push(value T) error
  Pop() (*T, error)
  Peek() (*T, error)
  IsEmpty() bool
}


func NewStack[T any]() Stack[T] {
  s := new(stack[T])  
  s.data = make([]T, 1<<4)
  s.head = 0
  s.capacity = 1 << 10
  s.growthFactor = 10
  return s
}

type stack[T any] struct {
  data []T
  head int
  capacity int
  growthFactor int
}

func (s *stack[T]) Push(value T) error {
  if s.head == len(s.data)-1 {
    s.growthFactor = s.growthFactor + 1
    s.data = append(s.data, make([]T, 1<<s.growthFactor)...)
  }

  s.data[s.head + 1] = value
  s.head = s.head + 1

  return nil
}

func (s *stack[T]) Pop() (*T, error) {
  if s.head == -1 {
    return nil, errors.New("Empty stack")
  }

  if s.growthFactor > 4 && s.head < 1<<(s.growthFactor - 1) {
    s.growthFactor = s.growthFactor - 1
    s.data = s.data[:(1<<s.growthFactor + (1<<s.growthFactor-1))]
  }

  s.head = s.head - 1
  return &s.data[s.head + 1], nil
}

func (s *stack[T]) Peek() (*T, error) {
  if s.head == -1 {
    return nil, errors.New("Empty stack")
  }

  return &s.data[s.head], nil
}

func (s *stack[T]) IsEmpty() bool {
  return s.head == -1
}
