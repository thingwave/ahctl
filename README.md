# Eclipse Arrowhead Control application

## Introduction
Eclipse Arrowhead Control app, or ahctl, is a small helper tool to interact with an Eclipse Arrowhead local cloud. It can be used to check to availability of several core systems, as well as list systems and services in the ServiceRegistry.

## Build

### AMD64
To compile for 64-bit Intel based platforms, issue the following command:
```
ahctl$ make all
```

### ARM64
To compile for 64-bit ARM based platforms (such as Raspberry Pi), issue the following command:
```
ahctl$ make all-arm64
```

## Usage

Please that some commands require a sysop certificate when using CERTIFICATE mode!

### ServiceRegistry availability
To test if the ServiceRegistry is available, issue the following command (change the IP address and HTTP/HTTPS depending on the local cloud's configuration):
```
ahctl$ ./ahctl --sr=http://192.168.1.10:8443/serviceregistry
Calling http://192.168.1.10:8443/serviceregistry/echo
Got it!
```

### Get all registered systems
To get all registred systems from the ServiceRegistry, issue the following command:
```
ahctl$ ./ahctl  --sr=http://192.168.1.10:8443/serviceregistry --cmd=list-all-systems
{
  "Data": [
    {
      "Id": 1,
      "SystemName": "serviceregistry",
      "Address": "192.168.1.10",
      "Port": 8443,
      "AuthenticationInfo": "",
      "CreatedAt": "2022-09-24T15:43:27Z",
      "UpdatedAt": "2023-01-14T10:30:08Z"
    },
    {
      "Id": 2,
      "SystemName": "datamanager",
      "Address": "192.168.1.10",
      "Port": 8461,
      "AuthenticationInfo": "",
      "CreatedAt": "2022-09-24T15:43:36Z",
      "UpdatedAt": "2023-01-15T00:44:17Z"
    },
    {
      "Id": 9,
      "SystemName": "authorization",
      "Address": "192.168.1.10",
      "Port": 8445,
      "AuthenticationInfo": "",
      "CreatedAt": "2022-09-26T09:51:12Z",
      "UpdatedAt": "2023-01-15T12:16:39Z"
    },
    {
      "Id": 11,
      "SystemName": "serviceregistry",
      "Address": "192.168.1.10",
      "Port": 8443,
      "AuthenticationInfo": "",
      "CreatedAt": "2023-01-14T10:30:41Z",
      "UpdatedAt": "2023-01-15T13:16:33Z"
    },
    {
      "Id": 15,
      "SystemName": "orchestrator",
      "Address": "192.168.1.10",
      "Port": 8441,
      "AuthenticationInfo": "",
      "CreatedAt": "2023-01-14T18:02:17Z",
      "UpdatedAt": "2023-01-14T18:02:17Z"
    }
  ],
  "Count": 5
}
```

### Supported commands
Below is a list of the currently supported commands. More will be added in the future.

#### sr-echo
This command tries to get the "Got it!" response from the ServiceRegistry's /echo endpoint.

#### list-all-systems
This command lists all systems stored in the ServiceRegistry.

#### list-all-services
This command gets the list of all ServiceDefinitions stored in the ServiceRegistry.

#### or-echo
This command tries to get the "Got it!" response from the Orchestrator's /echo endpoint. The 
address of the Orchestrator is automatically queried from the ServiceRegistry.

#### au-echo
This command tries to get the "Got it!" response from the Authorization system's /echo endpoint. The 
address of the Authorization system is automatically queried from the ServiceRegistry.

#### dm-echo
This command tries to get the "Got it!" response from the DataManagers's /echo endpoint. The 
address of the DataManager is automatically queried from the ServiceRegistry.

