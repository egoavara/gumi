package gumi

import "image"

type SingleStructure struct {
	parent GUMI
	child  GUMI
}
func (s *SingleStructure) GUMIInit() {
	s.child.GUMIInit()
}
func (s *SingleStructure) born(gumi GUMI) {
	s.parent = gumi
}
func (s *SingleStructure) breed(gumi []GUMI) {
	if len(gumi) > 0 {
		s.child = gumi[0]
	}
}
func (s *SingleStructure) Parent() GUMI {
	return s.parent
}
func (s *SingleStructure) Childrun() []GUMI {
	return []GUMI{s.child}
}

type MultipleStructure struct {
	parent GUMI
	child  []GUMI
}

func (s *MultipleStructure) GUMIInit() {
	for _ , v := range s.child{
		v.GUMIInit()
	}
}
func (s *MultipleStructure) born(gumi GUMI) {
	s.parent = gumi
}
func (s *MultipleStructure) breed(gumi []GUMI) {
	s.child = gumi
}
func (s *MultipleStructure) Parent() GUMI {
	return s.parent
}
func (s *MultipleStructure) Childrun() []GUMI {
	res := make([]GUMI, len(s.child))
	for i, v := range s.child {
		res[i] = v
	}
	return res
}

type VoidStructure struct {
	parent GUMI
}
func (s *VoidStructure) GUMIInit() {
}
func (s *VoidStructure) born(gumi GUMI) {
	s.parent = gumi
}
func (s *VoidStructure) breed(gumi []GUMI) {
}
func (s *VoidStructure) Parent() GUMI {
	return s.parent
}
func (s *VoidStructure) Childrun() []GUMI {
	return nil
}

type boundStore struct {
	bound image.Rectangle
}
type styleStore struct {
	style *Style
}
