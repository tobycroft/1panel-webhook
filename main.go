package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/config/app_conf"
	"main.go/route"
	"os"
	"time"
)

func init() {
	time.Local = app_conf.TimeZone
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
	s, err := os.Stat("./exec/")
	if err != nil {
		os.Mkdir("./exec", 0755)
	} else if s.IsDir() {
		os.Mkdir("./exec", 0755)
	}
}

func main() {
	fmt.Println(os.Args)
	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":888")

}
