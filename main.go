package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	handler "github.com/Abhishek1833/link_in_bio/handler"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func init() {
	setUpViper()
	//models.ConnectDb()
	registerDatabase()
}
func main() {
	r := mux.NewRouter()
	/*r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/signup", handler.Register)*/
	r.HandleFunc("/addlink", handler.AddLink)
	r.HandleFunc("/delete", handler.DeleteLink)
	r.HandleFunc("/registeruser", handler.GetUrl)
	r.HandleFunc("/fetch", handler.GetLink)
	log.Fatal(http.ListenAndServe(":3102", r))
}

func registerDatabase() {
	runmode := cast.ToString(viper.Get("runmode"))
	mysql := viper.Get(runmode + ".mysql").(map[string]interface{})
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
	log.Println("conf", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Kolkata")
	orm.Debug = true

}

//set up config file from conf folder
func setUpViper() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
}
