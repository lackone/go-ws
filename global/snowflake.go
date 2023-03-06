package global

import "github.com/bwmarrin/snowflake"

var (
	SnowflakeNode *snowflake.Node
)

func InitSnowflakeNode() error {
	if SnowflakeNode == nil {
		var err error
		SnowflakeNode, err = snowflake.NewNode(SnowflakeSetting.NodeId)
		if err != nil {
			return err
		}
	}
	return nil
}
