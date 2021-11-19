package errorno

// Errno 定义业务错误码
type Errno struct {
	Code    int
	Message string
}

// Err 定义业务错误
type Err struct {
	Code    int       // 错误码
	Message string    // 展示给用户看的
	Errored []error   // 保存内部错误信息
	stack   []uintptr // 错误跟踪
}
