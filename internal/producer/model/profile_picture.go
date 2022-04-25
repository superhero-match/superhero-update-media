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

package model

import "fmt"

var (
	ErrInvalidPosition    = fmt.Errorf("the position of the profile picture is invalid")
	ErrURLIsEmpty         = fmt.Errorf("profile picture url is empty")
	ErrSuperheroIDIsEmpty = fmt.Errorf("profile picture superhero id is empty")
	ErrCreatedAtIsEmpty   = fmt.Errorf("profile picture createdAt id is empty")
)

// ProfilePicture holds profile picture data.
type ProfilePicture struct {
	SuperheroID string `json:"superheroId"`
	URL         string `json:"url"`
	Position    int64  `json:"position"`
	CreatedAt   string `json:"createdAt"`
}

// Validate validates ProfilePicture data.
func (pp ProfilePicture) Validate() error {
	if pp.Position < 0 {
		return ErrInvalidPosition
	}

	if len(pp.URL) == 0 {
		return ErrURLIsEmpty
	}

	if len(pp.SuperheroID) == 0 {
		return ErrSuperheroIDIsEmpty
	}

	if len(pp.CreatedAt) == 0 {
		return ErrCreatedAtIsEmpty
	}

	return nil
}
