package idgen

import (
	"github.com/bwmarrin/snowflake"
)

// 定义ID生成器接口，未来可以根据需要实现不同的ID生成器
type Generator interface {
	NextID() uint64
}

type Snowflake struct {
	Node *snowflake.Node
}

func NewSnowflake(nodeID int64) *Snowflake {
	node, _ := snowflake.NewNode(nodeID)
	return &Snowflake{Node: node}
}

func (s *Snowflake) NextID() uint64 {
	return uint64(s.Node.Generate().Int64())
}

// sonyflake
type Sonyflake struct {
	Node *snowflake.Node
}

func NewSonyflake(nodeID uint64) *Sonyflake {
	node, _ := snowflake.NewNode(int64(nodeID))
	return &Sonyflake{Node: node}
}

func (s *Sonyflake) NextID() uint64 {
	return uint64(s.Node.Generate())
}
