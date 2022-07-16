package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Length     int
	FrontPoint *ListItem
	BackPoint  *ListItem
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FrontPoint
}

func (l *list) Back() *ListItem {
	return l.BackPoint
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.Length++
	var newL ListItem
	newL.Value = v
	if l.FrontPoint != nil {
		newL.Next = l.FrontPoint
		l.FrontPoint.Prev = &newL
		l.FrontPoint = &newL
	} else {
		l.FrontPoint = &newL
		l.BackPoint = &newL
	}
	return &newL
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.Length++
	var newL ListItem
	newL.Value = v
	if l.BackPoint != nil {
		newL.Prev = l.BackPoint
		l.BackPoint.Next = &newL
		l.BackPoint = &newL
	} else {
		l.FrontPoint = &newL
		l.BackPoint = &newL
	}
	return &newL
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	l.Length--
	if i.Next == nil {
		i.Prev.Next = nil
	} else if i.Prev == nil {
		i.Next.Prev = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	return new(list)
}
