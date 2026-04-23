package content

import (
	"strconv"
	"strings"
)

type NavItem struct {
	Key   string
	Label string
	Href  string
}

type ServiceCard struct {
	Title   string
	Summary string
	Href    string
	Tone    string
}

type PictureSlot struct {
	Label    string
	ThumbSrc string
	FullSrc  string
	Alt      string
	Caption  string
}

func (slot PictureSlot) DisplayLabel() string {
	switch {
	case strings.TrimSpace(slot.Caption) != "":
		return strings.TrimSpace(slot.Caption)
	case strings.TrimSpace(slot.Label) != "":
		return strings.TrimSpace(slot.Label)
	case strings.TrimSpace(slot.Alt) != "":
		return strings.TrimSpace(slot.Alt)
	default:
		return "Photo"
	}
}

func (slot PictureSlot) EffectiveAlt() string {
	if strings.TrimSpace(slot.Alt) != "" {
		return strings.TrimSpace(slot.Alt)
	}
	return slot.DisplayLabel()
}

func (slot PictureSlot) EffectiveThumbSrc() string {
	if strings.TrimSpace(slot.ThumbSrc) != "" {
		return strings.TrimSpace(slot.ThumbSrc)
	}
	if strings.TrimSpace(slot.FullSrc) != "" {
		return strings.TrimSpace(slot.FullSrc)
	}
	return ""
}

func (slot PictureSlot) EffectiveFullSrc() string {
	if strings.TrimSpace(slot.FullSrc) != "" {
		return strings.TrimSpace(slot.FullSrc)
	}
	if strings.TrimSpace(slot.ThumbSrc) != "" {
		return strings.TrimSpace(slot.ThumbSrc)
	}
	return ""
}

type RentalOption struct {
	Key         string
	Title       string
	RateLabel   string
	LaundryNote string
	Tone        string
	Features    []string
	Pictures    []PictureSlot
}

type ServiceCategory struct {
	Title    string
	Tone     string
	Points   []string
	Pictures []PictureSlot
}

type ConstructionCard struct {
	Title    string
	Tone     string
	Points   []string
	Pictures []PictureSlot
}

type PerformanceNight struct {
	Day      string
	Singers  []string
	Keyboard string
}

type SiteContent struct {
	CompanyName       string
	OfficePhone       string
	OfficeHours       string
	BarPhone          string
	BarHours          string
	LunchWindow       string
	ServiceCards      []ServiceCard
	Rentals           []RentalOption
	CleaningGroups    []ServiceCategory
	ConstructionCards []ConstructionCard
	ConstructionSteps []string
	BarSchedule       []PerformanceNight
	BarPictures       []PictureSlot
}

func NewSiteContent() SiteContent {
	return SiteContent{
		CompanyName: "GN&N Company",
		OfficePhone: "680-488-2307",
		OfficeHours: "Monday to Saturday · 8am-5pm",
		BarPhone:    "680-488-5711",
		BarHours:    "Thursday to Sunday · 10pm-2am",
		LunchWindow: "11:30am-12:30pm",
		ServiceCards: []ServiceCard{
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
		},
		Rentals: []RentalOption{
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
		},
		CleaningGroups: []ServiceCategory{
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
		},
		ConstructionCards: []ConstructionCard{
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
		},
		ConstructionSteps: []string{
			"Book an inspection and describe the issue.",
			"Receive a scope, estimate, and expected timeline.",
			"Approve either contract work or a one-time repair.",
			"Coordinate the work window directly with GN&N.",
		},
		BarSchedule: []PerformanceNight{
			{
				Day:      "Thursday",
				Singers:  []string{"Lisa Sandei", "Jackie Franz"},
				Keyboard: "Maslyn",
			},
			{
				Day:      "Friday",
				Singers:  []string{"Lisa Sandei", "Jackie Franz", "Sasa Naruo"},
				Keyboard: "Maslyn",
			},
			{
				Day:      "Saturday",
				Singers:  []string{"Lisa Sandei", "Jackie Franz", "Sasa Naruo"},
				Keyboard: "Brandon",
			},
			{
				Day:      "Sunday",
				Singers:  []string{"Jackie Franz", "Sasa Naruo"},
				Keyboard: "Brandon",
			},
		},
		BarPictures: pictureSlots("Bar stage photo", 4),
	}
}

func pictureSlots(prefix string, count int) []PictureSlot {
	if count < 1 {
		return nil
	}

	slots := make([]PictureSlot, 0, count)
	for index := 0; index < count; index++ {
		slots = append(slots, PictureSlot{
			Label:   prefix + " " + strconv.Itoa(index+1),
			Alt:     prefix + " " + strconv.Itoa(index+1),
			Caption: prefix + " " + strconv.Itoa(index+1),
		})
	}
	return slots
}

func PrimaryNavItems() []NavItem {
	return []NavItem{
		{Key: "home", Label: "Home", Href: "/"},
		{Key: "rental", Label: "Space Rental", Href: "/space-rental"},
		{Key: "cleaning", Label: "Cleaning Express", Href: "/cleaning-express"},
		{Key: "construction", Label: "Construction", Href: "/construction"},
		{Key: "bar", Label: "Bar", Href: "/bar"},
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
