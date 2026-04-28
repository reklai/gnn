package content

func barSchedule() []PerformanceNight {
	return []PerformanceNight{
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
	}
}

func barPictures() []PictureSlot {
	return pictureSlots("Bar stage photo", 4)
}
