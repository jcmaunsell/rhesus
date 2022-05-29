package repl

import (
	"bufio"
	"fmt"
	"github.com/jcmaunsell/rhesus/lexer"
	"github.com/jcmaunsell/rhesus/token"
	"io"
	"log"
)

const PROMPT = "> "

type REPL interface {
	Start(in io.Reader, out io.Writer)
}

type repl struct {
	logger *log.Logger
}

func New(stderr io.Writer) REPL {
	return &repl{log.New(stderr, "", log.Ldate|log.Ltime|log.Llongfile)}
}

func (r *repl) Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprintf(out, PROMPT); err != nil {
			r.logger.Println(fmt.Errorf("could not print prompt: %w", err))
			return
		}
		notDone := scanner.Scan()
		if !notDone {
			if err := scanner.Err(); err != nil {
				r.logger.Println(fmt.Errorf("could not scan input: %w", err))
			}
			return
		}

		l := lexer.New(scanner.Text())

		for tok := l.NextToken(); tok != token.EndOfFile; tok = l.NextToken() {
			if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
				r.logger.Println(fmt.Errorf("could not print token: %w", err))
				return
			}
		}
	}
}
