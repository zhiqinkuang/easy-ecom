package service

import (
	"errors"
	"github.com/zhiqinkuang/easy-ecom/repository"
	"regexp"
)

// UserMsg 的结构体
type UserMsgService struct {
	UserId  string
	Name    string
	Phone   string
	Avatar  string
	Address string
}

func NewUserMsgService(userId, name, phone, avatar, address string) *UserMsgService {
	return &UserMsgService{
		UserId:  userId,
		Name:    name,
		Phone:   phone,
		Avatar:  avatar,
		Address: address,
	}
}

// Create service
func CreateUserMsg(userId, name, phone, avatar, address string) error {
	return NewUserMsgService(userId, name, phone, avatar, address).add()
}

func UpdateUserMsg(userId, name, phone, avatar, address string) error {
	return NewUserMsgService(userId, name, phone, avatar, address).update()
}

func QueryUserMsg(userId string) (*repository.UserMsg, error) {
	um, err := repository.NewUserMsgDaoInstance().GetUserMsg(userId)
	if err != nil {
		return nil, err
	}
	return um, nil
}

func (f *UserMsgService) update() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.upUserMsgToDB(); err != nil {
		return err
	}
	return nil
}

func (f *UserMsgService) add() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.addUserMsgToDB(); err != nil {
		return err
	}
	return nil
}

func (f *UserMsgService) upUserMsgToDB() error {
	umsg := repository.UserMsg{
		UserId:  f.UserId,
		Name:    f.Name,
		Phone:   f.Phone,
		Avatar:  f.Avatar,
		Address: f.Address,
	}
	// 参数写入
	if err := repository.NewUserMsgDaoInstance().UpdateUMsg(umsg); err != nil {
		return err
	}

	return nil
}

func (f *UserMsgService) addUserMsgToDB() error {
	umsg := repository.UserMsg{
		UserId:    f.UserId,
		Name:      f.Name,
		Phone:     f.Phone,
		Avatar:    f.Avatar,
		Address:   f.Address,
		MsgStatus: int64(1),
	}
	// 参数写入
	if err := repository.NewUserMsgDaoInstance().CreateUMsg(umsg); err != nil {
		return err
	}
	return nil
}

// 验证输入数据的有效性
func (f *UserMsgService) checkParam() error {
	// 验证每个字段是否为非空字符串
	if f.UserId == "" {
		return errors.New("UserId cannot be empty")
	}
	if f.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if f.Avatar == "" {
		return errors.New("Avatar cannot be empty")
	}
	if f.Address == "" {
		return errors.New("Address cannot be empty")
	}

	// 验证 Phone 是否为合法电话
	if f.Phone == "" {
		return errors.New("Phone cannot be empty")
	}
	if !isValidPhone(f.Phone) {
		return errors.New("invalid Phone format")
	}
	return nil
}

// 验证 phone 格式是否有效
func isValidPhone(phone string) bool {
	// 假设电话是国际格式，如 "+1234567890" 或国内格式 "1234567890"
	phoneRegex := `^\+?[0-9]{10,15}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
