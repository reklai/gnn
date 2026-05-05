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
				"Available as contract work or one-time request",
			},
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
		},
	}
}

func cleaningCoverage() string {
	return "Residential properties, Commercial spaces, and Government agencies"
}

func cleaningPictures() []PictureSlot {
	return pictureSlots("Cleaning service photo", 4)
}
