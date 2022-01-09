/*
  Copyright (C) 2019 - 2022 MWSOFT
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
package aws

import (
	a "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/superhero-match/superhero-register-media/internal/config"
)

// AWS interface defines AWS methods.
type AWS interface {
	PutObject(buffer []byte, key string) error
}

// aws holds all AWS related data.
type aws struct {
	Session             *session.Session
	SuperheroesS3Bucket string
	ContentEncoding     string
	ContentType         string
}

// NewAWS configures and returns AWS.
func NewAWS(cfg *config.Config) (AWS, error) {
	s, err := session.NewSession(&a.Config{
		Region: a.String(cfg.Aws.Region),
	})
	if err != nil {
		return nil, err
	}

	return &aws{
		Session:             s,
		SuperheroesS3Bucket: cfg.Aws.SuperheroesS3Bucket,
		ContentEncoding:     cfg.Aws.ContentEncoding,
		ContentType:         cfg.Aws.ContentType,
	}, nil
}
