package model

/////////old struct
// type Children struct {
// 	Path      string `json:"path"`
// 	Name      string `json:"name"`
// 	Component string `json:"component"`
// 	Meta      struct {
// 		Title       string   `json:"title"`
// 		IsLink      string   `json:"isLink"`
// 		IsHide      bool     `json:"isHide"`
// 		IsKeepAlive bool     `json:"isKeepAlive"`
// 		IsAffix     bool     `json:"isAffix"`
// 		IsIframe    bool     `json:"isIframe"`
// 		Roles       []string `json:"roles"`
// 		Icon        string   `json:"icon"`
// 	} `json:"meta"`
// 	Redirect string `json:"redirect,omitempty"`
// 	Children []Children `json:"children,omitempty"`
// }
// type Routes []struct { //r[0].Children 有效值...
// 	Path      string `json:"path"`
// 	Name      string `json:"name"`
// 	Component string `json:"component"`
// 	Redirect  string `json:"redirect"`
// 	Meta      struct {
// 		IsKeepAlive bool `json:"isKeepAlive"`
// 	} `json:"meta"`
// 	Children []Children `json:"children,omitempty"`
// }
////end old type

type Menus struct {
	M_id      int64  `json:"id"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Meta      struct {
		Title       string   `json:"title"`
		IsLink      string   `json:"isLink"`
		IsHide      bool     `json:"isHide"`
		IsKeepAlive bool     `json:"isKeepAlive"`
		IsAffix     bool     `json:"isAffix"`
		IsIframe    bool     `json:"isIframe"`
		Roles       []string `json:"roles"`
		Icon        string   `json:"icon"`
	} `json:"meta"`
	Redirect string  `json:"redirect,omitempty"`
	Children []Menus `json:"children,omitempty"`
}
