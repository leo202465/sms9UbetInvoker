package sms9UbetInvoker

import (
	"fmt"
	"log"
	"testing"
)

const shopId = "49361"
const ApiUrl = "https://sms-api-staging.9ubet.com"
const AESKey = "VecsTwFCcg034ZsK8Oo7t0nWouxHJPgh"
const MD5Key = "9ahzccEk9EvKTPD67PC2Vp95Yz8sVMr5"

func TestInvoker_RegAccount(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.RegAccount("012345671")
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v", respObj.GetData())
}

func TestInvoker_GetAccountBalance(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.GetAccountBalance("012345671")
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_UserRecharge(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.UserRecharge("012345671", 100)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_UserWithdraw(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.UserWithdraw("012345671", 100)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_GetGameList(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.GetGameList(0)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}
