package producer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/superhero-match/superhero-update-media/internal/producer/model"
	"testing"
)

var shouldGenerateEncodeError = false

func mockPublishUpdateProfilePicture(producer *kafka.Writer, pp model.ProfilePicture) error {
	err := pp.Validate()
	if err != nil {
		return err
	}

	var sb bytes.Buffer

	var encoderValue interface{}
	encoderValue = pp

	if shouldGenerateEncodeError {
		encoderValue = make(chan int)
	}

	err = json.NewEncoder(&sb).Encode(encoderValue)
	if err != nil {
		return fmt.Errorf("encoder error")
	}

	return nil
}

func TestProducer_UpdateProfilePicture(t *testing.T) {
	tests := []struct {
		mockProducer            producer
		pp                      model.ProfilePicture
		willGenerateEncodeError bool
		shouldReturnError       bool
		expected                error
	}{
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "test-id",
				URL:         "https://www.test-url.com",
				Position:    0,
				CreatedAt:   "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       false,
			expected:                nil,
		},
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "",
				URL:         "https://www.test-url.com",
				Position:    0,
				CreatedAt:   "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture superhero id is empty"),
		},
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "test-id",
				URL:         "",
				Position:    0,
				CreatedAt:   "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture url is empty"),
		},
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "test-id",
				URL:         "https://www.test-url.com",
				Position:    -1,
				CreatedAt:   "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the position of the profile picture is invalid"),
		},
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "test-id",
				URL:         "https://www.test-url.com",
				Position:    0,
				CreatedAt:   "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture createdAt id is empty"),
		},
		{
			mockProducer: producer{
				Producer:             nil,
				updateProfilePicture: mockPublishUpdateProfilePicture,
			},
			pp: model.ProfilePicture{
				SuperheroID: "test-id",
				URL:         "https://www.test-url.com",
				Position:    0,
				CreatedAt:   "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: true,
			shouldReturnError:       true,
			expected:                fmt.Errorf("encoder error"),
		},
	}

	for _, test := range tests {
		shouldGenerateEncodeError = false

		if test.willGenerateEncodeError {
			shouldGenerateEncodeError = true
		}

		err := test.mockProducer.UpdateProfilePicture(test.pp)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if !test.shouldReturnError && err != nil {
			t.Fatal(err)
		}
	}
}
