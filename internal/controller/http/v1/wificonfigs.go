package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/open-amt-cloud-toolkit/console/internal/entity/dto"
	"github.com/open-amt-cloud-toolkit/console/internal/usecase/wificonfigs"
	"github.com/open-amt-cloud-toolkit/console/pkg/logger"
)

type WirelessConfigRoutes struct {
	t wificonfigs.Feature
	l logger.Interface
}

func newWirelessConfigRoutes(handler *gin.RouterGroup, t wificonfigs.Feature, l logger.Interface) {
	r := &WirelessConfigRoutes{t, l}

	h := handler.Group("/wirelessconfigs")
	{
		h.GET("", r.get)
		h.GET(":profileName", r.getByName)
		h.POST("", r.insert)
		h.PATCH("", r.update)
		h.DELETE(":profileName", r.delete)
	}
}

func (r *WirelessConfigRoutes) get(c *gin.Context) {
	var odata OData
	if err := c.ShouldBindQuery(&odata); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	items, err := r.t.Get(c.Request.Context(), odata.Top, odata.Skip, "")
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - getCount")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	if odata.Count {
		count, err := r.t.GetCount(c.Request.Context(), "")
		if err != nil {
			r.l.Error(err, "http - wireless configs - v1 - getCount")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		countResponse := dto.WirelessConfigCountResponse{
			Count: count,
			Data:  items,
		}

		c.JSON(http.StatusOK, countResponse)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func (r *WirelessConfigRoutes) getByName(c *gin.Context) {
	profileName := c.Param("profileName")

	config, err := r.t.GetByName(c.Request.Context(), profileName, "")
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			r.l.Error(err, "wireless Config "+profileName+" not found")
			errorResponse(c, http.StatusNotFound, "database problems")
		} else {
			r.l.Error(err, "http - wireless configs - v1 - getByName")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		return
	}

	c.JSON(http.StatusOK, config)
}

func (r *WirelessConfigRoutes) insert(c *gin.Context) {
	var config dto.WirelessConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	version, err := r.t.Insert(c.Request.Context(), &config)
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - insert")

		// var pgErr *pgconn.PgError
		// if errors.As(err, &pgErr) {
		// 	if pgErr.Code == postgres.UniqueViolation {
		// 		errorResponse(c, http.StatusBadRequest, pgErr.Message)
		// 	}

		// 	return
		// }

		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusCreated, version)
}

func (r *WirelessConfigRoutes) update(c *gin.Context) {
	var config dto.WirelessConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	configs, err := r.t.Update(c.Request.Context(), &config)
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - update")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, configs)
}

func (r *WirelessConfigRoutes) delete(c *gin.Context) {
	configName := c.Param("profileName")

	configs, err := r.t.Delete(c.Request.Context(), configName, "")
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - delete")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusNoContent, configs)
}
