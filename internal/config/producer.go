/*
  Copyright (C) 2019 - 2020 MWSOFT
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

// Producer holds the configuration values for the Kafka producer.
type Producer struct {
	Brokers      []string `env:"KAFKA_BROKERS" default:"[192.168.0.105:9092]"`
	Topic        string   `env:"KAFKA_UPDATE_MUNICIPALITY_PROFILE_PICTURE_TOPIC" default:"update.municipality.profilepicture"`
	BatchSize    int      `env:"KAFKA_BATCH_SIZE" default:"1"`
	BatchTimeout int      `env:"KAFKA_BATCH_TIMEOUT" default:"10"`
}