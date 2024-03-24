package main

import (
	"gorm.io/gorm"
	"llm_training_management_system/internal/orm"
	"llm_training_management_system/rpcs"
	"log"
	"net"
	"net/rpc"
)

var db *gorm.DB
var userOrm *orm.UserOrm

func main() {
	db = orm.InitDB()
	userOrm = orm.NewUserOrm(db)
	//user, err := userOrm.GetByUsername("admin")
	//if err != nil {
	//	log.Println("ssssss")
	//	log.Println(err)
	//}
	//log.Println(user)
	//userOrm.Create(&orm.User{Username: "admin", Password: "admin"})
	//user := userOrm.GetById(0)
	beatService := new(rpcs.HeartbeatService)
	rpc.Register(beatService)
	listener, err := net.Listen("tcp", ":16116")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	log.Println("Server is running")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
