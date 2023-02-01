package manage

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/enum"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	manageRes "main.go/model/manage/response"
	"strconv"
	"time"
)

type ManageOrderService struct {
}

// CheckDone 修改订单状态为配货成功
func (m *ManageOrderService) CheckDone(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 2, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功无法执行出库操作")
		}
	}
	return
}

// CheckOut 出库
func (m *ManageOrderService) CheckOut(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() && order.OrderStatus != enum.ORDER_PACKAGED.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 3, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功或配货完成无法执行出库操作")
		}
	}
	return
}

func (m *ManageOrderService) UpdateOrder(req manageReq.OrderStatusParam) (err error) {
	//var adminUserToken manage.MallAdminUserToken
	//err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	//if err != nil {
	//	return errors.New("不存在的token  " + err.Error())
	//}
	//根据订单号 查询出 这个订单的持有人 user_id ,更新订单状态成功后，需要修改订单持有人的余额 (订单总价 * 120%)

	var order manage.MallOrder
	err = global.GVA_DB.Where("order_no =? ", req.OrderNo).First(&order).Error
	if err != nil {
		fmt.Println(00000000000000000000)
		return errors.New("订单不存在  " + err.Error())
	}

	//查询订单持有用户
	var user manage.MallUser
	err = global.GVA_DB.Where("user_id =? ", order.UserId).First(&user).Error
	if err != nil {
		fmt.Println(3333333333333)
		return errors.New("订单持有用户不存在： " + err.Error())
	}
	//更新订单状态
	err = global.GVA_DB.Where("order_no = ?", req.OrderNo).Updates(&manage.MallOrder{
		OrderStatus: 5,
	}).Error

	if err != nil {
		fmt.Println(434343434343)
		return errors.New("更新订单状态失败  " + err.Error())
	}
	money, e := strconv.Atoi(req.OrderMoney)
	if e != nil {
		fmt.Println(44444444444444)
		return errors.New("填写的金额有误  " + err.Error())
	}

	//更新余额
	user.UserMoney = money + user.UserMoney
	err = global.GVA_DB.Where("user_id = ?", order.UserId).Updates(&user).Error

	if err != nil {
		fmt.Println(5555555555555)
		err = global.GVA_DB.Where("order_no = ?", req.OrderNo).Updates(&manage.MallOrder{
			OrderStatus: 4,
		}).Error

		return
	}
	fmt.Println(7676767676)
	return
}

// CloseOrder 商家关闭订单
func (m *ManageOrderService) CloseOrder(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus == enum.ORDER_SUCCESS.Code() || order.OrderStatus < 0 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: enum.ORDER_CLOSED_BY_JUDGE.Code(), UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单不能执行关闭操作")
		}
	}
	return
}

// GetMallOrder 根据id获取MallOrder记录
func (m *ManageOrderService) GetMallOrder(id string) (err error, newBeeMallOrderDetailVO manageRes.NewBeeMallOrderDetailVO) {
	var newBeeMallOrder manage.MallOrder
	if err = global.GVA_DB.Where("order_id = ?", id).First(&newBeeMallOrder).Error; err != nil {
		return
	}
	var orderItems []manage.MallOrderItem
	if err = global.GVA_DB.Where("order_id = ?", newBeeMallOrder.OrderId).Find(&orderItems).Error; err != nil {
		return
	}
	//获取订单项数据
	if len(orderItems) > 0 {
		var newBeeMallOrderItemVOS []manageRes.NewBeeMallOrderItemVO
		copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
		copier.Copy(&newBeeMallOrderDetailVO, &newBeeMallOrder)

		_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.OrderStatus)
		_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.PayType)
		newBeeMallOrderDetailVO.OrderStatusString = OrderStatusStr
		newBeeMallOrderDetailVO.PayTypeString = payTapStr
		newBeeMallOrderDetailVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
	}
	return
}

// GetMallOrderInfoList 分页获取MallOrder记录
func (m *ManageOrderService) GetMallOrderInfoList(info request.PageInfo, orderNo string, orderStatus string, token string) (err error, list interface{}, total int64) {
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的token  " + err.Error()), list, total
	}
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallOrder{})
	if orderNo != "" {
		db.Where("order_no", orderNo)
	}
	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功(申请回购，) 5.回购完成(已经退款) -1.手动关闭 -2.超时关闭 -3.商家关闭
	if orderStatus != "" {
		status, _ := strconv.Atoi(orderStatus)
		db.Where("order_status", status)
	}
	var mallOrders []manage.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if adminUserToken.AgentId == "8888" { //8888是最高管理权限
		err = db.Limit(limit).Offset(offset).Order("update_time desc").Find(&mallOrders).Error
	} else {
		err = db.Limit(limit).Offset(offset).Where("agent_id", adminUserToken.AgentId).Order("update_time desc").Find(&mallOrders).Error
	}

	return err, mallOrders, total
}

// 获取订单记录 v2-包括订单详情信息
func (m *ManageOrderService) GetMallOrderInfoListV2(info request.PageInfo, orderNo string, orderStatus string, token string) (err error, list interface{}, total int64) {
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的token  " + err.Error()), list, total
	}
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallOrder{})
	if orderNo != "" {
		db.Where("order_no", orderNo)
	}
	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功(申请回购，) 5.回购完成(已经退款) -1.手动关闭 -2.超时关闭 -3.商家关闭
	if orderStatus != "" {
		status, _ := strconv.Atoi(orderStatus)
		db.Where("order_status", status)
	}

	var mallOrders []manage.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if adminUserToken.AgentId == "8888" { //8888是最高管理权限
		err = db.Limit(limit).Offset(offset).Order("update_time desc").Find(&mallOrders).Error
	} else {
		err = db.Limit(limit).Offset(offset).Where("agent_id", adminUserToken.AgentId).Order("update_time desc").Find(&mallOrders).Error
	}

	//获取订单详情

	for _, value := range mallOrders {
		var orderItems []manage.MallOrderItem
		if err = global.GVA_DB.Where("order_id = ?", value.OrderId).Find(&orderItems).Error; err != nil {
			return
		}
		//获取订单项数据
		//if len(orderItems) > 0 {
		//	var newBeeMallOrderItemVOS []manageRes.NewBeeMallOrderItemVO
		//	copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
		//	copier.Copy(&newBeeMallOrderDetailVO, &value)
		//
		//	_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.OrderStatus)
		//	_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.PayType)
		//	newBeeMallOrderDetailVO.OrderStatusString = OrderStatusStr
		//	newBeeMallOrderDetailVO.PayTypeString = payTapStr
		//	newBeeMallOrderDetailVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
		//}

	}

	return err, mallOrders, total
}

// 获取回购记录
func (m *ManageOrderService) GetMallOrderBuyBackList(info request.PageInfo, orderNo string, orderStatus string, token string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的token  " + err.Error()), list, total
	}
	// 创建db
	db := global.GVA_DB.Model(&manage.MallOrder{})
	if orderNo != "" {
		db.Where("order_no", orderNo)
	}
	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功(申请回购，) 5.回购完成(已经退款) -1.手动关闭 -2.超时关闭 -3.商家关闭
	if orderStatus != "" {
		status, _ := strconv.Atoi(orderStatus)
		db.Where("order_status", status)
	}
	var mallOrders []manage.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if adminUserToken.AgentId == "8888" { //8888是最高管理权限
		err = db.Limit(limit).Offset(offset).Order("update_time desc").Find(&mallOrders).Error
	} else {
		err = db.Limit(limit).Offset(offset).Where("agent_id", adminUserToken.AgentId).Order("update_time desc").Find(&mallOrders).Error
	}

	return err, mallOrders, total
}
