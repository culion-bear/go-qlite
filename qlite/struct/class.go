package class

import (
	"qlite/node"
)

var scale	=	[9]int{0x1, 0x10, 0x100, 0x1000, 0x10000, 0x100000, 0x1000000, 0x10000000, 0x100000000}

var bytes	=	[16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

var errType = []byte("type is not exist")

type Message struct {
	Type			byte
	IsWrote			bool
	String			[]byte
	List			[][]byte
	Node			node.Node
	Number			int
}

func NewError(err []byte) Message{
	return Message{
		Type:    '!',
		IsWrote: false,
		String:  err,
	}
}

func (m *Message) ToString() []byte{
	switch m.Type {
	case '+', '!':
		str := make([]byte, len(m.String))
		copy(str, m.String)
		return str
	case '-':
		return m.listToString()
	case '?':
		str := make([]byte, len(m.Node.GetKey()))
		copy(str, m.Node.GetKey())
		return str
	case '*':
		return m.intToByte(m.Number)
	default:
		m.Type = '!'
		return errType
	}
}

func (m *Message) IsString() bool{
	return m.Type == '+'
}

func (m *Message) IsList() bool{
	return m.Type == '-'
}

func (m *Message) IsNode() bool{
	return m.Type == '?'
}

func (m *Message) IsError() bool{
	return m.Type == '!'
}

func (m *Message) IsNumber() bool{
	return m.Type == '*'
}

func (m *Message) listToString() []byte{
	str := make([]byte, 0)
	for k, v := range m.List{
		if k != 0{
			str = append(str, ',')
		}
		str = append(str, v...)
	}
	return str
}

func (m *Message) intToByte(length int) []byte{
	s := make([]byte, 0)

	s = append(s, bytes[length % 16])
	length /= 16

	for length != 0{
		s = append(s, bytes[length % 16])
		length /= 16
	}

	return s
}

func (m Message) Print() []byte{
	switch m.Type {
	case '+':
		return append([]byte{'+'}, m.toString(m.String)...)
	case '-':
		return append([]byte{'-'}, m.toList()...)
	case '?':
		return append([]byte{'?'}, m.toString(m.Node.GetKey())...)
	case '*':
		return append([]byte{'*'}, m.intToByte(m.Number)...)
	case '!':
		return append([]byte{'!'}, m.toString(m.String)...)
	default:
		return append([]byte{'!'}, m.toString(errType)...)
	}
}

func (m *Message) toString(msg []byte) []byte{
	return append(append(m.intToByte(len(msg)), ';'), msg...)
}

func (m *Message) toList() []byte{
	str := m.intToByte(len(m.List))
	str = append(str, ';')
	for _, v := range m.List{
		str = append(str, append(append(m.intToByte(len(v)), ';'), v...)...)
	}
	return str
}