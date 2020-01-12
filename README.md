# Ayo Les Core API

AyoLesCore is backend core api for simple online course with grahpql
* run via docker images 

```

docker run -d -p 8000:8000 renosyah/ayolescore /bin/main --config=.staging.toml

```

* how to run via local fo golang dev

```

del main.exe && go build main.go && main.exe --config=.server.toml

```



* note

dont forget if you on windows enable allow access to docker in firewall rule: C:\Program Files\Docker\Docker\resources\com.docker.backend.exe

## Requirement

* Golang [https://golang.org](https://golang.org)
* CockroachDB [https://www.cockroachlabs.com](https://www.cockroachlabs.com)
* Git Cmd [https://git-scm.com/downloads](https://git-scm.com/downloads)


## For Database

* import file sql schema : 

```
/sql/ayoles_db.sql

```

* import file sql seed : 

```
/sql/seed.sql

```



## Client side with android kotlin

[github.com/renosyah/AyoLesApp](https://github.com/renosyah/AyoLesApp)