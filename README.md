# gqlgen-validator-sample

[gqlgen](https://github.com/99designs/gqlgen)と[validator](https://github.com/go-playground/validator)を使っていい感じに手を抜くサンプル

# Generate GraphQL code

```
go run ./cmd/gqlgenerate
```

# Run

```
go run ./cmd/server
```

Then open playground

```
open http://localhost:8080/
```

# Environment

| Env  | Example | Required | Default |
|------|---------|----------|---------|
| PORT | 8080    | true     | 8080    |
