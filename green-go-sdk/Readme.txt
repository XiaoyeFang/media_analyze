环境配置：
#1 安装好相应的Go语言环境（https://golang.org/dl/）
#2 将src中的uuid和greensdksample包放在自己的GOPATH中的src中或者将该项目根目录加入到GOPATH中(github上也有很多开源uuid的Go实现方案，可以根据自己喜好进行选择）
#3 将imageSyncScanSample.go中的accessKeyId accessKeySecret替换为自己阿里云账户的AK （查看网址：https://ak-console.aliyun.com/#/accesskey）
#4 在绿网API文档中配置好相应服务的path、scene，示例中是调用绿网的同步鉴黄服务（https://help.aliyun.com/document_detail/28427.html）

运行方式：
Go语言的运行方式一般分为两种：IDE和命令行，在此推荐使用IDEA + Go Plugin，具体的Go插件安装方法不再赘述
#1 IDE
IDE方式在按照以上4步的环境配置后，直接点击main方法的运行按钮，就会在控制台看到结果

#2 命令行
命令行方式需要在项目根目录中运行go build或go install命令编译项目源代码，得到可执行文件后，再运行可执行文件；
也可执行使用go run <file> 解释执行
