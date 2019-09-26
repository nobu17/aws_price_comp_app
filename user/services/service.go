package services

import (
	"common/log"
	"fmt"
	"user/repositories"
)

// alertService serivce
type userService struct {
	logger      log.LoggerImpl
	repository  repositories.UserInfoImpl
	grepository repositories.GroupInfoImpl
}

// NewUserService constructor
func NewUserService(logger log.LoggerImpl, repository repositories.UserInfoImpl, grepository repositories.GroupInfoImpl) ServiceImpl {
	return &userService{
		logger:      logger,
		repository:  repository,
		grepository: grepository,
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

// GetUserInfo impl
func (u *userService) PutItemGroup(req PutItemGroupInputModel) (PutItemGroupOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start Service:PutItemGroup:", req)

	var itemGroup = make([]repositories.ItemGroup, 0)
	for _, group := range req.GroupList {
		itemGroup = append(itemGroup, repositories.NewItemGroup(group.GroupID, group.GroupName))
	}
	var input = repositories.PutItemGroupRequest{UserID: req.UserID, GroupList: itemGroup}
	res, err := u.repository.PutItemGroup(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end Service:PutItemGroup:")
		return PutItemGroupOutputModel{}, err
	}
	// put a result
	var output = PutItemGroupOutputModel{}
	output.SuccessItemGroupList = res.SuccessItemGroupList
	output.FailedPutGroupList = res.FailedPutGroupList

	u.logger.LogWriteWithMsgAndObj(log.Info, "end Service:DeleteItemGroup:", output)

	return output, nil
}

// GetUserInfo impl
func (u *userService) DeleteItemGroup(req DeleteInputModel) (DeleteOutputModel, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start Service:DeleteItemGroup:", req)

	// delete group master at first
	var input = repositories.NewDeleteItemGroupRequest(req.UserID, req.GroupIDList)
	res, err := u.repository.DeleteItemGroup(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end Service:DeleteItemGroup:")
		return DeleteOutputModel{}, err
	}
	var output = DeleteOutputModel{}
	output.FailedDeleteGroupList = res.FailedItemGroupList
	// filter only success groupID list
	input = repositories.NewDeleteItemGroupRequest(req.UserID, res.SuccessItemGroupList)

	// then delete item master
	gres, err := u.grepository.DeleteItemGroup(input)
	if err != nil {
		u.logger.LogWrite(log.Error, "repository retrun error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end Service:DeleteItemGroup:")
		return DeleteOutputModel{}, err
	}
	// set a result
	output.SuccessItemGroupList = gres.SuccessItemGroupList
	output.FailedDeleteItemList = gres.FailedItemGroupList
	u.logger.LogWriteWithMsgAndObj(log.Info, "end Service:DeleteItemGroup:", output)

	return output, nil
}
