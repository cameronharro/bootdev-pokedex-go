package pokeapi

type BulkLocationArea struct {
	Count    int
	Next     *string
	Previous *string
	Results  []NamedAPIResource
}

type NamedAPIResource struct {
	Name string
	Url  string
}

type LocationArea struct {
	Id                   int
	Name                 string
	GameIndex            int
	EncounterMethodRates []EncounterMethodRate
	Location             NamedAPIResource
	Names                []Name
	PokemonEncounters    []PokemonEncounter
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource
	VersionDetails  []EncounterVersionDetails
}

type EncounterVersionDetails struct {
	Rate    int
	Version NamedAPIResource
}

type Name struct {
	Name     string
	Language NamedAPIResource
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource
	VersionDetails []VersionEncounterDetail
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource
	MaxChance        int
	EncounterDetails []Encounter
}

type Encounter struct {
	MinLevel        int
	MaxLevel        int
	ConditionValues []NamedAPIResource
	Chance          int
	Method          NamedAPIResource
}
