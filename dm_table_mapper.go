package xorm

import (
    "xorm.io/core"
)

// 定义一个新的Mapper，在SnakeMapper的基础上添加前缀
type PrefixedSnakeMapper struct {
    core.SnakeMapper
    prefix string
}

// 重写Obj2Table方法以添加前缀
func (mapper PrefixedSnakeMapper) Obj2Table(name string) string {
    // 首先使用SnakeMapper转换结构体名称到表名
    tableName := mapper.SnakeMapper.Obj2Table(name)
    // 然后添加指定的前缀
    return mapper.prefix + tableName
}