package controllers

const (
	ERROR_OK                 = 0  //处理成功
	ERROR_FALSE              = -1 //处理失败
	ERROR_CONFLICT           = 1
	ERROR_MAC_FORMAT         = 30001 //格式错误
	ERROR_MAC_OWN            = 30002 //为自有mac
	ERROR_MAC_NOTMATCH       = 30003 //为自有mac
	ERROR_PARAM_NOT_COMPLETE = 40001 //参数不完整
	ERROR_GET_DATA_WRONG     = 40002 //获取数据失败
	ERROR_WRONG_PARAM        = 40003 //参数信息错误
	ERROR_NODATA             = 40004 //无数据
)
