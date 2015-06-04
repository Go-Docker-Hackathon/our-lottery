/*
*需求：根据金数据表单的提交数据建立一份名单(每行一个人名，可以手动建立)
*然后编写一个程序用于抽奖。
*功能：
*1.基于WEB，便于演示；
*2.有启停按钮，点击启动按钮则开始极速循环遍历名单(人名或序号)
*3.需要实时列有中奖名单
*4.选出10人(可配置)后启停按钮不可用，抽奖结束
*5.程序需纯go开发，放在github里
*@author xialingsc
*@data 2015-06-04
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//	"math/rand"
	//	"time"
)

const (
	LUCK_PERSON_NUMBER = 10
)

var (
	luckPersonList []Person //获奖列表
	allPersonList  []Person //参与抽奖人初始列表

)

type MainController struct {
	beego.Controller
}

type Person struct {
	SerialNumber int    `json:"serial,omitempty"` //序号
	Name         string `json:"name,omitempty"`   //姓名
}

//用于通过json返回web页面
type ScrollTempResult struct {
	TQueuePersonList []Person //放参与抽奖人数列表
	LPersonList      []Person //放中奖人数列表

}

//type Team struct {
//	Person
//	number int
//}

//27个参赛人员+8个志愿者
func init() {
	allPersonList = []Person{Person{1, "张三"}, Person{2, "李..."}, Person{3, "王..."}, Person{4, "赵..."}, Person{5, "钱.."}, Person{6, "徐.."}, Person{7, "相.."}, Person{8, "毛..."}, Person{9, "杜..."}, Person{10, "a"}, Person{11, "b"}, Person{12, "c"}, Person{13, "d"}, Person{14, "e"}, Person{15, "f"}, Person{16, "g"}, Person{17, "h"}, Person{18, "i"}, Person{19, "g"}, Person{20, "k"}, Person{21, "l"}, Person{22, "m"}, Person{23, "n"}, Person{24, "o"}, Person{25, "p"}, Person{26, "q"}, Person{27, "r"}, Person{28, "s"}, Person{29, "t"}, Person{30, "u"}, Person{31, "v"}, Person{32, "w"}, Person{33, "x"}, Person{34, "y"}, Person{35, "z"}}
	luckPersonList = []Person{}
}

//放入获奖列表
//func Push(p Person) []Person {
//	//	i := len(LuckPersonList)
//	luckPersonList = append(luckPersonList, p)
//	//		fmt.Errorf(" parameter p is nil ....")
//	return luckPersonList
//}

func (this *MainController) Push() {
	var p Person
	var sTemp ScrollTempResult
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &p)
	fmt.Println(p)
	luckPersonList = append(luckPersonList, p)
	if err != nil {
		this.Data["json"] = -1
	} else {
		sTemp.TQueuePersonList = GetQueuePersonList()
		//fmt.Println(len(sTemp.TQueuePersonList))
		sTemp.LPersonList = luckPersonList
		this.Data["json"] = sTemp
	}
	this.ServeJson()

}

//测试展现列表
func DisplayLuckPersonList() {
	if len(luckPersonList) > 0 {
		for i := 0; i < len(luckPersonList); i++ {
			fmt.Println(luckPersonList[i].Name)
		}
	} else {
		fmt.Println("获奖列表为空")
	}
}

func DisplayList(list []Person) {
	if len(list) > 0 {
		for i := 0; i < len(list); i++ {
			fmt.Println(list[i].Name)
		}
	} else {
		fmt.Println("获奖列表为空")
	}
}

//获取参与抽奖的列表,排除已抽中人员
func GetQueuePersonList() []Person {
	var tempQueuePersonList []Person = []Person{} //变化列表
	if len(luckPersonList) == 0 {                 //没有中奖人员
		return allPersonList
	} else { //若有中奖人员，则排除已中奖人员
		for i := 0; i < len(allPersonList); i++ {
			flag := false
			for j := 0; j < len(luckPersonList); j++ {
				if luckPersonList[j].SerialNumber == allPersonList[i].SerialNumber {
					flag = false
					break
				} else {
					flag = true
				}
			}
			if flag {
				tempQueuePersonList = append(tempQueuePersonList, allPersonList[i])
			}
		}
		return tempQueuePersonList

	}
}

func (c *MainController) Get() {

	c.Data["QueuePersonList"] = GetQueuePersonList()
	c.Data["LuckPersonList"] = luckPersonList
	c.Data["SetLuckPersonNumber"] = LUCK_PERSON_NUMBER
	c.TplNames = "index.tpl"
}

////获取随机数
//func GetRandomNumber(min int, max int) int {
//	//	r := rand.New(rand.NewSource(35))
//	rand.Seed(time.Now().UTC().UnixNano())
//	//fmt.Println((r.Intn(100) * (max - min + 1)))
//	return (min + rand.Intn(max-min))
//	//	return (r.Intn(35))
//}
