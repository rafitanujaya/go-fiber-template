package functionCallerInfo

type FunctionCaller string

const (
	UserRepositoryCreate   FunctionCaller = "userRepository.Create"
	UserServiceRegister    FunctionCaller = "userService.Register"
	UserControllerRegister FunctionCaller = "userController.Register"
)
