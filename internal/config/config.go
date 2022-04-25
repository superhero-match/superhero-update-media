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

	"github.com/jinzhu/configor"
)

// Config holds the configuration.
type Config struct {
	App      *App
	Aws      *Aws
	Producer *Producer
}

// NewConfig returns the configuration.
func NewConfig() (*Config, error) {
	configFile := "config.yml"

	if len(os.Getenv("TEST_CONFIG")) > 0 {
		configFile = os.Getenv("TEST_CONFIG")
	}

	var cfg Config

	if err := configor.Load(&cfg, configFile); err != nil {
		return nil, err
	}

	return &cfg, nil
}
