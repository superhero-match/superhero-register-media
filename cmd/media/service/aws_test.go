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

package service

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/superhero-match/superhero-register-media/internal/aws"
	"testing"
)

var shouldGenerateEncodeError = false

var (
	ErrDataBufferIsEmpty  = fmt.Errorf("data buffer passed into PutObject is empty or nil")
	ErrS3BucketKeyIsEmpty = fmt.Errorf("s3 bucket key passed into PutObject is empty")
)

type mockAws struct {
	putObject func(buffer []byte, key string) error
}

func mockUploadObjectToS3(buffer []byte, key string) error {
	if buffer == nil || len(buffer) == 0 {
		return ErrDataBufferIsEmpty
	}

	if len(key) == 0 {
		return ErrS3BucketKeyIsEmpty
	}

	return nil
}

func (m mockAws) PutObject(buffer []byte, key string) error {
	return m.putObject(buffer, key)
}

func TestService_PutObject(t *testing.T) {
	mAws := mockAws{
		putObject: mockUploadObjectToS3,
	}

	mockService := &service{
		AWS: mAws,
	}

	buffer, err := b64.StdEncoding.DecodeString(aws.TestImgBase64)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		buffer            []byte
		key               string
		shouldReturnError bool
		expected          error
	}{
		{
			buffer:            buffer,
			key:               "test-key",
			shouldReturnError: false,
			expected:          nil,
		},
		{
			buffer:            nil,
			key:               "test-key",
			shouldReturnError: true,
			expected:          fmt.Errorf("data buffer passed into PutObject is empty or nil"),
		},
		{
			buffer:            buffer,
			key:               "",
			shouldReturnError: true,
			expected:          fmt.Errorf("s3 bucket key passed into PutObject is empty"),
		},
	}

	for _, test := range tests {
		err = mockService.PutObject(test.buffer, test.key)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if !test.shouldReturnError && err != nil {
			t.Fatal(err)
		}
	}
}
