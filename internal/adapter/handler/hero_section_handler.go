package handler

import (
	"compro/config"
	"compro/internal/adapter/handler/request"
	"compro/internal/adapter/handler/response"
	"compro/internal/core/domain/entity"
	"compro/internal/core/service"
	"compro/utils/conv"
	"compro/utils/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type HeroSectionHandlerInterface interface {
	CreateHeroSection(c echo.Context) error
	FetchAllHeroSection(c echo.Context) error
	FetchByIDHeroSection(c echo.Context) error
	EditByIDHeroSection(c echo.Context) error
	DeleteByIDHeroSection(c echo.Context) error
}

type heroSectionHandler struct {
	heroSectionService service.HeroSectionServiceInterface
}

// CreateHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) CreateHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateHeroSection - 1: Unauthorized")
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(&req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
	}

	err = h.heroSectionService.CreateHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

// DeleteByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) DeleteByIDHeroSection(c echo.Context) error {
	panic("unimplemented")
}

// EditByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) EditByIDHeroSection(c echo.Context) error {
	panic("unimplemented")
}

// FetchAllHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchAllHeroSection(c echo.Context) error {
	panic("unimplemented")
}

// FetchByIDHeroSection implements HeroSectionHandlerInterface.
func (h *heroSectionHandler) FetchByIDHeroSection(c echo.Context) error {
	panic("unimplemented")
}

func NewHeroSectionHandler(c *echo.Echo, cfg *config.Config, heroSectionService service.HeroSectionServiceInterface) HeroSectionHandlerInterface {
	heroHandler := &heroSectionHandler{
		heroSectionService: heroSectionService,
	}

	mid := middleware.NewMiddleware(cfg)

	heroApp := c.Group("hero-sections")
	adminApp := heroApp.Group("/admin", mid.CheckToken())
	adminApp.GET("", heroHandler.FetchAllHeroSection)
	adminApp.POST("", heroHandler.CreateHeroSection)
	adminApp.GET("/:id", heroHandler.FetchByIDHeroSection)
	adminApp.PUT("/:id", heroHandler.EditByIDHeroSection)
	adminApp.DELETE("/:id", heroHandler.DeleteByIDHeroSection)
	return heroHandler
}
