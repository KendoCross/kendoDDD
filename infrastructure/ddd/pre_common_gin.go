package ddd

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/KendoCross/kendoDDD/infrastructure/errorext"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCase method change string to snakecase
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//CQRS 契合Command的,统一进行路由
func PreCommandHandles(command string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		ctx := context.Background()
		// if claims, has := c.Get("claims"); has {
		// 	ctx = context.WithValue(ctx, "claims", claims)
		// }

		err = HandCommand(ctx, data, command)
		if err != nil {
			var listErr *errorext.ListMsgError
			if errors.As(err, &listErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": listErr.MsgMap})
				c.Abort()
				return
			}

			var validatErr validator.ValidationErrors
			if errors.As(err, &validatErr) {
				list := make(map[string]string)
				for _, err := range validatErr {
					list[strings.ToLower(toSnakeCase(err.Field()))] = errorext.ValidationErrorToText(err)
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": list})
				c.Abort()
				return
			}

			var codeErr *errorext.CodeError
			if errors.As(err, &codeErr) {
				c.JSON(http.StatusConflict, gin.H{"code": codeErr.Code, "msg": codeErr.Msg})
				c.Abort()
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error)
			return
		}

		c.Status(http.StatusOK)
	}
}

func PreCommandDeals(command string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		ctx := context.Background()
		// if claims, has := c.Get("claims"); has {
		// 	ctx = context.WithValue(ctx, "claims", claims)
		// }
		result, err := DealCommand(ctx, data, command)
		if err != nil {

			var listErr *errorext.ListMsgError
			if errors.As(err, &listErr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": listErr.MsgMap})
				c.Abort()
				return
			}

			var validatErr validator.ValidationErrors
			if errors.As(err, &validatErr) {
				list := make(map[string]string)
				for _, err := range validatErr {
					list[strings.ToLower(toSnakeCase(err.Field()))] = errorext.ValidationErrorToText(err)
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": list})
				c.Abort()
				return
			}

			var codeErr *errorext.CodeError
			if errors.As(err, &codeErr) {
				c.JSON(http.StatusConflict, gin.H{"code": codeErr.Code, "msg": codeErr.Msg})
				c.Abort()
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
