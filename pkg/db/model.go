//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package db

import (
	"time"

	"github.com/algotiqa/core/msg"
)

//=============================================================================
//===
//=== Entities
//===
//=============================================================================

type Event struct {
	Id         uint           `json:"id" gorm:"primaryKey"`
	Username   string         `json:"username"`
	EventDate  time.Time      `json:"eventDate"`
	Level      msg.EventLevel `json:"level"`
	Title      string         `json:"title"`
	Message    string         `json:"message"`
	Parameters []byte         `json:"parameters"`
}

//=============================================================================
//===
//=== Table names
//===
//=============================================================================

func (Event) TableName() string { return "event" }

//=============================================================================
