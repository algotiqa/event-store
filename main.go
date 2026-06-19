//=============================================================================
//===
//=== Copyright (C) 2025-present Andrea Carboni
//===
//=== This source code is licensed under the Elastic License 2.0 (ELv2) available at:
//=== https://github.com/algotiqa/docs/blob/main/LICENSE.md
//=== By using this file, you agree to the terms and conditions of that license.
//=============================================================================

package main

import (
	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/core/boot"
	"github.com/algotiqa/core/dbms"
	"github.com/algotiqa/core/msg"
	"github.com/algotiqa/core/req"
	"github.com/algotiqa/event-store/pkg/app"
	"github.com/algotiqa/event-store/pkg/core/messaging"
	"github.com/algotiqa/event-store/pkg/service"

	"log/slog"
)

//=============================================================================

const component = "event-store"

//=============================================================================

func main() {
	cfg := &app.Config{}
	boot.ReadConfig(component, cfg)
	logger := boot.InitLogger(component, &cfg.Application)
	engine := boot.InitEngine(logger, &cfg.Application)
	initClients()
	dbms.InitDatabase(&cfg.Database)
	auth.InitAuthentication(&cfg.Authentication)
	msg.InitMessaging(&cfg.Messaging)
	service.Init(engine, cfg, logger)
	messaging.InitMessageListener()
	boot.RunHttpServer(engine, &cfg.Application)
}

//=============================================================================

func initClients() {
	slog.Info("Initializing clients...")
	req.AddDefaultClient("ca.crt", "server.crt", "server.key")
}

//=============================================================================
