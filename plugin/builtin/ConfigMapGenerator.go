// +build plugin

package main

import (
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
	"sigs.k8s.io/kustomize/pkg/types"
)

type plugin struct {
	ldr     ifc.Loader
	rf      *resmap.Factory
	options types.GeneratorOptions
	args    types.ConfigMapArgs
}

var KustomizePlugin plugin

func (p *plugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, k ifc.Kunstructured) (err error) {
	p.ldr = ldr
	p.rf = rf
	p.args.GeneratorArgs, err = resmap.GeneratorArgsFromKunstruct(k)
	return
}

func (p *plugin) Generate() (resmap.ResMap, error) {
	argsList := make([]types.ConfigMapArgs, 1)
	argsList[0] = p.args
	return p.rf.NewResMapFromConfigMapArgs(p.ldr, &p.options, argsList)
}
