package pages

import (
	"gnn/internal/components"
	"gnn/internal/content"
)

func HomeHeroActions(site content.SiteContent) []content.Action {
	return []content.Action{
		{Label: "Explore rentals", Href: "/space-rental"},
		{Label: "Call office", Href: components.PhoneHref(site.OfficePhone)},
	}
}

func RentalHeroActions(site content.SiteContent) []content.Action {
	return []content.Action{
		{Label: "Ask about units", Href: "#rental-contact"},
		{Label: "Call office", Href: components.PhoneHref(site.OfficePhone)},
	}
}

func CleaningHeroActions(site content.SiteContent) []content.Action {
	return []content.Action{
		{Label: "Request a quote", Href: "#cleaning-contact"},
		{Label: "Call office", Href: components.PhoneHref(site.OfficePhone)},
	}
}

func ConstructionHeroActions(site content.SiteContent) []content.Action {
	return []content.Action{
		{Label: "Book inspection", Href: "#construction-contact"},
		{Label: "Call office", Href: components.PhoneHref(site.OfficePhone)},
	}
}

func BarHeroActions(site content.SiteContent) []content.Action {
	return []content.Action{
		{Label: "See tonight's lineup", Href: "#bar-schedule"},
		{Label: "Call the bar", Href: components.PhoneHref(site.BarPhone)},
	}
}
