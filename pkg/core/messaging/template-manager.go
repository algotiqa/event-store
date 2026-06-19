//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package messaging

import (
	"os"

	"github.com/algotiqa/core"
	"gopkg.in/yaml.v3"
)

//=============================================================================

type Template struct {
	Code    string
	Level   string
	Title   string
	Message string
}

//=============================================================================

var templates = make(map[string]*Template)

//=============================================================================

func initTemplates() {
	var content map[string]interface{}

	file, err := os.ReadFile("config/event-templates.yaml")
	core.ExitIfError(err)

	err = yaml.Unmarshal(file, &content)
	core.ExitIfError(err)

	buildTemplateMap("", content)
}

//=============================================================================

func buildTemplateMap(path string, currMap map[string]interface{}) {
	title, ok := currMap["title"]
	message, _ := currMap["message"]
	level, _ := currMap["level"]

	if ok {
		templates[path] = &Template{
			Code:    path,
			Level:   level.(string),
			Title:   title.(string),
			Message: message.(string),
		}
	} else {
		for key, value := range currMap {
			subPath := path + "." + key

			if path == "" {
				subPath = key
			}

			buildTemplateMap(subPath, value.(map[string]interface{}))
		}
	}
}

//=============================================================================
