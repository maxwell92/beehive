package error

type Error struct {
	LogMsg string
	ErrMsg string
}

const (
	EOK int32 = 0

	//TODO: 修改为itoa形式的枚举
	EMYSQL_CONN   int32 = 1001
	EMYSQL_QUERY  int32 = 1002
	EMYSQL_INSERT int32 = 1003
	EMYSQL_DELETE int32 = 1004
	EMYSQL_UPDATE int32 = 1005
)

var Errors = map[int32]*Error{

	EOK: &Error{
		LogMsg: "OK",
		ErrMsg: "操作成功",
	},

	EMYSQL_CONN: &Error{
		LogMsg: "MySQL Connection Error",
		ErrMsg: "MySQL连接错误",
	},
	EMYSQL_QUERY: &Error{
		LogMsg: "MySQL Connection Error",
		ErrMsg: "MySQL查询错误",
	},
	EMYSQL_INSERT: &Error{
		LogMsg: "MySQL Connection Error",
		ErrMsg: "MySQL插入错误",
	},
	EMYSQL_UPDATE: &Error{
		LogMsg: "MySQL Connection Error",
		ErrMsg: "MySQL更新错误",
	},
	EMYSQL_DELETE: &Error{
		LogMsg: "MySQL Connection Error",
		ErrMsg: "MySQL删除错误",
	},
}
