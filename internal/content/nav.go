package content

func PrimaryNavItems() []NavItem {
	return []NavItem{
		{Key: "home", Label: "Home", Href: "/"},
		{Key: "rental", Label: "Space Rental", Href: "/space-rental"},
		{Key: "cleaning", Label: "Cleaning Express", Href: "/cleaning-express"},
		{Key: "construction", Label: "Construction", Href: "/construction"},
		{Key: "security", Label: "Security", Href: "/security"},
		{Key: "bar", Label: "Bar", Href: "/bar"},
	}
}
