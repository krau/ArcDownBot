package bot

import (
	"arcdownbot/config"
	"arcdownbot/dao"
	"fmt"
	"os"
	"time"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

func UploadVersion(version string) error {
	versionModel, err := dao.GetVersionByVersion(version)
	if err != nil {
		return err
	}

	if versionModel.Uploaded {
		fmt.Println("Version already uploaded: ", version)
		return nil
	}

	osFile, err := os.Open(versionModel.FilePath)
	if err != nil {
		return err
	}
	defer osFile.Close()
	caption := fmt.Sprintf("Arcaea v%s\nTime (Asia/Shanghai): %s", version, versionModel.CreatedAt.In(time.FixedZone("Asia/Shanghai", 8*3600)).Format("2006-01-02 15:04:05"))
	msg, err := Bot.SendDocument(
		&telego.SendDocumentParams{
			ChatID:                      MainChannelChatID,
			Document:                    telegoutil.File(telegoutil.NameReader(osFile, "arcaea_"+version+".apk")),
			Caption:                     caption,
			DisableContentTypeDetection: true,
		},
	)
	if err != nil {
		return err
	}
	versionModel.Uploaded = true
	versionModel.FileID = msg.Document.FileID
	versionModel.MessageID = msg.MessageID
	if err := dao.UpdateVersion(versionModel); err != nil {
		return err
	}
	for i, username := range config.Cfg.Usernames {
		if i == 0 {
			continue
		}
		_, err := Bot.SendDocument(
			telegoutil.Document(
				telegoutil.Username(username),
				telegoutil.FileFromID(msg.Document.FileID),
			).WithCaption(caption).WithDisableContentTypeDetection(),
		)
		if err != nil {
			fmt.Println("Error sending document to ", username)
		}
	}
	return nil
}
