/********************************************************************************
 * Copyright (c) 2022 ThingWave AB
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License 2.0 which is available at
 * http://www.eclipse.org/legal/epl-2.0.
 *
 * SPDX-License-Identifier: EPL-2.0
 *
 * Contributors:
 *   ThingWave AB - implementation
 ********************************************************************************/

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

type GroupedDTO struct {
	ServicesGroupedBySystems []ServicesGroupedBySystem `json:"servicesGroupedBySystems"`
	//ServicesGroupedByServiceDefinition []xxx `json:"servicesGroupedByServiceDefinition"`
	AutoCompleteData AutoCompleteDataDTO `json:"autoCompleteData"`
}

type ServicesGroupedBySystem struct {
	SystemId   int64               `json:"systemId"`
	SystemName string              `json:"systemName"`
	Address    string              `json:"address"`
	Port       int                 `json:"port"`
	Services   []ServiceQueryEntry `json:"services"`
}

type AutoCompleteDataDTO struct {
	ServiceList   []EntryDTO    `json:"serviceList"`
	SystemList    []ProviderDTO `json:"systemList"`
	InterfaceList []EntryDTO    `json:"interfaceList"`
}

type EntryDTO struct {
	Id    int64  `json:"id"`
	Value string `json:"value"`
}


type OrchestratorStoreListResponseDTO struct {
	count int64  `json:"count"`
	data []OrchestratorStoreResponseDTO `json:"data"`
}

type OrchestratorStoreResponseDTO struct {
	Id int64 `json:"id"`
	ServiceDefinition ServiceDefinitionDTO `json:"serviceDefinition"`
	ConsumerSystem ProviderDTO `json:"consumerSystem"`
	Foreign bool `json:"foreign"`
	ProviderSystem ProviderDTO `json:"providerSystem"`
	ProviderCloud CloudResponseDTO `json:"providerCloud"`
	ServiceInterface InterfaceEntry `json:"serviceInterface"`
	Priority int64 `json:"priority"`
	Attribute map[string]string `json:"attribute"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CloudResponseDTO struct {
        Id int64 `json:"id"`
        Operator string `json:"operator"`
        Name string `json:"name"`
        AuthenticationInfo string `json:"authenticationInfo"`
        Secure bool `json:"secure"`
        Neighbor bool `json:"neighbor"`
        OwnCloud bool `json:"ownCloud"`
        CreatedAt string `json:"createdAt"`
        UpdatedAt string `json:"updatedAt"`
}
