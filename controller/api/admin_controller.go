package api

import (
	"golang-course-registration/common/exception"
	"golang-course-registration/controller/dto"
	"golang-course-registration/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	lectureService service.LectureService
}

func NewAdminController(lectureService service.LectureService) *AdminController {
	return &AdminController{lectureService: lectureService}
}

func (c *AdminController) RegisterRoutes(group *echo.Group) {
	group.POST("/lectures", c.CreateLecture)
	group.GET("/lectures", c.ListLectures)
	group.DELETE("/lectures/:id", c.DeleteLecture)
}

// CreateLecture 강좌 등록
func (c *AdminController) CreateLecture(ctx echo.Context) error {
	var req dto.CreateLectureRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(exception.ErrInvalidRequestBody))
	}

	_, err := c.lectureService.Create(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, successResponse(map[string]string{"message": "강좌가 등록되었습니다"}))
}

// ListLectures 강좌 목록 조회
func (c *AdminController) ListLectures(ctx echo.Context) error {
	lectures, err := c.lectureService.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(exception.ErrLectureListFailed))
	}
	return ctx.JSON(http.StatusOK, successResponse(lectures))
}

// DeleteLecture 강좌 삭제
func (c *AdminController) DeleteLecture(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, errorResponse("유효하지 않은 강좌 번호입니다"))
	}

	err = c.lectureService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, successResponse(map[string]string{"message": "강좌가 삭제되었습니다"}))
}
