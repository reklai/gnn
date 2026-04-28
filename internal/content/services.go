package content

func serviceCards() []ServiceCard {
	return []ServiceCard{
		{
			Title:   "Space Rental",
			Summary: "Browse studios, two-bedroom apartments, and two-bedroom houses.",
			Href:    "/space-rental",
			Tone:    "harbor",
		},
		{
			Title:   "Cleaning Express",
			Summary: "Janitorial and landscaping for residential, commercial, and government properties.",
			Href:    "/cleaning-express",
			Tone:    "grove",
		},
		{
			Title:   "Construction",
			Summary: "Repair work for homes, walls, roofs, pipes, floors, and furniture.",
			Href:    "/construction",
			Tone:    "sand",
		},
		{
			Title:   "Bayside Bar",
			Summary: "Late-night live band from Thursday to Sunday with singers and keyboard lineup.",
			Href:    "/bar",
			Tone:    "night",
		},
	}
}
