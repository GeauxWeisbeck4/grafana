package api

import (
	"errors"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/plugins/backendplugin"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/query"
	"github.com/grafana/grafana/pkg/tsdb/legacydata"
	"github.com/grafana/grafana/pkg/util"
	"github.com/grafana/grafana/pkg/web"
)

func (hs *HTTPServer) handleQueryMetricsError(err error) *response.NormalResponse {
	if errors.Is(err, models.ErrDataSourceAccessDenied) {
		return response.Error(http.StatusForbidden, "Access denied to data source", err)
	}
	if errors.Is(err, models.ErrDataSourceNotFound) {
		return response.Error(http.StatusNotFound, "Data source not found", err)
	}
	var badQuery *query.ErrBadQuery
	if errors.As(err, &badQuery) {
		return response.Error(http.StatusBadRequest, util.Capitalize(badQuery.Message), err)
	}

	if errors.Is(err, backendplugin.ErrPluginNotRegistered) {
		return response.Error(http.StatusNotFound, "Plugin not found", err)
	}

	return response.Error(http.StatusInternalServerError, "Query data error", err)
}

// QueryMetricsV2 returns query metrics.
// POST /api/ds/query   DataSource query w/ expressions
func (hs *HTTPServer) QueryMetricsV2(c *models.ReqContext) response.Response {
	reqDTO := dtos.MetricRequest{}
	if err := web.Bind(c.Req, &reqDTO); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}

	reqDTO.HTTPRequest = c.Req

	resp, err := hs.queryDataService.QueryData(c.Req.Context(), c.SignedInUser, c.SkipCache, reqDTO, true)
	if err != nil {
		return hs.handleQueryMetricsError(err)
	}
	return hs.toJsonStreamingResponse(resp)
}

// QueryMetrics returns query metrics
// POST /api/tsdb/query
//nolint: staticcheck // legacydata.DataResponse deprecated
//nolint: staticcheck // legacydata.DataQueryResult deprecated
// Deprecated: use QueryMetricsV2 instead.
func (hs *HTTPServer) QueryMetrics(c *models.ReqContext) response.Response {
	reqDto := dtos.MetricRequest{}
	if err := web.Bind(c.Req, &reqDto); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}

	sdkResp, err := hs.queryDataService.QueryData(c.Req.Context(), c.SignedInUser, c.SkipCache, reqDto, false)
	if err != nil {
		return hs.handleQueryMetricsError(err)
	}

	legacyResp := legacydata.DataResponse{
		Results: map[string]legacydata.DataQueryResult{},
	}

	for refID, res := range sdkResp.Responses {
		dqr := legacydata.DataQueryResult{
			RefID: refID,
		}

		if res.Error != nil {
			dqr.Error = res.Error
		}

		if res.Frames != nil {
			dqr.Dataframes = legacydata.NewDecodedDataFrames(res.Frames)
		}

		legacyResp.Results[refID] = dqr
	}

	statusCode := http.StatusOK
	for _, res := range legacyResp.Results {
		if res.Error != nil {
			res.ErrorString = res.Error.Error()
			legacyResp.Message = res.ErrorString
			statusCode = http.StatusBadRequest
		}
	}

	return response.JSON(statusCode, &legacyResp)
}

func (hs *HTTPServer) toJsonStreamingResponse(qdr *backend.QueryDataResponse) response.Response {
	statusWhenError := http.StatusBadRequest
	if hs.Features.IsEnabled(featuremgmt.FlagDatasourceQueryMultiStatus) {
		statusWhenError = http.StatusMultiStatus
	}

	statusCode := http.StatusOK
	for _, res := range qdr.Responses {
		if res.Error != nil {
			statusCode = statusWhenError
		}
	}

	return response.JSONStreaming(statusCode, qdr)
}
