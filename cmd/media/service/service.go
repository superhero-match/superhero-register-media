/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package service

import (
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/superhero-match/superhero-register-media/internal/aws"
	"github.com/superhero-match/superhero-register-media/internal/config"
	"go.uber.org/zap"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	Session             *session.Session
	Logger              *zap.Logger
	TimeFormat          string
	SuperheroesS3Bucket string
	CdnURL              string
	ImageExtension      string
	ContentType         string
	ContentEncoding     string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	s, err := aws.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Service{
		Session:             s,
		Logger:              logger,
		TimeFormat:          cfg.App.TimeFormat,
		SuperheroesS3Bucket: cfg.Aws.SuperheroesS3Bucket,
		CdnURL:              cfg.Aws.CdnURL,
		ImageExtension:      cfg.Aws.ImageExtension,
		ContentType:         cfg.Aws.ContentType,
		ContentEncoding:     cfg.Aws.ContentEncoding,
	}, nil
}
