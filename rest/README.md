# REST API sample

go run main.go

curl http://localhost:3333
curl http://localhost:3333/articles

curl -H 'user_role:member'  http://localhost:3333/articles
curl -H 'user_role:owner'  http://localhost:3333/admin

curl -H 'user_role:admin'  http://localhost:3333/admin
