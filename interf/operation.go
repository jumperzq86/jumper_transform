package interf

//定义打包和加密类型

//操作接口
type Operation interface {
	Operate(direct int8, input interface{}, output interface{}) (bool, error)
}
