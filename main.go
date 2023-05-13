package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory("./pkg/app")

	if err != nil {
		fmt.Printf("Error obteniendo la refenrencia del directorio host: %s", err)
		os.Exit(1)
	}

	cn, err := client.Container().
		Build(src).
		Publish(ctx, "allfait/automation-tilt-dagger:latest")

	if err != nil {
		fmt.Printf("Error creando y empujando el contenedor: %s", err)
		os.Exit(1)
	}

	fmt.Print("Contenedor creado y pusheado: %s", cn)

	if err != nil {
		return err
	}

	return nil
}
