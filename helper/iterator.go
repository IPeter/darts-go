package helper

import (
	"container/list"
)

type Iterator struct {
	e *list.Element
}

func (p *Iterator) Valid() bool {
	return p.e != nil
}

func (p *Iterator) Value() (int, int) {
	pe := p.e.Value.(*Element)
	return pe.k, pe.v
}

func (p *Iterator) Next() {
	p.e = p.e.Next()
}

type Element struct {
	k, v int
}

type ListMap struct {
	m map[int]*list.Element
	l *list.List
}

func NewListMap() *ListMap {
	return &ListMap{
		m: make(map[int]*list.Element),
		l: list.New(),
	}
}

func (p *ListMap) Set(k, v int) {
	e, ok := p.m[k]
	if ok {
		e.Value.(*Element).v = v
	} else {
		p.m[k] = p.l.PushBack(&Element{k, v})
	}
}

func (p *ListMap) Get(k int) (int, bool) {
	e, ok := p.m[k]
	if !ok {
		return 0, false
	}
	return e.Value.(*Element).v, true
}

func (p *ListMap) Delete(k int) {
	e, ok := p.m[k]
	if ok {
		delete(p.m, k)
		p.l.Remove(e)
	}
}

func (p *ListMap) Iterate() Iterator {
	return Iterator{p.l.Front()}
}
