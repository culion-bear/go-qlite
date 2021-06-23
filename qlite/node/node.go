package node

type NewNode func ([]byte) Node

type Node interface {
	ToHash(int) int				//将key转为hash值
	GetKey() []byte				//获取key
	SetKey([]byte)				//设置key
	GetType() []byte			//获取类别
	Compare(Node) bool			//与node的key比较
	CompareKey([]byte) bool		//直接与key比较
}
