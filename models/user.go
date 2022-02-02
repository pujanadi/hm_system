package user

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          string `orm: "pk" json: "id"`
	DisplayName string `orm: "display_name json: "display_name"`
	UserLevel   string `orm: "level" json: "level"`
	LoginFlag   string `orm: "is_login: json: "is_login"`
}

type Auth struct {
	userid   string `orm: "pk" json: "userid"`
	password string `orm: password`
}

func init() {
	prefix := "fm_"
	orm.RegisterModelWithPrefix(prefix, new(User))
	orm.RegisterModelWithPrefix(prefix, new(Auth))

	orm.RegisterDataBase("default", "mysql", "root:root@/FP_HMS?charset=utf8")
}
