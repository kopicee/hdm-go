```sh
# Run tests
make test

# Run app
make up
```

```sh
# Hit API
curl 'http://localhost:3000/api/hotels'

# Filter by one or more hotel IDs
curl 'http://localhost:3000/api/hotels?id=SjyX'
curl 'http://localhost:3000/api/hotels?id=SjyX&id=f8c9' 

# Filter by one or more destinations
curl 'http://localhost:3000/api/hotels?destination=1122'
curl 'http://localhost:3000/api/hotels?destination=1122&destination=5432'

# Filter by hotel ID and destination
curl 'http://localhost:3000/api/hotels?id=SjyX&destination=1122'
```
