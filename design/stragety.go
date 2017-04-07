package design

import (
	"errors"
	"fmt"
	"math"
)

func CheckError() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，bug
		}
	}()
}

//收费
type CashSuper interface {
	AcceptCash(money float64) float64
}

//正常收费
type CashNormal struct {
}

func (c CashNormal) AcceptCash(money float64) float64 {
	return money
}

//打折收费
type CashDiscount struct {
	Discount float64
}

func (c CashDiscount) AcceptCash(money float64) float64 {
	return c.Discount * money
}

//满减如满300减100 moneycondition为300 moneyreturn为100
type CashReturn struct {
	MoneyCondition float64
	MoneyReturn    float64
}

func (c CashReturn) AcceptCash(money float64) float64 {
	if money > c.MoneyCondition {
		result := money - math.Floor(money/c.MoneyCondition)*c.MoneyReturn
		return result
	}
	return money
}

type CashContext struct {
	Cs CashSuper
}

func GetCashContext(what string, data ...float64) (ctx CashContext, err error) {
	//不定参数解决struct的初始化传参问题
	var mycs CashSuper
	switch what {
	case "normal":
		mycs = CashNormal{}
	case "discount":
		if len(data) < 1 {
			return ctx, errors.New("need discount ")
		}
		mycs = CashDiscount{Discount: data[0]}
	case "return":
		if len(data) < 2 {
			return ctx, errors.New("need moneycondition and moneyreturn ")
		}
		mycs = CashReturn{MoneyCondition: data[0], MoneyReturn: data[1]}
	default:
		mycs = CashNormal{}
	}
	return CashContext{Cs: mycs}, nil
}
func (c CashContext) GetResult(money float64) float64 {
	return c.Cs.AcceptCash(money)
}
