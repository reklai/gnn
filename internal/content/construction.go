package content

func constructionCards() []ConstructionCard {
	return []ConstructionCard{
		{
			Title: "Renovate and Repair Scope",
			Tone:  "harbor",
			Points: []string{
				"Wall, floor, and roof patching",
				"Pipe fixes and utility repair",
				"Furniture fixes and touch-up work",
			},
			Pictures: pictureSlots("Renovation photo", 2),
		},
		{
			Title: "Property Coverage",
			Tone:  "grove",
			Points: []string{
				"Residential properties",
				"Commercial spaces",
				"Government agencies",
			},
			Pictures: pictureSlots("Coverage photo", 2),
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

func constructionSteps() []string {
	return []string{
		"Book an inspection and describe the issue.",
		"Receive a scope, estimate, and expected timeline.",
		"Approve either contract work or a one-time repair.",
		"Coordinate the work window directly with GN&N.",
	}
}
