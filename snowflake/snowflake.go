package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type IdGenertor struct {
	ID   int64
	node *snowflake.Node
}

func NewIdGenertor(id int64) IdGenertor {
	node, _ := snowflake.NewNode(id)
	return IdGenertor{ID: id, node: node}
}

func (i IdGenertor) Uint() uint {
	return uint(i.node.Generate().Int64())
}

func (i IdGenertor) Int64() int64 {
	return i.node.Generate().Int64()
}

func (i IdGenertor) String() string {
	return fmt.Sprintf("%d", i.node.Generate().Int64())
}
