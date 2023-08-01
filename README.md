<div align="center">

![cloudmate logo](https://avatars.githubusercontent.com/u/69299682?s=200&v=4)

# GORM gqlgen relay

<small style="opacity: 0.7;">by Cloudmate</small>

---

![Cloudmate](https://img.shields.io/badge/Cloudmate-FFFFFF?style=for-the-badge&logoColor=black&logo=data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBzdGFuZGFsb25lPSJubyI/Pgo8IURPQ1RZUEUgc3ZnIFBVQkxJQyAiLS8vVzNDLy9EVEQgU1ZHIDIwMDEwOTA0Ly9FTiIKICJodHRwOi8vd3d3LnczLm9yZy9UUi8yMDAxL1JFQy1TVkctMjAwMTA5MDQvRFREL3N2ZzEwLmR0ZCI+CjxzdmcgdmVyc2lvbj0iMS4wIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiB3aWR0aD0iMjAwLjAwMDAwMHB0IiBoZWlnaHQ9IjIwMC4wMDAwMDBwdCIgdmlld0JveD0iMCAwIDIwMC4wMDAwMDAgMjAwLjAwMDAwMCIKIHByZXNlcnZlQXNwZWN0UmF0aW89InhNaWRZTWlkIG1lZXQiPgoKPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMC4wMDAwMDAsMjAwLjAwMDAwMCkgc2NhbGUoMC4xMDAwMDAsLTAuMTAwMDAwKSIKZmlsbD0iIzAwMDAwMCIgc3Ryb2tlPSJub25lIj4KPHBhdGggZD0iTTg3NyAxODkwIGMtMTU0IC0yOCAtMzE5IC0xMDQgLTQzNCAtMjAwIC0xNDMgLTEyMCAtMjY4IC0zMzAgLTI5MAotNDg2IC03IC01MSAtMjEgLTM3IDE5NyAtMTg4IDkxIC02MyAyMDUgLTE0MyAyNTUgLTE3OCAyMTIgLTE0OSAyODcgLTE5OCAzMDgKLTE5OCAxNCAwIDcwIDQ4IDE2NiAxNDIgMTgwIDE3OCAyMDMgMjA1IDE5NSAyMzIgLTMgMTIgLTEzIDI3IC0yMiAzNCAtMzggMjkKLTYyIDE1IC0xOTcgLTExOCAtOTYgLTk1IC0xMzkgLTEzMCAtMTU3IC0xMzAgLTI1IDAgLTU4MiAzODEgLTYwMiA0MTIgLTcgMTIKLTUgMzEgNyA3MCAyOCA5MCA4NyAxNzggMTgyIDI3MyA3MyA3NCAxMDYgOTkgMTc3IDEzNCAxNzAgODQgMzM0IDEwNyA1MDIgNzIKMjAzIC00MiAzNzMgLTE1NCA0ODYgLTMyMSA0MCAtNjAgOTQgLTE3NiAxMDUgLTIyNyA3IC0zNCAzNyAtNjMgNjUgLTYzIDkgMAoyNyAxMSAzOSAyNCAyOCAzMCAyNSA2MCAtMTkgMTY2IC04NSAyMTEgLTI1NSAzOTAgLTQ1NSA0ODEgLTExMyA1MSAtMTk5IDcwCi0zMzUgNzQgLTY5IDIgLTE0NyAwIC0xNzMgLTV6Ii8+CjxwYXRoIGQ9Ik02NjcgMTQxOCBjLTE0NSAtMTQwIC0xNTYgLTE1OSAtMTE3IC0xOTggNDAgLTQwIDYyIC0yOCAyMDEgMTA3IDcxCjY5IDEzMiAxMzcgMTM1IDE1MCA5IDM0IC0yMCA3MyAtNTQgNzMgLTIyIDAgLTU2IC0yNyAtMTY1IC0xMzJ6Ii8+CjxwYXRoIGQ9Ik04OTIgMTQwOCBjLTE5OSAtMTg5IC0yMzUgLTIyNSAtMjQzIC0yNDYgLTE3IC00MyAyOSAtOTAgNzQgLTc2IDEyCjQgODggNzIgMTcwIDE1MiAxNTUgMTU0IDE2NiAxNzMgMTI3IDIxMiAtMzIgMzIgLTU3IDI0IC0xMjggLTQyeiIvPgo8cGF0aCBkPSJNOTMyIDEyMjIgYy0xMzggLTEzNCAtMTcyIC0xNzIgLTE3MiAtMTk0IDAgLTM1IDI0IC01OCA2MiAtNTggMjUgMAo1MSAyMSAxNjIgMTMwIDc4IDc2IDE0MCAxMzAgMTUyIDEzMCAxMSAwIDgyIC00NCAxNTkgLTk4IDc3IC01NCAyMDYgLTE0NCAyODYKLTE5OSA4MCAtNTYgMTUxIC0xMDkgMTU5IC0xMTcgMTEgLTE0IDkgLTI3IC0xNCAtODYgLTc2IC0xOTMgLTI0OCAtMzYzIC00NDUKLTQzOSAtMjkxIC0xMTIgLTY0NCAtMTQgLTg1NCAyMzYgLTQ5IDU5IC0xMzAgMjEzIC0xNDIgMjY5IC0xMiA2MCAtODEgODkKLTExMyA0OCAtNDYgLTYzIDc1IC0zMTQgMjIyIC00NTkgMzAxIC0yOTcgNzY0IC0zNDkgMTA5OCAtMTI0IDgzIDU2IDE4OSAxNTkKMjQ0IDIzNiA3OCAxMDcgMTY0IDMxOCAxNDkgMzYzIC03IDIwIC03NDAgNTMwIC03NjIgNTMwIC0xMCAwIC05NiAtNzYgLTE5MQotMTY4eiIvPgo8cGF0aCBkPSJNMTMwOCA5NDIgYy05IC00IC04MiAtNzEgLTE2MiAtMTUwIC0xMzIgLTEzMCAtMTQ2IC0xNDcgLTE0NiAtMTc4IDAKLTQyIDM1IC02NCA3NSAtNDggMzcgMTQgMzE1IDI5MyAzMTUgMzE2IDAgMTkgLTQyIDY5IC01NyA2NyAtNCAwIC0xNiAtMyAtMjUKLTd6Ii8+CjxwYXRoIGQ9Ik0xMjgzIDY5OCBjLTczIC03MyAtMTM2IC0xNDAgLTEzOSAtMTUwIC0xMCAtMzIgMjIgLTY4IDU5IC02OCAyOCAwCjUwIDE4IDE2NSAxMzEgMTA3IDEwNyAxMzIgMTM3IDEzMiAxNjAgMCAzMCAtMzEgNTkgLTY0IDU5IC0xMSAwIC03MyAtNTQgLTE1MwotMTMyeiIvPgo8L2c+Cjwvc3ZnPgo=)
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

</div>

---

## Install

```sh
go get -u github.com/cloudmatelabs/gorm-gqlgen-relay
```

## About

Using [Relay] for **GORM + gqlgen** projects

## Prepare

1. Must set `autobind` value in gqlgen.yml
2. Must set `schema` value in gqlgen.yml

## Steps

1. Write Edge and Connection schema  
    see: [example/graph/schema/todo.graphql](./example/graph/schema/todo.graphql)
2. Write Filter schema  
   see: [example/graph/schema/todo.graphql](./example/graph/schema/todo.graphql)
3. Import connection and edge  
   see: [example/graph/model/todo.go](./example/graph/model/todo.go)
4. generate code  
   see: [example/generate.go](./example/generate.go)
5. Using paginate resolver  
   see: under usage or [example/graph/todo.resolvers.go](./example/graph/todo.resolvers.go)

## Usage

### resolver

see: [example/graph/todo.resolvers.go](./example/graph/todo.resolvers.go)

```go
import (
  "github.com/cloudmatelabs/gorm-gqlgen-relay/relay"
  customContext "github.com/juunini/gorm-custom-context"
)

func (r *queryResolver) Todos(ctx context.Context, first *int, after *string, last *int, before *string, orderBy map[string]interface{}, where *model.TodoFilter) (*relay.Connection[model.Todo], error) {
	context := customContext.GetContext(ctx)
	db := context.Database.Preload("User")

	return relay.Paginate[model.Todo](db, where, orderBy, relay.PaginateOption{
		First:      first,
		After:      after,
		Last:       last,
		Before:     before,
		Table:      "todos",
		PrimaryKey: "id",
	})
}
```

### model

```go
import "github.com/cloudmatelabs/gorm-gqlgen-relay/relay"

type TodoEdge = relay.Edge[Todo]
type TodoConnection = relay.Connection[Todo]
```

### schema(graphql)

```graphql
type Query {
  todos(
    first: Int
    after: String
    last: Int
    before: String
    orderBy: Map
    where: TodoFilter
  ): TodoConnection!
}

type TodoEdge {
  node: Todo!
  cursor: String!
}

type TodoConnection {
  totalCount: Int!
  edges: [TodoEdge!]!
  pageInfo: PageInfo!
}

input TodoFilter {
  id: IDFilter
  text: StringFilter
  done: BooleanFilter
}
```

### generate.go

```go
package main

//go:generate go run -mod=mod github.com/cloudmatelabs/gorm-gqlgen-relay
//go:generate go run -mod=mod github.com/99designs/gqlgen
```

## Paginate Option

| Name | Type | Description |
| --- | --- | --- |
| First | int | The number of items to return |
| Last | int | The number of reversed items to return |
| After | string | A cursor for use in pagination, which is a base-64 encoded string that points to a specific page in the dataset. |
| Before | string | A cursor for use in pagination, which is a base-64 encoded string that points to a specific page in the dataset. |
| Table | string | Table name (optional) |
| PrimaryKey | string | Primary key name (optional) |

## Execute example

```sh
git init gorm-gqlgen-relay
cd gorm-gqlgen-relay

git config core.sparseCheckout true
git remote add -f origin https://github.com/cloudmatelabs/gorm-gqlgen-relay.git

echo "example" >> .git/info/sparse-checkout
git pull origin main

cd example
go run server.go
```

[Relay]: https://relay.dev/
