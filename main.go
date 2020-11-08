package main

import (
	"fmt"
	"github.com/alexsuslov/go-graphiql"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ivanovladyslav/waves-backend/m/config"
)

type User struct {
	Username string
	Password string
	Email	 string
}

func createSchema(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type UserResolver struct {
	username string
	password string
	email string
}

func (ur *UserResolver) Username() string {
	return ur.username
}

func (ur *UserResolver) Password() string {
	return ur.password
}

func (ur *UserResolver) Email() string {
	return ur.email
}

type Resolver struct{}

func (q *Resolver) Info() (string, error) {
	return "strisdffdng", nil
}

type SignUpArgs struct {
	Username string
	Password string
	Email string
}

func (q *Resolver) SignUp(args SignUpArgs) (*UserResolver, error) {
	return &UserResolver{
		username: args.Username,
		password: args.Password,
		email: args.Email,
	}, nil
}

func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(schemaString, resolver)
	if err != nil {
		panic(err)
	}

	return parsedSchema
}

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Unable to load env")
	}

	db := pg.Connect(&pg.Options{
		User: cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})
	defer db.Close()

	models := []interface{}{
		(*User)(nil),
	}

	err = createSchema(db, models)
	if err != nil {
		panic(err)
	}

	schema := parseSchema("./schema.graphql", &Resolver{})
	http.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})
	http.HandleFunc("/api", graphiql.ServeGraphiQL)

	fmt.Println("Snappy is on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
