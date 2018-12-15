package query

import (
	"github.com/ericr/llvmq/config"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/robertkrimen/otto"
	"time"
	"log"
)

type Query struct {
	SourcePath string
	logger     *log.Logger
	config     *config.Config
	module     *ir.Module
	vm         *otto.Otto
}

func New(path string, conf *config.Config) (*Query, error) {
	conf.Logger.Printf("Parsing %s", path)
	
	parseStart := time.Now()

	m, err := asm.ParseFile(path)
	if err != nil {
		return nil, err
	}
	conf.Logger.Printf("Parsed in %fs", time.Now().Sub(parseStart).Seconds())

	q := &Query{
		SourcePath: path,
		logger:     conf.Logger,
		config:     conf,
		module:     m,
		vm:         otto.New(),
	}
	q.buildRootObj()

	return q, nil
}

func (q *Query) Exec(query string) (otto.Value, error) {
	q.logger.Print("Executing query\n")

	execStart := time.Now()
	val, err := q.vm.Run(query)

	q.logger.Printf("Query executed in %fs", time.Now().Sub(execStart).Seconds())

	return val, err
}

func (q *Query) buildRootObj() {
	q.vm.Set("ll", q.module)
}