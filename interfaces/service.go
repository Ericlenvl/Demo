package interfaces

type UserInfo struct {
	UserId string
}

type UserService interface {
	GetUserInfo(id string) (userInfos []UserInfo,err error)
}
