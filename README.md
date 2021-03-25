<p align="center">
  <a href="https://opensource.org/licenses/MPL-2.0"
    ><img
      src="https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg"
      alt="License: MPL 2.0"
  /></a>
  <a href="https://github.com/magodo/terrassist/actions"
    ><img
      src="https://img.shields.io/github/workflow/status/magodo/terrassist/Provider?label=workflow&style=flat-square"
      alt="GitHub Actions workflow status"
  /></a>
  <a href="https://github.com/magodo/terrassist/releases"
    ><img
      src="https://img.shields.io/github/release/magodo/terrassist.svg?style=flat-square"
      alt="Relase status"/></a
  ><br />
  <a href="https://twitter.com/magodo_"
    ><img
      src="https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Ftwitter.com%2Fmagodo_"
      alt="Follow me on Twitter"
  /></a>
</p>

# Important Disclaimer

I am a Microsoft employee, but this is not an official Microsoft product nor an endorsed product. Purely a project for fun. 

# Terrassist

A small tool to assist Terraform Provider developers to generate `flatten/expand` functions.

## Install

For Go v1.16+: 

```
go install github.com/magodo/terrassist@latest
```

Otherwise:

```
go get github.com/magodo/terrassist
```

## Usage

This tool is used to generate Go code that defines the Terraform Provider `flatten/expand` function for a particular Go type (e.g. a structure). Normally, the target type is defined in some SDK that your provider mainly depends on (e.g. [terraform-providers/terraform-provider-azurerm](https://github.com/terraform-providers/terraform-provider-azurerm) depends on [Azure/azure-sdk-for-go](https://github.com/azure/azure-sdk-for-go)).

To generate the code, users will need to run this tool under the provider codebase folder, where the codebase is technically a **Go module that has added the dependency towards the target SDK module**. Then follow the usage of this tool as below, to run the tool:

```
terrassist <pkg> <type>
  -j    Ignore struct field that has json tag "-" specified
  -p    Whether to generate for the pointer of the specified type?
```

## Example

Under the cloned [terraform-providers/terraform-provider-aws](https://github.com/terraform-providers/terraform-provider-aws) repository root folder, run:

```
terrassist github.com/aws/aws-sdk-go/service/elasticache CreateCacheParameterGroupInput
```

It will then output:

```go
package main

import (
        elasticache "github.com/aws/aws-sdk-go/service/elasticache"
        utils "types/utils"
)

func expandCreateCacheParameterGroupInput(input []interface{}) elasticache.CreateCacheParameterGroupInput {
        if len(input) == 0 || input[0] == nil {
                return elasticache.CreateCacheParameterGroupInput{}
        }
        b := input[0].(map[string]interface{})
        output := elasticache.CreateCacheParameterGroupInput{
                CacheParameterGroupFamily: utils.String(b["cache_parameter_group_family"].(string)),
                CacheParameterGroupName:   utils.String(b["cache_parameter_group_name"].(string)),
                Description:               utils.String(b["description"].(string)),
        }
        return output
}
func flattenCreateCacheParameterGroupInput(input elasticache.CreateCacheParameterGroupInput) []interface{} {
        var cacheParameterGroupFamily string
        if input.CacheParameterGroupFamily != nil {
                cacheParameterGroupFamily = *input.CacheParameterGroupFamily
        }
        var cacheParameterGroupName string
        if input.CacheParameterGroupName != nil {
                cacheParameterGroupName = *input.CacheParameterGroupName
        }
        var description string
        if input.Description != nil {
                description = *input.Description
        }
        return []interface{}{map[string]interface{}{
                "cache_parameter_group_family": cacheParameterGroupFamily,
                "cache_parameter_group_name":   cacheParameterGroupName,
                "description":                  description,
        }}
}
```

Note: 

- The `package` name will be set to the package name of the folder you are running this tool.
- The namings used in the Terraform `ResourceData`/schema is the snake cased form of the corresponding field.
- There is a hypothetic `types/utils` package involved. The reason is that most popular providers (e.g. aws, azurerm/azuread, etc.) has their own `utils` package that provides functions that convert a basic type to its pointer. So we are not reinventing this, rather users need only replace the import path to using the one that is available in your provider codebase.
