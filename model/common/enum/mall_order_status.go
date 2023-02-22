package enum

type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "Waiting to pay" //"待支付"
	case 1:
		return 1, "Paid for, awaiting buyback" //"待回购" //(已支付)
	case 2:
		return 2, "Distribution is complete" // "配货完成"
	case 3:
		return 3, "Product is out of stock" // "出库成功"
	case 4:
		return 4, "Transaction complete..." // 已经支付,并且确认回购 ,交易完成
	case 5:
		return 5, "Already paid back" //已经回款
	case -1:
		return -1, "Manual shutdown" //"手动关闭"
	case -2:
		return -2, "Timeout Shutdown" //"超时关闭"
	case -3:
		return -3, "Closing Merchant" // "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}
