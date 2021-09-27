/*
1. 编写proto
server/common/helloworld.proto

2. 生成pb代码:
cd server
protoc --go_out=plugins=grpc:./src/common/in ./common/in/helloworld.proto
或：../bin/protoc --go_out=plugins=grpc:./src/common/in --plugin=protoc-gen-go=../bin/protoc-gen-go.exe ./common/in/helloworld.proto
执行以上命令生成server/src/common/in/helloworld.pb.go

3. 实现客户端
在server/src/admin/controller/sayhello.go中编写代码调用服务端SayHello方法

4. 实现服务端
在server/src/innerserver/sayhello/main.go中实现

5. 调用
启动服务端：go run server/src/innerserver/sayhello/main.go
启动客户端：go run server/src/admin/main.go
浏览器访问：http://localhost:8000/grpc/sayhello?name=alex

*/
package controller

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"gm/admin/controller/output"
	pb "gm/common/in"
)

const (
	address     = "localhost:8001"
	defaultName = "world"
)

type C2S_SayHello struct {
	Name string `from:"name"`
}

type S2C_SayHello struct {
	Code    uint32
	Message string
}

func SayHello(c *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	cl := pb.NewGreeterClient(conn)

	//req := &C2S_SayHello{}
	name := c.Query("name")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cl.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}

	output.Response(c, S2C_SayHello{
		Code:    0,
		Message: r.Message,
	})
}
