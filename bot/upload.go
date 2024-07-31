package bot

import (
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
	msg, err := Bot.SendDocument(
		&telego.SendDocumentParams{
			ChatID:                      ChannelChatID,
			Document:                    telegoutil.File(telegoutil.NameReader(osFile, "arcaea_"+version+".apk")),
			Caption:                     fmt.Sprintf("Arcaea v%s\nTime (Asia/Shanghai): %s", version, versionModel.CreatedAt.In(time.FixedZone("Asia/Shanghai", 8*3600)).Format("2006-01-02 15:04:05")),
			DisableContentTypeDetection: true,
		},
	)
	if err != nil {
		return err
	}
	versionModel.Uploaded = true
	versionModel.FileID = msg.Document.FileID
	if err := dao.UpdateVersion(versionModel); err != nil {
		return err
	}
	return nil
}
