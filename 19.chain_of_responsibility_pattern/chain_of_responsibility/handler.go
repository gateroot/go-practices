package chain_of_responsibility

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	Request(int)
}

type ConcreteHandlerA struct {
	next   Handler
	rangeS int
	rangeE int
	name   string
}

func NewConcreteHandlerA(name string, rangeS, rangeE int) Handler {
	return &ConcreteHandlerA{name: name, rangeS: rangeS, rangeE: rangeE}
}

func (h *ConcreteHandlerA) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandlerA) Request(req int) {
	if req >= h.rangeS && req <= h.rangeE {
		fmt.Println(h.name, ": 私が成敗してくれるわ！")
	} else if h.next != nil {
		fmt.Println(h.name, ": 私では無理じゃ！次の者出合えー")
		h.next.Request(req)
	} else {
		fmt.Println(h.name, ": 次の者がおりませぬー")
	}
}

type ConcreteHandlerB struct {
	next Handler
	name string
}

func NewConcreteHandlerB(name string) Handler {
	return &ConcreteHandlerB{name: name}
}

func (h *ConcreteHandlerB) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *ConcreteHandlerB) Request(req int) {
	if req%2 == 0 {
		fmt.Println(h.name, ": ようこそ！偶数は私のテリトリーです。")
	} else if h.next != nil {
		fmt.Println(h.name, ": 無念！偶数以外は・・・！次の方お願いします。")
		h.next.Request(req)
	} else {
		fmt.Println(h.name, ": 無念！偶数以外は・・・！次もいないー")
	}
}
