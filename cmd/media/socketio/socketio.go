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
package socketio

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	socketio "github.com/googollee/go-socket.io"
	"github.com/superhero-match/superhero-register-media/cmd/media/service"
	"github.com/superhero-match/superhero-register-media/internal/config"
)

// SocketIO holds all the data related to Socket.IO.
type SocketIO struct {
	Service *service.Service
}

// NewSocketIO returns new value of type SocketIO.
func NewSocketIO(cfg *config.Config) (*SocketIO, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &SocketIO{
		Service: srv,
	}, nil
}

// NewSocketIOServer returns Socket.IO server.
func (socket *SocketIO) NewSocketIOServer() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}

	server.OnConnect("/", func(c socketio.Conn) error {
		log.Println("New client connected")

		return nil
	})

	server.OnEvent("/", "onUploadMainProfilePicture", func(c socketio.Conn, superheroID string, picture string) {
		log.Println("onUploadMainProfilePicture event raised...")

		buffer, err := b64.StdEncoding.DecodeString(picture)
		if err != nil {
			log.Println(err)
		}

		t := time.Now().UTC()
		hours := strings.ReplaceAll(t.Format(socket.Service.TimeFormat), ":", "_")
		date := strings.ReplaceAll(hours, "-", "_")
		final := strings.ReplaceAll(date, "T", "_")

		uid := strings.ReplaceAll(uuid.New().String(), "-", "")

		key := fmt.Sprintf(
			"%s/%s_%s.%s",
			superheroID,
			uid,
			final,
			socket.Service.ImageExtension,
		)

		_, err = s3.New(socket.Service.Session).PutObject(&s3.PutObjectInput{
			Bucket:          aws.String(socket.Service.SuperheroesS3Bucket),
			Key:             aws.String(key),
			Body:            bytes.NewReader(buffer),
			ContentLength:   aws.Int64(int64(len(buffer))),
			ContentEncoding: aws.String(socket.Service.ContentEncoding),
			ContentType:     aws.String(socket.Service.ContentType),
		})
		if err != nil {
			log.Println(err)
		}

		url := fmt.Sprintf(
			"https://%s/%s",
			socket.Service.CdnURL,
			key,
		)

		c.Emit("mainProfilePictureURL", url)
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Println("OnDisconnect event raised...", reason)
	})

	return server, nil
}
