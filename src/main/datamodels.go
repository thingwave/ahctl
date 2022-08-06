package main

type SystemList struct {
	Data  []System
	Count int
}

type System struct {
	Id                 int
	SystemName         string
	Address            string
	Port               int
	AuthenticationInfo string
	CreatedAt          string
	UpdatedAt          string
}

type ServiceDefinitionList struct {
	Data  []ServiceDefinition `json:"data"`
	Count int                 `json:"count"`
}

type ServiceDefinition struct {
	Id                int    `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

type ServiceQueryRequest struct {
	InterfaceRequirements        []string           `json:"interfaceRequirements"`
	MaxVersionRequirement        *int               `json:"maxVersionRequirement,omitempty"`
	MinVersionRequirement        *int               `json:"minVersionRequirement,omitempty"`
	MetadataRequirements         *map[string]string `json:"MetadataRequirements,omitempty"`
	PingProviders                *bool              `json:"pingProviders,omitempty"`
	SecurityRequirements         []string           `json:"securityRequirements,omitempty"`
	ServiceDefinitionRequirement string             `json:"serviceDefinitionRequirement"`
	VersionRequirement           *int               `json:"versionRequirement,omitempty"`
}

type ServiceQueryResponse struct {
	ServiceQueryData []ServiceQueryEntry `json:"serviceQueryData"`
	UnfilteredHits   int                 `json:"unfilteredHits"`
}

type ServiceQueryEntry struct {
	Id                int64                `json:"id"`
	EndOfValidity     *string              `json:"endOfValidity,omitempty"`
	Interfaces        []InterfaceEntry     `json:"interfaces"`
	Metadata          *map[string]string   `json:"metadata,omitempty"`
	Provider          ProviderDTO          `json:"provider"`
	Secure            string               `json:"secure"`
	ServiceDefinition ServiceDefinitionDTO `json:"serviceDefinition"`
	ServiceUri        string               `json:"serviceUri"`
	CreatedAt         string               `json:"createdAt"`
	UpdatedAt         string               `json:"updatedAt"`
	Version           *int                 `json:"version,omitempty"`
}

type ServiceDefinitionDTO struct {
	Id                int64  `json:"id"`
	ServiceDefinition string `json:"serviceDefinition"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

type InterfaceEntry struct {
	Id            int64  `json:"id"`
	InterfaceName string `json:"interfaceName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type ProviderDTO struct {
	Id                 int64              `json:"id"`
	Address            string             `json:"address"`
	AuthenticationInfo *string            `json:"authenticationInfo,omitempty"`
	Metadata           *map[string]string `json:"metadata,omitempty"`
	Port               int                `json:"port"`
	SystemName         string             `json:"systemName"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
}
