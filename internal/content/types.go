package content

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
	Lead     string
	Points   []string
	Pictures []PictureSlot
}

type PerformanceNight struct {
	Day      string
	Singers  []string
	Keyboard string
}
