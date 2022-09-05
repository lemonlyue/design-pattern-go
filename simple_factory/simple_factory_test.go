package simple_factory

import "testing"

func TestGooglePayPay(t *testing.T) {
	factory := PayFactory("Google")
	channel := factory.Pay()
	if channel != "Google Pay" {
		t.Fatal("TestGooglePayPay test fail")
	}
}

func TestPayPalPay(t *testing.T) {
	factory := PayFactory("PayPal")
	channel := factory.Pay()
	if channel != "PayPal Pay" {
		t.Fatal("TestGooglePayPay test fail")
	}
}
