package dto

type PetInputDTO struct {
	ID                      string   `json:"id"`
	OrgID                   string   `json:"org_id"`
	Name                    string   `json:"name"`
	About                   string   `json:"about"`
	Age                     string   `json:"age"`
	Size                    string   `json:"size"`
	EnergyLevel             string   `json:"energy_level"`
	DependencyLevel         string   `json:"dependency_level"`
	Environment             string   `json:"enviroment"`
	Photos                  []string `json:"photos"`
	RequirementsForAdoption []string `json:"requiriments_for_adoption"`
}

type PetOutputDTO struct {
	ID                      string   `json:"id"`
	Org                     string   `json:"organization"`
	Name                    string   `json:"name"`
	About                   string   `json:"about"`
	Age                     string   `json:"age"`
	Size                    string   `json:"size"`
	EnergyLevel             string   `json:"energy_level"`
	DependencyLevel         string   `json:"dependency_level"`
	Environment             string   `json:"enviroment"`
	Photos                  []string `json:"photos"`
	RequirementsForAdoption []string `json:"requiriments_for_adoption"`
}
