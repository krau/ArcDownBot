package common

import (
	"arcdownbot/config"
	"fmt"
	"path/filepath"
	"time"

	"github.com/imroc/req/v3"
)

var Client *req.Client

func init() {
	c := req.C().ImpersonateChrome().SetCommonRetryCount(2).EnableDebugLog()
	Client = c
}

func DownloadFile(url string, filename string) (string, error) {
	filePath := filepath.Join(config.Cfg.CacheDir, filename)
	_, err := Client.R().
		SetOutputFile(filePath).
		SetDownloadCallbackWithInterval(func(info req.DownloadInfo) {
			if info.Response != nil {
				fmt.Printf("downloaded %.2f%%\n", float64(info.DownloadedSize)/float64(info.Response.ContentLength)*100.0)
			}
		}, 5*time.Second).
		Get(url)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
