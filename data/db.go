package data

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

//定义 MongoDB 连接字符串
const(
	host="localhost:27017"
	source="localhost"
	user=""//用户，本地未设所以为空
	pass=""//数据库密码，同样是未设置
)

//定义 Mongo Session
var globalS *mgo.Session
var MongoUrl=host+"/"+source

//初始化连接 MongoDB
func init()  {
	dialInfo:=&mgo.DialInfo{
		Addrs:[]string{host},
		Source:source,
		Username:user,
		Password:pass,
	}
	session,err:=mgo.DialWithInfo(dialInfo)
	if err!=nil {
		log.Fatal("Create Session Error",err)
	}
	fmt.Println("MongoDB Has Connected")
	globalS=session
}

//连接 MongoDB 返回一个 Session 会话和一个集合 c
func connect(db,collection string) (*mgo.Session,*mgo.Collection) {
	s:=globalS.Copy()
	c:=s.DB(db).C(collection)
	return s,c
}

//插入
func Insert(db,collection string,docs ...interface{}) error {
	ms,c:=connect(db,collection)
	defer ms.Close()
	return c.Insert(docs...)
}

//查找某一个
func FindOne(db,collection string,query,selector,result interface{})  error{
	ms,c:=connect(db,collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

//是否存在
func IsExist(db,collection string,query interface{})  bool{
	ms,c:=connect(db,collection)
	defer ms.Close()
	count,_:=c.Find(query).Count()
	return count>0
}

//查找所有
func FindAll(db,collection string,query,selector,result interface{}) error {
	ms,c:=connect(db,collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

//更新
func Update(db,collection string,query,update interface{}) error {
	ms,c:=connect(db,collection)
	defer ms.Close()
	return c.Update(query,update)
}

//删除
func Remove(db,collection string,query interface{}) error {
	ms,c:=connect(db,collection)
	defer ms.Close()
	return c.Remove(query)
}