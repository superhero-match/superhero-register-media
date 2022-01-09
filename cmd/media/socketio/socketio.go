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
package socketio

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-register-media/cmd/media/service"
	"github.com/superhero-match/superhero-register-media/internal/config"
)

// SocketIO holds all the data related to Socket.IO.
type SocketIO struct {
	Service        service.Service
	Logger         *zap.Logger
	TimeFormat     string
	CdnURL         string
	ImageExtension string
}

// NewSocketIO returns new value of type SocketIO.
func NewSocketIO(cfg *config.Config) (*SocketIO, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &SocketIO{
		Service:        srv,
		Logger:         logger,
		TimeFormat:     cfg.App.TimeFormat,
		CdnURL:         cfg.Aws.CdnURL,
		ImageExtension: cfg.Aws.ImageExtension,
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
		hours := strings.ReplaceAll(t.Format(socket.TimeFormat), ":", "_")
		date := strings.ReplaceAll(hours, "-", "_")
		final := strings.ReplaceAll(date, "T", "_")

		uid := strings.ReplaceAll(uuid.New().String(), "-", "")

		key := fmt.Sprintf(
			"%s/%s_%s.%s",
			superheroID,
			uid,
			final,
			socket.ImageExtension,
		)

		err = socket.Service.PutObject(buffer, key)
		if err != nil {
			log.Println(err)
		}

		url := fmt.Sprintf(
			"https://%s/%s",
			socket.CdnURL,
			key,
		)

		c.Emit("mainProfilePictureURL", url)
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Println("OnDisconnect event raised...", reason)
	})

	return server, nil
}
