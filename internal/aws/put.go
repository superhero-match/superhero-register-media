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
	"bytes"

	a "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// PutObject adds new object to S3.
func (aws *aws) PutObject(buffer []byte, key string) error {
	_, err := s3.New(aws.Session).PutObject(&s3.PutObjectInput{
		Bucket:          a.String(aws.SuperheroesS3Bucket),
		Key:             a.String(key),
		Body:            bytes.NewReader(buffer),
		ContentLength:   a.Int64(int64(len(buffer))),
		ContentEncoding: a.String(aws.ContentEncoding),
		ContentType:     a.String(aws.ContentType),
	})

	return err
}
