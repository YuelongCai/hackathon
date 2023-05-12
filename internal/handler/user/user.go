package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"

	userD "hackson/internal/data/user"
	"hackson/internal/handler"
	userS "hackson/internal/service/user"
)

// Handler for user
type Handler struct {
	userService *userS.Service
}

// NewHandler .
func NewHandler(us *userS.Service) *Handler {
	return &Handler{userService: us}
}

// HandleBehavior .
func (s *Handler) HandleBehavior(c *gin.Context) {
	var req userD.Behavior
	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	err = s.userService.HandleBehavior(&req)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, nil)
}

// ListBadgetAssets .
func (s *Handler) ListBadgetAssets(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	category, _ := c.GetQuery("category")
	badgetList, err := s.userService.ListBadgetAssets(userID, category)
	if err != nil {
		handler.BadRequest(c, err.Error())
		return
	}
	handler.SuccessRequest(c, badgetList)
	return
}
