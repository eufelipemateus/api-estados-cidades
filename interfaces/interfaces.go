package interfaces

type Estado struct {
	Nome  string `json:"nome"`
	Sigla string `json:"sigla"`
}

type Cidade struct {
	Nome  string `json:"nome"`
	Estado string `json:"estado"`
}
