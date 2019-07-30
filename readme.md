运行步骤：

1.1go get github.com/kardianos/govendor

1.2go get github.com/gin-gonic/gin

2.cd src

3.govendor init

4.govendor fetch github.com/gin-gonic/gin@v1.4

5.curl https://raw.githubusercontent.com/gin-gonic/gin/master/examples/basic/main.go > main.go

6.go run main.go

目录结构

src --> go文件

templates --> html模版文件

测试地址：

http://localhost:9090/sayHello/dong

http://localhost:9090/index
