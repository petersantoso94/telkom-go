package user

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type DAO struct {
	db *gorm.DB
}

func (d *DAO) GetUser(opt *GetUserOption) (*GetUserResponse, error) {
	var users []TelinUser
	q := d.db.Model(&TelinUser{}).Where(&TelinUser{Email: opt.Email}).Find(&users)

	if gorm.IsRecordNotFoundError(q.Error) || q.RowsAffected == 0 {
		return nil, &NotFoundError{}
	} else if q.Error != nil {
		return nil, &InternalError{message: q.Error.Error()}
	}

	response:= &GetUserResponse{}
	for _,user := range users {
		response.Users = append(response.Users,&User{
			Email: user.Email,
			LockIP: user.LockIP,
			Password: user.Password,
			Position: user.Position,
		})
	}

	return response, nil
}

func (d *DAO) CreateUser(opt *User) error{
	return nil
}

func (d *DAO) UpdateUser(opt *User) error{
	return nil
}

func (d *DAO) DeleteUser(opt *DeleteUserOption) error{
	return nil
}

func CreateDAO(db *gorm.DB) *DAO {
	if err := db.AutoMigrate(&TelinUser{}).Error; err != nil {
		log.WithError(err).Fatal("failed to migrate table schema")
		return nil
	}

	return &DAO{
		db: db,
	}
}