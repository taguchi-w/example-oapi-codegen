package main

import (
	"context"
	"fmt"
	"io"

	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
	"github.com/taguchi-w/example-oapi-codegen/pkg/util"
)

func main() {
	ctx := context.Background()
	client, err := api.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	todoId := "1" // TODOのID

	// post
	resp, err := client.PostTodos(ctx, api.PostTodosJSONRequestBody{
		Id:      todoId,
		Subject: "subject a",
		Body:    "body",
	})
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Todo:", resp.Status)
	fmt.Println(string(body))

	// get all
	resp, err = client.GetTodos(ctx)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Get Todos:", resp.Status)
	fmt.Println(string(body))

	// update
	updateReqBody := api.UpdateTodoPartialJSONRequestBody{
		Subject: util.P("subject a (update)"),
		Body:    util.P("body(update)"),
	}
	resp, err = client.UpdateTodoPartial(ctx, todoId, updateReqBody)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Todo:", resp.Status)
	fmt.Println(string(body))

	// delete
	resp, err = client.DeleteTodo(ctx, todoId)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete Todo:", resp.Status)
	fmt.Println(string(body))
}
