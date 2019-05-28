package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

func Unmarshal(b bytes.Buffer) (error, *Api) {
	var model Model
	err := json.Unmarshal(b.Bytes(), &model)
	if err != nil {
		return err, nil
	}

	api := &Api{
		Version:     model.Version,
		BasePath:    model.Path,
		Paths:       make(map[string]Command),
		Definitions: make(map[string]Definition),
	}

	for _, c := range model.Commands {
		api.Paths[c.Name] = c
	}

	for _, d := range model.Definitions {
		api.Definitions[d.Name] = d
	}

	return nil, api
}

type Api struct {
	Version     string
	BasePath    string
	Paths       map[string]Command
	Definitions map[string]Definition
}

func (a *Api) Convert() *spec.Swagger {
	c := &converter{}
	return c.consume(a)
}

type Model struct {
	Version     string       `json:"version"`
	Path        string       `json:"path"`
	Commands    []Command    `json:"commands"`
	Definitions []Definition `json:"objects"`
}

type Command struct {
	Name                 string                 `json:"name"`
	Type                 string                 `json:"type"`
	Definition           string                 `json:"request"`
	Response             Response               `json:"response"`
	Description          string                 `json:"description"`
	Documented           bool                   `json:"documented"`
	Internal             bool                   `json:"internal"`
	Deprecated           bool                   `json:"deprecated"`
	AdditionalProperties map[string]interface{} `json:""`
}

func (c *Command) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if c.AdditionalProperties == nil {
					c.AdditionalProperties = make(map[string]interface{})
				}
				c.AdditionalProperties[k] = additionalProperties
			}
		case "name":
			if v != nil {
				var nameData map[string]interface{}
				err = json.Unmarshal(*v, &nameData)
				if err != nil {
					return err
				}

				if name, ok := nameData["web"].(string); ok {
					c.Name = name
					continue
				}

				return errors.New(fmt.Sprintf("Could not find 'name' in name data %v", v))
			}
		case "type":
			commandType, err := unmarshalString(v)
			if err != nil {
				return err
			}
			c.Type = *commandType
		case "request":
			definition, err := unmarshalString(v)
			if err != nil {
				return err
			}
			c.Definition = *definition
		case "response":
			if v != nil {
				var response Response
				err = json.Unmarshal(*v, &response)
				if err != nil {
					return err
				}
				c.Response = response
			}
		case "description":
			description, err := unmarshalString(v)
			if err != nil {
				return err
			}
			c.Description = *description
		case "documented":
			documented, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			c.Documented = *documented
		case "internal":
			internal, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			c.Internal = *internal
		case "deprecated":
			deprecated, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			c.Deprecated = *deprecated
		}
	}

	return nil
}

type Response struct {
	Success ResponseObject `json:"on-success"`
	Failure ResponseObject `json:"on-failure"`
}

func (r *Response) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			continue
		case "on-success":
			if v != nil {
				var success ResponseObject
				err = json.Unmarshal(*v, &success)
				if err != nil {
					return err
				}
				r.Success = success
			}
		case "on-failure":
			if v != nil {
				var failure ResponseObject
				err = json.Unmarshal(*v, &failure)
				if err != nil {
					return err
				}
				r.Failure = failure
			}
		}
	}

	return nil
}

type ResponseObject struct {
	Code string `json:"status-code"`
	Name string `json:"object"`
}

func (r *ResponseObject) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}

	if v, ok := m["web"]; ok {
		if v != nil {
			var webResponse map[string]interface{}
			err = json.Unmarshal(*v, &webResponse)
			if err != nil {
				return err
			}
			if code, ok := webResponse["status-code"].(string); ok {
				r.Code = code
			}
			if object, ok := webResponse["object"].(map[string]interface{}); ok {
				if objectName, ok := object["object-name"].(string); ok {
					r.Name = objectName
				}
			}

			return nil
		}
	}

	return errors.New(fmt.Sprintf("Could not find 'web' response object in %v", m))
}

type Definition struct {
	Name                 string                 `json:"name"`
	ObjectType           string                 `json:"objectType"`
	RequestType          string                 `json:"requestType"`
	Show                 bool                   `json:"show"`
	Common               bool                   `json:"common"`
	Description          string                 `json:"description"`
	Fields               []Field                `json:"fields"`
	UnderMoreFields      []Field                `json:"under-more-fields"`
	RequiredFields       []Field                `json:"required-fields"`
	AdditionalProperties map[string]interface{} `json:""`
}

func (d *Definition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if d.AdditionalProperties == nil {
					d.AdditionalProperties = make(map[string]interface{})
				}
				d.AdditionalProperties[k] = additionalProperties
			}
		case "name":
			name, err := unmarshalString(v)
			if err != nil {
				return err
			}
			d.Name = *name
		case "objectType":
			objectType, err := unmarshalString(v)
			if err != nil {
				return err
			}
			d.ObjectType = *objectType
		case "requestType":
			requestType, err := unmarshalString(v)
			if err != nil {
				return err
			}
			d.RequestType = *requestType
		case "show":
			show, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			d.Show = *show
		case "common":
			common, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			d.Common = *common
		case "description":
			description, err := unmarshalString(v)
			if err != nil {
				return err
			}
			d.Description = *description
		case "fields":
			if v != nil {
				var fields []Field
				err = json.Unmarshal(*v, &fields)
				if err != nil {
					return err
				}
				d.Fields = fields
			}
		case "under-more-fields":
			if v != nil {
				var moreFields []Field
				err = json.Unmarshal(*v, &moreFields)
				if err != nil {
					return err
				}
				d.UnderMoreFields = moreFields
			}
		case "required-fields":
			if v != nil {
				var requiredFields []Field
				err = json.Unmarshal(*v, &requiredFields)
				if err != nil {
					return err
				}
				d.RequiredFields = requiredFields
			}
		}

	}

	return nil
}

type TypeOrElementType interface {
	GetName() string
	GetValidValues() []string
	GetObjectName() string
	GetElementType() ElementType
}

type Type struct {
	Name        string      `json:"name"`
	ValidValues []string    `json:"valid-values"`
	ObjectName  string      `json:"object-name"`
	ElementType ElementType `json:"element-type"`
}

func (t Type) GetName() string {
	return t.Name
}

func (t Type) GetValidValues() []string {
	return uniqueNonEmptyElementsOf(t.ValidValues)
}

func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}

func (t Type) GetObjectName() string {
	return t.ObjectName
}

func (t Type) GetElementType() ElementType {
	return t.ElementType
}

type ElementType struct {
	Name       string `json:"name"`
	ObjectName string `json:"object-name"`
}

func (t ElementType) GetName() string {
	return t.Name
}

func (t ElementType) GetValidValues() []string {
	return nil
}

func (t ElementType) GetObjectName() string {
	return t.ObjectName
}

func (t ElementType) GetElementType() ElementType {
	return ElementType{}
}

type Field struct {
	Name                 string                 `json:"name"`
	Description          string                 `json:"description"`
	Types                []Type                 `json:"types"`
	UnderMore            bool                   `json:"under-more"`
	Deprecated           bool                   `json:"deprecated"`
	FieldOrder           int                    `json:"field-order"`
	Default              string                 `json:"default-value"`
	Required             bool                   `json:"required"`
	Alternatives         []Field                `json:"field-alternatives"`
	InRequest            bool                   `json:"field-in-request"`
	AdditionalProperties map[string]interface{} `json:""`
}

func (f *Field) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if f.AdditionalProperties == nil {
					f.AdditionalProperties = make(map[string]interface{})
				}
				f.AdditionalProperties[k] = additionalProperties
			}
		case "name":
			name, err := unmarshalString(v)
			if err != nil {
				return err
			}
			f.Name = *name
		case "description":
			description, err := unmarshalString(v)
			if err != nil {
				return err
			}
			f.Description = *description
		case "types":
			var types []Type
			err = json.Unmarshal(*v, &types)
			if err != nil {
				return err
			}
			f.Types = types
		case "show":
			more, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			f.UnderMore = *more
		case "deprecated":
			deprecated, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			f.Deprecated = *deprecated
		case "field-order":
			order, err := unmarshalInt(v)
			if err != nil {
				return err
			}
			f.FieldOrder = *order
		case "default-value":
			defaultValue, err := unmarshalString(v)
			if err != nil {
				return err
			}
			f.Default = *defaultValue
		case "required":
			reqired, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			f.Required = *reqired
		case "field-alternatives":
			var fields []Field
			err = json.Unmarshal(*v, &fields)
			if err != nil {
				return err
			}
			f.Alternatives = fields
		case "field-in-request":
			inRequest, err := unmarshalBool(v)
			if err != nil {
				return err
			}
			f.InRequest = *inRequest
		}

	}

	return nil
}

func unmarshalString(v *json.RawMessage) (*string, error) {
	if v != nil {
		var value *string
		err := json.Unmarshal(*v, &value)
		return value, err
	}

	return nil, errors.New(fmt.Sprintf("Cannot unmarshal %v to string.", v))
}

func unmarshalInt(v *json.RawMessage) (*int, error) {
	if v != nil {
		var value *int
		err := json.Unmarshal(*v, &value)
		return value, err
	}

	return nil, errors.New(fmt.Sprintf("Cannot unmarshal %v to int.", v))
}

func unmarshalBool(v *json.RawMessage) (*bool, error) {
	if v != nil {
		var value *bool
		err := json.Unmarshal(*v, &value)
		return value, err
	}

	return nil, errors.New(fmt.Sprintf("Cannot unmarshal %v to bool.", v))
}
