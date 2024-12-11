package dto

type PetInputDTO struct {
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
	OrgWhatspp              string   `json:"organization_whatsapp"`
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

type OrganizationInputDTO struct {
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Email    string `json:"email"`
	CEP      string `json:"cep"`
	City     string `json:"city"`
	Whatsapp string `json:"whatsapp"`
	Password string `json:"password"`
}

type OrganizationOutputDTO struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Owner    string         `json:"owner"`
	Email    string         `json:"email"`
	CEP      string         `json:"cep"`
	City     string         `json:"city"`
	Whatsapp string         `json:"whatsapp"`
	Pets     []PetOutputDTO `json:"pets"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
