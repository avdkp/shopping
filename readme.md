Assumptions
- creating users/admin is totally public, anyone can create an admin user and access admin APIs
- There is only 1 cart for 1 user
- /add-to-cart and /remove-from-cart will return partial success if adding/remove of at least 1 item fails.
- There is no mechanism to eventually unlock resources in case of failures/timeouts etc
- There is no logger, so in case of failures, errors have just been printed on screen.
- Assumption has been made, that admin can not access APIs intended for regular userRole.


Setup
- `make build`

Run
- `make run`

Run Test
- `make test`