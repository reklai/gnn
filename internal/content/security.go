package content

func securityDetails() ServiceCategory {
	return ServiceCategory{
		Title: "Security Services",
		Tone:  "night",
		Points: []string{
			"Security services for residential, commercial, and government agencies",
			"Security guards are trained, experienced, and reliable",
			"Security-related equipment costs are covered as part of the service package",
			"Contract flexibility ranges from short-term to long-term arrangements",
		},
	}
}

func securityPictures() []PictureSlot {
	return pictureSlots("Security service photo", 4)
}
