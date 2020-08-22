package py

import (
	"github.com/plutov/paypal/v3"
)

var payCli *paypal.Client

func Init() {
	c, err := paypal.NewClient("AR2A6Ltm3BGecZxUoOfpSpelfEHAf0t3K0ENJCzasnCLNbssEyW70jKQJUDkaxy2CvyKasfQjwsz__Hk", "EMjBU4PwigeDvinU7gIk9NqzX4KjXL16YBV75HzztLXuDeBKiPQXKk6h2HJnsd9UdF3NgMQvHxCuFNLa", paypal.APIBaseSandBox)
	if err != nil {
		panic(err)
	}
	_, err = c.GetAccessToken()
	if err != nil {
		panic(err)
	}
	payCli = c
}

func GetDbCli() *paypal.Client {
	return payCli
}
