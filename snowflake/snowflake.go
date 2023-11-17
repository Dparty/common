package snowflake

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func NewIdGenertor(id int64) *snowflake.Node {
	node, _ = snowflake.NewNode(id)
	return node
}
