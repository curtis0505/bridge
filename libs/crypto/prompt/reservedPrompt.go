package prompt

import (
	"fmt"

	"github.com/klaytn/klaytn/console"
)

type ReservedPromptInput struct {
	useReserved bool
	inputs      []string
}

func NewReservedPromptInput(inputs ...string) *ReservedPromptInput {
	return &ReservedPromptInput{
		inputs:      inputs,
		useReserved: len(inputs) > 0,
	}
}

func (p *ReservedPromptInput) PopFront() (string, bool) {
	if len(p.inputs) == 0 {
		return "", false
	}
	input := p.inputs[0]
	p.inputs = p.inputs[1:]

	if len(p.inputs) == 0 {
		p.useReserved = false
	}
	return input, true
}

func (p *ReservedPromptInput) PushBack(inputs ...string) {
	p.inputs = append(p.inputs, inputs...)

	if len(p.inputs) > 0 {
		p.useReserved = true
	}
}

func (p *ReservedPromptInput) PromptInput(prompt string) (string, error) {
	if p.useReserved == false {
		return console.Stdin.PromptInput(prompt)
	}
	fmt.Println(prompt)
	input, ok := p.PopFront()
	if ok == false {
		return "", fmt.Errorf("empty reserved prompt inputs")
	}
	return input, nil
}

func (p *ReservedPromptInput) PromptPassword(prompt string) (string, error) {
	if p.useReserved == false {
		return console.Stdin.PromptPassword(prompt)
	}
	fmt.Println(prompt)
	input, ok := p.PopFront()
	if ok == false {
		return "", fmt.Errorf("empty reserved prompt inputs")
	}
	return input, nil
}
