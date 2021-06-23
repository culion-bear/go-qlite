package trie

import (
	"errors"
	class "qlite/struct"
)

type Manager struct {
	head	*node
}

func New() *Manager{
	return &Manager{head:newNode()}
}

func (m *Manager) Push(msg []byte, f Option) error{
	if len(msg) > 10 || len(msg) == 0{
		return errors.New(string(msg) + "option is too long or too short")
	}
	return m.head.push(msg, f, 0)
}

func (m *Manager) Work(opt []byte, list []class.Message) class.Message{
	return m.head.work(opt, list, 0, len(opt))
}