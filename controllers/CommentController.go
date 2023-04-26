package controllers

import (
	"final-project-hacktiv8/helpers"
	"final-project-hacktiv8/models"
	"final-project-hacktiv8/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CommentController struct {
	CommentService *services.CommentService
}

type CommentDto struct {
	Message string `json:"message" form:"message" binding:"required"`
}

type CommentResponse struct {
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
	Message string `json:"comment"`
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var (
		data = c.MustGet("data").(map[string]any)

		commentDto CommentDto

		err error

		contentType = helpers.GetContentType(c)
	)

	if contentType == helpers.AppJson {
		if err = c.ShouldBindJSON(&commentDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		if err = c.ShouldBind(&commentDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	userData := data["user"].(jwt.MapClaims)
	photoData := data["photo"].(models.Photo)

	result, err := cc.CommentService.CreateComment(uint(userData["id"].(float64)), photoData.ID, commentDto.Message)

	if err != nil {
		if err = c.ShouldBind(&commentDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	response := CommentResponse{
		PhotoID: result.PhotoID,
		UserID:  result.UserID,
		Message: result.Message,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

func (cc *CommentController) GetAllComments(c *gin.Context) {
	var (
		userData = c.MustGet("userData").(jwt.MapClaims)

		response []CommentResponse

		err error
	)

	comments, err := cc.CommentService.GetAllComments(uint(userData["id"].(float64)))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, value := range comments {
		response = append(response, CommentResponse{
			PhotoID: value.PhotoID,
			UserID:  value.UserID,
			Message: value.Message,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": response,
	})
}

func (cc *CommentController) GetComment(c *gin.Context) {
	var (
		commentResponse CommentResponse

		data = c.MustGet("dataComment").(models.Comment)
	)

	commentResponse = CommentResponse{
		PhotoID: data.PhotoID,
		UserID:  data.UserID,
		Message: data.Message,
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": commentResponse,
	})

}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var (
		err error

		data = c.MustGet("dataComment").(models.Comment)

		commentDto CommentDto

		contentType = helpers.GetContentType(c)
	)

	if contentType == helpers.AppJson {
		if err = c.ShouldBindJSON(&commentDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		if err = c.ShouldBind(&commentDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	_, err = cc.CommentService.UpdateComment(data.ID, models.Comment{
		Message: commentDto.Message,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data sucessfully updated",
	})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	var (
		err error

		data = c.MustGet("dataComment").(models.Comment)
	)

	err = cc.CommentService.DeleteComment(data.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data sucesfully deleted",
	})
}
