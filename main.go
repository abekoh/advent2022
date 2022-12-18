package main

import "fmt"

func genericFilter[T any](sl []T, test func(T) bool) (res []T) {
	for _, e := range sl {
		if test(e) {
			res = append(res, e)
		}
	}
	return res
}

type User struct {
	Name string
	Age  int
}

type UserList []User

func (ul UserList) Filter(test func(User) bool) UserList {
	return genericFilter(ul, test)
}

func main() {
	users := UserList{
		User{Name: "Alice", Age: 22},
		User{Name: "Bob", Age: 10},
		User{Name: "Carol", Age: 38},
		User{Name: "Dan", Age: 18},
	}
	filtered := users.Filter(func(u User) bool {
		return u.Age >= 20
	})
	fmt.Printf("filtered: %+v\n", filtered)
}
