// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package GraphQL_Service

import (
	"fmt"
	"io"
	"strconv"
)

// A Film resource is a single film.
type Film struct {
	ID           string      `json:"id"`
	Title        string      `json:"title"`
	EpisodeID    *int        `json:"episode_id"`
	OpeningCrawl *string     `json:"opening_crawl"`
	Director     *string     `json:"director"`
	Producer     *string     `json:"producer"`
	ReleaseDate  *string     `json:"release_date"`
	Species      []*Specie   `json:"species"`
	Starships    []*Starship `json:"starships"`
	Vehicles     []*Vehicle  `json:"vehicles"`
	Characters   []*People   `json:"characters"`
	Planets      []*Planet   `json:"planets"`
}

// A connection to a list of items.
type FilmConnection struct {
	PageInfo   PageInfo   `json:"pageInfo"`
	Edges      []FilmEdge `json:"edges"`
	TotalCount int        `json:"totalCount"`
}

// An edge in a connection.
type FilmEdge struct {
	Node   *Film  `json:"node"`
	Cursor string `json:"cursor"`
}

// Information about pagination in a connection.
type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

// A People resource is an individual person or character within the Star Wars universe.
type People struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	BirthYear *string     `json:"birth_year"`
	EyeColor  *string     `json:"eye_color"`
	Gender    *string     `json:"gender"`
	HairColor *string     `json:"hair_color"`
	Height    *string     `json:"height"`
	Mass      *string     `json:"mass"`
	SkinColor *string     `json:"skin_color"`
	Homeworld *Planet     `json:"homeworld"`
	Films     []*Film     `json:"films"`
	Species   []*Specie   `json:"species"`
	Starships []*Starship `json:"starships"`
	Vehicles  []*Vehicle  `json:"vehicles"`
}

// A connection to a list of items.
type PeopleConnection struct {
	PageInfo   PageInfo     `json:"pageInfo"`
	Edges      []PeopleEdge `json:"edges"`
	TotalCount int          `json:"totalCount"`
}

// An edge in a connection.
type PeopleEdge struct {
	Node   *People `json:"node"`
	Cursor string  `json:"cursor"`
}

// A Planet resource is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Diameter       *string   `json:"diameter"`
	RotationPeriod *string   `json:"rotation_period"`
	OrbitalPeriod  *string   `json:"orbital_period"`
	Gravity        *string   `json:"gravity"`
	Population     *string   `json:"population"`
	Climate        *string   `json:"climate"`
	Terrain        *string   `json:"terrain"`
	SurfaceWater   *string   `json:"surface_water"`
	Residents      []*People `json:"residents"`
	Films          []*Film   `json:"films"`
}

// A connection to a list of items.
type PlanetConnection struct {
	PageInfo   PageInfo     `json:"pageInfo"`
	Edges      []PlanetEdge `json:"edges"`
	TotalCount int          `json:"totalCount"`
}

// An edge in a connection.
type PlanetEdge struct {
	Node   *Planet `json:"node"`
	Cursor string  `json:"cursor"`
}

// A Species resource is a type of person or character within the Star Wars Universe.
type Specie struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Classification  *string    `json:"classification"`
	Designation     *string    `json:"designation"`
	AverageHeight   *string    `json:"average_height"`
	AverageLifespan *string    `json:"average_lifespan"`
	EyeColors       *string    `json:"eye_colors"`
	HairColors      *string    `json:"hair_colors"`
	SkinColors      *string    `json:"skin_colors"`
	Language        *string    `json:"language"`
	Homeworld       *Planet    `json:"homeworld"`
	Vehicle         []*Vehicle `json:"Vehicle"`
	Films           []*Film    `json:"films"`
	People          []*People  `json:"People"`
}

// A connection to a list of items.
type SpecieConnection struct {
	PageInfo   PageInfo     `json:"pageInfo"`
	Edges      []SpecieEdge `json:"edges"`
	TotalCount int          `json:"totalCount"`
}

// An edge in a connection.
type SpecieEdge struct {
	Node   *Specie `json:"node"`
	Cursor string  `json:"cursor"`
}

// A Starship resource is a single transport craft that has hyperdrive capability.
type Starship struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Model                *string   `json:"model"`
	StarshipClass        *string   `json:"starship_class"`
	Manufacturer         *string   `json:"manufacturer"`
	CostInCredits        *string   `json:"cost_in_credits"`
	Length               *string   `json:"length"`
	Crew                 *string   `json:"crew"`
	Passengers           *string   `json:"passengers"`
	MaxAtmospheringSpeed *string   `json:"max_atmosphering_speed"`
	HyperdriveRating     *string   `json:"hyperdrive_rating"`
	MGLT                 *string   `json:"MGLT"`
	CargoCapacity        *string   `json:"cargo_capacity"`
	Consumables          *string   `json:"consumables"`
	Films                []*Film   `json:"films"`
	Pilots               []*People `json:"pilots"`
}

// A connection to a list of items.
type StarshipConnection struct {
	PageInfo   PageInfo       `json:"pageInfo"`
	Edges      []StarshipEdge `json:"edges"`
	TotalCount int            `json:"totalCount"`
}

// An edge in a connection.
type StarshipEdge struct {
	Node   *Starship `json:"node"`
	Cursor string    `json:"cursor"`
}

// A Vehicle resource is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	Model                *string       `json:"model"`
	VehicleClass         *VehicleClass `json:"vehicle_class"`
	Manufacturer         *string       `json:"manufacturer"`
	Length               *string       `json:"length"`
	CostInCredits        *string       `json:"cost_in_credits"`
	Crew                 *string       `json:"crew"`
	Passengers           *string       `json:"passengers"`
	MaxAtmospheringSpeed *string       `json:"max_atmosphering_speed"`
	CargoCapacity        *string       `json:"cargo_capacity"`
	Consumables          *string       `json:"consumables"`
	Films                []*Film       `json:"films"`
	Pilots               []*People     `json:"pilots"`
}

// A connection to a list of items.
type VehicleConnection struct {
	PageInfo   PageInfo      `json:"pageInfo"`
	Edges      []VehicleEdge `json:"edges"`
	TotalCount int           `json:"totalCount"`
}

// An edge in a connection.
type VehicleEdge struct {
	Node   *Vehicle `json:"node"`
	Cursor string   `json:"cursor"`
}

// The genders of people in the Star Wars universe.
type Gender string

const (
	// Does not have a gender.
	GenderNa Gender = "NA"
	// Male.
	GenderMale Gender = "MALE"
	// Female.
	GenderFemale Gender = "FEMALE"
)

func (e Gender) IsValid() bool {
	switch e {
	case GenderNa, GenderMale, GenderFemale:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The class of vehicle in the Star Wars universe.
type VehicleClass string

const (
	// Wheeled.
	VehicleClassWheeled VehicleClass = "WHEELED"
	// Repulsorcraft.
	VehicleClassRepulsocraft VehicleClass = "REPULSOCRAFT"
)

func (e VehicleClass) IsValid() bool {
	switch e {
	case VehicleClassWheeled, VehicleClassRepulsocraft:
		return true
	}
	return false
}

func (e VehicleClass) String() string {
	return string(e)
}

func (e *VehicleClass) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VehicleClass(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VehicleClass", str)
	}
	return nil
}

func (e VehicleClass) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
