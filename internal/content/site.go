package content

import "strings"

type Action struct {
	Label string
	Href  string
}

type NavItem struct {
	Key   string
	Label string
	Href  string
}

type Stat struct {
	Value string
	Label string
}

type FeaturePanel struct {
	Title  string
	Detail string
	Tone   string
}

type ServiceCard struct {
	Key        string
	Title      string
	Summary    string
	Href       string
	Tone       string
	Highlights []string
}

type RentalOption struct {
	Key         string
	Title       string
	Subtitle    string
	RateLabel   string
	LaundryNote string
	Tone        string
	Features    []string
	Gallery     []FeaturePanel
}

type ServiceCategory struct {
	Title   string
	Summary string
	Tone    string
	Points  []string
}

type ConstructionCard struct {
	Title   string
	Summary string
	Tone    string
	Points  []string
}

type PerformanceNight struct {
	Day      string
	Singers  []string
	Keyboard string
	Note     string
}

type InquiryFormConfig struct {
	Scope       string
	Title       string
	Intro       string
	SubmitLabel string
	TargetID    string
}

type SiteContent struct {
	CompanyName        string
	Tagline            string
	OfficePhone        string
	OfficeHours        string
	BarPhone           string
	BarHours           string
	LunchWindow        string
	HomeStats          []Stat
	RentalStats        []Stat
	ServiceCards       []ServiceCard
	HomePanels         []FeaturePanel
	RentalPanels       []FeaturePanel
	CleaningPanels     []FeaturePanel
	ConstructionPanels []FeaturePanel
	BarGallery         []FeaturePanel
	Rentals            []RentalOption
	CleaningGroups     []ServiceCategory
	ConstructionCards  []ConstructionCard
	ConstructionSteps  []string
	BarSchedule        []PerformanceNight
}

func NewSiteContent() SiteContent {
	return SiteContent{
		CompanyName: "GN&N Company",
		Tagline:     "Rentals, Cleaning, Construction, and Bar from one local business to yours",
		OfficePhone: "680-488-2307",
		OfficeHours: "Monday to Saturday · 8am-5pm",
		BarPhone:    "680-488-5711",
		BarHours:    "Thursday to Sunday · 10pm-2am",
		LunchWindow: "11:30am-12:30pm",
		HomeStats: []Stat{
			{Value: "4 Services under one brand"},
			{Value: "3 different rental property types"},
			{Value: "6 days the offie line is open"},
			{Value: "4 nights of live band each week"},
		},
		ServiceCards: []ServiceCard{
			{
				Key:     "rental",
				Title:   "Space Rental",
				Summary: "Browse Studios, Two Bedroom Apartments, and Two Bedroom Houses + Yard",
				Href:    "/space-rental",
				Tone:    "harbor",
				Highlights: []string{
					"Studios includes the following: one bedroom + shower room + shared kitchen place",
					"Two Bedroom Apartments includes the following: two bedroom + bathroom + living room + kitchen",
					"Two Bedroom Houses includes the following: two bedroom + bathroom + living room + kitchen + backyard",
					"Affordable rental costs",
				},
			},
			{
				Key:     "cleaning",
				Title:   "Cleaning Express",
				Summary: "Janitorial and Landscaping coverage for residential, commercial, and government agencies",
				Href:    "/cleaning-express",
				Tone:    "grove",
				Highlights: []string{
					"Recurring or One-Time Work",
					"Janitorial and Landscaping",
					"Pressure Wash",
					"Pricing and time flexible based on a per-inspection basis",
				},
			},
			{
				Key:     "construction",
				Title:   "Construction",
				Summary: "Repair work for walls, roofs, pipes, and furniture across multiple property",
				Href:    "/construction",
				Tone:    "ember",
				Highlights: []string{
					"Residential, commercial, and government agency scope",
					"Repair and renovation work",
					"Pricing and time flexible based on a per-inspection basis",
				},
			},
			{
				Key:     "bar",
				Title:   "Bayside Bar",
				Summary: "Late-Night live band from Thursday to Sunday with a stable singer lineup and rotating keyboard musicians",
				Href:    "/bar",
				Tone:    "night",
				Highlights: []string{
					"Open 10pm-2am",
					"Live Band",
				},
			},
		},
		HomePanels: []FeaturePanel{
			{Title: "Space Rental", Detail: "Studios, apartments, and houses grouped by layout.", Tone: "harbor"},
			{Title: "Cleaning Express", Detail: "Janitorial and landscaping support.", Tone: "grove"},
			{Title: "Construction", Detail: "Renovation work from walls and roofs to pipes and furniture.", Tone: "ember"},
			{Title: "Bayside Bar", Detail: "Live band Thursday to Sunday, 10pm-2am.", Tone: "night"},
		},
		RentalPanels: []FeaturePanel{
			{Title: "Studios", Detail: "Shared kitchen setup with compact everyday living.", Tone: "harbor"},
			{Title: "Apartments", Detail: "Two-bedroom apartments with complete living essentials.", Tone: "sand"},
			{Title: "Houses", Detail: "Two-bedroom houses with yard space included.", Tone: "grove"},
		},
		CleaningPanels: []FeaturePanel{
			{Title: "Floor care", Detail: "Waxing, buffing, sweeping, mopping, and vacuum detail work.", Tone: "grove"},
			{Title: "Exterior reset", Detail: "Water blasting, trash pickup, and site presentation cleanup.", Tone: "harbor"},
			{Title: "Landscaping", Detail: "Grass cutting, planting, trimming, and property edge maintenance.", Tone: "sand"},
		},
		ConstructionPanels: []FeaturePanel{
			{Title: "Walls & roofs", Detail: "Patch work and renovation support for aging structures.", Tone: "ember"},
			{Title: "Pipe fixes", Detail: "Repair-first scope for utility issues discovered on inspection.", Tone: "harbor"},
			{Title: "Furniture repairs", Detail: "Practical restoration work that closes out interior issues.", Tone: "sand"},
		},
		BarGallery: []FeaturePanel{
			{Title: "Thursdays", Detail: "Lisa Sandei and Jackie Franz open the week.", Tone: "night"},
			{Title: "Weekend crowd", Detail: "Friday and Saturday bring the fullest singer lineup.", Tone: "ember"},
			{Title: "Late close", Detail: "Every bar night runs through 2am.", Tone: "harbor"},
		},
		Rentals: []RentalOption{
			{
				Key:         "studio",
				Title:       "Studio Units",
				Subtitle:    "Compact living with a shared-kitchen setup.",
				RateLabel:   "Call for the latest monthly rate",
				LaundryNote: "No on-site laundry. A laundry service is about two minutes away.",
				Tone:        "harbor",
				Features: []string{
					"Single Bedroom Aparment",
					"All units have their own shower rooms",
					"Shared kitchen commonplace for all studio tenants",
					"Best fit for practical, lower-footprint living environment",
				},
				Gallery: []FeaturePanel{
					{Title: "Shared kitchen", Detail: "Central prep and cooking area.", Tone: "harbor"},
					{Title: "Compact studio", Detail: "Simple footprint designed for everyday use.", Tone: "sand"},
					{Title: "Shower access", Detail: "Private or shared depending on unit.", Tone: "grove"},
				},
			},
			{
				Key:         "apartment",
				Title:       "Two-Bedroom Apartments",
				Subtitle:    "Full apartment layouts with private essentials.",
				RateLabel:   "Call for the latest monthly rate",
				LaundryNote: "No on-site laundry. Nearby service remains the closest option.",
				Tone:        "sand",
				Features: []string{
					"Two bedrooms plus living room",
					"Private shower room and toilet",
					"Kitchen included in every apartment",
					"Better suited for roommates or small families",
				},
				Gallery: []FeaturePanel{
					{Title: "Living room", Detail: "Shared lounge space inside the unit.", Tone: "sand"},
					{Title: "Two bedrooms", Detail: "Separated sleeping areas.", Tone: "harbor"},
					{Title: "Kitchen + bath", Detail: "Private utility spaces built into the layout.", Tone: "grove"},
				},
			},
			{
				Key:         "house",
				Title:       "Two-Bedroom Houses",
				Subtitle:    "Standalone homes with additional outdoor space.",
				RateLabel:   "Current house pricing by inquiry",
				LaundryNote: "Built-in laundry washing machine",
				Tone:        "grove",
				Features: []string{
					"Two bedrooms, living room, kitchen, shower room, and toilet",
					"Private yard space included",
					"Good fit for tenants who want more separation and room, complete package experience",
				},
				Gallery: []FeaturePanel{
					{Title: "House layout", Detail: "Private rooms and fuller family footprint.", Tone: "grove"},
					{Title: "Outdoor yard", Detail: "Extra breathing room beyond the interior.", Tone: "sand"},
					{Title: "Complete setup", Detail: "Kitchen, bath, and living space in one home.", Tone: "ember"},
				},
			},
		},
		CleaningGroups: []ServiceCategory{
			{
				Title:   "Janitorial",
				Summary: "Detailed cleaning work priced after inspection, whether you need a recurring contract or a one-time reset.",
				Tone:    "harbor",
				Points: []string{
					"Waxing and buffing tile floors",
					"Residential, commercial, and government agency cleaning",
					"Water blasting and pressure washing",
					"Sweep, mop, vacuum, and wall-cleaning routines",
					"Trash removal and post-job reset",
				},
			},
			{
				Title:   "Landscaping",
				Summary: "Property-edge and grounds maintenance that keeps outdoor spaces presentable and operational.",
				Tone:    "grove",
				Points: []string{
					"Cutting grass and keeping grounds neat",
					"Planting trees and flowers",
					"Trimming and shaping work",
					"Picking up trash across the property",
					"Inspection-first scope for contract or one-off service",
				},
			},
		},
		ConstructionCards: []ConstructionCard{
			{
				Title:   "Renovation & Repair Scope",
				Summary: "Fixes move from obvious structural wear to practical everyday repairs after site review.",
				Tone:    "ember",
				Points: []string{
					"Wall and roof patching",
					"Pipe fixes and utility repair",
					"Furniture fixes and touch-up work",
				},
			},
			{
				Title:   "Property Types",
				Summary: "Work is available for the same mix of clients already reflected across the business.",
				Tone:    "harbor",
				Points: []string{
					"Residential properties",
					"Commercial spaces",
					"Government agency jobs",
				},
			},
			{
				Title:   "Inspection-Led Planning",
				Summary: "Pricing and timing are shaped after seeing the actual site conditions, not guessed up front.",
				Tone:    "sand",
				Points: []string{
					"Scope confirmed during inspection",
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
				Keyboard: "Maslyn or Brandon depending on schedule",
				Note:     "Opening night lineup with a smaller singer rotation.",
			},
			{
				Day:      "Friday",
				Singers:  []string{"Lisa Sandei", "Jackie Franz", "Sasa Naruo"},
				Keyboard: "Maslyn or Brandon depending on schedule",
				Note:     "Full singer lineup for the weekend crowd.",
			},
			{
				Day:      "Saturday",
				Singers:  []string{"Lisa Sandei", "Jackie Franz", "Sasa Naruo"},
				Keyboard: "Maslyn or Brandon depending on schedule",
				Note:     "Peak-night entertainment with the same full singer list.",
			},
			{
				Day:      "Sunday",
				Singers:  []string{"Jackie Franz", "Sasa Naruo"},
				Keyboard: "Maslyn or Brandon depending on schedule",
				Note:     "Late-week closer with a tighter singer set.",
			},
		},
	}
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

func HomeInquiryForm() InquiryFormConfig {
	return InquiryFormConfig{
		Scope:       "General Inquiry",
		Title:       "Start with a quick inquiry",
		Intro:       "Leave a name, callback number, and a short note. GN&N can route you to the right service line.",
		SubmitLabel: "Send inquiry",
		TargetID:    "home-inquiry-response",
	}
}

func RentalInquiryForm() InquiryFormConfig {
	return InquiryFormConfig{
		Scope:       "Space Rental",
		Title:       "Ask about units",
		Intro:       "Use this form for current rates, viewings, and unit questions.",
		SubmitLabel: "Request rental follow-up",
		TargetID:    "rental-inquiry-response",
	}
}

func CleaningInquiryForm() InquiryFormConfig {
	return InquiryFormConfig{
		Scope:       "Cleaning Express",
		Title:       "Request a cleaning inspection",
		Intro:       "Pricing and timing are set after inspection, so the fastest next step is a callback request.",
		SubmitLabel: "Request cleaning quote",
		TargetID:    "cleaning-inquiry-response",
	}
}

func ConstructionInquiryForm() InquiryFormConfig {
	return InquiryFormConfig{
		Scope:       "Construction",
		Title:       "Book a construction inspection",
		Intro:       "Share the problem area and GN&N can follow up with the next inspection window.",
		SubmitLabel: "Request inspection",
		TargetID:    "construction-inquiry-response",
	}
}
