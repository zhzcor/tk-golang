package errorno

var (
	//基础配置错误
	RoleNotExistsError = Errno{
		Code:    20001,
		Message: "角色不存在",
	}
	AdminNotExistsError = Errno{
		Code:    20002,
		Message: "管理员不存在",
	}
	RoleExistsError = Errno{
		Code:    20003,
		Message: "角色已存在",
	}
	PermissionExistsError = Errno{
		Code:    20004,
		Message: "权限已存在",
	}
	AdminNotfoundError = Errno{
		Code:    20005,
		Message: "管理员不存在",
	}
	PidNotfoundError = Errno{
		Code:    20006,
		Message: "父级不存在",
	}
)
