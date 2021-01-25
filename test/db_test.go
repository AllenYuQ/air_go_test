package test

import (
	"fmt"
	"task5/models"
	"testing"
)

//进行单元测试时，需要将setting/setting.go的init() 方法中
//Cfg, err = ini.Load("conf/app.ini") 改为 Cfg, err = ini.Load("../conf/app.ini")
func TestDB(t *testing.T) {
	//测试数据库连接是否成功 username 用户名 password 密码
	isExist := models.CheckUser("admin", "admin123")
	fmt.Println(isExist)
	//测试增加用户
	data := make(map[string]interface{})
	data["id"] = 3
	data["username"] = "test"
	data["realname"] = "testrealname"
	data["password"] = "test123"
	data["email"] = "2969141711@qq.com"
	data["phone"] = "18816209310"
	data["sex"] = 1
	data["roleId"] = 1
	canAdd := models.AddUser(data)
	fmt.Println(canAdd)
	//列出所有的用户
	var users []models.User
	users = models.ListUsers()
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestMlmet(t *testing.T) {
	mlmets := models.ListMlmets()
	fmt.Println(len(mlmets))
	for i, mlmet := range mlmets {
		fmt.Println(i)
		fmt.Println(mlmet)
	}
	fmt.Println(mlmets)
}

func TestMlout(t *testing.T) {
	mlouts := models.ListMloutsBetweenInterval("24")
	fmt.Println(mlouts)
}

func TestListPosition(t *testing.T) {
	positions := models.ListPostions()
	fmt.Println(positions)
	fmt.Println(len(positions))
	for _, str := range positions {
		switch str {
		case positions[0]:
			fmt.Println(positions)
			break
		default:
			fmt.Println("1")
		}
	}
}

func TestMloutTypeList(t *testing.T) {
	mloutTypes := models.ListMloutTypesBetweenInterval("2")
	for i := 0; i < len(mloutTypes); i++ {
		fmt.Println(mloutTypes[i])
	}
}

func TestSubstancesList(t *testing.T) {
	substances := models.ListSubstances()
	fmt.Println(substances)
}

func TestRechargeHistory(t *testing.T) {
	fmt.Println(models.RechargeHistory("111"))
}

func TestRecharge(t *testing.T) {
	fmt.Println(models.Recharge("1", "1"))
}
