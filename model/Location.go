package model

import "time"

type Location struct {
	ID          int64  `json:"id" example:"23" gorm:"primaryKey"`
	UID         int64  `json:"uid" example:"12"` // gorm:foreignKey
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"`

	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
}

// UpdateLocation ... Update location
func UpdateLocation(id int64, coordinates string) (err error) {
	//user.ID = id
	//err = config.Config.DB.Model(&user).
	//	Select("name", "coordinates").
	//	Updates(userMap).Error
	//
	//err = config.Config.DB.Save(user).Error
	//if err != nil {
	//	byteErr, _ := json.Marshal(err)
	//	var newError modelErrors.GormErr
	//	err = json.Unmarshal(byteErr, &newError)
	//	if err != nil {
	//		return
	//	}
	//	switch newError.Number {
	//	case 1062:
	//		err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
	//		return
	//	default:
	//		err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
	//	}
	//}
	//
	//err = config.Config.DB.Where("id = ?", id).First(&user).Error
	return
}
