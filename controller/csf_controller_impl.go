package controller

import (
	"isustrategisService/model/web"
	"isustrategisService/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CsfControllerImpl struct {
	CsfService service.CsfService
}

func NewCsfControllerImpl(csfService service.CsfService) *CsfControllerImpl {
	return &CsfControllerImpl{CsfService: csfService}
}

func (controller *CsfControllerImpl) Create(c echo.Context) error {
	request := web.CsfCreateRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.CsfService.Create(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Create Csf",
		Data:   response,
	})

}

func (controller *CsfControllerImpl) Update(c echo.Context) error {
	request := web.CsfUpdateRequest{}
	err := c.Bind(&request) // Tambahkan ini untuk bind JSON request ke struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	request.Id = idInt

	response, err := controller.CsfService.Update(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update Csf",
		Data:   response,
	})
}

func (controller *CsfControllerImpl) Delete(c echo.Context) error {
	csfId := c.Param("csfId")
	csfIdInt, err := strconv.Atoi(csfId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
	err = controller.CsfService.Delete(c.Request().Context(), csfIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Csf",
		Data:   nil,
	})
}

func (controller *CsfControllerImpl) FindById(c echo.Context) error {
	csfId := c.Param("csfId")
	csfIdInt, err := strconv.Atoi(csfId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
	response, err := controller.CsfService.FindById(c.Request().Context(), csfIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Csf By Id",
		Data:   response,
	})
}

func (controller *CsfControllerImpl) FindAll(c echo.Context) error {
	pohonId := c.Param("pohonId")
	pohonIdInt, err := strconv.Atoi(pohonId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.CsfService.FindAll(c.Request().Context(), pohonIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All Csf",
		Data:   response,
	})
}
