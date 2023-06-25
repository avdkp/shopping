Assumptions
- creating users/admin is totally public, anyone can create an admin user and access admin APIs.
- There is only 1 cart for 1 user.
- /add-to-cart and /remove-from-cart will return partial success if adding/remove of at least 1 item fails.
- There is no mechanism to eventually unlock resources in case of failures/timeouts etc.
- There is no logger, so in case of failures, errors have just been printed on screen.
- Assumption has been made, that admin can not access APIs intended for regular userRole.


Setup
- `make build`

Run
- `make run`

Run Test
- `make test`

APIs
1. Create User
   ```
   curl --location 'localhost:8080/public/users' \
    --header 'Content-Type: application/json' \
    --data '{
    "username": "avd",
    "password": "12345",
    "role": "admin"
    }'
   ```

2. Login
   ```
   curl --location 'localhost:8080/public/login' \
    --header 'Content-Type: application/json' \
    --data '{
    "username": "avd",
    "password": "12345"
    }'
   ```

3. Add Items to inventory
   ```
    curl --location 'localhost:8080/admin/add-items' \
    --header 'Content-Type: application/json' \
    --header 'Auth-Token: ycILu8lshx' \
    --data '[{
    "name": "apple",
    "description": "a fruit"
    },{
    "name": "bat",
    "description": "a mammal"
    },
    {
    "name": "elephant",
    "description": "an animal"
    }]'
    ```
4. List all Items in inventory
   ```
    curl --location 'localhost:8080/all-items' \
    --header 'Content-Type: application/json' \
    --header 'Auth-Token: PvsuP1Y8ZR' \
    --data ''
    ```

5. Adding itemIds to the cart
   ```
    curl --location --request PATCH 'localhost:8080/add-to-cart' \
    --header 'Content-Type: application/json' \
    --header 'Auth-Token: PvsuP1Y8ZR' \
    --data '[1,2]'
    ```
   
6. Remove Items from the Cart
   ```
    curl --location --request PATCH 'localhost:8080/remove-from-cart' \
    --header 'Content-Type: application/json' \
    --header 'Auth-Token: PvsuP1Y8ZR' \
    --data '[10]'
    ```