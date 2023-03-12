package generated_image

import (
	"github.com/snowflake-server/src/common"
)

type GeneratedImage struct {
	common.Model
	UserID uint   `json:"userId"`
	Prompt string `json:"prompt"`
	Hash   string `json:"hash"`
}

const (
	generatedImagePrefix = "generated_image:"
)

func GetGeneratedImageByID(id uint) (*GeneratedImage, error) {
	var gi GeneratedImage
	if err := gi.GetByID(id, generatedImagePrefix, &gi); err != nil {
		return nil, err
	}
	return &gi, nil
}

func GetGeneratedImageByHash(hash string) (*GeneratedImage, error) {
	var gi GeneratedImage
	err := gi.GetByColumn("hash", hash, generatedImagePrefix, &gi)
	if err != nil {
		return nil, err
	}
	return &gi, nil
}

func UpsertGeneratedImage(gi *GeneratedImage) error {
	if err := gi.UpsertObject(gi, generatedImagePrefix); err != nil {
		return err
	}
	return nil
}

func DeleteGeneratedImage(id uint) error {
	var gi GeneratedImage
	if err := gi.DeleteObject(id, generatedImagePrefix); err != nil {
		return err
	}
	return nil
}
