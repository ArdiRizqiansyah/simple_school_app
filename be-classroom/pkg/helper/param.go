package helper

import (
	"be-classroom/pkg/errs"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context, key string) (int, errs.Error) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequestError(fmt.Sprintf("parameter '%s' has to be a number", key))
	}

	return id, nil
}

// GetQueryPagination is a helper function to get query pagination from request
func GetQueryPagination(c *gin.Context) (int, int, errs.Error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		return 0, 0, errs.NewBadRequestError("query 'page' has to be a number")
	}

	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	if err != nil {
		return 0, 0, errs.NewBadRequestError("query 'per_page' has to be a number")
	}

	return page, perPage, nil
}
