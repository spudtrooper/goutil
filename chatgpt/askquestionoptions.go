// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package chatgpt

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type AskQuestionOption struct {
	f func(*askQuestionOptionImpl)
	s string
}

func (o AskQuestionOption) String() string { return o.s }

type AskQuestionOptions interface {
	InferResultType() bool
	HasInferResultType() bool
	Model() string
	HasModel() bool
	Verbose() bool
	HasVerbose() bool
}

func AskQuestionInferResultType(inferResultType bool) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		opts.has_inferResultType = true
		opts.inferResultType = inferResultType
	}, fmt.Sprintf("chatgpt.AskQuestionInferResultType(bool %+v)", inferResultType)}
}
func AskQuestionInferResultTypeFlag(inferResultType *bool) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		if inferResultType == nil {
			return
		}
		opts.has_inferResultType = true
		opts.inferResultType = *inferResultType
	}, fmt.Sprintf("chatgpt.AskQuestionInferResultType(bool %+v)", inferResultType)}
}

func AskQuestionModel(model string) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		opts.has_model = true
		opts.model = model
	}, fmt.Sprintf("chatgpt.AskQuestionModel(string %+v)", model)}
}
func AskQuestionModelFlag(model *string) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		if model == nil {
			return
		}
		opts.has_model = true
		opts.model = *model
	}, fmt.Sprintf("chatgpt.AskQuestionModel(string %+v)", model)}
}

func AskQuestionVerbose(verbose bool) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("chatgpt.AskQuestionVerbose(bool %+v)", verbose)}
}
func AskQuestionVerboseFlag(verbose *bool) AskQuestionOption {
	return AskQuestionOption{func(opts *askQuestionOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("chatgpt.AskQuestionVerbose(bool %+v)", verbose)}
}

type askQuestionOptionImpl struct {
	inferResultType     bool
	has_inferResultType bool
	model               string
	has_model           bool
	verbose             bool
	has_verbose         bool
}

func (a *askQuestionOptionImpl) InferResultType() bool    { return a.inferResultType }
func (a *askQuestionOptionImpl) HasInferResultType() bool { return a.has_inferResultType }
func (a *askQuestionOptionImpl) Model() string            { return or.String(a.model, "gpt-3.5-turbo") }
func (a *askQuestionOptionImpl) HasModel() bool           { return a.has_model }
func (a *askQuestionOptionImpl) Verbose() bool            { return a.verbose }
func (a *askQuestionOptionImpl) HasVerbose() bool         { return a.has_verbose }

func makeAskQuestionOptionImpl(opts ...AskQuestionOption) *askQuestionOptionImpl {
	res := &askQuestionOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAskQuestionOptions(opts ...AskQuestionOption) AskQuestionOptions {
	return makeAskQuestionOptionImpl(opts...)
}
