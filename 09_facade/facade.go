package facade

// IUser 用户接口
type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

// IUserFacade 门面模式
type IUserFacade interface {
	LoginOrRegister(phone int, code int) (*User, error)
}

// User 用户
type User struct {
	Name string
}

func newUser(name string) IUser {
	return User{Name: name}
}

// Login 登录
func (u User) Login(phone int, code int) (*User, error) {
	// 校验操作 ...
	return &User{Name: "test login"}, nil
}

// Register 注册
func (u User) Register(phone int, code int) (*User, error) {
	// 校验操作 ...
	// 创建用户
	return &User{Name: "test register"}, nil
}

// UserService facade struct
type UserService struct {
	User IUser
}

func NewUserService(name string) UserService {
	return UserService{
		User: newUser(name),
	}
}

// LoginOrRegister 登录或注册
func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.User.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.User.Register(phone, code)
}
