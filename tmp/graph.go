package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()
	router.Path("/api").HandlerFunc(handleGraphQL)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	handler := c.Handler(router)
	logrus.Fatal(http.ListenAndServe(":80", handler))
}

func handleGraphQL(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ip := strings.Split(r.RemoteAddr, ":")[0]
	args := make(map[string]interface{})

	args["ip"] = ip

	g := make(map[string]interface{})

	if r.Method == http.MethodGet {
		g["query"] = r.URL.Query().Get("query")
		result := executeQuery(g, schema, args)
		json.NewEncoder(w).Encode(result)
	}

	if r.Method == http.MethodPost {
		data, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(data, &g)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}

		result := executeQuery(g, schema, args)
		json.NewEncoder(w).Encode(result)
	}
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func executeQuery(query map[string]interface{}, schema graphql.Schema, args map[string]interface{}) *graphql.Result {

	params := graphql.Params{
		Schema:        schema,
		RequestString: query["query"].(string),
	}

	if query["variables"] != nil {
		params.VariableValues = query["variables"].(map[string]interface{})
	}

	if len(args) > 0 {
		_context := context.Background()
		for key, value := range args {
			_context = context.WithValue(_context, key, value)
		}
		params.Context = _context
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		fmt.Println(result.Errors)
	}
	return result
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"query": query,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"update": update,
	},
})

var query = &graphql.Field{
	Type:        graphql.String,
	Description: "A Query Action",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)

		return name, nil
	},
}

var update = &graphql.Field{
	Type:        graphql.String,
	Description: "A Update Action",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)

		return name, nil
	},
}

