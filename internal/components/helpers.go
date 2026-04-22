package components

import (
	"net/url"
	"strings"
)

func PhoneHref(number string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "", ".", "")
	return "tel:" + replacer.Replace(number)
}

func NavLinkClass(activeSection, itemKey string) string {
	if activeSection == itemKey {
		return "nav-link is-active"
	}
	return "nav-link"
}

func ActionClass(primary bool) string {
	if primary {
		return "button"
	}
	return "button button-secondary"
}

func RentalFilterLinkClass(activeFilter, filterKey string) string {
	if activeFilter == filterKey {
		return "filter-link is-active"
	}
	return "filter-link"
}

func InquiryTargetSelector(id string) string {
	return "#" + id
}

func InquiryStatusClass(isSuccess bool) string {
	if isSuccess {
		return "status-card is-success"
	}
	return "status-card is-error"
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

func InquiryAction() string {
	return "/partials/inquiry"
}

func SchedulePartialHref() string {
	return "/partials/bar-schedule"
}

func CommaList(items []string) string {
	return strings.Join(items, ", ")
}
