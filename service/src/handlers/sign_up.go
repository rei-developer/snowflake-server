package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/snowflake-server/src/common"
	"github.com/snowflake-server/src/novelai"
	"github.com/snowflake-server/src/response"
)

var validHairColorOptions = []string{
	"white_hair",
	"brown_hair",
	"red_hair",
	"pink_hair",
	"orange_hair",
	"yellow_hair",
	"Blonde_hair",
	"light_green_hair",
	"green_hair",
	"sky_blue_hair",
	"blue_hair",
	"purple_hair",
	"black_hair",
}

type requestDrawFirstLover struct {
	Name string `json:"name" `
	//validate:"required,regexp=^([a-zA-Z]{2,12}|[가-힣]{2,6})$"
	Race      int    `json:"race" validate:"required,min=1,max=2"`
	Sex       int    `json:"sex" validate:"required,min=1,max=2"`
	Age       int    `json:"age" validate:"required,min=18,max=25"`
	HairColor string `json:"hairColor" validate:"required,eqfield=ValidHairColor"`
	HairShape string `json:"hairShape"`
	HairStyle string `json:"hairStyle"`
	Face      string `json:"face"`
	Eyes      string `json:"eyes"`
	Nose      string `json:"nose"`
	Mouth     string `json:"mouth"`
	Ears      string `json:"ears"`
	Body      string `json:"body"`
	Breast    string `json:"breast"`

	ValidHairColor string `json:"-"`
}

func (r *requestDrawFirstLover) Validate() error {
	if !common.ContainsString(validHairColorOptions, r.ValidHairColor) {
		return fmt.Errorf("validation failed: invalid HairColor")
	}
	return nil
}

func HandleDrawFirstLover(payload []byte, outgoing chan []byte) {
	var req requestDrawFirstLover
	if err := json.Unmarshal(payload, &req); err != nil {
		fmt.Printf("unmarshalling failed: %v\n", err)
		return
	}

	req.ValidHairColor = req.HairColor

	if err := common.ValidateStruct(req); err != nil {
		fmt.Printf("validation failed: %v\n", err)
		return
	}

	if err := req.Validate(); err != nil {
		fmt.Printf("validation failed: %v\n", err)
		return
	}

	values := []string{
		"1 girl",
		req.HairColor,
		req.HairShape,
		req.HairStyle,
		req.Face,
		req.Eyes,
		req.Nose,
		req.Mouth,
		req.Ears,
		req.Body,
		req.Breast,
		"black school_uniform, white blouse, red ribbon, forest",
	}

	input := strings.Join(values, ", ")

	hash := novelai.GenerateImage(input)

	response.SendMessage(outgoing, 134, map[string]interface{}{"hash": hash})
}
