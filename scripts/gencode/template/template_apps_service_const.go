package template

var AppsServiceConstTemplate = ParseTemplate(`
package service

const (
	ERROR_CODE_{{.UpperServiceName}}_GRPC_SERVICE_FAILURE                         int32 = 1
    ERROR_CODE_{{.UpperServiceName}}_QUERY_DB_FAILED                              int32 = 2
	ERROR_CODE_{{.UpperServiceName}}_INSERT_VALUE_FAILED                          int32 = 3
	ERROR_CODE_{{.UpperServiceName}}_UPDATE_VALUE_FAILED                          int32 = 4
	ERROR_CODE_{{.UpperServiceName}}_REDIS_GET_FAILED                             int32 = 5
	ERROR_CODE_{{.UpperServiceName}}_REDIS_SET_FAILED                             int32 = 6
)

const (
	ERROR_{{.UpperServiceName}}_GRPC_SERVICE_FAILURE                         = "服务故障"
	ERROR_{{.UpperServiceName}}_QUERY_DB_FAILED                              = "查询失败"
	ERROR_{{.UpperServiceName}}_INSERT_VALUE_FAILED                          = "数据入库失败"
	ERROR_{{.UpperServiceName}}_UPDATE_VALUE_FAILED                          = "更新Value失败"
	ERROR_{{.UpperServiceName}}_REDIS_GET_FAILED                             = "读取redis缓存失败"
	ERROR_{{.UpperServiceName}}_REDIS_SET_FAILED                             = "设置redis缓存失败"
)
`)
