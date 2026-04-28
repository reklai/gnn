package content

func cleaningGroups() []ServiceCategory {
	return []ServiceCategory{
		{
			Title: "Janitorial Cleaning",
			Tone:  "harbor",
			Points: []string{
				"Sweep, mop, vacuum, and wall-cleaning routines",
				"Bathroom and office deep cleaning",
				"Waxing and buffing tile floors",
				"Water blasting and pressure washing concrete",
				"Residential, commercial, and government coverage",
				"Available as contract work or one-time request",
			},
			Pictures: pictureSlots("Janitorial photo", 4),
		},
		{
			Title: "Landscaping",
			Tone:  "grove",
			Points: []string{
				"Grass cutting and trimming",
				"Planting trees and flowers",
				"Trash pickup across the property",
				"Available as contract work or one-time request",
			},
			Pictures: pictureSlots("Landscaping photo", 4),
		},
	}
}
