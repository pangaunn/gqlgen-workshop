package main_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/pangaunn/gqlgen-workshop/graph"
	"github.com/pangaunn/gqlgen-workshop/graph/generated"
	"github.com/stretchr/testify/require"
)

var c *client.Client

func setupClient() *client.Client {
	r := graph.NewResolver()
	config := generated.Config{Resolvers: &r}
	s := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	return client.New(s)
}

func TestMain(m *testing.M) {
	c = setupClient()
	m.Run()
}

func TestQueryStarShip(t *testing.T) {
	t.Run("X-Wing", func(t *testing.T) {
		var resp struct {
			Starship struct {
				Name string
			}
		}
		c.MustPost(`
		{ 
			starship(id:"3001") {
				name
			}
		}`, &resp)

		require.Equal(t, "X-Wing", resp.Starship.Name)
	})

	t.Run("length", func(t *testing.T) {
		var resp struct {
			Starship struct {
				Name    string
				Length  float64
				History [][]int
			}
		}
		c.MustPost(`
		{ 
			starship(id:"3001") {
				name
				history
				length
			}
		}`, &resp)

		require.Equal(t, "X-Wing", resp.Starship.Name)
		require.Equal(t, 12.5, resp.Starship.Length)
		require.Len(t, resp.Starship.History, 4)
		require.Len(t, resp.Starship.History[0], 2)
	})

	t.Run("Imperialshuttle FOOT length", func(t *testing.T) {
		var resp struct {
			Starship struct {
				Name   string
				Length float64
			}
		}
		c.MustPost(`
		{ 
			starship(id:"3003") {
				name
				length(unit: FOOT)
			}
		}`, &resp)

		require.Equal(t, "Imperial shuttle", resp.Starship.Name)
		require.Equal(t, 65.6168, resp.Starship.Length)
	})
}

func TestQueryHuman(t *testing.T) {
	t.Run("luke", func(t *testing.T) {
		var resp struct {
			Human struct {
				Name string
			}
		}
		c.MustPost(`
		{
			human(id:"1000") {
				name
			}
		}`, &resp)

		require.Equal(t, "Luke Skywalker", resp.Human.Name)
	})

	t.Run("vader with starship", func(t *testing.T) {
		var resp struct {
			Human struct {
				Name      string
				Starships []struct {
					Name   string
					Length float64
				}
			}
		}
		c.MustPost(`
		{
			human(id:"1001") {
				name
				starships {
					name
				}
			}
		}`, &resp)

		require.Equal(t, "Darth Vader", resp.Human.Name)
		require.Equal(t, "TIE Advanced x1", resp.Human.Starships[0].Name)
	})

	t.Run("leia with friend", func(t *testing.T) {
		var resp struct {
			Human struct {
				Name    string
				Friends []struct {
					Name string
				}
			}
		}
		c.MustPost(`
		{
			human(id:"1003") {
				name
				friends {
					name
				}
			}
		}`, &resp)

		require.Equal(t, "Leia Organa", resp.Human.Name)
		require.Equal(t, "Luke Skywalker", resp.Human.Friends[0].Name)
	})
}

func TestQueryDriod(t *testing.T) {
	t.Run("R2-D2", func(t *testing.T) {
		var resp struct {
			Droid struct {
				Name            string
				PrimaryFunction string
			}
		}
		c.MustPost(`
		{
			droid(id:"2001") {
				name
				primaryFunction
			}
		}`, &resp)

		require.Equal(t, "R2-D2", resp.Droid.Name)
		require.Equal(t, "Astromech", resp.Droid.PrimaryFunction)
	})

	t.Run("C3PO with friend", func(t *testing.T) {
		var resp struct {
			Droid struct {
				Name    string
				Friends []struct {
					Name string
				}
			}
		}
		c.MustPost(`
		{
			droid(id:"2000") {
				name
				friends {
					name
				}
			}
		}`, &resp)

		require.Equal(t, "C-3PO", resp.Droid.Name)
		require.Equal(t, "Luke Skywalker", resp.Droid.Friends[0].Name)
	})
}

func TestQueryHero(t *testing.T) {
	t.Run("default episode", func(t *testing.T) {
		var resp struct {
			Hero struct {
				Name string
			}
		}
		c.MustPost(`
		{
			hero {
				... on Droid {
					name
				}
			}
		}`, &resp)
		require.Equal(t, "R2-D2", resp.Hero.Name)
	})

	t.Run("Episode=EMPIRE", func(t *testing.T) {
		var resp struct {
			Hero struct {
				Name   string
				Height float64
			}
		}
		c.MustPost(`
		{ 
			hero(episode:EMPIRE) { 
				... on Human { 
					height(unit:METER)
					name
				}
			} 
		}`, &resp)
		require.Equal(t, 1.72, resp.Hero.Height)
		require.Equal(t, "Luke Skywalker", resp.Hero.Name)
	})
}

func TestQueryCharacter(t *testing.T) {
	t.Run("R2D2", func(t *testing.T) {
		var resp struct {
			Character struct {
				Name     string
				Typename string `json:"__typename"`
			}
		}
		c.MustPost(`
		{
			character(id:"2001") {
				name,
				__typename
			}
		}`, &resp)
		require.Equal(t, "R2-D2", resp.Character.Name)
		require.Equal(t, "Droid", resp.Character.Typename)
	})

	t.Run("missins character", func(t *testing.T) {
		var resp struct {
			Character *struct{ Name string }
		}
		c.MustPost(`
		{
			character(id:"2002") {
				name
			}
		}`, &resp)
		require.Nil(t, resp.Character)
	})
}

func TestQuerySearch(t *testing.T) {
	t.Run("Luk", func(t *testing.T) {
		var resp struct {
			Search []struct {
				Name string
			}
		}
		c.MustPost(`
		{
			search(text:"Luk") {
				... on Human {
					name
				}
			}
		}`, &resp)

		require.Equal(t, "Luke Skywalker", resp.Search[0].Name)
		require.Equal(t, 1, len(resp.Search))
	})
	t.Run("r", func(t *testing.T) {
		var resp struct {
			Search []struct {
				Name string
			}
		}
		c.MustPost(`
		{
			search(text:"r") {
				... on Human {
					name
				}
				... on Starship {
					name
				}
			}
		}`, &resp)

		expectedName := []string{"Wilhuff Tarkin", "Leia Organa", "Luke Skywalker", "Darth Vader", "Imperial shuttle"}
		matchCount := 0
		for _, s := range resp.Search {
			for _, e := range expectedName {
				if s.Name == e {
					matchCount++
				}
			}
		}

		require.Equal(t, 5, matchCount)
		require.Equal(t, 5, len(resp.Search))
	})
	t.Run("Lukes with starships", func(t *testing.T) {
		var resp struct {
			Search []struct{ Starships []struct{ Name string } }
		}
		c.MustPost(`{ search(text:"Luke") { ... on Human { starships { name } } } }`, &resp)

		require.Equal(t, "X-Wing", resp.Search[0].Starships[0].Name)
		require.Equal(t, "Imperial shuttle", resp.Search[0].Starships[1].Name)
	})
}

func TestMutationCreateReviewAndQueryReview(t *testing.T) {
	t.Run("createReview", func(t *testing.T) {
		var resp struct {
			createReview struct{ Time string }
		}

		c.MustPost(`
			mutation {
				createReview(episode: NEWHOPE, review:{stars:1, commentary:"i love it", time: "2019-10-12T07:20:50.52Z"})  {
					time
				}
			}`, &resp)

		require.NotEqual(t, resp.createReview.Time, 1)
	})

	t.Run("NEWHOPE's review", func(t *testing.T) {
		var resp struct {
			Reviews []struct {
				Stars      int
				Commentary string
			}
		}
		c.MustPost(`
			{ 
				reviews(episode: NEWHOPE) {
					stars
					commentary
				}
			}`, &resp)

		require.Equal(t, 1, resp.Reviews[0].Stars)
		require.Equal(t, "i love it", resp.Reviews[0].Commentary)
	})

	t.Run("NEWHOPE's review since 2020-01-01", func(t *testing.T) {
		var resp struct {
			Reviews []struct {
				Stars      int
				Commentary string
			}
		}
		c.MustPost(`
			{ 
				reviews(episode: NEWHOPE, since: "2020-01-01T07:20:50.52Z") {
					stars
					commentary
				}
			}`, &resp)

		require.Equal(t, 0, len(resp.Reviews))
	})
}

func TestQueryWithFriendsConnection(t *testing.T) {
	t.Run("driod with friendsConnection", func(t *testing.T) {
		var resp struct {
			Droid struct {
				FriendsConnection struct {
					Friends []struct {
						Name string
					}
				}
			}
		}
		c.MustPost(`
		{
			droid(id:"2001") {
				friendsConnection {
					friends {
						name
					}
				}
			}
		}`, &resp)

		require.Equal(t, "Luke Skywalker", resp.Droid.FriendsConnection.Friends[0].Name)
		require.Equal(t, "Han Solo", resp.Droid.FriendsConnection.Friends[1].Name)
		require.Equal(t, "Leia Organa", resp.Droid.FriendsConnection.Friends[2].Name)
	})

	t.Run("human with friendsConnection edges", func(t *testing.T) {
		var resp struct {
			Human struct {
				FriendsConnection struct {
					Edges []struct {
						Cursor string
						Node   struct {
							Name string
						}
					}
				}
			}
		}
		c.MustPost(`
		{ 
			human(id:"1000") {
				friendsConnection {
					edges {
						cursor
						node {
							name
						}
					}
				}
			}
		}`, &resp)

		require.Equal(t, "Y3Vyc29yMQ==", resp.Human.FriendsConnection.Edges[0].Cursor)
		require.Equal(t, "Han Solo", resp.Human.FriendsConnection.Edges[0].Node.Name)
		require.Equal(t, "Y3Vyc29yMg==", resp.Human.FriendsConnection.Edges[1].Cursor)
		require.Equal(t, "Leia Organa", resp.Human.FriendsConnection.Edges[1].Node.Name)
		require.Equal(t, "Y3Vyc29yMw==", resp.Human.FriendsConnection.Edges[2].Cursor)
		require.Equal(t, "C-3PO", resp.Human.FriendsConnection.Edges[2].Node.Name)
	})
}
