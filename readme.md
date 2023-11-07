## Start Server Need two step
#### 1. Start Docker Environment
```
docker compose -f docker-compose.db.yml up
```

#### 2. Select Environment File To Start Golang
> go run main.go ./env/dev/.env.dev

### Command To Migrate
```

go run ./pkg/database/script/migration.go ./env/dev/.env.todo

```


- The TODO application can show the `LIST` of tasks with the following requirements
    - Can sort the data by `Title` or `Date` or `Status` fields
    - Can search the data by `Title` or `Description` fields
- The TODO application can `UPDATE` a task with the following requirements
    - Can update a task by `ID` field
    - Can update `Title`, `Description`, `Date`, `Image`, and `Status` fields corresponding to the requirements from the `CREATE` feature

# Api -> Get All Todo (Maybe Pagination) : No Params Required
# Api -> Get Single Todo (For Update) : Need Todo Id Params

# Api -> Sorting Todo : Need Text Query
# Api -> Searching Todo : Need Text Query

# Api -> Create Todo : Need Req Body
# Api -> Update Todo : Need Req Body

# APi -> Delete Todo : Need Todo Id Params