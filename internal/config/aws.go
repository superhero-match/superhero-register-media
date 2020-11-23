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
package config

type Aws struct {
	Region              string `env:"AWS_REGION" yaml:"region" default:"us-west-2"`
	SuperheroesS3Bucket string `env:"AWS_SUPERHEROES_S3_BUCKET" yaml:"superheroes_s3_bucket" default:"superheroes-pictures"`
	CdnURL              string `env:"AWS_CDN_URL" yaml:"cdn_url" default:"d3pfwtk1dtfl92.cloudfront.net"`
	ImageExtension      string `env:"AWS_IMAGE_EXTENSION" yaml:"image_extension" default:"jpg"`
	ContentType         string `env:"AWS_IMAGE_CONTENT_TYPE" yaml:"content_type" default:"image/jpg"`
	ContentEncoding     string `env:"AWS_IMAGE_CONTENT_ENCODING" yaml:"content_encoding" default:"base64"`
}
