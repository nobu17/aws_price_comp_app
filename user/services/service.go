package services

import (
	"common/log"
	"fmt"
	"user/repositories"
)

// alertService serivce
type userService struct {
	logger     log.LoggerImpl
	repository repositories.UserInfoImpl
}

// NewUserService constructor
func NewUserService(logger log.LoggerImpl, repository repositories.UserInfoImpl) ServiceImpl {
	return &userService{
		logger:     logger,
		repository: repository,
	}
}

// GetUserInfo impl
func (u *userService) GetUserInfo(req GetInputModel) (GetOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start Service:GetUserInfo:", req)

	var input = repositories.NewGetRequest(req.UserID, req.Password)
	res, err := u.repository.GetUserInfo(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end Service:GetUserInfo:")
		return GetOutputModel{}, err
	}

	user := NewUserInfo(res.UserInfo.UserID, res.UserInfo.Name, res.UserInfo.Mail)

	list := make([]ItemGroup, 0)
	for _, item := range res.ItemGroupList {
		list = append(list, NewItemGroup(item.GroupID, item.GroupName))
	}

	var output = GetOutputModel{UserInfo: user, ItemGroupList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end Service:GetUserInfo:", output)

	return output, nil
}
