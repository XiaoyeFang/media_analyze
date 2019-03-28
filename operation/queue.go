package operation

import (
	"fmt"
	pb "pure-media/protos"
)

/*
异步返回返回信息 请求node接口
10.19 改为同步返回
*/
func AsynUrlQueue(downloadUrl []string) (reply *pb.ListImgReply, err error) {
	reply = &pb.ListImgReply{}
	//限制最大并发任务数
	sem := make(chan int, 10)
	defer close(sem)
	c := make(chan *pb.ImgReply, len(downloadUrl))
	defer close(c)

	for i := 0; i < len(downloadUrl); i++ {
		sem <- 1
		go func(url string) {
			msg, err := GetproDominant(url)
			if err != nil {
				fmt.Errorf("GetproDominant %s \n", err)
			}
			<-sem
			c <- msg
		}(downloadUrl[i])
	}

	for i := 0; i < len(downloadUrl); i++ {
		reply.ImgReply = append(reply.ImgReply, <-c)
	}

	return reply, err
}

func UploadUrlQueue(downloadUrl *pb.ListUploadFileRequest) (reply *pb.ListUploadFileReply, err error) {
	//限制最大并发任务数
	sem := make(chan int, 10)
	defer close(sem)
	c := make(chan *pb.UploadFileReply, len(downloadUrl.UpRequest))
	defer close(c)
	reply = &pb.ListUploadFileReply{}

	for i := 0; i < len(downloadUrl.UpRequest); i++ {
		sem <- 1
		go func(req *pb.UploadFileRequest) {
			msg, err := SaveFileInfo(req.FileUrl, req.FileType, req.Prefix, req.Bucket, req.Tags)
			if err != nil {
				fmt.Errorf("GetproDominant %s \n", err)
			}
			<-sem
			c <- msg
		}(downloadUrl.UpRequest[i])
	}

	for i := 0; i < len(downloadUrl.UpRequest); i++ {
		reply.UpReply = append(reply.UpReply, <-c)
	}

	return reply, err
}
