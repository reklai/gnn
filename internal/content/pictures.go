package content

import (
	"strconv"
	"strings"
)

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

func pictureSlots(prefix string, count int) []PictureSlot {
	if count < 1 {
		return nil
	}

	slots := make([]PictureSlot, 0, count)
	for index := 0; index < count; index++ {
		label := prefix + " " + strconv.Itoa(index+1)
		slots = append(slots, PictureSlot{
			Label:   label,
			Alt:     label,
			Caption: label,
		})
	}
	return slots
}
