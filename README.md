```go
package main
// define your env variables
const shopId = "{YOUR_SHOP_ID}"
const AESKey = "{YOUR_AES_KEY}"
const MD5Key = "{YOUR_MD5_KEY}"
const ApiUrl = "{YOUR_API_URL}"

// new invoker
invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

// do anything
// example, register a mobile
invoker.RegAccount("012345678")
```

### APIList
```go
// join mobile to system
RegAccount(_mobile string)

// payment
UserRecharge(_mobile string, _price float64)

// withdraw
UserWithdraw(_mobile string, _price float64)

// get game list
// page from 0 to 1000, 20 records / page
GetGameList(_page uint32)

// bet to system
// mobile: who bet
// gameId: which game id from get game list api
// content: the sms content
UserBet(_mobile string, _gameId string, _content string)


// get user recharge logs
/**
type GetUserRechargeLogsQuery struct {
	StartTime string // Start Time, format: "YYYY-MM-DD HH:ii:ss"
	EndTime   string // End Time, format: "YYYY-MM-DD HH:ii:ss"
	Page      uint32 // Page id, from 0 to 100, default 0
	Count     uint32 // Count of per page, default 20, can not be 0, and cant larger than 100
}
 */
GetUserRechargeLogs(_mobile string, _queryOpt GetUserRechargeLogsQuery)
```

All Done, Good Job!!
