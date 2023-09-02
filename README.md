# Go API client for salad-client

The SaladCloud Public API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0-alpha.56
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.salad.com](https://support.salad.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import salad-client "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), salad-client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), salad-client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), salad-client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), salad-client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.salad.com/api/public*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContainerGroupsAPI* | [**ContainerGroupInstanceReallocate**](docs/ContainerGroupsAPI.md#containergroupinstancereallocate) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/reallocate | Reallocate container group instance to another node
*ContainerGroupsAPI* | [**ContainerGroupInstanceRecreate**](docs/ContainerGroupsAPI.md#containergroupinstancerecreate) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/recreate | Recreate container on a node
*ContainerGroupsAPI* | [**ContainerGroupInstanceRestart**](docs/ContainerGroupsAPI.md#containergroupinstancerestart) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/restart | Restart container on a node
*ContainerGroupsAPI* | [**CreateContainerGroup**](docs/ContainerGroupsAPI.md#createcontainergroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers | Create a Container Group
*ContainerGroupsAPI* | [**DeleteContainerGroup**](docs/ContainerGroupsAPI.md#deletecontainergroup) | **Delete** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Delete a Container Group
*ContainerGroupsAPI* | [**GetContainerGroup**](docs/ContainerGroupsAPI.md#getcontainergroup) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Get a Container Group
*ContainerGroupsAPI* | [**ListContainerGroupInstances**](docs/ContainerGroupsAPI.md#listcontainergroupinstances) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances | List Container Group Instances
*ContainerGroupsAPI* | [**ListContainerGroups**](docs/ContainerGroupsAPI.md#listcontainergroups) | **Get** /organizations/{organization_name}/projects/{project_name}/containers | List Container Groups
*ContainerGroupsAPI* | [**RestartContainerGroup**](docs/ContainerGroupsAPI.md#restartcontainergroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/restart | Restart a Container Group
*ContainerGroupsAPI* | [**StartContainerGroup**](docs/ContainerGroupsAPI.md#startcontainergroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start | Start a Container Group
*ContainerGroupsAPI* | [**StopContainerGroup**](docs/ContainerGroupsAPI.md#stopcontainergroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop | Stop a Container Group
*ContainerGroupsAPI* | [**UpdateContainerGroup**](docs/ContainerGroupsAPI.md#updatecontainergroup) | **Patch** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Update a Container Group
*OrganizationDataAPI* | [**ListGpuClasses**](docs/OrganizationDataAPI.md#listgpuclasses) | **Get** /organizations/{organization_name}/gpu-classes | List the GPU Classes
*QuotasAPI* | [**GetQuotas**](docs/QuotasAPI.md#getquotas) | **Get** /organizations/{organization_name}/quotas | Get Quotas
*RecipeDeploymentsAPI* | [**CreateRecipeDeployment**](docs/RecipeDeploymentsAPI.md#createrecipedeployment) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments | Create a Recipe Deployment
*RecipeDeploymentsAPI* | [**DeleteRecipeDeployment**](docs/RecipeDeploymentsAPI.md#deleterecipedeployment) | **Delete** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Delete a Recipe Deployment
*RecipeDeploymentsAPI* | [**GetRecipeDeployment**](docs/RecipeDeploymentsAPI.md#getrecipedeployment) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Get a Recipe Deployment
*RecipeDeploymentsAPI* | [**ListRecipeDeploymentInstances**](docs/RecipeDeploymentsAPI.md#listrecipedeploymentinstances) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/instances | List Recipe Deployment Instances
*RecipeDeploymentsAPI* | [**ListRecipeDeployments**](docs/RecipeDeploymentsAPI.md#listrecipedeployments) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments | List Recipe Deployments
*RecipeDeploymentsAPI* | [**RestartDeployedRecipe**](docs/RecipeDeploymentsAPI.md#restartdeployedrecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/restart | Restart a Deployed Recipe
*RecipeDeploymentsAPI* | [**StartDeployedRecipe**](docs/RecipeDeploymentsAPI.md#startdeployedrecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/start | Start a Deployed Recipe
*RecipeDeploymentsAPI* | [**StopDeployedRecipe**](docs/RecipeDeploymentsAPI.md#stopdeployedrecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/stop | Stop a Deployed Recipe
*RecipeDeploymentsAPI* | [**UpdateRecipeDeployment**](docs/RecipeDeploymentsAPI.md#updaterecipedeployment) | **Patch** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Update a Recipe Deployment
*RecipesAPI* | [**GetRecipe**](docs/RecipesAPI.md#getrecipe) | **Get** /organizations/{organization_name}/recipes/{recipe_name} | Get a Recipe
*RecipesAPI* | [**ListRecipes**](docs/RecipesAPI.md#listrecipes) | **Get** /organizations/{organization_name}/recipes | List Recipes


## Documentation For Models

 - [Container](docs/Container.md)
 - [ContainerGroup](docs/ContainerGroup.md)
 - [ContainerGroupInstanceStatusCount](docs/ContainerGroupInstanceStatusCount.md)
 - [ContainerGroupInstances](docs/ContainerGroupInstances.md)
 - [ContainerGroupInstancesInstancesInner](docs/ContainerGroupInstancesInstancesInner.md)
 - [ContainerGroupList](docs/ContainerGroupList.md)
 - [ContainerGroupNetworking](docs/ContainerGroupNetworking.md)
 - [ContainerGroupProbe](docs/ContainerGroupProbe.md)
 - [ContainerGroupProbeExec](docs/ContainerGroupProbeExec.md)
 - [ContainerGroupProbeGrpc](docs/ContainerGroupProbeGrpc.md)
 - [ContainerGroupProbeHttp](docs/ContainerGroupProbeHttp.md)
 - [ContainerGroupProbeTcp](docs/ContainerGroupProbeTcp.md)
 - [ContainerGroupState](docs/ContainerGroupState.md)
 - [ContainerGroupStatus](docs/ContainerGroupStatus.md)
 - [ContainerGroupsQuotas](docs/ContainerGroupsQuotas.md)
 - [ContainerLogging](docs/ContainerLogging.md)
 - [ContainerLoggingNewRelic](docs/ContainerLoggingNewRelic.md)
 - [ContainerLoggingSplunk](docs/ContainerLoggingSplunk.md)
 - [ContainerLoggingTcp](docs/ContainerLoggingTcp.md)
 - [ContainerNetworkingProtocol](docs/ContainerNetworkingProtocol.md)
 - [ContainerProbeHttpScheme](docs/ContainerProbeHttpScheme.md)
 - [ContainerResourceRequirements](docs/ContainerResourceRequirements.md)
 - [ContainerRestartPolicy](docs/ContainerRestartPolicy.md)
 - [CountryCode](docs/CountryCode.md)
 - [CreateContainer](docs/CreateContainer.md)
 - [CreateContainerGroup](docs/CreateContainerGroup.md)
 - [CreateContainerGroupNetworking](docs/CreateContainerGroupNetworking.md)
 - [CreateContainerGroupProblemType](docs/CreateContainerGroupProblemType.md)
 - [CreateContainerRegistryAuthentication](docs/CreateContainerRegistryAuthentication.md)
 - [CreateContainerRegistryAuthenticationAwsEcr](docs/CreateContainerRegistryAuthenticationAwsEcr.md)
 - [CreateContainerRegistryAuthenticationBasic](docs/CreateContainerRegistryAuthenticationBasic.md)
 - [CreateContainerRegistryAuthenticationDockerHub](docs/CreateContainerRegistryAuthenticationDockerHub.md)
 - [CreateContainerRegistryAuthenticationGcpGcr](docs/CreateContainerRegistryAuthenticationGcpGcr.md)
 - [CreateRecipeDeployment](docs/CreateRecipeDeployment.md)
 - [CreateRecipeDeploymentProblemType](docs/CreateRecipeDeploymentProblemType.md)
 - [CreateRecipeNetworking](docs/CreateRecipeNetworking.md)
 - [DeleteContainerGroupProblemType](docs/DeleteContainerGroupProblemType.md)
 - [DeleteRecipeDeploymentProblemType](docs/DeleteRecipeDeploymentProblemType.md)
 - [GetContainerGroupProblemType](docs/GetContainerGroupProblemType.md)
 - [GetRecipeDeploymentProblemType](docs/GetRecipeDeploymentProblemType.md)
 - [GetRecipeProblemType](docs/GetRecipeProblemType.md)
 - [GpuClass](docs/GpuClass.md)
 - [GpuClassesList](docs/GpuClassesList.md)
 - [HttpHeadersInner](docs/HttpHeadersInner.md)
 - [ListContainerGroupsProblemType](docs/ListContainerGroupsProblemType.md)
 - [ListGpuClassesProblemType](docs/ListGpuClassesProblemType.md)
 - [ListRecipeDeploymentsProblemType](docs/ListRecipeDeploymentsProblemType.md)
 - [ListRecipesProblemType](docs/ListRecipesProblemType.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [Quotas](docs/Quotas.md)
 - [ReallocateContainerGroupInstanceProblemType](docs/ReallocateContainerGroupInstanceProblemType.md)
 - [Recipe](docs/Recipe.md)
 - [RecipeDeployment](docs/RecipeDeployment.md)
 - [RecipeDeploymentInstances](docs/RecipeDeploymentInstances.md)
 - [RecipeDeploymentInstancesInstancesInner](docs/RecipeDeploymentInstancesInstancesInner.md)
 - [RecipeDeploymentList](docs/RecipeDeploymentList.md)
 - [RecipeList](docs/RecipeList.md)
 - [RecipeNetworking](docs/RecipeNetworking.md)
 - [RecipeNetworkingProtocol](docs/RecipeNetworkingProtocol.md)
 - [RecipeResources](docs/RecipeResources.md)
 - [RecipesQuotas](docs/RecipesQuotas.md)
 - [RecreateContainerGroupInstanceProblemType](docs/RecreateContainerGroupInstanceProblemType.md)
 - [RestartContainerGroupInstanceProblemType](docs/RestartContainerGroupInstanceProblemType.md)
 - [RestartContainerGroupProblemType](docs/RestartContainerGroupProblemType.md)
 - [RestartRecipeDeploymentProblemType](docs/RestartRecipeDeploymentProblemType.md)
 - [StartContainerGroupProblemType](docs/StartContainerGroupProblemType.md)
 - [StartRecipeDeploymentProblemType](docs/StartRecipeDeploymentProblemType.md)
 - [StopContainerGroupProblemType](docs/StopContainerGroupProblemType.md)
 - [StopRecipeDeploymentProblemType](docs/StopRecipeDeploymentProblemType.md)
 - [UpdateContainerGroup](docs/UpdateContainerGroup.md)
 - [UpdateContainerGroupProblemType](docs/UpdateContainerGroupProblemType.md)
 - [UpdateRecipeDeployment](docs/UpdateRecipeDeployment.md)
 - [UpdateRecipeDeploymentProblemType](docs/UpdateRecipeDeploymentProblemType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Salad-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Salad-Api-Key and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Salad-Api-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@salad.com

