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

var validHairShapeOptions = []string{
	"short_hair",
	"very_short_hair",
	"absurdly_short_hair",
	"medium_hair",
	"long_hair",
	"very_long_hair",
	"absurdly_long_hair",
	"straight_hair",
	"curly_hair",
	"wavy_hair",
	"spiked_hair",
	"flipped_hair",
	"bluntbangs",
	"pointy_hair",
	"messy_hair",
	"Blunt_bangs",
	"parted_bangs",
	"crossed_bangs",
	"swept_bangs",
	"braided_bangs",
	"hair_over_one_eye",
	"hair_over_eyes",
	"hair_between_eyes",
	"asymmetrical_bangs",
}

var validHairStyleOptions = []string{
	"pixie_cut",
	"undercut",
	"cornrows",
	"dreadlocks",
	"braid",
	"braided_bangs",
	"front_braid",
	"side_braid",
	"french_braid",
	"crown_braid",
	"single_braid",
	"multiple_braids",
	"twin_braids",
	"tri_braids",
	"quad_braids",
	"hair_bun",
	"braided_bun",
	"single_hair_bun",
	"double_bun",
	"triple_bun",
	"cone_hair_bun",
	"doughnut_hair_bun",
	"heart_hair_bun",
	"one_side_up",
	"two_side_up",
	"ponytail",
	"folded_ponytail",
	"front_ponytail",
	"high_ponytail",
	"short_ponytail",
	"side_ponytail",
	"split_ponytail",
	"twintails",
	"low_twintails",
	"short_twintails",
	"uneven_twintails",
	"tri_tails",
	"quad_tails",
	"quin_tails",
}

var validFaceOptions = []string{
	"",
	"small_face",
	"small_medium_face",
	"medium_face",
	"large_face",
	"huge_face",
	"sharp_face",
	"pointed_face",
	"peanut_face",
	"round_face",
	"egg_shaped_face",
}

var validEyesOptions = []string{
	"gray_eyes",
	"white_eyes",
	"brown_eyes",
	"red_eyes",
	"pink_eyes",
	"orange_eyes",
	"yellow_eyes",
	"golden_eyes",
	"light_green_eyes",
	"green_eyes",
	"sky_blue_eyes",
	"blue_eyes",
	"purple_eyes",
	"black_eyes",
}

var validNoseOptions = []string{
	"",
	"small_nose",
	"long_nose",
	"slender_nose",
	"heavy_nose",
	"broad_nose",
	"high_end_nose",
	"wild_boogged_nose",
	"aquiline_nose",
}

var validMouthOptions = []string{
	"",
	"small_mouth",
	"big_mouth",
	"thin_lips",
	"thick_lips",
}

var validBodyOptions = []string{
	"slender",
}

var validBreastOptions = []string{
	"flat_chest",
	"small_breasts",
	"small_medium_breasts",
	"medium_breasts",
	"large_breasts",
	"huge_breasts",
	"gigantic_breasts",
}

type requestDrawFirstLover struct {
	//validate:"required,regexp=^([a-zA-Z]{2,12}|[가-힣]{2,6})$"
	Name           string `json:"name" `
	Race           int    `json:"race" validate:"required,min=1,max=2"`
	Sex            int    `json:"sex" validate:"required,min=1,max=2"`
	Age            int    `json:"age" validate:"required,min=18,max=25"`
	HairColor      string `json:"hairColor" validate:"required,eqfield=ValidHairColor"`
	HairShape      string `json:"hairShape"`
	HairStyle      string `json:"hairStyle"`
	Face           string `json:"face"`
	Eyes           string `json:"eyes"`
	Nose           string `json:"nose"`
	Mouth          string `json:"mouth"`
	Ears           string `json:"ears"`
	Body           string `json:"body"`
	Breast         string `json:"breast"`
	ValidHairColor string `json:"-"`
	ValidHairShape string `json:"-"`
	ValidHairStyle string `json:"-"`
	ValidFace      string `json:"-"`
	ValidEyes      string `json:"-"`
	ValidNose      string `json:"-"`
	ValidMouth     string `json:"-"`
	ValidBody      string `json:"-"`
	ValidBreast    string `json:"-"`
}

func (r *requestDrawFirstLover) Validate() error {
	if !common.ContainsString(validHairColorOptions, r.ValidHairColor) {
		return fmt.Errorf("validation failed: invalid HairColor")
	}
	if !common.ContainsString(validHairShapeOptions, r.ValidHairShape) {
		return fmt.Errorf("validation failed: invalid HairShape")
	}
	if !common.ContainsString(validHairStyleOptions, r.ValidHairStyle) {
		return fmt.Errorf("validation failed: invalid HairStyle")
	}
	if !common.ContainsString(validFaceOptions, r.ValidFace) {
		return fmt.Errorf("validation failed: invalid Face")
	}
	if !common.ContainsString(validEyesOptions, r.ValidEyes) {
		return fmt.Errorf("validation failed: invalid Eyes")
	}
	if !common.ContainsString(validNoseOptions, r.ValidNose) {
		return fmt.Errorf("validation failed: invalid Nose")
	}
	if !common.ContainsString(validMouthOptions, r.ValidMouth) {
		return fmt.Errorf("validation failed: invalid Mouth")
	}
	if !common.ContainsString(validBodyOptions, r.ValidBody) {
		return fmt.Errorf("validation failed: invalid Body")
	}
	if !common.ContainsString(validBreastOptions, r.ValidBreast) {
		return fmt.Errorf("validation failed: invalid Breast")
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
	req.ValidHairShape = req.HairShape
	req.ValidHairStyle = req.HairStyle
	req.ValidFace = req.Face
	req.ValidEyes = req.Eyes
	req.ValidNose = req.Nose
	req.ValidMouth = req.Mouth
	req.ValidBody = req.Body
	req.ValidBreast = req.Breast

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

	response.SendMessage(outgoing, 1, map[string]interface{}{"hash": hash})
}
