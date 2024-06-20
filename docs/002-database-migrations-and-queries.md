# Database migrations and queries

Table of contents:
1. [Creating a new migration](#creating-a-new-migration)
2. [Making SQL queries and generating `Go` code](#making-sql-queries-and-generating-`go`-code)

## Creating a new migration

To create a new database migration, add a new SQL file to `sql/migrations`. Your file name should be of the format `<identifier>_<filename>.sql`, where the identifier should be the latest migration's identified incremented by 1. For example, if the latest migration is `008_some_migration.sql`, then your new migration should be named `009_<filename>.sql`

Your migration file should have the following layout
```sql
-- +goose Up

-- Write your UP migration here...

-- +goose Down

-- Write your DOWN migration here...

```
<details>
<summary>Example migration</summary>

```sql
-- +goose Up

CREATE TABLE my_table (
    id UUID PRIMARY KEY,
    column_1 TEXT NOT NULL,
    column_2 TEXT NOT NULL
);

-- +goose Down

DROP TABLE my_table;
```
</details>

To apply your migration, first make sure the local database is running (else use `make db-up`), then run the following to apply the UP migration
```
make db-migrate-up
```

To apply the DOWN migration, run
```
make db-migrate-down
```

## Making SQL queries and generating `Go` code

Create a new SQL file to `sql/queries` or simply add to the existing SQL files. Your queries should be of the following format

```sql
-- name: <YourQueryName> :<Output>
-- Write your SQL query here...
```

where `<Output>` accepts the following:

- `one`: return one record
- `many`: return multiple records
- `exec`: return none

<details>
<summary>Example queries</summary>
  
  ```sql
  -- name: GetAllRecords :many
  SELECT * FROM my_table

  -- name: CreateRecord :one
  INSERT INTO my_table (id, column_1, column_2)
  VALUES ($1, $2, $3)
  RETURNING *;

  -- name: DeleteRecord :exec
  DELETE FROM my_table WHERE id = $1;
  ```

</details>

For more information, please consult the [sqlc documentation](https://docs.sqlc.dev/en/stable/index.html)
