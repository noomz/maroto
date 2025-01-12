package core

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Maroto interface {
	RegisterHeader(rows ...Row) error
	RegisterFooter(rows ...Row) error
	AddRows(rows ...Row)
	AddRow(rowHeight float64, cols ...Col) Row
	AddPages(pages ...Page)
	GetStructure() *node.Node[Structure]
	Generate() (Document, error)
}

type Document interface {
	GetBytes() []byte
	GetBase64() string
	Save(file string) error
	GetReport() *metrics.Report
	Merge([]byte) error
}

type Node interface {
	SetConfig(config *entity.Config)
	GetStructure() *node.Node[Structure]
}

type Component interface {
	Node
	Render(provider Provider, cell *entity.Cell)
}

type Col interface {
	Node
	Add(components ...Component) Col
	GetSize() int
	WithStyle(style *props.Cell) Col
	Render(provider Provider, cell entity.Cell, createCell bool)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
	WithStyle(style *props.Cell) Row
	Render(provider Provider, cell entity.Cell)
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetRows() []Row
	GetNumber() int
	SetNumber(number int, total int)
	Render(provider Provider, cell entity.Cell)
}
