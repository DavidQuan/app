package app

import (
	"testing"

	"github.com/murlokswarm/log"
)

type Hello struct {
	Greeting  string
	BadMarkup bool
}

func (h *Hello) OnInputChange(a OnChangeArg) {
	h.Greeting = a.Value
	Render(h)
}

func (h *Hello) Render() string {
	return `
<div>
    Hello, 
    <span>
        {{if .Greeting}}
            {{html .Greeting}}
        {{else}}
            World
        {{end}}
    </span>
    <input onchange="@OnInputChange" />

	{{if .BadMarkup}}<div></span>{{end}}
</div>
    `
}

func TestRun(t *testing.T) {
	OnFinalize = func() {
		log.Info("OnFinalize called")
	}

	Run()
	Run()
	Finalize()
	Finalize()
}

func TestRender(t *testing.T) {
	hello := &Hello{}

	ctx := NewZeroContext("rendering")
	defer ctx.Close()

	ctx.Mount(hello)
	hello.Greeting = "Maxence"
	Render(hello)
}

func TestRenderPanicCompoCtxError(t *testing.T) {
	defer func() { recover() }()

	hello := &Hello{}
	Render(hello)
	t.Error("should panic")
}

func TestRenderPanicCompoBadMarkup(t *testing.T) {
	defer func() { recover() }()

	hello := &Hello{}

	ctx := NewZeroContext("rendering")
	defer ctx.Close()

	ctx.Mount(hello)
	hello.BadMarkup = true
	Render(hello)
	t.Error("should panic")
}

func TestMenu(t *testing.T) {
	t.Log(Menu())
}

func TestDock(t *testing.T) {
	t.Log(Dock())
}