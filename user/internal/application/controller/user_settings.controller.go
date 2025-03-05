package controller

import (
	"reflect"

	"github.com/bhtoan2204/user/internal/application/command"
	"github.com/bhtoan2204/user/internal/application/middleware"
	"github.com/bhtoan2204/user/internal/application/query"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type UserSettingController struct {
	commandBus *command.CommandBus
	queryBus   *query.QueryBus
}

func NewUserSettingController(commandBus *command.CommandBus, queryBus *query.QueryBus, r *gin.RouterGroup) *UserSettingController {
	ctrl := &UserSettingController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	controllerName := reflect.TypeOf(ctrl).Elem().Name()
	meter := otel.Meter("user-server-meter")
	serverAttribute := attribute.String("controller", controllerName)
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"user_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := middleware.NewInstrumentedHandler(requestCount, commonLabels)

	r.GET("", instrument(ctrl.GetUserSettings))
	r.POST("", instrument(ctrl.UpdateUserSettings))

	return ctrl
}

func (controller *UserSettingController) GetUserSettings(c *gin.Context) {
	c.JSON(200, nil)
}

func (controller *UserSettingController) UpdateUserSettings(c *gin.Context) {
	c.JSON(200, nil)
}
