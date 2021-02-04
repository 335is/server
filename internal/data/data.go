package data

import (
	"errors"
	"strconv"
	"strings"
)

// Band lists the members, name, and other info
type Band struct {
	Name    string   `yaml:"name"`
	Members []Member `yaml:"members"`
	Year    int      `yaml:"year"`
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
	beatles := Band{
		Name: "Beatles",
		Members: []Member{
			{
				Name: "John Lennon",
				Instruments: []string{
					"vocals",
					"guitar",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Paul McCartney",
				Instruments: []string{
					"vocals",
					"bass",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "George Harrison",
				Instruments: []string{
					"vocals",
					"guitar",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Ringo Starr",
				Instruments: []string{
					"drums",
					"vocals",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Billy Preston",
				Instruments: []string{
					"keyboards",
				},
				Founder: false,
				Current: false,
			},
		},
		Year: 1957,
	}

	dishwalla := Band{
		Name: "Dishwalla",
		Members: []Member{
			{
				Name: "J.R. Richards",
				Instruments: []string{
					"vocals",
					"guitar",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Rodney Browning",
				Instruments: []string{
					"guitar",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Scott Alexander",
				Instruments: []string{
					"bass",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "George Pendergast",
				Instruments: []string{
					"drums",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Jim Wood",
				Instruments: []string{
					"keyboards",
				},
				Founder: false,
				Current: false,
			},
			{
				Name: "Pete Maloney",
				Instruments: []string{
					"drums",
				},
				Founder: false,
				Current: false,
			},
			{
				Name: "Justin Fox",
				Instruments: []string{
					"vocals",
				},
				Founder: false,
				Current: true,
			},
		},
		Year: 1994,
	}

	eagles := Band{
		Name: "Eagles",
		Members: []Member{
			{
				Name: "Glenn Frey",
				Instruments: []string{
					"vocals",
					"guitar",
					"keyboards",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Don Henley",
				Instruments: []string{
					"drums",
					"vocals",
					"guitar",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Randy Meisner",
				Instruments: []string{
					"bass",
					"vocals",
					"guitar",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Bernie Leadon",
				Instruments: []string{
					"guitar",
					"vocals",
					"pedal steel guitar",
					"mandolin",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Don Felder",
				Instruments: []string{
					"guitar",
					"background vocals",
					"pedal steel guitar",
					"mandolin",
				},
				Founder: false,
				Current: false,
			},
			{
				Name: "Joe Walsh",
				Instruments: []string{
					"guitar",
					"background vocals",
					"keyboards",
				},
				Founder: false,
				Current: true,
			},
			{
				Name: "Timothy B. Schmit",
				Instruments: []string{
					"bass",
					"vocals",
				},
				Founder: false,
				Current: true,
			},
		},
		Year: 1971,
	}

	stp := Band{
		Name: "Stone Temple Pilots",
		Members: []Member{
			{
				Name: "Scott Weiland",
				Instruments: []string{
					"vocals",
				},
				Founder: true,
				Current: false,
			},
			{
				Name: "Dean DeLeo",
				Instruments: []string{
					"guitar",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Robert DeLeo",
				Instruments: []string{
					"bass",
					"vocals",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Eric Kretz",
				Instruments: []string{
					"drums",
				},
				Founder: true,
				Current: true,
			},
			{
				Name: "Chester Bennington",
				Instruments: []string{
					"vocals",
				},
				Founder: false,
				Current: false,
			},
			{
				Name: "Jeff Gutt",
				Instruments: []string{
					"vocals",
				},
				Founder: false,
				Current: true,
			},
		},
		Year: 1989,
	}

	bands = append(bands, beatles, dishwalla, eagles, stp)
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
