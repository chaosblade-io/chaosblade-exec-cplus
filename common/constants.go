package common

const (
	TargetName                = "cplus"
	DelayActionName           = "delay"
	ReturnErrorDataActionName = "return"
	VariableModifyActionName  = "modify"
	FileKeyword               = "file "
	SetEnvLdLibraryPath       = "set env LD_LIBRARY_PATH "
	SleepMilliseconds         = 1000
	CorePoolSize              = 1
	MaximumPoolSize           = 1
	KeepAliveTime             = 0
	BinName                   = "chaosblade-exec-cplus"
)

const (
	InitializationScript        = "shell_initialization.sh"
	RemoveProcessScript         = "shell_remove_process.sh"
	ResponseDelayScript         = "shell_response_delay.sh"
	ResponseDelayAttachScript   = "shell_response_delay_attach.sh"
	BreakAndReturnScript        = "shell_break_and_return.sh"
	BreakAndReturnAttachScript  = "shell_break_and_return_attach.sh"
	ModifyVariableScript        = "shell_modify_variable.sh"
	ModifyVariableAttachScript  = "shell_modify_variable_attach.sh"
	CheckProcessIdScript        = "shell_check_process_id.sh"
	CheckProcessDuplicateScript = "shell_check_process_duplicate.sh"
)
