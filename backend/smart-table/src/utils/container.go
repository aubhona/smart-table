package utils

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

const DiContainerName = "di-container"

func GetFromContainer[T any](ctx context.Context) (T, error) {
	var zeroValue T

	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return zeroValue, errors.New("failed getting gin context")
	}

	container, ok := ginCtx.Get(DiContainerName)
	if !ok {
		return zeroValue, errors.New("failed getting di container from gin context")
	}

	diContainer, ok := container.(*dig.Container)
	if !ok {
		return zeroValue, errors.New("failed getting di container from gin context")
	}

	var dependency T

	err := diContainer.Invoke(func(dep T) {
		dependency = dep
	})

	if err != nil {
		return zeroValue, err
	}

	return dependency, nil
}
