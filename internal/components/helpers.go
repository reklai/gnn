package components

import (
	"net/url"
	"strings"
)

const assetVersion = "20260423-31"

func StaticAssetHref(path string) string {
	return path + "?v=" + assetVersion
}

func StylesheetHref() string {
	return StaticAssetHref("/static/styles.css")
}

func GalleryScriptHref() string {
	return StaticAssetHref("/static/gallery.js")
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
