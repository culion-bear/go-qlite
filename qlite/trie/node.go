package trie

import (
	"errors"
	class "qlite/struct"
)

type Option func ([]class.Message) class.Message

type node struct {
	list	[26]*node
	value	Option
}

var errOptionNotFound = class.Message{
	Type:    '!',
	IsWrote: false,
	String:  []byte("option is not found"),
}


var errOptionIllegal = class.Message{
	Type:    '!',
	IsWrote: false,
	String:  []byte("option is illegal"),
}

func newNode() *node{
	return &node{}
}

func (m *node) push(msg []byte, f Option, k int) error{
	if k == len(msg){
		if m.value != nil{
			return errors.New(string(msg) + " option is exist")
		}
		m.value = f
		return nil
	}
	key := msg[k] - 'a'
	if key >= 26 || key < 0{
		return errors.New(string(msg) + " option is illegal")
	}
	if m.list[key] == nil{
		m.list[key] = newNode()
	}
	return m.list[key].push(msg, f, k + 1)
}

func (m *node) work(opt []byte, msg []class.Message, k, length int) class.Message{
	if k == length{
		if m.value == nil{
			return errOptionNotFound
		}
		return m.value(msg)
	}
	key := opt[k] - 'a'
	if key >= 26 || key < 0{
		return errOptionIllegal
	}
	if m.list[key] == nil {
		return errOptionNotFound
	}
	return m.list[key].work(opt, msg, k + 1, length)
}