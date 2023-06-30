package repositories

import (
	"IofIPOS/quota/internal/configs"
	"IofIPOS/shared/gormext"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	log *logrus.Logger
	DB  *gorm.DB
}

func NewRepository(log *logrus.Logger, conf *configs.Configs) (*Repository, error) {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Pass, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DB)
	fmt.Println(dsn)
	db, err := gormext.Open(dsn)
	if err != nil {
		log.WithError(err).Fatal("can not load repository configs")
		return nil, err
	}
	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx, gormext.UUIDExtension); err != nil {
			return err
		}
		if err = tx.AutoMigrate(
			new(User),
			new(UserQuota),
			new(Object),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("can not migration database")
		return nil, err
	}
	return &Repository{log: log, DB: db}, nil
}
