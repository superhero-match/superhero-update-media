package service

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/superhero-match/superhero-update-media/internal/producer/model"
	"testing"
)

var shouldGenerateEncodeError = false

var (
	ErrDataBufferIsEmpty  = fmt.Errorf("data buffer passed into PutObject is empty or nil")
	ErrS3BucketKeyIsEmpty = fmt.Errorf("s3 bucket key passed into PutObject is empty")
)

type MockProducer interface {
	Close() error
	UpdateProfilePicture(pp model.ProfilePicture) error
}

type mockProducer struct {
	updateProfilePicture func(pp model.ProfilePicture) error
}

func (m *mockProducer) Close() error {
	return nil
}

func (m *mockProducer) UpdateProfilePicture(pp model.ProfilePicture) error {
	return m.updateProfilePicture(pp)
}

func mockPublishUpdateProfilePicture(pp model.ProfilePicture) error {
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

type MockAWS interface {
	PutObject(buffer []byte, key string) error
}

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
	mockProd := &mockProducer{
		updateProfilePicture: mockPublishUpdateProfilePicture,
	}

	mAws := mockAws{
		putObject: mockUploadObjectToS3,
	}

	mockService := &service{
		Producer: mockProd,
		AWS:      mAws,
	}

	buffer, err := b64.StdEncoding.DecodeString(testImgBase64)
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

func TestService_UpdateProfilePicture(t *testing.T) {
	mockProd := &mockProducer{
		updateProfilePicture: mockPublishUpdateProfilePicture,
	}

	mAws := mockAws{
		putObject: mockUploadObjectToS3,
	}

	mockService := &service{
		Producer: mockProd,
		AWS:      mAws,
	}

	tests := []struct {
		superheroID             string
		url                     string
		position                int64
		createdAt               string
		willGenerateEncodeError bool
		shouldReturnError       bool
		expected                error
	}{
		{
			superheroID:             "test-id",
			url:                     "https://www.test-url.com",
			position:                0,
			createdAt:               "2022-04-25T12:00:00",
			willGenerateEncodeError: false,
			shouldReturnError:       false,
			expected:                nil,
		},
		{
			superheroID:             "",
			url:                     "https://www.test-url.com",
			position:                0,
			createdAt:               "2022-04-25T12:00:00",
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture superhero id is empty"),
		},
		{
			superheroID:             "test-id",
			url:                     "",
			position:                0,
			createdAt:               "2022-04-25T12:00:00",
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture url is empty"),
		},
		{
			superheroID:             "test-id",
			url:                     "https://www.test-url.com",
			position:                -1,
			createdAt:               "2022-04-25T12:00:00",
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the position of the profile picture is invalid"),
		},
		{
			superheroID:             "test-id",
			url:                     "https://www.test-url.com",
			position:                0,
			createdAt:               "",
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("profile picture createdAt id is empty"),
		},
		{
			superheroID:             "test-id",
			url:                     "https://www.test-url.com",
			position:                0,
			createdAt:               "2022-04-25T12:00:00",
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

		err := mockService.UpdateProfilePicture(test.superheroID, test.url, test.position, test.createdAt)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if !test.shouldReturnError && err != nil {
			t.Fatal(err)
		}
	}
}
