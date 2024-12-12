# steppingDisk

Service to accept a database query, and distribute it to multiple to database connections. I'm basically building this as a resume project. To show my skills with Go, Networking, etc.

### Goals
1. Accept one or more connections to distribute queries.
2. Support multiple databases, like MySQL, Postgres, MongoDB...
3. Encrypt queries as they pass through application.
4. Support queues via memory to handle many requests big or small.
5. Have live loading during migration to this service to never miss a query.