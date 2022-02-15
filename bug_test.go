package bug

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, entClient *ent.Client) {
	schema := NewSchema(entClient)
	srv := handler.NewDefaultServer(schema)
	graphQLClient := client.New(srv)

	for i := 0; i < 100; i++ {
		entClient.User.Create().SetName(fmt.Sprintf("Test%d", i)).SetAge(i + 10).SaveX(context.Background())
	}

	type users struct {
		Users struct {
			Edges []struct {
				Node struct {
					ID string
				}
			}
		}
	}

	t.Run("basic", func(t *testing.T) {
		var resp users
		graphQLClient.MustPost(`
query {
	users(first: 10) {
		edges {
			node {
				id
			}
		}
	}
}
`, &resp)
		gid := resp.Users.Edges[0].Node.ID
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
	users(first: 1, after: "MTMKMQ==" ) {
		edges {
			node {
				id
			}
		}
	}
}
`, &resp)
		gid := resp.Users.Edges[0].Node.ID
		id, err := base64.URLEncoding.DecodeString(gid)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("gid=%v id=%q\n", gid, string(id))
	})
}
