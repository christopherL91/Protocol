package Protocol

type Menus map[string]Menu

type Menu struct {
	Menu    []string `json:"menu"`
	Login   []string `json:"log-in"`
	Invalid []string `json:"invalid-input"`
}

type MenuConfig struct {
	Menus
}

type Message struct {
	Number   uint16 `codec:",omitempty"` //3byte
	Pass     uint16 `codec:",omitempty"` //3 byte
	LoggedIn bool   `codec:",omitempty"` // 1 byte
}
