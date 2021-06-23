package db

import (
	"qlite/node"
	class "qlite/struct"
)

type Database interface {
	Set			([]byte, []byte)	class.Message
	SetX		([]byte, []byte)	class.Message
	SetNode		(node.Node)			class.Message
	SetNodeX	(node.Node)			class.Message
	Get			([]byte)			class.Message
	Type		([]byte)			class.Message
	Exists		([]byte)			class.Message
	Delete		([]byte)			class.Message
	DeleteList	([][]byte)			class.Message
	Rename		([]byte, []byte)	class.Message
	RenameX		([]byte, []byte)	class.Message
}