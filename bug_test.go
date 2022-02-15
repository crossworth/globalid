package bug

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent/enttest"
)

func TestGlobalID(t *testing.T) {
	entClient := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer entClient.Close()

	schema := NewSchema(entClient)
	srv := handler.NewDefaultServer(schema)
	graphQLClient := client.New(srv)

	for i := 0; i < 100; i++ {
		entClient.User.Create().SetName(fmt.Sprintf("Test%d", i)).SetAge(i + 10).SaveX(context.Background())
	}

	type users struct {
		Users struct {
			Nodes []struct {
				ID string
			}
		}
	}

	t.Run("basic", func(t *testing.T) {
		var resp users
		graphQLClient.MustPost(`
query {
	users(first: 1) {
		nodes {
			id
		}
	}
}
`, &resp)
		gid := resp.Users.Nodes[0].ID
		id, err := base64.URLEncoding.DecodeString(gid)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("gid=%v id=%q\n", gid, string(id))
	})

	t.Run("where", func(t *testing.T) {
		var resp users
		graphQLClient.MustPost(`
query {
	users(first: 1, after: "MTMKMDI0ODI0NTEtMDk2MS00ODIyLTlkZDgtMjA1ZWJhYjY5ZDhj" ) {
		nodes {
			id
		}
	}
}
`, &resp)
		gid := resp.Users.Nodes[0].ID
		id, err := base64.URLEncoding.DecodeString(gid)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("gid=%v id=%q\n", gid, string(id))
	})
}
