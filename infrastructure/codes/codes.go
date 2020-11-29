package codes

const (
	SUCCESS         = 1  //成功
	SUCCESS_Partial = 2  //部分成功（批量操作）
	Fail            = -1 //失败

	ServerError   = 10001 //本服务异常
	UpServerError = 10002 //上游服务异常
	DbError       = 10101 //数据库（持久化）异常
	NotFound      = 10404 //找不到

	//登录相关业务
	//
	CreateTokenErr   = 60101
	CaptchaErr       = 60102
	LoginNameExisted = 60103
	PhoneExisted     = 60104
	WxInfoSet        = 60105
)

const ()
