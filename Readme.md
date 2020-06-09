#### 搭建微服务框架
1. 第一步setting里面先集成gomodule模块
2. 怎么引进来go-micro框架呢？？？
   * 直接上代码吗？能不能像spring一样，直接写业务代码，
   对，逆向工程，先编写那个proto文件，然后生成代码
3. 逆向工程命令，我记得装这个工具的时候还得装go的插件，也不知道这个micro是什么时候装上的
   * D:\protoc\protoc-3.11.2-win32\bin\protoc.ex
   e -I . user/user.proto --go_out=plugins=micro:. 
4. 应该是升级了，我按照github上的例子的，我逆向的是一个文件，之前shippy也是一个文件，但是github是两个文件，
还好，内容是一致的
5. 奇葩，goland项目，先创建项目，然后生成gomod ，然后x包下面的东西就下载不下来了，本来的话啊，我已经设置了七牛的代理了啊
   * 更奇葩的是，没想到，当我准备新建一个项目待mod的时候，又出现了mod文件，然后都正常了
6. api作为对外的入口，webservice不知道是做什么用的，但是利用gin框架就能实现外部访问8080端口，然后api里面的方法利用rpc的客户端可以调用user服务，实现注册和登录
7. 关键就是所有的程序docker化之后，互相之间如何通信，以及外部如何访问里面的程序，我用的是win7加docker Windows加goland开发，虽然docker-compose
创建了一个network，但是并不意味着容器之间是可以通过localhost进行通信的，这一点和k8s是完全不一样的，使用docker-compose之后，省略的是容器的ip地址，
我可以直接使用service的name作为ip，进行通信，比如，Addr:"redis:6379",并不是localhost
8. 整个项目使用go-micro作为基础架构，整合etcd作为服务注册与发现，整合kafaka作为消息队列，使用gin作为网关对外暴露端口，
没有使用Makefile，直接使用goland进行交叉编译，然后利用Dockerfile构建镜像，docker-compose运行镜像，生成容器