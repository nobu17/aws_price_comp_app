package controllers

import (
	"common/log"
	"errors"
	"fmt"
	"strings"
	"user/services"
)

// UserController controller
type UserController struct {
	logger  log.LoggerImpl
	service services.ServiceImpl
}

// NewUserController constroctor
func NewUserController() UserController {
	container := Container{}
	ser := container.NewService()
	log := container.NewLogger()
	return UserController{
		logger:  log,
		service: ser,
	}
}

// GetUserInfo get user
func (u *UserController) GetUserInfo(req GetRequest) (GetResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start UserController:GetUserInfo", req)

	err := u.validateGetReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:GetUserInfo")
		return GetResponce{}, err
	}

	var inputModel = services.NewGetInputModel(req.UserID, req.Password)
	res, err := u.service.GetUserInfo(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:GetUserInfo")
		return GetResponce{}, err
	}

	user := NewUserInfo(res.UserInfo.UserID, res.UserInfo.Name, res.UserInfo.Mail)

	list := make([]ItemGroup, 0)
	for _, item := range res.ItemGroupList {
		list = append(list, NewItemGroup(item.GroupID, item.GroupName))
	}

	var respo = GetResponce{UserInfo: user, ItemGroupList: list}
	u.logger.LogWriteWithMsgAndObj(log.Info, "end UserController:GetUserInfo:", respo)
	return respo, nil
}

// PutItemGroup put group for specified user user
func (u *UserController) PutItemGroup(req PutGroupRequest) (PutGroupResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start UserController:PutItemGroup", req)

	err := u.validatePutGroupReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:PutItemGroup")
		return PutGroupResponce{}, err
	}
	var itemGroup = make([]services.ItemGroup, 0)
	for _, group := range req.GroupList {
		itemGroup = append(itemGroup, services.NewItemGroup(group.GroupID, group.GroupName))
	}

	var inputModel = services.PutItemGroupInputModel{ UserID: req.UserID, GroupList: itemGroup}
	res, err := u.service.PutItemGroup(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:PutItemGroup")
		return PutGroupResponce{}, err
	}

	var respo = PutGroupResponce{ SuccessItemGroupList: res.SuccessItemGroupList,  FailedPutGroupList: res.FailedPutGroupList }
	u.logger.LogWriteWithMsgAndObj(log.Info, "end UserController:PutItemGroup:", respo)
	return respo, nil
}

// DeleteItemGroup delete group for specified user user
func (u *UserController) DeleteItemGroup(req DeleteGroupRequest) (DeleteGroupResponce, error) {
	u.logger.LogWriteWithMsgAndObj(log.Info, "start UserController:DeleteItemGroup", req)

	err := u.validateDeleteGroupReq(req)
	if err != nil {
		u.logger.LogWrite(log.Info, "input error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:DeleteItemGroup")
		return DeleteGroupResponce{}, err
	}

	var inputModel = services.NewDeleteInputModel(req.UserID, req.GroupIDList)
	res, err := u.service.DeleteItemGroup(inputModel)
	if err != nil {
		u.logger.LogWrite(log.Info, "servie error:"+fmt.Sprint(err))
		u.logger.LogWrite(log.Info, "end UserController:DeleteItemGroup")
		return DeleteGroupResponce{}, err
	}

	var respo = NewDeleteGroupResponce(res.SuccessItemGroupList, res.FailedDeleteGroupList, res.FailedDeleteItemList)
	u.logger.LogWriteWithMsgAndObj(log.Info, "end UserController:DeleteItemGroup:", respo)
	return respo, nil
}


func (u *UserController) validatePutGroupReq(req PutGroupRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupList == nil || len(req.GroupList) == 0 {
		return errors.New("GroupList is empty")
	}
	// dupicated check
	dup := make(map[string]string)
	for _, item := range req.GroupList {
		if(strings.TrimSpace(item.GroupID) == "") {
			return errors.New("empty groupID is included")
		}
		if(strings.TrimSpace(item.GroupName) == "") {
			return errors.New("empty groupName is included")
		}
		_, existed := dup[item.GroupID]
		if existed {
			return errors.New("Same Product ID set:" + item.GroupID)
		}
		dup[item.GroupID] = ""
	}

	return nil
}

func (u *UserController) validateDeleteGroupReq(req DeleteGroupRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.GroupIDList == nil || len(req.GroupIDList) == 0 {
		return errors.New("GroupIDList is empty")
	}
	// dupicated check
	dup := make(map[string]string)
	for _, item := range req.GroupIDList {
		if(strings.TrimSpace(item) == "") {
			return errors.New("empty groupID is included")
		}
		_, existed := dup[item]
		if existed {
			return errors.New("Same Product ID set:" + item)
		}
		dup[item] = ""
	}

	return nil
}


func (u *UserController) validateGetReq(req GetRequest) error {
	if req.UserID == "" || strings.TrimSpace(req.UserID) == "" {
		return errors.New("UserID is empty")
	}
	if req.Password == "" || strings.TrimSpace(req.Password) == "" {
		return errors.New("Password is empty")
	}

	return nil
}
