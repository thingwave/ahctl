package main


type SystemList struct {
  Data []System 
  Count int
}

type System struct {
    Id int
    SystemName string
    Address string
    Port int
    AuthenticationInfo string
    CreatedAt string
    UpdatedAt string
}

type ServiceDefinitionList struct {
  Data []ServiceDefinition `json:"data"`
  Count int `json:"count"`
}

type ServiceDefinition struct {
  Id int `json:"id"`
  ServiceDefinition string `json:"serviceDefinition"`
  CreatedAt string `json:"createdAt"`
  UpdatedAt string `json:"updatedAt"`
}