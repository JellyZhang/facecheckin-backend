package serializer

import "facecheckin/model"

// User 用户序列化器
type User struct {
	PhoneNumber  string `json:"phone_number"`
	UserName     string `json:"user_name"`
	Sex          uint8  `json:"sex"`
	Organization string `json:"organization"`
	Face         string `json:"face"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		PhoneNumber:  user.PhoneNumber,
		UserName:     user.UserName,
		Sex:          user.Sex,
		Face:         user.Face,
		Organization: user.Organization,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
