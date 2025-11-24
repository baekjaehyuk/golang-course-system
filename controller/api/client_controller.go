package api

import (
	"golang-course-registration/common/exception"
	"golang-course-registration/controller/dto"
	"golang-course-registration/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClientController struct {
	studentService    service.StudentService
	lectureService    service.LectureService
	enrollmentService service.EnrollmentService
}

func NewClientController(
	studentService service.StudentService,
	lectureService service.LectureService,
	enrollmentService service.EnrollmentService,
) *ClientController {
	return &ClientController{
		studentService:    studentService,
		lectureService:    lectureService,
		enrollmentService: enrollmentService,
	}
}

func (c *ClientController) RegisterRoutes(group *echo.Group) {
	group.POST("/students", c.CreateStudent)

	group.GET("/lectures", c.ListLectures)

	group.POST("/enrollments", c.Enroll)
	group.GET("/enrollments/:studentId", c.ListEnrollmentsByStudent)
	group.DELETE("/enrollments/:studentId/:lectureId", c.CancelEnrollment)
}

// CreateStudent 학생 등록
func (c *ClientController) CreateStudent(ctx echo.Context) error {
	var req dto.CreateStudentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(exception.ErrInvalidRequestBody))
	}

	student, err := c.studentService.Register(req.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, successResponse(student))
}

// ListLectures 강좌 목록 조회
func (c *ClientController) ListLectures(ctx echo.Context) error {
	lectures, err := c.lectureService.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(exception.ErrLectureListFailed))
	}
	return ctx.JSON(http.StatusOK, successResponse(lectures))
}

// Enroll 수강신청
func (c *ClientController) Enroll(ctx echo.Context) error {
	var req dto.EnrollRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(exception.ErrInvalidRequestBody))
	}

	enrollment, err := c.enrollmentService.Enroll(req.StudentID, req.LectureID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, successResponse(enrollment))
}

// ListEnrollmentsByStudent 학생의 수강신청 내역 조회
func (c *ClientController) ListEnrollmentsByStudent(ctx echo.Context) error {
	studentID, _ := strconv.Atoi(ctx.Param("studentId"))
	lectures, err := c.enrollmentService.ListByStudent(studentID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, successResponse(lectures))
}

// CancelEnrollment 수강신청 취소
func (c *ClientController) CancelEnrollment(ctx echo.Context) error {
	studentIDStr := ctx.Param("studentId")
	lectureIDStr := ctx.Param("lectureId")

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil || studentID <= 0 {
		return ctx.JSON(http.StatusBadRequest, errorResponse(exception.ErrStudentIDNotNumber))
	}

	lectureID, err := strconv.Atoi(lectureIDStr)
	if err != nil || lectureID <= 0 {
		return ctx.JSON(http.StatusBadRequest, errorResponse(exception.ErrLectureIDInvalid))
	}

	err = c.enrollmentService.Cancel(studentID, lectureID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, successResponse("수강신청이 취소되었습니다"))
}
