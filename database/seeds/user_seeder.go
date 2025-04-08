package seeds

import (
	"compro/internal/core/domain/model"
	"compro/utils/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	hash, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	admin := model.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: hash,
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("failed to seed admin user")
	} else {
		log.Info().Msg("Admin user has been seeded")
	}
}
