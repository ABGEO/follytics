package loader

import (
	"fmt"
	"strings"

	"ariga.io/atlas-provider-gorm/gormschema"

	"github.com/abgeo/follytics/db/sql"
	"github.com/abgeo/follytics/internal/model"
)

type Loader interface {
	Load() (string, error)
}

type SchemaLoader struct {
	models []interface{}
}

var _ Loader = (*SchemaLoader)(nil)

func NewSchemaLoader() *SchemaLoader {
	return &SchemaLoader{
		models: []interface{}{
			&model.Event{},
			&model.JobState{},
			&model.User{},
		},
	}
}

func (l *SchemaLoader) Load() (string, error) {
	sb := &strings.Builder{}

	sb.WriteString(sql.GetSchema())

	if err := l.loadModels(sb); err != nil {
		return "", fmt.Errorf("failed to load models: %w", err)
	}

	return sb.String(), nil
}

func (l *SchemaLoader) loadModels(sb *strings.Builder) error {
	statements, err := gormschema.New("postgres").Load(l.models...)
	if err != nil {
		return fmt.Errorf("failed to load gorm schema: %w", err)
	}

	sb.WriteString(statements)

	return nil
}
