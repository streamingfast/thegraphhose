package thegraph

type Manifest struct {
	SpecVersion string `yaml:"specVersion"`
	Description string `yaml:"description"`
	Repository  string `yaml:"repository"`
	Schema      struct {
		File string `yaml:"file"`
	} `yaml:"schema"`
	DataSources []EthereumContractDataSourceEntity         `yaml:"dataSources"`
	Templates   []EthereumContractDataSourceTemplateEntity `yaml:"templates"`
}

type EthereumContractDataSourceEntity struct {
	Kind    string                        `yaml:"kind"`
	Network string                        `yaml:"network"`
	Name    string                        `yaml:"name"`
	Source  EthereumContractSourceEntity  `yaml:"source"`
	Mapping EthereumContractMappingEntity `yaml:"mapping"`
}

type EthereumContractDataSourceTemplateEntity struct {
	Kind    string                                         `yaml:"kind"`
	Network string                                         `yaml:"network"`
	Name    string                                         `yaml:"name"`
	Source  EthereumContractDataSourceTemplateSourceEntity `yaml:"source"`
	Mapping EthereumContractMappingEntity                  `yaml:"mapping"`
}

type EthereumContractSourceEntity struct {
	Address    string `yaml:"address"`
	Abi        string `yaml:"abi"`
	StartBlock int    `yaml:"startBlock"`
}

type EthereumContractDataSourceTemplateSourceEntity struct {
	Abi string `yaml:"abi"`
}

type EthereumContractMappingEntity struct {
	Kind       string   `yaml:"kind"`
	APIVersion string   `yaml:"apiVersion"`
	Language   string   `yaml:"language"`
	File       string   `yaml:"file"`
	Entities   []string `yaml:"entities"`
	Abis       []struct {
		Name string `yaml:"name"`
		File string `yaml:"file"`
	} `yaml:"abis"`
	EventHandlers []struct {
		Event   string `yaml:"event"`
		Handler string `yaml:"handler"`
	} `yaml:"eventHandlers"`
}
