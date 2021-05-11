package main

import "C"
import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pelletier/go-toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库Monitor对象
type Monitor struct {
	Id           int    `gorm:"column:Id"`
	MemTotal     string `gorm:"column:MemTotal"`
	MemFree      string `gorm:"column:MemFree"`
	MemAvailable string `gorm:"column:MemAvailable"`
	Type         string `gorm:"column:Type"`
	Total        string `gorm:"column:Total"`
	Used         string `gorm:"column:Used"`
	Free         string `gorm:"column:Free"`
	UsedPercent  string `gorm:"column:UsedPercent"`
	Uptime       string `gorm:"column:Uptime"`
	Payload1     string `gorm:"column:Payload1"`
	Payload5     string `gorm:"column:Payload5"`
	Payload15    string `gorm:"column:Payload15"`
}

func (Monitor) TableName() string {
	return "Monitor"
}

// 数据库PC对象
type PC struct {
	Id     int `gorm:"primaryKey"`
	CPU    string
	Core   int
	Thread int
}

func (PC) TableName() string {
	return "PC"
}

// 定义连接用的参数
var (
	opts      *mqtt.ClientOptions
	conn      mqtt.Client
	token     mqtt.Token
	subscribe []string
	id        int
)

// 用于处理接收到信息的方法
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	message := byteString(msg.Payload()) // 将获取到的字节数组转为字符串
	strArr := strings.Fields(message)    // 将字符串按照空格分隔

	// gorm连接数据库
	dsn := "mqtt:186536_Wlj@tcp(47.104.253.11:3306)/mqtt?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if len(strArr) == 13 { // 如果分割出来的块数是13块
		id, _ := strconv.Atoi(strArr[0])
		// 初始化一行数据
		monitor := Monitor{
			Id:           id,
			MemTotal:     strArr[1],
			MemFree:      strArr[2],
			MemAvailable: strArr[3],
			Type:         strArr[4],
			Total:        strArr[5],
			Used:         strArr[6],
			Free:         strArr[7],
			UsedPercent:  strArr[8],
			Uptime:       strArr[9],
			Payload1:     strArr[9],
			Payload5:     strArr[10],
			Payload15:    strArr[11],
		}
		result := db.Create(&monitor)    // 插入数据到数据库
		fmt.Println(result.RowsAffected) // 返回插入的条目数
	} else if len(strArr) <= 6 { // 如果分割出来的块数是6块
		strArr = strings.Split(message, ",")
		id, _ := strconv.Atoi(strArr[0])
		core, _ := strconv.Atoi(strArr[2])
		thread, _ := strconv.Atoi(strArr[3])
		pc := PC{
			Id:     id,
			CPU:    strArr[1],
			Core:   core,
			Thread: thread,
		}
		result := db.Create(&pc)
		fmt.Println(result.RowsAffected)
	} else { // 如果分割出来的块数是其他块
		fmt.Println(message)
		panic("数据不正确")
	}
}

// byte转字符串方法
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//export ReadToml
func ReadToml() {
	config, err := toml.LoadFile("./config.toml")
	if err != nil {
		panic("没有读取toml文件")
	}

	// 配置连接到的订阅,从toml文件读取
	subscribe = config.GetArray("subscribe.topic").([]string)
}

//export MqttPublish
func MqttPublish() {

	// 配置连接的属性
	opts = mqtt.NewClientOptions().AddBroker("mqtt://47.104.253.11:1883").SetClientID("moniter_client")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)

	// 连接和排错
	conn = mqtt.NewClient(opts)
	if token = conn.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		panic(token.Error())
	}

	// 针对每个订阅,开启一个线程用来监听
	for _, value := range subscribe {
		go func(v string) {
			tokenTemp := conn.Subscribe(v, 1, nil)
			tokenTemp.Wait()
		}(value)
	}
}

func main() {
	//需要一个主函数使CGO编译包成为C共享库
	ReadToml()
	MqttPublish()
	for {
		<-time.After(time.Second * 30)
	}
}
