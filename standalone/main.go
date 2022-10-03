package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

func main() {

	modelParam, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("NewModelFromString: %s", err)
	}

	enforcer, err := casbin.NewEnforcer(modelParam, fileadapter.NewAdapter("policy.csv"))
	if err != nil {
		log.Fatalf("NewEnforcer: %s", err)
	}

	in := [][]any{
		{"Aさん", "file1", "read"},
		{"Aさん", "file1", "write"},
		{"Aさん", "file4", "read"},
		{"Bさん", "file1", "read"},
	}

	for _, v := range in {
		ok, err := enforcer.Enforce(v...)
		if err != nil {
			log.Fatalf("enforce: %s", err)
		}
		fmt.Printf("%v: %v\n", v, ok)
	}

}
