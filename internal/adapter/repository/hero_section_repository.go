package repository

import (
	"compro/internal/core/domain/entity"
	"compro/internal/core/domain/model"
	"context"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type HeroSectionInterface interface {
	CreateHeroSection(ctx *context.Context, req entity.HeroSectionEntity) error
	FetchAllHeroSection(ctx *context.Context) ([]entity.HeroSectionEntity, error)
	FetchByIDHeroSection(ctx *context.Context, req entity.HeroSectionEntity) error
	EditByIDHeroSection(ctx *context.Context, id int64) (*entity.HeroSectionEntity, error)
	DeleteByIDHeroSection(ctx *context.Context, id int64) error
}

type heroSection struct {
	DB *gorm.DB
}

// CreateHeroSection implements HeroSectionInterface.
func (h *heroSection) CreateHeroSection(ctx *context.Context, req entity.HeroSectionEntity) error {
	modelHeroSection := model.HeroSection{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  &req.PathVideo,
		Banner:     req.Banner,
	}

	if err = h.DB.Create(&modelHeroSection).Error; err != nil {
		log.Errorf("[REPOSITORY] CreateHeroSection - 1: %v", err)
		return err
	}

	return nil
}

// DeleteByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) DeleteByIDHeroSection(ctx *context.Context, id int64) error {
	panic("unimplemented")
}

// EditByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) EditByIDHeroSection(ctx *context.Context, id int64) (*entity.HeroSectionEntity, error) {
	panic("unimplemented")
}

// FetchAllHeroSection implements HeroSectionInterface.
func (h *heroSection) FetchAllHeroSection(ctx *context.Context) ([]entity.HeroSectionEntity, error) {
	panic("unimplemented")
}

// FetchByIDHeroSection implements HeroSectionInterface.
func (h *heroSection) FetchByIDHeroSection(ctx *context.Context, req entity.HeroSectionEntity) error {
	panic("unimplemented")
}

func NewHeroSectionRepository(DB *gorm.DB) HeroSectionInterface {
	return &heroSection{
		DB: DB,
	}
}
