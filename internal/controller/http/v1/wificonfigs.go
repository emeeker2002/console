package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/open-amt-cloud-toolkit/console/internal/entity"
	"github.com/open-amt-cloud-toolkit/console/internal/usecase/wificonfigs"
	"github.com/open-amt-cloud-toolkit/console/pkg/logger"
	"github.com/open-amt-cloud-toolkit/console/pkg/postgres"
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

type WirelessConfigCountResponse struct {
	Count int                     `json:"totalCount"`
	Data  []entity.WirelessConfig `json:"data"`
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

		countResponse := WirelessConfigCountResponse{
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
		if err.Error() == postgres.NotFound {
			r.l.Error(err, "wireless Config "+profileName+" not found")
			errorResponse(c, http.StatusNotFound, "Config not found")
		} else {
			r.l.Error(err, "http - wireless configs - v1 - getByName")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		return
	}

	c.JSON(http.StatusOK, config)
}

func (r *WirelessConfigRoutes) insert(c *gin.Context) {
	var config entity.WirelessConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	_, err := r.t.Insert(c.Request.Context(), &config)
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - insert")

		if unique, errMsg := postgres.CheckUnique(err); !unique {
			errorResponse(c, http.StatusBadRequest, errMsg)
		} else {
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		return
	}

	storedConfig, err := r.t.GetByName(c.Request.Context(), config.ProfileName, "")
	if err != nil {

		if err.Error() == postgres.NotFound {
			r.l.Error(err, "wifi profile "+config.ProfileName+" not found")
			errorResponse(c, http.StatusNotFound, "wifi profile not found")
		} else {
			r.l.Error(err, "http - v1 - getByName")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		return
	}

	c.JSON(http.StatusCreated, storedConfig)
}

func (r *WirelessConfigRoutes) update(c *gin.Context) {
	var config entity.WirelessConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	successfulUpdate, err := r.t.Update(c.Request.Context(), &config)
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - update")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	storedConfig, err := r.t.GetByName(c.Request.Context(), config.ProfileName, "")
	if err != nil {

		if err.Error() == postgres.NotFound {
			r.l.Error(err, "wireless config "+config.ProfileName+" not found")
			errorResponse(c, http.StatusNotFound, "wireless config not found")
		} else {
			r.l.Error(err, "http - v1 - getByName")
			errorResponse(c, http.StatusInternalServerError, "database problems")
		}

		return
	}

	if !successfulUpdate {
		r.l.Error(err, "http - wireless configs - v1 - update")
		errorResponse(c, http.StatusBadRequest, "wireless config unchanged")

		return
	}

	c.JSON(http.StatusOK, storedConfig)
}

func (r *WirelessConfigRoutes) delete(c *gin.Context) {
	configName := c.Param("profileName")

	deleteSuccessful, err := r.t.Delete(c.Request.Context(), configName, "")
	if err != nil {
		r.l.Error(err, "http - wireless configs - v1 - delete")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	if !deleteSuccessful {
		r.l.Error(err, "http - wireless configs - v1 - delete")
		errorResponse(c, http.StatusNotFound, "wireless config not found")
	}

	c.JSON(http.StatusNoContent, deleteSuccessful)
}
