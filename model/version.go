package model

import (
	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	Version   string `json:"version"`
	Uploaded  bool   `json:"uploaded" gorm:"default:false"`
	MessageID int64  `json:"message_id"`
	FileID    string `json:"file_id"`
	FilePath  string `json:"file_path"` // 本地文件路径, 用于上传和删除
}
