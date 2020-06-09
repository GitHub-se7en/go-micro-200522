package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/micro/go-micro/broker"
	"go-micro-200522/user/database"
	"go-micro-200522/user/proto"
	"go-micro-200522/user/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	userCreateTopic = "com.weiguo.srv.user.created"
	userLoginTopic  = "com.weiguo.srv.user.login"
)

type UserHandler struct {
	//创建一个空的类型，用来升级成接口类型
}

func (u *UserHandler) CreateUser(ctx context.Context, user *proto.User, resp *proto.Response) error {
	generatePassword, err := util.GeneratePassword(user.Password, 1)
	if err != nil {
		raven.CaptureError(err, nil)
		return fmt.Errorf("generate password error: %v", err)
	}
	//查库，看看这个用户是不是存在的
	collection := database.DB.Collection("user")
	docs := bsonx.Doc{}
	collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&docs)
	if len(docs) !=0{
		resp.Success = false
		resp.Msg = "用户已存在"
		return nil
	}
	//新建一个用户,把加密后的密码给这个user
	user.Password = generatePassword
	result, err := collection.InsertOne(context.TODO(), &user)
	if err!=nil{
		return err
	}
	resp.Success = true
	resp.Msg = "success"

	//利用消息队列发布消息
	id,_ := result.InsertedID.(primitive.ObjectID)
	bytes, err := json.Marshal(user)
	if err != nil {
		raven.CaptureError(err,nil)
	}
	message := broker.Message{
		Header: map[string]string{"userId": id.Hex()},
		Body:   bytes,
	}
	err = broker.Publish(userCreateTopic, &message)
	if err != nil {
		raven.CaptureError(err,nil)
	}
	return nil
}

func (u *UserHandler) UserLogin(ctx context.Context, req *proto.LoginRequest, user *proto.LoginResponse) error {
	docs := bsonx.Doc{}
	err := database.DB.Collection("user").FindOne(context.TODO(), bson.D{{"username", req.Username}}).Decode(&docs)
	if err != nil {
		return err
	}
	validatePassword := util.ValidatePassword(docs.Lookup("password").String(), req.Password)
    if !validatePassword{
    	return fmt.Errorf("密码错误")
	}
	userId := docs.Lookup("_id").ObjectID().Hex()
	//封装返回报文
	user.UserId = userId
	user.Nickname = docs.Lookup("nickname").StringValue()
	user.Username = docs.Lookup("username").StringValue()
	user.Phone = docs.Lookup("phone").StringValue()
	i, _ := docs.Lookup("birthday").Int64OK()
	user.Birthday = i
	bytes, err := json.Marshal(&user)
	if err != nil {
		raven.CaptureError(err,nil)
	}
	message := broker.Message{
		Header: map[string]string{"userId": userId},
		Body:   bytes,
	}
	err = broker.Publish(userLoginTopic, &message)
	if err != nil {
		raven.CaptureError(err,nil)
	}
	return nil
}
