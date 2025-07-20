package controller

import (
	"isustrategisService/model/web"
	"isustrategisService/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IntermediateControllerImpl struct {
	intermediateService service.IntermediateService
}

func NewIntermediateControllerImpl(intermediateService service.IntermediateService) *IntermediateControllerImpl {
	return &IntermediateControllerImpl{intermediateService: intermediateService}
}

// @Summary Create Intermediate
// @Description Create Intermediate
// @Tags Intermediate
// @Accept json
// @Produce json
// @Param intermediate body web.IntermediateCreateRequest true "Intermediate"
// @Success 201 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /intermediate [post]
func (controller *IntermediateControllerImpl) Create(c echo.Context) error {
	request := web.IntermediateCreateRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.intermediateService.Create(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Create Intermediate",
		Data:   response,
	})
}

// @Summary Update Intermediate
// @Description Update Intermediate
// @Tags Intermediate
// @Accept json
// @Produce json
// @Param id path int true "Intermediate ID"
// @Param intermediate body web.IntermediateUpdateRequest true "Intermediate"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /intermediate/:id [put]
func (controller *IntermediateControllerImpl) Update(c echo.Context) error {
	request := web.IntermediateUpdateRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	request.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.intermediateService.Update(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update Intermediate",
		Data:   response,
	})
}

// @Summary Delete Intermediate
// @Description Delete Intermediate
// @Tags Intermediate
// @Accept json
// @Produce json
// @Param id path int true "Intermediate ID"
// @Success 200 {object} nil
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /intermediate/:pohon_id [delete]
func (controller *IntermediateControllerImpl) Delete(c echo.Context) error {
	pohonId, err := strconv.Atoi(c.Param("pohon_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	err = controller.intermediateService.Delete(c.Request().Context(), pohonId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Intermediate",
		Data:   nil,
	})
}

// @Summary Find Intermediate By Id
// @Description Find Intermediate By Id
// @Tags Intermediate
// @Accept json
// @Produce json
// @Param id path int true "Intermediate ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /intermediate/detail/:id [get]
func (controller *IntermediateControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.intermediateService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Intermediate",
		Data:   response,
	})
}

// @Summary Find All Intermediate
// @Description Find All Intermediate
// @Tags Intermediate
// @Accept json
// @Produce json
// @Param tahun path string true "Tahun"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /intermediate/:tahun [get]
func (controller *IntermediateControllerImpl) FindAll(c echo.Context) error {
	tahun := c.Param("tahun")

	response, err := controller.intermediateService.FindAll(c.Request().Context(), tahun)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All Intermediate",
		Data:   response,
	})
}
