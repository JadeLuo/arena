package training

import (
	"fmt"
	"time"

	"github.com/kubeflow/arena/pkg/apis/types"
	"github.com/kubeflow/arena/pkg/argsbuilder"
)

type ScaleInETJobBuilder struct {
	args      *types.ScaleInETJobArgs
	argValues map[string]interface{}
	argsbuilder.ArgsBuilder
}

func NewScaleInETJobBuilder() *ScaleInETJobBuilder {
	args := &types.ScaleInETJobArgs{
		ScaleETJobArgs: types.ScaleETJobArgs{
			Timeout: 60,
			Retry:   0,
			Count:   1,
		},
	}
	return &ScaleInETJobBuilder{
		args:        args,
		argValues:   map[string]interface{}{},
		ArgsBuilder: argsbuilder.NewScaleInETJobArgsBuilder(args),
	}
}

// Name is used to set job name,match option --name
func (b *ScaleInETJobBuilder) Name(name string) *ScaleInETJobBuilder {
	if name != "" {
		b.args.Name = name
	}
	return b
}

// Retry is used to set retry times
func (b *ScaleInETJobBuilder) Retry(count int) *ScaleInETJobBuilder {
	if count > 0 {
		b.args.Retry = count
	}
	return b
}

// Retry is used to set retry times
func (b *ScaleInETJobBuilder) Count(count int) *ScaleInETJobBuilder {
	if count > 0 {
		b.args.Count = count
	}
	return b
}

// Timeout is used to set timeout seconds
func (b *ScaleInETJobBuilder) Timeout(timeout time.Duration) *ScaleInETJobBuilder {
	b.argValues["timeout"] = &timeout
	return b
}

// Script is used to set scale script
func (b *ScaleInETJobBuilder) Script(s string) *ScaleInETJobBuilder {
	if s != "" {
		b.args.Script = s
	}
	return b
}

// Envs is used to set envs
func (b *ScaleInETJobBuilder) Envs(envs map[string]string) *ScaleInETJobBuilder {
	items := []string{}
	for key, value := range envs {
		items = append(items, fmt.Sprintf("%v=%v", key, value))
	}
	b.argValues["env"] = items
	return b
}

// Build is used to build the job
func (b *ScaleInETJobBuilder) Build() (*Job, error) {
	for key, value := range b.argValues {
		b.AddArgValue(key, value)
	}
	if err := b.PreBuild(); err != nil {
		return nil, err
	}
	if err := b.ArgsBuilder.Build(); err != nil {
		return nil, err
	}
	return NewJob(b.args.Name, types.ETTrainingJob, b.args), nil
}
