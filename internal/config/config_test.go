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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	// App configuration.
	assert.Equal(t, ":7000", cfg.App.Port, "The port should be :7000.")
	assert.Equal(t, "2006-01-02T15:04:05", cfg.App.TimeFormat, "The time format should be 2006-01-02T15:04:05.")

	// AWS configuration.
	assert.Equal(t, "us-west-2", cfg.Aws.Region, "The AWS region should be us-west-2.")
	assert.Equal(t, "superheroes-pictures", cfg.Aws.SuperheroesS3Bucket, "The AWS bucket should be superheroes-pictures.")
	assert.Equal(t, "d3pfwtk1dtfl92.cloudfront.net", cfg.Aws.CdnURL, "The AWS cdn url should be d3pfwtk1dtfl92.cloudfront.net.")
	assert.Equal(t, "jpg", cfg.Aws.ImageExtension, "The AWS image extension should be jpg.")
	assert.Equal(t, "image/jpg", cfg.Aws.ContentType, "The AWS content type should be image/jpg.")
	assert.Equal(t, "base64", cfg.Aws.ContentEncoding, "The AWS content encoding should be base64.")
}
