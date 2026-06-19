//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package service

import (
	"log/slog"

	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/core/auth/roles"
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/event-store/pkg/app"
	"github.com/gin-gonic/gin"
)

//=============================================================================

func Init(router *gin.Engine, cfg *app.Config, logger *slog.Logger) {

	ctrl := auth.NewOidcController(cfg.Authentication.Authority, req.GetDefaultClient(), logger, cfg)

	router.GET("/api/event/v1/events", ctrl.Secure(getEvents, roles.Admin_User))
}

//=============================================================================
