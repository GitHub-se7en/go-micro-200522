package database

//func TestA(t *testing.T)  {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test"))
//	if err != nil {
//		raven.CaptureError(err, nil)
//		log.Println(err)
//	}
//	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
//	err = client.Ping(ctx, readpref.Primary())
//	log.Println(err)
//	collection := client.Database("test").Collection("numbers")
//	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
//	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
//	if err != nil {
//		log.Println(err)
//	}
//	id := res.InsertedID
//	log.Println(id)
//}
