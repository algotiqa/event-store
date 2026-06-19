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
	"github.com/algotiqa/core/req"
	"gorm.io/gorm"
)

//=============================================================================

func GetEventsByUser(tx *gorm.DB, username string) (*[]Event, error) {
	var list []Event

	filter := map[string]any{}
	filter["username"] = username

	res := tx.
		Where(filter).
		Order("event_date desc").
		Find(&list)

	if res.Error != nil {
		return nil, req.NewServerErrorByError(res.Error)
	}

	return &list, nil
}

//=============================================================================

func AddEvent(tx *gorm.DB, e *Event) error {
	return tx.Create(e).Error
}

//=============================================================================
