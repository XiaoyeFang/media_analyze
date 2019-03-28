// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

package gogrpc

import (
	"errors"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"pure-media/config"
	"pure-media/operation"
	pb "pure-media/protos"
)

const (
	FAILURE = "failure"
)

type PureMedia struct {
}

// 上传文件
//UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadReply, error)
// 图片分析
//ImgAnalyze(ctx context.Context, in *ImgRequest, opts ...grpc.CallOption) (*ImgReply, error)
// 队列图片分析
//ListImgAnalyze(ctx context.Context, in *ListImgRequest, opts ...grpc.CallOption) (*ListImgReply, error)

func (*PureMedia) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (reply *pb.UploadFileReply, err error) {
	reply = &pb.UploadFileReply{}
	glog.Infof("UploadFileRequest %s \n", req)
	reply, err = operation.SaveFileInfo(req.FileUrl, req.FileType, req.Bucket, req.Prefix, req.Tags)
	if err != nil {
		glog.Errorf("operation.SaveFileInfo %s \n", err)
		return reply, err

	}
	glog.Infof("reply = %v \n", reply)
	return reply, err
}

func (*PureMedia) ListUploadFile(ctx context.Context, req *pb.ListUploadFileRequest) (reply *pb.ListUploadFileReply, err error) {
	reply = &pb.ListUploadFileReply{}
	glog.V(5).Infof("UploadFileRequest %s \n", req)
	reply, err = operation.UploadUrlQueue(req)
	if err != nil {
		glog.Errorf("operation.SaveFileInfo %s \n", err)
		return reply, err

	}
	glog.V(5).Infoln("reply =", reply)
	return reply, err
}

func (*PureMedia) ImgAnalyze(ctx context.Context, req *pb.ImgRequest) (reply *pb.ImgReply, err error) {
	reply = &pb.ImgReply{}
	glog.V(5).Infof("req = %v \n", req)
	reply, err = operation.GetproDominant(req.DownloadUrl)
	if err != nil {
		if err.Error() == config.HTTPERRORCODE {
			glog.Errorf("url %s , httpError %s \n", req.DownloadUrl, err)
		} else if err.Error() == config.IMAGEERRORCODE {
			glog.Errorf("url %s , imageError %s \n", req.DownloadUrl, err)
		} else {
			glog.Errorf("url %s , Error %s \n", req.DownloadUrl, err)
			err = errors.New(config.IMAGEERRORCODE)
		}

		return reply, err

	}
	glog.V(5).Infof("reply = %v \n", reply)
	return reply, err
}

func (*PureMedia) ListImgAnalyze(ctx context.Context, req *pb.ListImgRequest) (reply *pb.ListImgReply, err error) {
	reply = &pb.ListImgReply{}
	downloadUrl := make([]string, 0)
	for _, url := range req.Request {
		downloadUrl = append(downloadUrl, url.DownloadUrl)
	}
	reply, err = operation.AsynUrlQueue(downloadUrl)
	if err != nil {
		glog.Errorf("operation.UrlQueue %s\n", err)
	}
	return reply, err
}

func (*PureMedia) YoutubeVideo(ctx context.Context, req *pb.YoutubeVideoRequest) (reply *pb.YoutubeVideoReply, err error) {
	glog.V(5).Infof("videoUrl : %s \n", req.Link)
	videoMsg, err := operation.GetYoutubeWatchReply(req.Link)
	if err != nil {
		if err.Error() != config.INVALID_URL {
			glog.Errorf("url %s,err %v \n", req.Link, err)
		}
		return nil, err
	}
	reply = &pb.YoutubeVideoReply{
		Id:               videoMsg.VideoId,
		Title:            videoMsg.Title,
		ScreenYoutubeUrl: videoMsg.ScreenYoutubeUrl,
		ScreenFid:        videoMsg.ScreenFid,
		LengthSeconds:    videoMsg.LengthSeconds,
		KeyWords:         videoMsg.KeyWords,
	}
	glog.V(5).Infof("reply : %v \n", reply)
	return reply, err
}

func (*PureMedia) ImgCensor(ctx context.Context, req *pb.ImgCensorRequest) (*pb.ImgCensorReply, error) {
	// swagger:route GET /users/{id} users getSingleUser
	//
	// get a user by userID
	//
	// This will show a user info
	//
	//     Responses:
	//       200: UserResponse

	return &pb.ImgCensorReply{}, nil
}

func (*PureMedia) AddImgtoBuckets(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	result, err := operation.AuditPorn(req.ImgUrl, req.ImgName)
	if err != nil {
		glog.Errorf("AuditPorn err %s", err)
		return &pb.AddReply{
			Result: FAILURE,
		}, err
	}
	return &pb.AddReply{
		Result: result,
	}, nil

}

func (*PureMedia) YoutubeDownload(ctx context.Context, req *pb.DLRequest) (*pb.DLReply, error) {

	fid, err := operation.DownloadVideo(req.VideoUrl, req.Format)
	if err != nil {
		glog.Errorf("url %s Download err %s \n", req.VideoUrl, err)
	}
	return &pb.DLReply{
		Fid: fid,
	}, nil
}
