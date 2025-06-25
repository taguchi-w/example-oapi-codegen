package main

import (
	"context"
	"encoding/json"
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
	// post
	resp, err := client.PostTodos(ctx, api.PostTodosJSONRequestBody{
		Subject: "subject a from request",
		Body:    "body from request",
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

	var todos []api.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		panic(err)
	}
	todoId := todos[0].Id

	// update
	updateReqBody := api.UpdateTodoPartialJSONRequestBody{
		Subject: util.P("subject a update from request"),
		Body:    util.P("body update from request"),
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
