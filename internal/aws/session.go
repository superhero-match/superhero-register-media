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
package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/superhero-match/superhero-register-media/internal/config"
)

// NewSession configures and returns AWS session.
func NewSession(cfg *config.Config) (*session.Session, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Aws.Region),
	})
	if err != nil {
		return nil, err
	}

	return s, nil
}
