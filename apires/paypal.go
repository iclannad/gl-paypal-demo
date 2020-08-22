package apires

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/plutov/paypal/v3"
	"gl-paypal-demo/py"
	"net/http"
)

//回调(可以利用上面的回调链接实现)
func PaypalCallback(orderId string) error {
	c := py.GetDbCli()
	ctor := paypal.CaptureOrderRequest{}

	order, err := c.CaptureOrder(orderId, ctor)
	if err != nil {
		fmt.Println("CaptureOrder", err)
		return err
	}
	//查看回调完成后订单状态是否支付完成。
	fmt.Println("(*order).Status", (*order).Status)
	if (*order).Status != "COMPLETED" {
		return errors.New("pay fail")
	}
	return nil
}

func Success(c echo.Context) error {
	fmt.Println("paypal success")
	token := c.QueryParam("token")
	PayerID := c.QueryParam("PayerID")
	fmt.Println(token, "   ", PayerID)

	fmt.Println("executing...")

	err := PaypalCallback(token)
	fmt.Println(err)

	return c.JSON(http.StatusOK, "ok")
}

func Cancel(c echo.Context) error {

	fmt.Println("paypal canel")

	return c.JSON(http.StatusOK, "ok")
}

func Create(cecho echo.Context) error {
	fmt.Println("paypal create order")
	c := py.GetDbCli()
	purchaseUnits := make([]paypal.PurchaseUnitRequest, 1)

	//
	items := make([]paypal.Item, 0)
	item := paypal.Item{
		Name: "流量包 1G",
		UnitAmount: &paypal.Money{
			Currency: "USD",
			Value:    "1",
		},
		Tax: &paypal.Money{
			Currency: "USD",
			Value:    "0.1",
		},
		Quantity: "2",
	}
	items = append(items, item)

	purchaseUnits[0] = paypal.PurchaseUnitRequest{
		Amount: &paypal.PurchaseUnitAmount{
			Currency: "USD",
			Value:    "2",
			Breakdown: &paypal.PurchaseUnitAmountBreakdown{
				ItemTotal: &paypal.Money{
					Currency: "USD",
					Value:    "2",
				},
				TaxTotal: &paypal.Money{
					Currency: "USD",
					Value:    "0.2",
				},
				Discount: &paypal.Money{
					Currency: "USD",
					Value:    "0.2",
				},
			},
		},

		Items: items,
	}
	payer := &paypal.CreateOrderPayer{}
	appContext := &paypal.ApplicationContext{
		ReturnURL: "http://45.76.223.152:8778/pay/success", //回调链接
		CancelURL: "http://45.76.223.152:8778/pay/cancel",
	}
	order, err := c.CreateOrder("CAPTURE", purchaseUnits, payer, appContext)
	if err != nil {
		fmt.Println("create order errors:", err)
	}

	for _, link := range order.Links {
		fmt.Println(link.Rel)
		fmt.Println(link.Method)
		fmt.Println(link.Href)
		fmt.Println(link.Enctype)
		fmt.Println(link.Description)
	}

	return cecho.JSON(http.StatusOK, "ok")
}
