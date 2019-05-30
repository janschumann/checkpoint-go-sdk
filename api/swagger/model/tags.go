package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

type Content struct {
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Name string `json:"name"`
	Commands []Command `json:"commands-data"`
	SubChapters []Chapter `json:"sub-chapters"`
}

type Tags map[string][]string

func (t *Tags) getTags() []string {
	data := *t

	result := make([]string, 0)
	for tag, _ := range data {
		result = append(result, tag)
	}
	sort.Strings(result)

	return result
}

func (t *Tags) getTagForCommand(command string) string {
	tags := *t
	for tag, cmds := range tags {
		for _, cmd := range cmds {
			if cmd == command {
				return tag
			}
		}
	}

	return "Unknown"
}

func UnmarshalTags(b bytes.Buffer) (*Tags, error) {
	var content Content
	err := json.Unmarshal(b.Bytes(), &content)
	if err != nil {
		return nil, err
	}

	tags := make(Tags)
	for _, c := range content.Chapters {
		if len(c.SubChapters) > 0 {
			for _, s := range c.SubChapters {
				name := fmt.Sprintf("%s: %s", c.Name, s.Name)
				commands := make([]string, 0)
				if len(s.SubChapters) > 0 {
					for _, ss := range s.SubChapters {
						if len(ss.Commands) > 0 {
							for _, cmd := range ss.Commands {
								commands = append(commands, cmd.Name)
							}
						}
					}
				}
				if len(s.Commands) > 0 {
					for _, cmd := range s.Commands {
						commands = append(commands, cmd.Name)
					}
				}
				tags[name] = commands
			}
			if len(c.Commands) > 0 {
				commands := make([]string, 0)
				for _, cmd := range c.Commands {
					commands = append(commands, cmd.Name)
				}
				tags[c.Name] = commands
			}
		}
	}

	return &tags, err
}

func (t Tags) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}

