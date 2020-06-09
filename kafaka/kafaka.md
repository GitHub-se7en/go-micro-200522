#### kafaka使用
**阿里云搭建kafaka**
1. 使用目录下面的docker-compose.yml文件，创建kafaka集群
    * kafaka对外暴露了两个端口，zookeeper对外暴露了一个端口，kafaka manager对外暴露一个端口
    * **注意**必须暴露阿里云的规则，否则manager会报timeout的异常
2. 使用kafaka-go连接工具连接kafaka
    * 不要使用confluent-kafka-go，这个要依赖c编译，试了好多次还是没成功
    * sarama 这个没有基本的文档
    * 之所以使用kafaka-go是因为这个代码写的太像go的rpc源码了，一开始我也觉得，不就是个连接工具吗，怎么这么麻烦，直到看到kafaka-go，看到里面的conn之后，整个人都舒服了好多
3. 使用kafaka-manager的注意事项
    * 在阿里云搭建起集群之后，kafaka-manager搭建之前，我已经用client添加了topic了，问题就出现在访问kafaka-manager的时候没有看到我创建的topic，
    访问kafaka-manager的时候，需要建立一个cluster，一开始我以为cluster是kafaka的概念，后来发现不是，这个cluster是和zookeeper一一对应的，但是这样的话，就没办法解释为什么没有我创建的topic，之后我刷新了一次，就出来了


**主要参考文章** [简书kafaka搭建](www.jianshu.com/p/e324ceabf494) 
