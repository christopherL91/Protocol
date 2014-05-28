package Protocol

type Menus map[string]Menu

type Menu struct {
	Menu []string `json:menu`
}

type MenuConfig struct {
	Menus
}
