# Dashboard API
    
    APIs to provide data to the dashboard of data analytical tool

### To setup the DB
    $ cd migration
    $ goose postgres <DB_URL> up
    $ cd ..

### To run the project

    Copy the config/env.example to config/development.env
    Add the necessary environment variables to config/development.env
    $ source config/development.env
    $ go run main.go

## API Specification

    1.  To get access tken
        POST /login
        Body
        {
            "username":"<username>"
            "password":"<password>"
        }

    2. Product distribution chart
        GET /getProductDistribution?product=<productName>
        Headers
        Authorization : Bearer <token>
    
    3. Distributor performance graph
        GET /getDistributorPerformance?distributor=<distributorName>
        Headers
        Authorization : Bearer <token>

    4.  Area Distribution Map
        GET /getAreaDistribution?place=<placeName>
        Headers
        Authorization : Bearer <token>

    5. Top N selling items in each place
        GET /getTopN?number=<N>&distributor=<distributorName>
        Headers
        Authorization : Bearer <token>

    Default user credentials
        username: user
        password: user