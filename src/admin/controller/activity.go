/*
通用类活动客户端
*/
package controller

import (
	"context"
	"fmt"
	"gm/admin/controller/output"
	pb "gm/common/in"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8001"
)

type S2C_Activity struct {
	Code uint32
	Cmd  uint32
	Ret  uint32
	Desc string
}

func CommonActivity(c *gin.Context) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	cl := pb.NewWebServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addList := make([]*pb.ActivityInfo, 1)
	addList[0] = &pb.ActivityInfo{
		Id:    1,
		Type:  1,
		Name:  "holiday",
		Start: 1500000000,
		End:   1500000080,
	}

	dels := []uint32{}

	act := &pb.Activity{
		Add:   addList,
		DelId: dels,
	}

	req := &pb.GMRequest{
		Cmd: 1,
		Sid: 10,
		Data: &pb.GMRequestByte{
			Activity: act,
		},
	}

	r, err := cl.GMCmd(ctx, req)
	if err != nil {
		log.Fatalf("could not GMCmd:%v", err)
	}

	fmt.Println(r)

	output.Response(c, S2C_Activity{
		Code: 0,
		Cmd:  r.Cmd,
		Ret:  r.Ret,
		Desc: r.Desc,
	})
}
