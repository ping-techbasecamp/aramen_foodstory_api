package branchHandler

import (
	"aramen-api/pkg/utils/any"
	branchService "aramen-api/srv/services/branch"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BranchHandlerInterface interface {
	GetBranches(c *gin.Context)
	GetBranchDetail(c *gin.Context)
}

type BranchHandler struct {
	branchService branchService.BranchService
}

func NewHandler(branchService branchService.BranchService) BranchHandler {
	return BranchHandler{branchService: branchService}
}

func (h BranchHandler) GetBranches(c *gin.Context) {
	branches, err := h.branchService.GetBranches()
	if err != nil {
		c.Status(http.StatusNotFound)
	}
	var response []string

	for i := range branches {
		fmt.Println(i)
		response = append(response, "2323")
	}
	c.Status(http.StatusOK)
}

func (h BranchHandler) GetBranchDetail(c *gin.Context) {
	req := c.Param("branchId")
	branchId := any.ParseInt(req)
	branch, err := h.branchService.GetBranchDetail(branchId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if branch == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
