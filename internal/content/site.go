package content

type SiteContent struct {
	CompanyName          string
	OfficePhone          string
	OfficeHours          string
	BarPhone             string
	BarHours             string
	LunchWindow          string
	ServiceCards         []ServiceCard
	Rentals              []RentalOption
	CleaningCoverage     string
	CleaningGroups       []ServiceCategory
	CleaningPictures     []PictureSlot
	ConstructionCoverage []string
	ConstructionCards    []ConstructionCard
	ConstructionPictures []PictureSlot
	ConstructionSteps    []string
	SecurityDetails      ServiceCategory
	SecurityPictures     []PictureSlot
	BarSchedule          []PerformanceNight
	BarPictures          []PictureSlot
}

func NewSiteContent() SiteContent {
	return SiteContent{
		CompanyName:          "G&N Company",
		OfficePhone:          "680-488-2307",
		OfficeHours:          "Monday to Saturday · 8am-5pm",
		BarPhone:             "680-488-5711",
		BarHours:             "Thursday to Sunday · 10pm-2am",
		LunchWindow:          "11:30am-12:30pm",
		ServiceCards:         serviceCards(),
		Rentals:              rentalOptions(),
		CleaningCoverage:     cleaningCoverage(),
		CleaningGroups:       cleaningGroups(),
		CleaningPictures:     cleaningPictures(),
		ConstructionCoverage: constructionCoverage(),
		ConstructionCards:    constructionCards(),
		ConstructionPictures: constructionPictures(),
		ConstructionSteps:    constructionSteps(),
		SecurityDetails:      securityDetails(),
		SecurityPictures:     securityPictures(),
		BarSchedule:          barSchedule(),
		BarPictures:          barPictures(),
	}
}
