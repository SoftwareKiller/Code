package service

import "fmt"

// UserService 是一个用户服务接口
type UserService interface {
	GetUserByID(id int) string
}

// UserServiceImpl 是 UserService 的具体实现
type UserServiceImpl struct{}

// GetUserByID 是 UserServiceImpl 的方法实现
func (s *UserServiceImpl) GetUserByID(id int) string {
	return fmt.Sprintf("User %d", id)
}

// UserController 是一个控制器，依赖于 UserService，依赖抽象
type UserController struct {
	UserService
}

// GetUserByID 是 UserController 的方法，通过依赖注入获取用户信息
func (c *UserController) GetUserByID(id int) string {
	return c.UserService.GetUserByID(id)
}

func IocDesign() {
	// 创建 UserServiceImpl 实例
	userService := &UserServiceImpl{}

	// 创建 UserController 实例，并通过依赖注入设置 UserService
	userController := &UserController{UserService: userService}

	// 使用 UserController 获取用户信息
	user := userController.GetUserByID(1)
	fmt.Println(user)
}
