package graph

import (
	"github.com/pangaunn/gqlgen-workshop/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	datasource RamDatasource
}

type RamDatasource struct {
	humans    map[string]model.Human
	droid     map[string]model.Droid
	starships map[string]model.Starship
	reviews   map[model.Episode][]*model.Review
}

func NewResolver() Resolver {
	r := Resolver{}

	ramDatasource := RamDatasource{}
	ramDatasource.humans = map[string]model.Human{
		"1000": {
			CharacterFields: model.CharacterFields{
				ID:        "1000",
				Name:      "Luke Skywalker",
				FriendIds: []string{"1002", "1003", "2000", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.72,
			Mass:         77,
			StarshipIds:  []string{"3001", "3003"},
		},
		"1001": {
			CharacterFields: model.CharacterFields{
				ID:        "1001",
				Name:      "Darth Vader",
				FriendIds: []string{"1004"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 2.02,
			Mass:         136,
			StarshipIds:  []string{"3002"},
		},
		"1002": {
			CharacterFields: model.CharacterFields{
				ID:        "1002",
				Name:      "Han Solo",
				FriendIds: []string{"1000", "1003", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.8,
			Mass:         80,
			StarshipIds:  []string{"3000", "3003"},
		},
		"1003": {
			CharacterFields: model.CharacterFields{
				ID:        "1003",
				Name:      "Leia Organa",
				FriendIds: []string{"1000", "1002", "2000", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			HeightMeters: 1.5,
			Mass:         49,
		},
		"1004": {
			CharacterFields: model.CharacterFields{
				ID:        "1004",
				Name:      "Wilhuff Tarkin",
				FriendIds: []string{"1001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope},
			},
			HeightMeters: 1.8,
			Mass:         0,
		},
	}

	ramDatasource.droid = map[string]model.Droid{
		"2000": {
			CharacterFields: model.CharacterFields{
				ID:        "2000",
				Name:      "C-3PO",
				FriendIds: []string{"1000", "1002", "1003", "2001"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			PrimaryFunction: "Protocol",
		},
		"2001": {
			CharacterFields: model.CharacterFields{
				ID:        "2001",
				Name:      "R2-D2",
				FriendIds: []string{"1000", "1002", "1003"},
				AppearsIn: []model.Episode{model.EpisodeNewhope, model.EpisodeEmpire, model.EpisodeJedi},
			},
			PrimaryFunction: "Astromech",
		},
	}

	ramDatasource.starships = map[string]model.Starship{
		"3000": {
			ID:   "3000",
			Name: "Millennium Falcon",
			History: [][]int{
				{1, 2},
				{4, 5},
				{1, 2},
				{3, 2},
			},
			Length: 34.37,
		},
		"3001": {
			ID:   "3001",
			Name: "X-Wing",
			History: [][]int{
				{6, 4},
				{3, 2},
				{2, 3},
				{5, 1},
			},
			Length: 12.5,
		},
		"3002": {
			ID:   "3002",
			Name: "TIE Advanced x1",
			History: [][]int{
				{3, 2},
				{7, 2},
				{6, 4},
				{3, 2},
			},
			Length: 9.2,
		},
		"3003": {
			ID:   "3003",
			Name: "Imperial shuttle",
			History: [][]int{
				{1, 7},
				{3, 5},
				{5, 3},
				{7, 1},
			},
			Length: 20,
		},
	}

	ramDatasource.reviews = map[model.Episode][]*model.Review{}

	r.datasource = ramDatasource
	return r
}
