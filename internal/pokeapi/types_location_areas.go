package pokeapi

type BulkLocationArea struct {
	Count    int
	Next     *string
	Previous *string
	Results  []namedAPIResource
}

type namedAPIResource struct {
	Name string
	Url  string
}

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []encounterMethodRate `json:"encounter_method_rates"`
	Location             namedAPIResource      `json:"location"`
	Names                []name                `json:"names"`
	PokemonEncounters    []pokemonEncounter    `json:"pokemon_encounters"`
}

type encounterMethodRate struct {
	EncounterMethod namedAPIResource          `json:"encounter_method"`
	VersionDetails  []encounterVersionDetails `json:"version_details"`
}

type encounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version namedAPIResource `json:"version"`
}

type name struct {
	Name     string           `json:"name"`
	Language namedAPIResource `json:"language"`
}

type pokemonEncounter struct {
	Pokemon        namedAPIResource         `json:"pokemon"`
	VersionDetails []versionEncounterDetail `json:"version_details"`
}

type versionEncounterDetail struct {
	Version          namedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []encounter      `json:"encounter_details"`
}

type encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []namedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          namedAPIResource   `json:"method"`
}
