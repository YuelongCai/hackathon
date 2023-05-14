package badge

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	badgeC "hackathon/internal/data/constant/badge"
	"hackathon/internal/handler"
	"hackathon/internal/model/rds/data"
	"hackathon/internal/service/badge"
	"hackathon/internal/util/handlerutil"
)

// Handler for badge
type Handler struct {
	badgeService *badge.Service
}

// NewHandler .
func NewHandler(bs *badge.Service) *Handler {
	return &Handler{badgeService: bs}
}

// ListRes for list dag response body
type ListRes struct {
	Badges []data.Badge `json:"badges"`
	Total  int64        `json:"totalItems"`
}

// List Badge
func (s *Handler) List(c *gin.Context) {
	filterClause, paramMap, orderClause, offset, limit, err := handlerutil.BuildQueryClause(c, "badge")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var badgeListRes ListRes
	badgeList, err := s.badgeService.ListByFilterAndOrder(filterClause, paramMap, orderClause, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	badgeListRes.Badges = badgeList
	totalCount, err := s.badgeService.CountByFilter(filterClause, paramMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	badgeListRes.Total = totalCount
	handler.SuccessRequest(c, badgeListRes)
}

// Register a Badge
func (s *Handler) Register(c *gin.Context) {
	var req data.Badge
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	err = s.badgeService.Register(&req)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, req)
}

// Change a Badge
func (s *Handler) Change(c *gin.Context) {
	var req data.Badge
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	err = s.badgeService.Change(&req)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, req)
}

// ListBadgeCategories .
func (s *Handler) ListBadgeCategories(c *gin.Context) {
	handler.SuccessRequest(c, badgeC.GetCategory())
}

// ListBadgeTriggerEvents .
func (s *Handler) ListBadgeTriggerEvents(c *gin.Context) {
	handler.SuccessRequest(c, badgeC.GetTriggerEvent())
}
