package content

func constructionCards() []ConstructionCard {
	return []ConstructionCard{
		{
			Title: "Construction Estimates & Repair Work",
			Tone:  "harbor",
			Lead:  "Estimations for construction-related jobs may include:",
			Points: []string{
				"Renovation and repair work for painting, walls, floors, roofs, fencing, and general property work",
				"Furniture repair, touch-up work, and practical property improvements",
				"Electrical services for homes and properties",
				"Plumbing repairs and installations for toilets, pipelines, and septic tanks",
			},
		},
		{
			Title: "Inspection Pricing Model",
			Tone:  "sand",
			Points: []string{
				"Pricing range confirmed after inspection",
				"Timeline adjusted to job size and complexity",
				"Available as contract work or one-time request",
			},
		},
	}
}

func constructionCoverage() []string {
	return []string{
		"Residential properties",
		"Commercial spaces",
		"Government agencies",
	}
}

func constructionPictures() []PictureSlot {
	return pictureSlots("Construction photo", 4)
}

func constructionSteps() []string {
	return []string{
		"Book an inspection and describe the issue.",
		"Receive a scope, estimate, and expected timeline.",
		"Approve either contract work or a one-time repair.",
		"Coordinate the work window directly with G&N.",
	}
}
