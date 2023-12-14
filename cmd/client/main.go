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
	petId := "1" // ペットのID

	// post
	resp, err := client.PostPets(ctx, api.PostPetsJSONRequestBody{
		Id:   petId,
		Name: "cat",
	})
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Pet:", resp.Status)
	fmt.Println(string(body))

	// get all
	resp, err = client.GetPets(ctx)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Get Pets:", resp.Status)
	fmt.Println(string(body))

	// update
	updateReqBody := api.UpdatePetPartialJSONRequestBody{
		Name: util.P("new cat"),
	}
	resp, err = client.UpdatePetPartial(ctx, petId, updateReqBody)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Pet:", resp.Status)
	fmt.Println(string(body))

	// delete
	resp, err = client.DeletePet(ctx, petId)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete Pet:", resp.Status)
	fmt.Println(string(body))
}
