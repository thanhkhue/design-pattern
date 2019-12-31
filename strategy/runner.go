package main

import "log"

type Runner interface {
	Run(string)
}

type RunStrategies struct {
	Runner Runner
}

func (r *RunStrategies) Execute(name string) {
	r.Runner.Run(name)
}

type Jog struct{}

func (j *Jog) Run(name string) {
	log.Printf("%s is now jogging at a comfortable pace.", name)
}

type Sprint struct{}

func (s *Sprint) Run(name string) {
	log.Printf("%s is now sprinting full speeed!", name)
}

type Marathon struct{}

func (m *Marathon) Run(name string) {
	log.Printf("%s is now running at a steady pace.", name)
}
