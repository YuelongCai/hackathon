package badget

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	badgetC "hackson/internal/data/constant/badget"
	"hackson/internal/handler"
	"hackson/internal/model/rds/data"
	"hackson/internal/service/badget"
	"hackson/internal/util/handlerutil"
)

// Handler for badget
type Handler struct {
	badgetService *badget.Service
}

// NewHandler .
func NewHandler(bs *badget.Service) *Handler {
	return &Handler{badgetService: bs}
}

// ListRes for list dag response body
type ListRes struct {
	Badgets []data.Badget `json:"badgets"`
	Total   int64         `json:"totalItems"`
}

// List Badget
func (s *Handler) List(c *gin.Context) {
	filterClause, paramMap, orderClause, offset, limit, err := handlerutil.BuildQueryClause(c, "badget")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var badgetListRes ListRes
	badgetList, err := s.badgetService.ListByFilterAndOrder(filterClause, paramMap, orderClause, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	badgetListRes.Badgets = badgetList
	totalCount, err := s.badgetService.CountByFilter(filterClause, paramMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	badgetListRes.Total = totalCount
	handler.SuccessRequest(c, badgetListRes)
}

// Register a Badget
func (s *Handler) Register(c *gin.Context) {
	var req data.Badget
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	err = s.badgetService.Register(&req)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, req)
}

// Change a Badget
func (s *Handler) Change(c *gin.Context) {
	var req data.Badget
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	err = s.badgetService.Change(&req)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, req)
}

// ListBadgetCategories .
func (s *Handler) ListBadgetCategories(c *gin.Context) {
	handler.SuccessRequest(c, badgetC.GetCategory())
}

// ListBadgetTriggerEvents .
func (s *Handler) ListBadgetTriggerEvents(c *gin.Context) {
	handler.SuccessRequest(c, badgetC.GetTriggerEvent())
}
