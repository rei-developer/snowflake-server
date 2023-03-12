package common

import (
	_ "image"
	_ "os"
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
	/*

		// Read the configuration file
		configData, err := ioutil.ReadFile("./config.yaml")
		if err != nil {
			return err
		}

		// Parse the configuration
		var config s3Config
		err = yaml.Unmarshal(configData, &config)
		if err != nil {
			return err
		}

		// Decode the base64 encoded image to a PNG image
		pngData, err := base64.StdEncoding.DecodeString(image)
		if err != nil {
			return fmt.Errorf("error decoding image data: %v", err)
		}

		// Convert the PNG image data to an image.Image object
		img, err := png.Decode(bytes.NewReader(pngData))
		if err != nil {
			return fmt.Errorf("error decoding png image: %v", err)
		}

		// Convert the image to WebP format
		var webpData bytes.Buffer
		options := webp.Options{Quality: 100}
		err = webp.Encode(&webpData, img, &options)
		if err != nil {
			return fmt.Errorf("error encoding webp image: %v", err)
		}

		// Upload the WebP data to S3
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
	*/
	return nil
}
