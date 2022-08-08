package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Declare Todos Struct here ...

// Declare Todos Global Variable ...

func main() {
	r := mux.NewRouter()

	// Create routes here ...

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

// Create FindTodos Function here ...

// Create GetTodo Function here ...

// Create CreateTodo Function here ...

// Create UpdateTodo Function here ...

// Create DeleteTodo Function here ...