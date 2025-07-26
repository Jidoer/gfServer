package vars

import(
	"github.com/gogf/gf/v2/net/ghttp"
)

var Session_ *ghttp.Session

func SetSession(s *ghttp.Session) {
	Session_ = s
}

func GetSession() *ghttp.Session {
	return Session_
}