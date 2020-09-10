# graphql-go-examples

- https://github.com/graphql-go/graphql/
- use sample data from [testutil.StarWarsSchema](https://github.com/graphql-go/graphql/blob/master/testutil/testutil.go)

## Run

```
go mod tidy
go run main.go
```

Open http://localhost:3000

```graphql
query {
  human (id: "1000") {
    name
    friends {
      name
    }
  }
}
```

## Using cURL
### Query

```bash
curl -XPOST \
     -H "Content-Type:application/graphql" \
     -d 'query { human (id: "1000") { name, friends { name }}}' \
     http://localhost:3000/graphql
```

### Introspection
```bash
curl -XPOST \
     -H 'Content-Type:application/graphql' \
     -d '{__schema { queryType { name, fields { name, description } }}}' \
     http://localhost:3000/graphql
```

## TODOs

- [ ] study [graphql/examples](https://github.com/graphql-go/graphql/tree/master/examples)
  - [ ] study [crud](https://github.com/graphql-go/graphql/tree/master/examples/crud)
- [ ] [Go Microservices blog series, part 14 - GraphQL](https://callistaenterprise.se/blogg/teknik/2018/05/07/go-blog-series-part14/)
