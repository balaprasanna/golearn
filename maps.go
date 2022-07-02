package main

import "fmt"

type Dict map[string]interface{}

func PrintMap(kv Dict, level int) {
	for key, value := range kv {
		if _, ok := value.(Dict); ok {
			fmt.Println(key, "-> level ->", level)
			PrintMap(value.(Dict), level+1)
		} else {
			fmt.Println(level, key, value)
		}
	}
}

func main() {
	students := make(Dict)
	students["bala"] = 29
	students["prasanna"] = 30
	students["contact"] = Dict{
		"mobile":   "+65-92339183",
		"landline": "06121234",
	}
	//PrintMap(students, 1)
	delete(students, "bala")
	PrintMap(students, 1)
	fmt.Println(len(students))

	key := "prasanna"
	if val, ok := students[key]; ok {
		fmt.Println(ok, val, key, "exist in map")
	} else {
		fmt.Println(ok, val, key, "doesn't exist in map")
	}
}
