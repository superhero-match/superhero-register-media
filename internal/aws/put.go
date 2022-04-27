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
	"fmt"

	a "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	ErrSuperheroesS3BucketIsEmpty = fmt.Errorf("aws SuperheroesS3Bucket value is empty in PutObject")
	ErrContentEncodingIsEmpty     = fmt.Errorf("aws ContentEncoding value is empty in PutObject")
	ErrContentTypeIsEmpty         = fmt.Errorf("aws ContentType value is empty in PutObject")
	ErrDataBufferIsEmpty          = fmt.Errorf("data buffer passed into PutObject is empty or nil")
	ErrS3BucketKeyIsEmpty         = fmt.Errorf("s3 bucket key passed into PutObject is empty")
)

// putObjectParams represents parameters for uploadObjectToS3 function.
type putObjectParams struct {
	client              *s3.S3
	superheroesS3Bucket string
	contentEncoding     string
	contentType         string
	buffer              []byte
	key                 string
}

// PutObject adds new object to S3.
func (aws *aws) PutObject(buffer []byte, key string) error {
	pop := putObjectParams{
		client:              aws.Client,
		superheroesS3Bucket: aws.SuperheroesS3Bucket,
		contentEncoding:     aws.ContentEncoding,
		contentType:         aws.ContentType,
		buffer:              buffer,
		key:                 key,
	}

	return aws.putObject(pop)
}

// uploadObjectToS3 adds new object to S3.
func uploadObjectToS3(pop putObjectParams) error {
	err := pop.validateParams()
	if err != nil {
		return err
	}

	_, err = pop.client.PutObject(&s3.PutObjectInput{
		Bucket:          a.String(pop.superheroesS3Bucket),
		Key:             a.String(pop.key),
		Body:            bytes.NewReader(pop.buffer),
		ContentLength:   a.Int64(int64(len(pop.buffer))),
		ContentEncoding: a.String(pop.contentEncoding),
		ContentType:     a.String(pop.contentType),
	})

	return err
}

// validateParams checks if parameters are valid.
func (pop putObjectParams) validateParams() error {
	if len(pop.superheroesS3Bucket) == 0 {
		return ErrSuperheroesS3BucketIsEmpty
	}

	if len(pop.contentEncoding) == 0 {
		return ErrContentEncodingIsEmpty
	}

	if len(pop.contentType) == 0 {
		return ErrContentTypeIsEmpty
	}

	if pop.buffer == nil || len(pop.buffer) == 0 {
		return ErrDataBufferIsEmpty
	}

	if len(pop.key) == 0 {
		return ErrS3BucketKeyIsEmpty
	}

	return nil
}
