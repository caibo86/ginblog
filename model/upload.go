package model

import (
	"context"
	"github.com/caibo86/ginblog/utils"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var QiniuServer = utils.QiniuServer

func UploadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := &storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := &storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	putExtra := &storage.PutExtra{}
	formUploader := storage.NewFormUploader(cfg)
	ret := &storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), ret, upToken, file, fileSize, putExtra)
	if err != nil {
		return "", err
	}
	url := QiniuServer + ret.Key
	return url, nil
}
