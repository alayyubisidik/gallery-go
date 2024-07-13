package exception

import (
	"gallery_go/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case validator.ValidationErrors:
				var errors []response.DetailError

				for _, fieldError := range e {
					errorDetail := response.DetailError{
						Field:   fieldError.Field(),
						Message: fieldError.Tag(),
					}
					errors = append(errors, errorDetail)
				}

				errResponse := response.ErrorResponse{
					Errors: errors,
				}

				ctx.JSON(http.StatusBadRequest, errResponse)
			case *ConflictError:
				errResponse := response.ErrorResponse{
					Errors: []response.DetailError{
						{
							Message: e.Error(),
						},
					},
				}

				ctx.JSON(http.StatusConflict, errResponse)
			case *UnAuthorizedError:
				errResponse := response.ErrorResponse{
					Errors: []response.DetailError{
						{
							Message: e.Error(),
						},
					},
				}

				ctx.JSON(http.StatusUnauthorized, errResponse)
			case *NotFoundError:
				errResponse := response.ErrorResponse{
					Errors: []response.DetailError{
						{
							Message: e.Error(),
						},
					},
				}

				ctx.JSON(http.StatusNotFound, errResponse)
			case *BadRequestError:
				errResponse := response.ErrorResponse{
					Errors: []response.DetailError{
						{
							Message: e.Error(),
						},
					},
				}

				ctx.JSON(http.StatusBadRequest, errResponse)
			default:
				errResponse := response.ErrorResponse{
					Errors: []response.DetailError{
						{
							Message: e.Error(),
						},
					},
				}
				ctx.JSON(http.StatusInternalServerError, errResponse)
			}

			ctx.Abort()
		}
	}
}
