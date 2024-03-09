package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"naufalzp": {
		AuthToken: "123456",
		Username:  "naufalzp",
	},
	"aurelwyt": {
		AuthToken: "654321",
		Username:  "aurelwyt",
	},
	"johndoe": {
		AuthToken: "123123",
		Username:  "johndoe",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"naufalzp": {
		Coins:    100,
		Username: "naufalzp",
	},
	"aurelwyt": {
		Coins:    200,
		Username: "aurelwyt",
	},
	"johndoe": {
		Coins:    300,
		Username: "johndoe",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
