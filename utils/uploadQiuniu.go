package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"main.go/global"

	"mime/multipart"
)

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File, fileSize int64) (int, string) {
	var AccessKey = global.GVA_CONFIG.Qiniu.AccessKey
	var SerectKey = global.GVA_CONFIG.Qiniu.SercetKey
	var Bucket = global.GVA_CONFIG.Qiniu.Bucket
	var ImgUrl = global.GVA_CONFIG.Qiniu.QiniuServer

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)

	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneBeimei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		code := 500
		return code, err.Error()
	}
	url := ImgUrl + ret.Key
	return 200, url
}
