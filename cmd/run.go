package cmd

import (
	"arcdownbot/bot"
	"arcdownbot/common"
	"arcdownbot/config"
	"arcdownbot/dao"
	"arcdownbot/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

func Run() {
	config.InitConfig()
	if err := dao.InitDB(); err != nil {
		fmt.Println("Error initializing database: ", err)
		os.Exit(1)
	}
	if err := bot.InitBot(); err != nil {
		fmt.Println("Error initializing bot: ", err)
		os.Exit(1)
	}
	go bot.Run()
	ticker := time.NewTicker(time.Duration(config.Cfg.Interval) * time.Second)
	for {
		execTask()
		<-ticker.C
	}
}

func execTask() {
	resp, err := common.Client.R().Get(common.ArcaeaAPIURL)
	if err != nil {
		fmt.Println("Error getting data from Arcaea API: ", err)
		return
	}
	var arcaeaResp *model.ArcaeaResponse
	if err := json.Unmarshal(resp.Bytes(), &arcaeaResp); err != nil {
		fmt.Println("Error unmarshalling Arcaea API response: ", err)
		return
	}
	fmt.Println("Arcaea API response: ", arcaeaResp)
	versionModel, err := dao.GetVersionByVersion(arcaeaResp.Value.Version)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Error getting version by version: ", err)
		return
	}
	if versionModel != nil {
		if versionModel.Uploaded {
			return
		}
	}
	filePath, err := common.DownloadFile(arcaeaResp.Value.Url, "arcaea_"+arcaeaResp.Value.Version+".apk")
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
	}
	fmt.Println("File downloaded: ", filePath)
	if versionModel == nil {
		versionModel = &model.Version{
			Version:  arcaeaResp.Value.Version,
			Uploaded: false,
			FilePath: filePath,
		}
		if err := dao.CreateVersion(versionModel); err != nil {
			fmt.Println("Error creating version: ", err)
			return
		}
		fmt.Println("New version created: ", arcaeaResp.Value.Version)
	} else {
		versionModel.Uploaded = false
		versionModel.FilePath = filePath
		if err := dao.UpdateVersion(versionModel); err != nil {
			fmt.Println("Error updating version: ", err)
			return
		}
		fmt.Println("Version updated: ", arcaeaResp.Value.Version)
	}
	fmt.Println("Uploading version: ", arcaeaResp.Value.Version)

	if err := bot.UploadVersion(arcaeaResp.Value.Version); err != nil {
		fmt.Println("Error uploading version: ", err)
		return
	}
	fmt.Println("Version uploaded: ", arcaeaResp.Value.Version)

	if err := common.PurgeFile(filePath); err != nil {
		fmt.Println("Error purging file: ", err)
		return
	}
	fmt.Println("File purged: ", filePath)
}
