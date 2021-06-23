package trie

import (
	db "qlite/database"
	class "qlite/struct"
)

type NewOptionTrie func(db.Database) (TrieManager, error)

type TrieManager interface {
	GetName() []byte
	GetIntroduce() []byte
	Work([]byte, []class.Message) class.Message
}