package components

import (
	"net/url"
	"strings"
)

const assetVersion = "20260423-71"

func StaticAssetHref(path string) string {
	return path + "?v=" + assetVersion
}

func StylesheetHref() string {
	return StaticAssetHref("/static/styles.css")
}

func GalleryScriptHref() string {
	return StaticAssetHref("/static/gallery.js")
}

func PhoneHref(number string) string {
	digits := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, number)

	if digits == "" {
		return "#"
	}

	if strings.HasPrefix(digits, "680") {
		return "tel:+" + digits
	}

	return "tel:" + digits
}

func NavLinkClass(activeSection, itemKey string) string {
	if activeSection == itemKey {
		return "nav-link is-active"
	}
	return "nav-link"
}

func RentalFilterLinkClass(activeFilter, filterKey string) string {
	if activeFilter == filterKey {
		return "filter-link is-active"
	}
	return "filter-link"
}

func RentalFilterHref(filter string) string {
	filter = strings.TrimSpace(filter)
	if filter == "" || filter == "all" {
		return "/space-rental"
	}
	return "/space-rental?type=" + url.QueryEscape(filter)
}

func RentalCatalogPartialHref(filter string) string {
	return "/partials/rentals?type=" + url.QueryEscape(filter)
}

func CommaList(items []string) string {
	return strings.Join(items, ", ")
}
