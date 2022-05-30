package repl

import (
	"bufio"
	"fmt"
	"github.com/jcmaunsell/rhesus/lexer"
	"github.com/jcmaunsell/rhesus/logger"
	"github.com/jcmaunsell/rhesus/parser"
	"github.com/jcmaunsell/rhesus/token"
	"github.com/sirupsen/logrus"
	"io"
)

const PROMPT = "> "

type REPL interface {
	Start(in io.Reader, out io.Writer)
}

type repl struct {
	log *logrus.Logger
}

func New() (REPL, error) {
	log := logger.Service()
	return &repl{log}, nil
}

func (r *repl) Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprintf(out, PROMPT); err != nil {
			r.log.Println(fmt.Errorf("could not print prompt: %w", err))
			return
		}
		notDone := scanner.Scan()
		if !notDone {
			if err := scanner.Err(); err != nil {
				r.log.WithError(err).Println(fmt.Errorf("could not scan input"))
			}
			return
		}

		input := scanner.Text()
		lex := lexer.New(input)

		for tok := lex.NextToken(); tok != token.EndOfFile; tok = lex.NextToken() {
			if tok.Type == token.ILLEGAL {
				r.log.WithField("character", tok.Literal).Error("User input contains illegal character.")
			}
			if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
				r.log.WithError(err).WithField("token", tok).Error("Could not print token.")
				return
			}
		}

		program, err := parser.New(lex).Parse()
		if err != nil {
			r.log.WithError(err).WithField("input", input).Error("Could not parse program.")
			return
		}
		if _, err := fmt.Fprintf(out, "%+v\n", program); err != nil {
			r.log.WithError(err).WithField("input", input).Error("Could not print parsed program.")
			return
		}
	}
}
