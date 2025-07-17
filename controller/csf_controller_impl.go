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

// @Summary Create Csf
// @Description Create Csf
// @Tags Csf
// @Accept json
// @Produce json
// @Param csf body web.CsfCreateRequest true "Csf"
// @Success 201 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /csf [post]
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

// @Summary Update Csf
// @Description Update Csf
// @Tags Csf
// @Accept json
// @Produce json
// @Param id path int true "Csf ID"
// @Param csf body web.CsfUpdateRequest true "Csf"
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

// @Summary Delete Csf
// @Description Delete Csf
// @Tags Csf
// @Accept json
// @Produce json
// @Param csfId path int true "Csf ID"
// @Success 200 {object} web.WebResponse
func (controller *CsfControllerImpl) Delete(c echo.Context) error {
	idPohon := c.Param("idPohon")
	idPohonInt, err := strconv.Atoi(idPohon)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}
	err = controller.CsfService.Delete(c.Request().Context(), idPohonInt)
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

// @Summary Find Csf By Id
// @Description Find Csf By Id
// @Tags Csf
// @Accept json
// @Produce json
// @Param csfId path int true "Csf ID"
// @Success 200 {object} web.WebResponse
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

// @Summary Find All Csf
// @Description Find All Csf
// @Tags Csf
// @Accept json
// @Produce json
// @Param pohonId path int true "Pohon ID"
// @Success 200 {object} web.WebResponse
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
