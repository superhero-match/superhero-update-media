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

type Aws struct {
	Region              string `env:"AWS_REGION" default:"us-west-2"`
	SuperheroesS3Bucket string `env:"AWS_SUPERHEROES_S3_BUCKET" default:"superheroes-pictures"`
	CdnURL              string `env:"AWS_CDN_URL" default:"d3pfwtk1dtfl92.cloudfront.net"`
	ImageExtension      string `env:"AWS_IMAGE_EXTENSION" default:"jpg"`
	ContentType         string `env:"AWS_IMAGE_CONTENT_TYPE" default:"image/jpg"`
	ContentEncoding     string `env:"AWS_IMAGE_CONTENT_ENCODING" default:"base64"`
}
