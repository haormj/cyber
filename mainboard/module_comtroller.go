package mainboard

import (
	"fmt"
	"plugin"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/component"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
)

type ModuleController struct {
	dagPaths []string
}

func NewModuleController(dagPaths []string) *ModuleController {
	return &ModuleController{
		dagPaths: dagPaths,
	}
}

func (c *ModuleController) Init() bool {
	return c.loadAll()
}

func (c *ModuleController) Clear() {

}

func (c *ModuleController) loadAll() bool {
	for _, dagPath := range c.dagPaths {
		if err := c.loadModuleFromPath(dagPath); err != nil {
			log.Logger.Error("load module from path error", "dagPath", dagPath, "err", err)
			return false
		}
	}

	return true
}

func (c *ModuleController) loadModuleFromPath(p string) error {
	dagConfig := &pb.DagConfig{}
	if err := common.GetProtoFromFile(p, dagConfig); err != nil {
		return err
	}

	return c.loadModuleFromConfig(dagConfig)
}

func (c *ModuleController) loadModuleFromConfig(config *pb.DagConfig) error {
	for _, moduleConfig := range config.ModuleConfig {
		if _, err := plugin.Open(moduleConfig.GetModuleLibrary()); err != nil {
			return err
		}

		for _, c := range moduleConfig.Components {
			className := c.GetClassName()
			v, ok := component.ComponentMap.Load(className)
			if !ok {
				return fmt.Errorf("not find %s", className)
			}

			start := v.(component.StartComponent)

			if !start(c.Config) {
				return fmt.Errorf("%s component initialize failed", className)
			}
		}

		for _, c := range moduleConfig.TimerComponents {
			className := c.GetClassName()
			v, ok := component.ComponentMap.Load(className)
			if !ok {
				return fmt.Errorf("not find %s", className)
			}

			start := v.(component.StartTimerComponent)

			if !start(c.Config) {
				return fmt.Errorf("%s timer component initialize failed", className)
			}
		}
	}

	return nil
}
