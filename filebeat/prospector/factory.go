package prospector

import (
	"github.com/elastic/beats/filebeat/registrar"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type Factory struct {
	outlet    Outlet
	registrar *registrar.Registrar
	beatDone  chan struct{}
}

func NewFactory(outlet Outlet, registrar *registrar.Registrar, beatDone chan struct{}) *Factory {
	return &Factory{
		outlet:    outlet,
		registrar: registrar,
		beatDone:  beatDone,
	}
}

func (r *Factory) Create(c *common.Config) (cfgfile.Runner, error) {

	p, err := NewProspector(c, r.outlet, r.registrar, r.beatDone)
	if err != nil {
		logp.Err("Error creating prospector: %s", err)
		return nil, err
	}

	return p, nil
}
