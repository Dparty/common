package snowflake

import "github.com/bwmarrin/snowflake"

type IdGenertor struct {
	ID   int64
	node *snowflake.Node
}

func NewIdGenertor(id int64) IdGenertor {
	node, _ := snowflake.NewNode(id)
	return IdGenertor{ID: id, node: node}
}
