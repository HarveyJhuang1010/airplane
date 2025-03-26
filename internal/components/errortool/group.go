package errortool

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	maxGroupCodeSequence = 1295
	maxErrorCodeSequence = 46655
)

type define struct {
	separated string
	codes     iCodeRepository
}

func (d *define) Codes() []customerError {
	keys := d.codes.Keys()
	sort.SliceStable(keys,
		func(i, j int) bool {
			return keys[i] < keys[j]
		})

	res := make([]customerError, len(keys))
	for i, v := range keys {
		if val, ok := d.codes.Get(v); ok {
			res[i] = *val
		} else {
			res[i] = customerError{code: v}
		}
	}

	return res
}

type group struct {
	*define
	parent   *group
	groupSeq *sequence
	codeSeq  *sequence
	code     string
}

func (g *group) Group() group {
	return g.group(fmt.Sprintf("%02s", strings.ToUpper(strconv.FormatInt(int64(g.groupSeq.Next()), 36))))
}

func (g *group) CustomGroup(code string) group {
	return g.group(code)
}

func (g *group) group(groupCode string) group {
	return group{
		parent:   g,
		define:   g.define,
		code:     groupCode,
		groupSeq: newSequence(0, maxGroupCodeSequence),
		codeSeq:  newSequence(0, maxErrorCodeSequence),
	}
}

func (g *group) Error(message string) Error {
	return g.error(fmt.Sprintf("%03s", strings.ToUpper(strconv.FormatInt(int64(g.codeSeq.Next()), 36))), message)
}

func (g *group) CustomError(code string, message string) Error {
	return g.error(code, message)
}

func (g *group) error(code string, message string) Error {
	errCode := g.makeErrorCode(g.parent, code)
	err := &customerError{
		code:    errCode,
		message: message,
	}

	g.define.codes.Add(errCode, err)
	return err
}

func (g *group) makeErrorCode(parent *group, code string) errorCode {
	if parent != nil {
		return parent.makeErrorCode(parent.parent, fmt.Sprintf("%s%s%s", g.code, g.define.separated, code))
	}

	return errorCode(code)
}
