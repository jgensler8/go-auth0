# Go API client for swagger

A Swagger definition for Auth0

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://OVERRIDE-ME-IN-CODE.auth0.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**UserinfoGet**](docs/DefaultApi.md#userinfoget) | **Get** /userinfo | 


## Documentation For Models

 - [UserInfo](docs/UserInfo.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```
	auth := context.WithValue(context.TODO(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```

## Author


