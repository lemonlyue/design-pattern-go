package simple_factory

import "fmt"

type Pay interface {
	Pay() string
}

type GooglePay struct {

}

func (*GooglePay) Pay() string {
	fmt.Println("Google pay success")
	return "Google Pay"
}

type PayPalPay struct {

}

func (*PayPalPay) Pay() string {
	fmt.Println("PayPal pay success")
	return "PayPal Pay"
}

func PayFactory(channel string) Pay {
	switch channel {
	case "Google":
		return &GooglePay{}
	case "PayPal":
		return &PayPalPay{}
	}

	return nil
}

