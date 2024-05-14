// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"time"

	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/domain"
	"github.com/chenmingyong0423/fnote/server/internal/visit_log/internal/repository"

	"github.com/chenmingyong0423/gkit/slice"
	"github.com/pkg/errors"
)

type IVisitLogService interface {
	CollectVisitLog(ctx context.Context, visitHistory domain.VisitHistory) error
	GetTodayViewCount(ctx context.Context) (int64, error)
	GetTodayUserViewCount(ctx context.Context) (int64, error)
	GetViewTendencyStats4PV(ctx context.Context, days int) ([]domain.TendencyData, error)
	GetViewTendencyStats4UV(ctx context.Context, days int) ([]domain.TendencyData, error)
	GetIpsByDate(ctx context.Context, start time.Time, end time.Time) ([]string, error)
}

var _ IVisitLogService = (*VisitLogService)(nil)

type VisitLogService struct {
	repo repository.IVisitLogRepository
}

func (s *VisitLogService) GetIpsByDate(ctx context.Context, start time.Time, end time.Time) ([]string, error) {
	visitHistories, err := s.repo.GetByDate(ctx, start, end)
	if err != nil {
		return nil, err
	}
	return slice.Map(visitHistories, func(_ int, vh domain.VisitHistory) string {
		return vh.Ip
	}), nil
}

func (s *VisitLogService) GetViewTendencyStats4UV(ctx context.Context, days int) ([]domain.TendencyData, error) {
	return s.repo.GetViewTendencyStats4UV(ctx, days)
}

func (s *VisitLogService) GetViewTendencyStats4PV(ctx context.Context, days int) ([]domain.TendencyData, error) {
	return s.repo.GetViewTendencyStats4PV(ctx, days)
}

func (s *VisitLogService) GetTodayUserViewCount(ctx context.Context) (int64, error) {
	return s.repo.CountOfTodayByIp(ctx)
}

func (s *VisitLogService) GetTodayViewCount(ctx context.Context) (int64, error) {
	return s.repo.CountOfToday(ctx)
}

func (s *VisitLogService) CollectVisitLog(ctx context.Context, visitHistory domain.VisitHistory) error {
	err := s.repo.Add(ctx, visitHistory)
	if err != nil {
		return errors.WithMessage(err, "s.repo.Add failed")
	}
	return nil
}

func NewVisitLogService(repo repository.IVisitLogRepository) *VisitLogService {
	return &VisitLogService{repo: repo}
}
