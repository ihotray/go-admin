// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Engine is the core component of goAdmin. It has two attributes.
// PluginList is an array of plugin. Adapter is the adapter of
// web framework context and goAdmin context. The relationship of adapter and
// plugin is that the adapter use the plugin which contains routers and
// controller methods to inject into the framework entity and make it work.
type Engine struct {
	PluginList []plugins.Plugin
	Adapter    adapter.WebFrameWork
	Config     config.Config
	services   service.List
}

// Default return the default engine instance.
func Default() *Engine {
	return &Engine{
		Adapter:  defaultAdapter,
		services: service.GetServices(),
	}
}

// Use enable the adapter.
func (eng *Engine) Use(router interface{}) error {
	if eng.Adapter == nil {
		panic("adapter is nil, import the default adapter or use AddAdapter method add the adapter")
	}
	return eng.Adapter.Use(router, eng.PluginList)
}

// AddPlugins add the plugins and initialize them.
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	for _, plug := range plugs {
		plug.InitPlugin(eng.services)
	}

	eng.PluginList = append(eng.PluginList, plugs...)
	return eng
}

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	eng.Config = config.Set(cfg)
	return eng
}

// AddConfigFromJSON set the global config from json file.
func (eng *Engine) AddConfigFromJSON(path string) *Engine {
	eng.Config = config.ReadFromJson(path)
	return eng
}

// AddAdapter add the adapter of engine.
func (eng *Engine) AddAdapter(ada adapter.WebFrameWork) *Engine {
	eng.Adapter = ada
	defaultAdapter = ada
	return eng
}

// defaultAdapter is the default adapter of engine.
var defaultAdapter adapter.WebFrameWork

// Register set default adapter of engine.
func Register(ada adapter.WebFrameWork) {
	if ada == nil {
		panic("adapter is nil")
	}
	defaultAdapter = ada
}

// Content call the Content method of defaultAdapter.
// If defaultAdapter is nil, it will panic.
func Content(ctx interface{}, panel types.GetPanelFn) {
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	defaultAdapter.Content(ctx, panel)
}
