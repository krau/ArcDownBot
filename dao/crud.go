package dao

import "arcdownbot/model"

func CreateVersion(version *model.Version) error {
	return db.Create(version).Error
}

func GetVersionByMessageID(messageID int64) (*model.Version, error) {
	var version model.Version
	err := db.Where("message_id = ?", messageID).First(&version).Error
	return &version, err
}

func GetVersionByVersion(version string) (*model.Version, error) {
	var v model.Version
	err := db.Where("version = ?", version).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func UpdateVersion(version *model.Version) error {
	return db.Save(version).Error
}

func DeleteVersion(version *model.Version) error {
	return db.Delete(version).Error
}
