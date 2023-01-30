// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// type key string

// func enrichContentx(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, key("request-id"), "1239")
// }

// func doSemething(ctx context.Context) {
// 	rID := ctx.Value("request-id")

// 	fmt.Println(rID)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Timed out")
// 			return
// 		default:
// 			fmt.Println("Doing Something")
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 	}

// }

// func main() {
// 	println("Go Context Example")
// 	// ctx := context.Background()
// 	// ctx = enrichContentx(ctx)

// 	// doSemething(ctx)

// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()
// 	ctx = enrichContentx(ctx)

// 	go doSemething(ctx)

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Oh no, we timed out")
// 		fmt.Println(ctx.Err())
// 	}

// 	time.Sleep(2 * time.Second)
// }

// // When working with context it is a best practice to place the first argument as the context itself.
/*
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func enrichContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
		err := ctx.Err()
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Println("Context not cancelled")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", enrichContext(http.DefaultServeMux))
}
*/
package main

import (
	"context"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type key string

func someFunc(ctx context.Context) context.Context {
	return context.WithValue(ctx, key("user"), User{
		Name: "John Doe",
		Age:  20,
	})
}

func main() {
	ctx := context.Background()
	ctx = someFunc(ctx)

	user := ctx.Value(key("user")).(User)
	fmt.Println(user)
}
