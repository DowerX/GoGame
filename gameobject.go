package main

import (
	"reflect"
)

type gameObject struct {
	Position   vector3f
	Rotation   vector3f
	Active     bool
	Components []component
}

func (g *gameObject) SetActive(a bool) {
	g.Active = a
}

func (g *gameObject) GetComponent(c component) component {
	typ := reflect.TypeOf(c)
	for _, v := range g.Components {
		if typ == reflect.TypeOf(v) {
			return v
		}
	}
	return nil
}

func (g *gameObject) AddComponent(c component) {
	for _, existing := range g.Components {
		if reflect.TypeOf(c) == reflect.TypeOf(existing) {
			return
		}
	}
	g.Components = append(g.Components, c)
}

func (g *gameObject) RemoveComponent(c component) {
	typ := reflect.TypeOf(c)
	for i, v := range g.Components {
		if typ == reflect.TypeOf(v) {
			g.Components[i] = g.Components[len(g.Components)-1]
			g.Components[len(g.Components)-1] = nil
			g.Components = g.Components[:len(g.Components)-1]
			return
		}
	}
}
