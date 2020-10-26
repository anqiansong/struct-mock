package mock

import (
	"reflect"
	"runtime"
	"strings"
	"struct-mock/render"
	"struct-mock/util"
)

type (
	InvokeFn      func(*render.MockContext) interface{}
	InvokeMatcher struct {
		m map[string]InvokeFn
	}
)

var (
	emptyInvokeFn = func(_ *render.MockContext) interface{} {
		return nil
	}
)

func NewInvokeMatcher() *InvokeMatcher {
	return &InvokeMatcher{
		m: make(map[string]InvokeFn),
	}
}

func (i *InvokeMatcher) Add(list ...InvokeFn) {
	for _, fn := range list {
		name := getFuncName(fn)
		if len(name) == 0 {
			continue
		}
		name = util.UnTitle(name)
		i.m[name] = fn
	}
}

func (i *InvokeMatcher) Match(name string) (InvokeFn, bool) {
	if name == "-" {
		return emptyInvokeFn, true
	}
	fn, ok := i.m[name]
	return fn, ok
}

func getFuncName(in interface{}) string {
	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Func {
		return ""
	}
	v := reflect.ValueOf(in)
	name := runtime.FuncForPC(v.Pointer()).Name()
	index := strings.LastIndex(name, ".")
	if index < 0 {
		return name
	}
	return name[index+1:]
}
