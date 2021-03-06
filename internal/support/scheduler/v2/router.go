// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"net/http"

	commonController "github.com/edgexfoundry/edgex-go/internal/pkg/v2/controller/http"
	schedulerController "github.com/edgexfoundry/edgex-go/internal/support/scheduler/v2/controller/http"

	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	v2Constant "github.com/edgexfoundry/go-mod-core-contracts/v2/v2"

	"github.com/gorilla/mux"
)

func LoadRestRoutes(r *mux.Router, dic *di.Container) {
	// v2 API routes
	// Common
	cc := commonController.NewV2CommonController(dic)
	r.HandleFunc(v2Constant.ApiPingRoute, cc.Ping).Methods(http.MethodGet)
	r.HandleFunc(v2Constant.ApiVersionRoute, cc.Version).Methods(http.MethodGet)
	r.HandleFunc(v2Constant.ApiConfigRoute, cc.Config).Methods(http.MethodGet)
	r.HandleFunc(v2Constant.ApiMetricsRoute, cc.Metrics).Methods(http.MethodGet)

	// Interval
	interval := schedulerController.NewIntervalController(dic)
	r.HandleFunc(v2Constant.ApiIntervalRoute, interval.AddInterval).Methods(http.MethodPost)
	r.HandleFunc(v2Constant.ApiIntervalByNameRoute, interval.IntervalByName).Methods(http.MethodGet)
	r.HandleFunc(v2Constant.ApiAllIntervalRoute, interval.AllIntervals).Methods(http.MethodGet)
	r.HandleFunc(v2Constant.ApiIntervalByNameRoute, interval.DeleteIntervalByName).Methods(http.MethodDelete)
	r.HandleFunc(v2Constant.ApiIntervalRoute, interval.PatchInterval).Methods(http.MethodPatch)

	// IntervalAction
	action := schedulerController.NewIntervalActionController(dic)
	r.HandleFunc(v2Constant.ApiIntervalActionRoute, action.AddIntervalAction).Methods(http.MethodPost)
}
