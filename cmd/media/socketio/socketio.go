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
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	socketio "github.com/googollee/go-socket.io"

	"github.com/superhero-match/superhero-update-media/cmd/media/service"
	"github.com/superhero-match/superhero-update-media/cmd/media/socketio/model"
	"github.com/superhero-match/superhero-update-media/internal/config"
)

// SocketIO holds all the data related to Socket.IO.
type SocketIO struct {
	Service        service.Service
	CdnURL         string
	ImageExtension string
	TimeFormat     string
}

// NewSocketIO returns new value of type SocketIO.
func NewSocketIO(cfg *config.Config) (*SocketIO, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &SocketIO{
		Service:        srv,
		CdnURL:         cfg.Aws.CdnURL,
		ImageExtension: cfg.Aws.ImageExtension,
		TimeFormat:     cfg.App.TimeFormat,
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

	server.OnEvent("/", "onUpdateProfilePicture", func(c socketio.Conn, superheroID string, picture string, position int64) {
		log.Println("onUpdateProfilePicture event raised...")

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

		err = socket.Service.UpdateProfilePicture(superheroID, url, position, t.Format(socket.TimeFormat))
		if err != nil {
			log.Println(err)
		}

		c.Emit("updateProfilePictureURL", model.Response{
			URL:      url,
			Position: position,
		})
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		log.Println("OnDisconnect event raised...", reason)
	})

	return server, nil
}
