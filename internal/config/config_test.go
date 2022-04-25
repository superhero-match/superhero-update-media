package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	// App configuration.
	assert.Equal(t, ":7100", cfg.App.Port, "The port should be :7100.")
	assert.Equal(t, "2006-01-02T15:04:05", cfg.App.TimeFormat, "The time format should be 2006-01-02T15:04:05.")

	// AWS configuration.
	assert.Equal(t, "us-west-2", cfg.Aws.Region, "The AWS region should be us-west-2.")
	assert.Equal(t, "superheroes-pictures", cfg.Aws.SuperheroesS3Bucket, "The AWS bucket should be superheroes-pictures.")
	assert.Equal(t, "d3pfwtk1dtfl92.cloudfront.net", cfg.Aws.CdnURL, "The AWS cdn url should be d3pfwtk1dtfl92.cloudfront.net.")
	assert.Equal(t, "jpg", cfg.Aws.ImageExtension, "The AWS image extension should be jpg.")
	assert.Equal(t, "image/jpg", cfg.Aws.ContentType, "The AWS content type should be image/jpg.")
	assert.Equal(t, "base64", cfg.Aws.ContentEncoding, "The AWS content encoding should be base64.")

	// Kafka producer.
	assert.Equal(t, "localhost:9092", cfg.Producer.Brokers, "The Kafka producer brokers should be localhost:9092.")
	assert.Equal(t, "update.municipality.profilepicture", cfg.Producer.Topic, "The Kafka producer topic should be update.municipality.profilepicture.")
	assert.Equal(t, 1, cfg.Producer.BatchSize, "The Kafka producer batch size should be 1.")
	assert.Equal(t, 10, cfg.Producer.BatchTimeout, "The Kafka producer batch timeout should be 10.")
}
