package model

import (
	"fmt"
	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type converter struct {
	model   *Api
	swagger *spec.Swagger
}

func (c *converter) consume(model *Api) *spec.Swagger {
	c.model = model

	swagger := &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			// everything is json
			Consumes: []string{"application/json"},
			Produces: []string{"application/json"},
			// always https
			Schemes: []string{"https"},
			Swagger: "2.0",
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Title:   "Check Point Management API",
					Version: model.Version,
				},
			},
			Host:                "example.com:443",
			BasePath:            fmt.Sprintf("/%s/%s", model.BasePath, model.Version),
			SecurityDefinitions: make(spec.SecurityDefinitions),
			Definitions:         make(spec.Definitions),
			Tags:                make([]spec.Tag, 0),
		},
	}
	c.swagger = swagger

	swagger.SecurityDefinitions["apiKey"] = &spec.SecurityScheme{
		SecuritySchemeProps: spec.SecuritySchemeProps{
			Description: "Session unique identifier as it returned by the login request",
			Type:        "apiKey",
			In:          "header",
			Name:        "X-chkp-sid",
		},
	}

	c.handlePaths()

	return c.swagger
}

func (c *converter) handlePaths() {
	pathItems := make(map[string]spec.PathItem)
	for k, i := range c.model.Paths {
		pathItems[fmt.Sprintf("/%s", k)] = c.createPathItem(i)
	}
	c.swagger.Paths = &spec.Paths{
		Paths: pathItems,
	}
}

func createParameters(definitionName string, definition Definition) []spec.Parameter {
	ref, _ := jsonreference.New(fmt.Sprintf("#/definitions/%s", definitionName))

	var parameters []spec.Parameter
	parameter := spec.Parameter{
		ParamProps: spec.ParamProps{
			Name:        definitionName,
			Description: definition.Description,
			Required:    true,
			In:          "body",
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Ref: spec.Ref{ref},
				},
			},
		},
	}

	return append(parameters, parameter)
}

func (c *converter) handleDefinition(definitionName string) (string, Definition) {
	definition := c.model.Definitions[definitionName]

	elem := strings.Split(definitionName, ".")
	// use the last element from the definition name
	definitionName = elem[len(elem)-1]
	// ref elements mus not contain `$`
	definitionName = strings.Replace(definitionName, "$", "", 1)

	// convert model to swagger definition, if not already done
	if _, ok := c.swagger.Definitions[definitionName]; !ok {
		fields := make(map[string]spec.Schema)
		required := make([]string, 0)

		c.handleFields(definition.RequiredFields, fields, &required)
		c.handleFields(definition.Fields, fields, &required)
		c.handleFields(definition.UnderMoreFields, fields, &required)

		c.swagger.Definitions[definitionName] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        spec.StringOrArray{"object"},
				Description: definition.Description,
				Required:    required,
				Properties:  fields,
			},
		}
	}

	return definitionName, definition
}

func (c *converter) handleFields(df []Field, fields map[string]spec.Schema, required *[]string) {
	for _, f := range df {
		if f.Required {
			*required = append(*required, f.Name)
		}
		fields[f.Name] = c.handleField(f)
		if f.Alternatives != nil && len(f.Alternatives) > 0 {
			for _, a := range f.Alternatives {
				fields[a.Name] = c.handleField(a)
			}
		}
	}
}

func (c *converter) handleField(f Field) spec.Schema {
	sProps := spec.SchemaProps{
		Description: f.Description,
	}
	c.handleFieldType(f.Types[0], &sProps)
	if f.Default != "" {
		sProps.Default = f.Default
	}
	return spec.Schema{
		SchemaProps: sProps,
	}
}

func (c *converter) handleFieldType(fType TypeOrElementType, sProps *spec.SchemaProps) {
	if fType.GetName() == "" {
		return
	}

	switch fType.GetName() {
	case "list":
		sProps.Type = spec.StringOrArray{"array"}

		itemProps := spec.SchemaProps{}
		c.handleFieldType(fType.GetElementType(), &itemProps)

		sProps.Items = &spec.SchemaOrArray{
			Schema: &spec.Schema{
				SchemaProps: itemProps,
			},
		}
	case "string":
		sProps.Type = spec.StringOrArray{"string"}
		if fType.GetValidValues() != nil {
			var enum []interface{}
			for _, v := range fType.GetValidValues() {
				enum = append(enum, v)
			}
			sProps.Enum = enum
		}
	case "object":
		sProps.Type = spec.StringOrArray{"object"}
		if fType.GetObjectName() != "java.lang.Object" {
			definitionName, _ := c.handleDefinition(fType.GetObjectName())

			ref, _ := jsonreference.New(fmt.Sprintf("#/definitions/%s", definitionName))
			sProps.Items = &spec.SchemaOrArray{
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.Ref{ref},
					},
				},
			}
		}
	case "integer":
		sProps.Type = spec.StringOrArray{"integer"}
		if fType.GetValidValues() != nil {
			var enum []interface{}
			for _, v := range fType.GetValidValues() {
				enum = append(enum, v)
			}
			sProps.Enum = enum
		}
	default:
		sProps.Type = spec.StringOrArray{fType.GetName()}
	}
}

func (c *converter) createPathItem(command Command) spec.PathItem {
	definitionName, definition := c.handleDefinition(command.Definition)

	item := spec.PathItem{
		PathItemProps: spec.PathItemProps{
			Post: &spec.Operation{
				OperationProps: spec.OperationProps{
					Description: command.Description,
					Deprecated:  command.Deprecated,
					Tags:        []string{c.createTag(definition.Name)},
				},
			},
		},
	}
	item.Post.Parameters = createParameters(definitionName, definition)
	responses, _ := c.createResponses(command.Response)
	item.Post.Responses = responses

	if command.Name != "login" {
		securityList := make([]map[string][]string, 0)
		optionList := make([]string, 0)
		sec := make(map[string][]string)
		sec["apiKey"] = optionList
		securityList = append(securityList, sec)
		item.Post.Security = securityList
	}

	return item
}

func (c *converter) createTag(definitionName string) string {
	elements := make([]string, 0)
	startAt := 6
	breakAt := startAt + 1
	for i, t := range strings.Split(definitionName, ".") {
		if i < startAt {
			continue
		}
		if t == "network_objects" {
			breakAt++
		}
		if t == "base" {
			breakAt++
		}
		if t != "objects" {
			elements = append(elements, t)
		}
		if i >= breakAt {
			break
		}
	}

	tag := strings.Join(elements, " ")

	for _, t := range c.swagger.Tags {
		if t.Name == tag {
			return tag
		}
	}

	c.swagger.Tags = append(c.swagger.Tags, spec.Tag{
		TagProps: spec.TagProps{
			Name: tag,
		},
	})

	return tag
}

func (c *converter) createResponses(responseDefinition Response) (*spec.Responses, error) {
	responses := make(map[int]spec.Response)

	definitionName, _ := c.handleDefinition(responseDefinition.Success.Name)
	ref, _ := jsonreference.New(fmt.Sprintf("#/definitions/%s", definitionName))
	code, err := strconv.Atoi(responseDefinition.Success.Code)
	if err != nil {
		return nil, errors.Errorf("Success response code is not integer: %s", err.Error())
	}
	responses[code] = spec.Response{
		ResponseProps: spec.ResponseProps{
			Description: definitionName,
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Ref: spec.Ref{ref},
				},
			},
		},
	}
	definitionName, _ = c.handleDefinition(responseDefinition.Failure.Name)
	ref, _ = jsonreference.New(fmt.Sprintf("#/definitions/%s", definitionName))
	failureResponse := spec.Response{
		ResponseProps: spec.ResponseProps{
			Description: definitionName,
			Schema: &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Ref: spec.Ref{ref},
				},
			},
		},
	}

	codes := strings.Split(responseDefinition.Failure.Code, ",")
	for _, c := range codes {
		code, err := strconv.Atoi(c)
		if err != nil {
			return nil, errors.Errorf("Failure response code is not integer: %s", err.Error())
		}
		responses[code] = failureResponse
	}

	return &spec.Responses{
		ResponsesProps: spec.ResponsesProps{
			StatusCodeResponses: responses,
		},
	}, nil
}