package g

const DUP_KEY_ERR = "Error 1062"
const (
	ReqFailedE = "Failed to request the backend service"
	ReqFailedC = 1001

	ParamEmptyE   = "Params cant be empty"
	ParamEmptyC   = 1002
	ParamInvalidE = "Params invalid"
	ParamInvalidC = 1003

	DatabaseE = "Access database error"
	DatabaseC = 1019

	AlreadyExistE     = "Target alread exist"
	AlreadyExistC     = 1030
	NotExistE         = "Target not exist"
	NotExistC         = 1031
	UserNotExistE     = "User not exist"
	UserNotExistC     = 1032
	UserNotExistOrPwE = "User not exist or password invalid"
	UserNotExistOrPwC = 1033

	NotExistOrAuthFailedE = "User not exist or has no permission"
	NotExistOrAuthFailedC = 1040

	ForbiddenE = "No permission"
	ForbiddenC = 1050

	ForbiddenViewerE = "You are the guest, no permissio to do so"
	ForbiddenViewerC = 1051

	OperateOnSelfE = "Cant operate on self"
	OperateOnSelfC = 1053

	NeedLoginE = "Session expires,need login"
	NeedLoginC = 1054

	TargetServiceNotFountE = "No application node founded"
	TargetServiceNotFountC = 1056

	NoUpdateHappenedE = "No need to update"
	NoUpdateHappenedC = 1057

	AccessLimitedE = "Access limited"
	AccessLimitedC = 1058

	OldPWErrorE = "Old password error or new pw and old pw is the same"
	OldPWErrorC = 1059

	PWAlphabetAndNumE = "Password can only be consisted of alphabet and numberic"
	PWAlphabetAndNumC = 1060

	ServerNotExistE = "Server is not exist"
	ServerNotExistC = 1061

	NoServerAvailableE = "No server available"
	NoServerAvailableC = 1062
)
