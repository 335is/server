package data

import (
	"errors"
	"strconv"
	"strings"
)

// Band lists the members, name, and other info
type Band struct {
	Name    string   `yaml:"name"`
	Year    int      `yaml:"year"`
	Members []Member `yaml:"members"`
}

// Member describes a band member
type Member struct {
	Name        string   `yaml:"name"`
	Instruments []string `yaml:"instruments"`
	Founder     bool     `yaml:"founder"`
	Current     bool     `yaml:"current"`
}

var (
	// ErrNotFound indicates the desired item was not found
	ErrNotFound = errors.New("not found")
)

var (
	bands []Band
)

func init() {
	fullcircle := Band{"Full Circle", 1998,
		[]Member{
			{"SLT", []string{"vocals", "guitar", "bass", "keyboards"}, true, true},
			{"DOB", []string{"vocals", "piano", "saxophone"}, true, true},
		},
	}

	beatles := Band{"Beatles", 1957,
		[]Member{
			{"John Lennon", []string{"vocals", "guitar"}, true, false},
			{"Paul McCartney", []string{"vocals", "bass"}, true, false},
			{"George Harrison", []string{"vocals", "guitar"}, true, false},
			{"Ringo Starr", []string{"drums", "vocals"}, true, false},
			{"Billy Preston", []string{"keyboards"}, false, false},
		},
	}

	dishwalla := Band{"Dishwalla", 1994,
		[]Member{
			{"J.R. Richards", []string{"vocals", "guitar"}, true, false},
			{"Rodney Browning", []string{"guitar"}, true, true},
			{"Scot Alexander", []string{"bass"}, true, true},
			{"George Pendergast", []string{"drums"}, true, true},
			{"Jim Wood", []string{"keyboards"}, false, false},
			{"Pete Maloney", []string{"drums"}, false, false},
			{"Justin Fox", []string{"vocals"}, false, true},
		},
	}

	eagles := Band{"Eagles", 1971,
		[]Member{
			{"Glenn Frey", []string{"vocals", "guitar", "keyboards"}, true, false},
			{"Don Henley", []string{"drums", "vocals", "guitar"}, true, true},
			{"Randy Meisner", []string{"bass", "vocals", "guitar"}, true, false},
			{"Bernie Leadon", []string{"guitar", "vocals", "pedal steel guitar", "mandolin"}, true, false},
			{"Don Felder", []string{"guitar", "pedal steel guitar", "mandolin", "background vocals"}, false, false},
			{"Joe Walsh", []string{"guitar", "background vocals", "keyboards"}, false, true},
			{"Timothy B. Schmit", []string{"bass", "vocals"}, false, true},
		},
	}

	stp := Band{"Stone Temple Pilots", 1989,
		[]Member{
			{"Scott Weiland", []string{"vocals"}, true, false},
			{"Robert DeLeo", []string{"bass", "vocals"}, true, true},
			{"Dean DeLeo", []string{"guitar"}, true, true},
			{"Eric Kretz", []string{"drums"}, true, true},
			{"Chester Bennington", []string{"vocals"}, false, false},
			{"Jeff Gutt", []string{"vocals"}, false, true},
		},
	}

	bands = append(bands, fullcircle, beatles, dishwalla, eagles, stp)
}

// GetBands returns a list of bands with all their info
func GetBands() ([]Band, error) {
	return bands, nil
}

// GetBandNames returns a list of band names
func GetBandNames() ([]string, error) {
	n := []string{}

	for _, b := range bands {
		n = append(n, b.Name)
	}

	return n, nil
}

// GetBand returns the info about the specified band
func GetBand(id string) (Band, error) {
	for i, b := range bands {
		if strings.EqualFold(id, b.Name) {
			return bands[i], nil
		}
	}

	return Band{}, ErrNotFound
}

// GetBandName returns the name of the specified band
func GetBandName(bandID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			return b.Name, nil
		}
	}

	return "", ErrNotFound
}

// GetBandMembers returns a list of members in the specified band
func GetBandMembers(bandID string) ([]Member, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			return b.Members, nil
		}
	}

	return []Member{}, ErrNotFound
}

// GetBandMember returns the specified member in the specified band
func GetBandMember(bandID string, memberID string) (Member, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					return m, nil
				}
			}

		}
	}

	return Member{}, ErrNotFound
}

// GetBandMemberName returns the name of the specified member in the specified band
func GetBandMemberName(bandID string, memberID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					return m.Name, nil
				}
			}

		}
	}

	return "", ErrNotFound
}

// GetBandMemberInstruments returns the name of the specified member in the specified band
func GetBandMemberInstruments(bandID string, memberID string) ([]string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					return m.Instruments, nil
				}
			}

		}
	}

	return []string{}, ErrNotFound
}

// GetBandMemberInstrument returns the instrument played by the specified member in the specified band
func GetBandMemberInstrument(bandID string, memberID string, instrumentID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					for _, i := range m.Instruments {
						if strings.EqualFold(instrumentID, i) {
							return i, nil
						}
					}
				}
			}

		}
	}

	return "", ErrNotFound
}

// GetBandMemberFounder returns if the specified member is a founder
func GetBandMemberFounder(bandID string, memberID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					return strconv.FormatBool(m.Founder), nil
				}
			}

		}
	}

	return "", ErrNotFound
}

// GetBandMemberCurrent returns if the specified member is currently in the band
func GetBandMemberCurrent(bandID string, memberID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			for _, m := range b.Members {
				if strings.EqualFold(memberID, m.Name) {
					return strconv.FormatBool(m.Current), nil
				}
			}

		}
	}

	return "", ErrNotFound
}

// GetBandYear returns the year the specified band formed
func GetBandYear(bandID string) (string, error) {
	for _, b := range bands {
		if strings.EqualFold(bandID, b.Name) {
			return strconv.Itoa(b.Year), nil
		}
	}

	return "", ErrNotFound
}
