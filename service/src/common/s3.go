package common

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	_ "image"
	"image/png"
	"io/ioutil"
	"os"
	_ "os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/chai2010/webp"
	"gopkg.in/yaml.v3"
)

type s3Config struct {
	S3 struct {
		AccessKeyId     string `yaml:"accessKeyId"`
		SecretAccessKey string `yaml:"secretAccessKey"`
		Region          string `yaml:"region"`
		Endpoint        string `yaml:"endpoint"`
		BucketName      string `yaml:"bucketName"`
	} `yaml:"s3"`
}

func UploadToS3(image string, keyName string) error {
	config, err := loadS3Config()
	if err != nil {
		panic(err)
	}

	pngData, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return fmt.Errorf("error decoding image data: %v", err)
	}

	img, err := png.Decode(bytes.NewReader(pngData))
	if err != nil {
		return fmt.Errorf("error decoding png image: %v", err)
	}

	var webpData bytes.Buffer
	options := webp.Options{Quality: 100}
	err = webp.Encode(&webpData, img, &options)
	if err != nil {
		return fmt.Errorf("error encoding webp image: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.S3.Region),
		Credentials: credentials.NewStaticCredentials(config.S3.AccessKeyId, config.S3.SecretAccessKey, ""),
		Endpoint:    aws.String(config.S3.Endpoint),
	})
	if err != nil {
		return fmt.Errorf("error creating session: %v", err)
	}

	s3Client := s3.New(sess)

	ctx := context.Background()

	_, err = s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(config.S3.BucketName),
		Key:         aws.String(keyName + ".webp"),
		Body:        bytes.NewReader(webpData.Bytes()),
		ContentType: aws.String("image/webp"),
	})

	if err != nil {
		return fmt.Errorf("error uploading image to S3: %v", err)
	}

	println("https://f002.backblazeb2.com/file/" + config.S3.BucketName + "/" + keyName + ".webp")

	return nil
}

func loadS3Config() (*s3Config, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(rootPath, "config.yaml")

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config s3Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
