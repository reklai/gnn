package content

import "strings"

func rentalOptions() []RentalOption {
	return []RentalOption{
		{
			Key:         "studio",
			Title:       "Studio Units",
			RateLabel:   "Call for the latest monthly rate",
			LaundryNote: "No on-site laundry. A laundry service is about two minutes away.",
			Tone:        "harbor",
			Features: []string{
				"Single bedroom apartment",
				"Private shower room",
				"Shared kitchen for studio tenants",
				"Best fit for practical, lower-footprint living",
			},
			Pictures: pictureSlots("Studio unit photo", 4),
		},
		{
			Key:         "apartment",
			Title:       "Two-Bedroom Apartments",
			RateLabel:   "Call for the latest monthly rate",
			LaundryNote: "No on-site laundry. Nearby service remains the closest option.",
			Tone:        "grove",
			Features: []string{
				"Two bedrooms plus living room",
				"Private shower room and toilet",
				"Kitchen included in every apartment",
				"Good fit for roommates or small families",
			},
			Pictures: pictureSlots("Apartment photo", 4),
		},
		{
			Key:         "house",
			Title:       "Two-Bedroom Houses",
			RateLabel:   "Current house pricing by inquiry",
			LaundryNote: "Built-in laundry washing machine",
			Tone:        "sand",
			Features: []string{
				"Two bedrooms, living room, kitchen, shower room, and toilet",
				"Private yard space included",
				"More separation and room for longer-stay living",
			},
			Pictures: pictureSlots("House photo", 4),
		},
	}
}

func RentalFilterItems() []NavItem {
	return []NavItem{
		{Key: "all", Label: "All Units"},
		{Key: "studio", Label: "Studios"},
		{Key: "apartment", Label: "2BR Apartments"},
		{Key: "house", Label: "2BR Houses"},
	}
}

func NormalizeRentalFilter(filter string) string {
	filter = strings.ToLower(strings.TrimSpace(filter))
	switch filter {
	case "", "all":
		return "all"
	case "studio", "apartment", "house":
		return filter
	default:
		return "all"
	}
}

func FilterRentalOptions(rentals []RentalOption, filter string) []RentalOption {
	filter = NormalizeRentalFilter(filter)
	if filter == "all" {
		return rentals
	}

	filteredRentals := make([]RentalOption, 0, len(rentals))
	for _, rental := range rentals {
		if rental.Key == filter {
			filteredRentals = append(filteredRentals, rental)
		}
	}
	return filteredRentals
}
