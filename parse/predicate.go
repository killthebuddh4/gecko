package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gecko/types"
)

func (p *Parser) predicate(parent *types.Expression) (*types.Expression, error) {
	left, err := p.equality(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isLogical) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: LOGICAL :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.equality(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"left":  types.Block{left},
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		left = &exp
	}

	return left, nil
}

func (p *Parser) equality(parent *types.Expression) (*types.Expression, error) {
	left, err := p.comparison(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isEquality) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: EQUALITY :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.comparison(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"left":  types.Block{left},
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		left = &exp
	}

	return left, nil
}

func (p *Parser) comparison(parent *types.Expression) (*types.Expression, error) {
	left, err := p.term(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isComparison) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: COMPARISON :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.term(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"left":  types.Block{left},
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		left = &exp
	}

	return left, nil
}

func (p *Parser) term(parent *types.Expression) (*types.Expression, error) {
	left, err := p.factor(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isTerm) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: TERM :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.factor(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"left":  types.Block{left},
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		left = &exp
	}

	return left, nil
}

func (p *Parser) factor(parent *types.Expression) (*types.Expression, error) {
	left, err := p.unary(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isFactor) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: FACTOR :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"left":  types.Block{left},
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		left = &exp
	}

	return left, nil
}

func (p *Parser) unary(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isUnary) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: UNARY :: ", p.previous().Text)
		}

		operator := types.Operator{
			Type:  p.previous().Text,
			Value: p.previous().Text,
		}

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		exp := types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments: map[string][]*types.Expression{
				"right": types.Block{right},
			},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}

		return &exp, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isAtom) {
		_, debug := os.LookupEnv("GECKO_DEBUG_PARSE")

		if debug {
			fmt.Println(":: PARSE :: ATOM :: ", p.previous().Text)
		}

		var operator types.Operator
		if isConstant(p.previous()) {
			operator = types.Operator{
				Type:  "string",
				Value: p.previous().Text,
			}
		} else if isNumber(p.previous()) {
			operator = types.Operator{
				Type:  "number",
				Value: p.previous().Text,
			}
		} else if isString(p.previous()) {
			operator = types.Operator{
				Type:  "string",
				Value: p.previous().Text,
			}
		} else if isTrue(p.previous()) {
			operator = types.Operator{
				Type:  "true",
				Value: "true",
			}
		} else if isFalse(p.previous()) {
			operator = types.Operator{
				Type:  "false",
				Value: "false",
			}
		} else if isNil(p.previous()) {
			operator = types.Operator{
				Type:  "nil",
				Value: "nil",
			}
		} else if isIdentifier(p.previous()) {
			operator = types.Operator{
				Type:  "identifier",
				Value: p.previous().Text,
			}
		} else {
			return nil, errors.New("expected atom but got <" + p.previous().Text + ">")
		}

		return &types.Expression{
			Parent:   nil,
			Operator: operator,
			Signature: &types.Signature{
				Parameters: make(map[string]string),
				Returns:    []*types.Expression{},
			},
			Arguments:    map[string][]*types.Expression{},
			Catches:      []types.Block{},
			Trajectories: []*types.Trajectory{},
		}, nil
	}

	return nil, errors.New(":: atom :: expected expression but got <" + p.read().Text + ">")
}
