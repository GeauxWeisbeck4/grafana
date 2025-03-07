package dashvertest

import (
	"context"

	dashver "github.com/grafana/grafana/pkg/services/dashboardversion"
)

type FakeDashboardVersionService struct {
	ExpectedDashboardVersion  *dashver.DashboardVersion
	ExpectedDashboardVersions []*dashver.DashboardVersion
	counter                   int
	ExpectedError             error
}

func NewDashboardVersionServiceFake() *FakeDashboardVersionService {
	return &FakeDashboardVersionService{}
}

func (f *FakeDashboardVersionService) Get(ctx context.Context, query *dashver.GetDashboardVersionQuery) (*dashver.DashboardVersion, error) {
	if len(f.ExpectedDashboardVersions) == 0 {
		return f.ExpectedDashboardVersion, f.ExpectedError
	}
	f.counter++
	return f.ExpectedDashboardVersions[f.counter-1], f.ExpectedError
}

func (f *FakeDashboardVersionService) DeleteExpired(ctx context.Context, cmd *dashver.DeleteExpiredVersionsCommand) error {
	return f.ExpectedError
}
