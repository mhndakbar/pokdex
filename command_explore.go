package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide an area name")
	}

	name := args[0]
	exploreResp, err := cfg.pokeapiClient.ExploreArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring area: %v...", exploreResp.Name)
	fmt.Println("\nFound Pokemons:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
