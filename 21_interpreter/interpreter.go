// Package interpreter 解释器模式
// 采用原课程的示例, 并且做了一下简化
// 假设我们现在有一个监控系统
// 现在需要实现一个告警模块，可以根据输入的告警规则来决定是否触发告警
// 告警规则支持 &&、>、< 3种运算符
// 其中 >、< 优先级比  && 更高
package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// AlertRule 告警规则
type AlertRule struct {
	expression IExpression
}

// NewAlertRule NewAlertRule
func NewAlertRule(rule string) (*AlertRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlertRule{expression: exp}, err
}

// Interpret 判断告警是否触发
func (r AlertRule) Interpret(stats map[string]float64) bool {
	return r.expression.Interpret(stats)
}

// IExpression 表达式接口
type IExpression interface {
	Interpret(stats map[string]float64) bool
}

// GreaterExpression >
type GreaterExpression struct {
	key   string
	value float64
}

// Interpret Interpret
func (g GreaterExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.value
}

// NewGreaterExpression NewGreaterExpression
func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

// LessExpression <
type LessExpression struct {
	key   string
	value float64
}

// Interpret Interpret
func (g LessExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v < g.value
}

// NewLessExpression NewLessExpression
func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("exp is invalid: %s", exp)
	}

	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}

// AndExpression &&
type AndExpression struct {
	expressions []IExpression
}

// Interpret Interpret
func (e AndExpression) Interpret(stats map[string]float64) bool {
	for _, expression := range e.expressions {
		if !expression.Interpret(stats) {
			return false
		}
	}
	return true
}

// NewAndExpression NewAndExpression
func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))

	for i, e := range exps {
		var expression IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("exp is invalid: %s", exp)
		}

		if err != nil {
			return nil, err
		}

		expressions[i] = expression
	}

	return &AndExpression{expressions: expressions}, nil
}
