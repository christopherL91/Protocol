package Protocol

import (
	"regexp"
)

var (
	LoginStatusCode    uint16 = 1
	LogoutStatusCode   uint16 = 2
	DepositStatusCode  uint16 = 3
	WithdrawStatusCode uint16 = 4
	BalanceStatusCode  uint16 = 5

	ResponseOK              uint16 = 0
	ResponseAlreadyLoggedIn uint16 = 1
	ResponseNotAccepted     uint16 = 2
	ResponseMoneyProblem    uint16 = 3

	CardnumberTest = regexp.MustCompile(`^(\d){4}$`) //4 digits.
	PassnumberTest = regexp.MustCompile(`^(\d){2}$`) //2 digits.
	MoneyTest      = regexp.MustCompile(`^(\d+)$`)   //at least one digit.
	ScratchTest    = regexp.MustCompile(`^(\d+)$`)   //at least one digit.
)

type Menus map[string]Menu

type Menu struct {
	Menu    []string `json:"menu"`
	Login   []string `json:"log-in"`
	Banner  string   `json:"banner"`
	Invalid []string `json:"invalid-input"`
}

type MenuConfig struct {
	Menus
}

type Message struct {
	_struct bool   `codec:",toarray"`
	Code    uint16 //3 byte
	Number  uint16 //3byte
	Payload uint16 //3 byte
	//Total packet weight: 9byte.
}
