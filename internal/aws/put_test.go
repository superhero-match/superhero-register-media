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
	b64 "encoding/base64"
	"fmt"
	"testing"
)

func mockUploadObjectToS3(pop putObjectParams) error {
	err := pop.validateParams()
	if err != nil {
		return err
	}

	return nil
}

func TestAws_PutObject(t *testing.T) {
	buffer, err := b64.StdEncoding.DecodeString(TestImgBase64)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		mockAws           *aws
		shouldReturnError bool
		expected          error
		buffer            []byte
		key               string
	}{
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "superheroes-pictures",
				ContentEncoding:     "base64",
				ContentType:         "image/jpg",
			},
			shouldReturnError: false,
			expected:          nil,
			buffer:            buffer,
			key:               "test-key",
		},
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "",
				ContentEncoding:     "base64",
				ContentType:         "image/jpg",
			},
			shouldReturnError: true,
			expected:          fmt.Errorf("aws SuperheroesS3Bucket value is empty in PutObject"),
			buffer:            buffer,
			key:               "test-key",
		},
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "superheroes-pictures",
				ContentEncoding:     "",
				ContentType:         "image/jpg",
			},
			shouldReturnError: true,
			expected:          fmt.Errorf("aws ContentEncoding value is empty in PutObject"),
			buffer:            buffer,
			key:               "test-key",
		},
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "superheroes-pictures",
				ContentEncoding:     "base64",
				ContentType:         "",
			},
			shouldReturnError: true,
			expected:          fmt.Errorf("aws ContentType value is empty in PutObject"),
			buffer:            buffer,
			key:               "test-key",
		},
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "superheroes-pictures",
				ContentEncoding:     "base64",
				ContentType:         "image/jpg",
			},
			shouldReturnError: true,
			expected:          fmt.Errorf("data buffer passed into PutObject is empty or nil"),
			buffer:            nil,
			key:               "test-key",
		},
		{
			mockAws: &aws{
				Session:             nil,
				Client:              nil,
				putObject:           mockUploadObjectToS3,
				SuperheroesS3Bucket: "superheroes-pictures",
				ContentEncoding:     "base64",
				ContentType:         "image/jpg",
			},
			shouldReturnError: true,
			expected:          fmt.Errorf("s3 bucket key passed into PutObject is empty"),
			buffer:            buffer,
			key:               "",
		},
	}

	for _, test := range tests {
		err = test.mockAws.PutObject(test.buffer, test.key)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if !test.shouldReturnError && err != nil {
			t.Fatal(err)
		}
	}
}
