package giteaissue

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	gitea "code.gitea.io/sdk/gitea"
)

type Widget struct {
	view.ScrollableWidget

	todos        []*gitea.Issue
	giteaClient *gitea.Client
	settings     *Settings
	err          error
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	widget.giteaClient, _ = gitea.NewClient(settings.domain, gitea.SetToken(settings.apiKey))
	
	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	todos, err := widget.getTodos()
	widget.todos = todos
	widget.err = err
	widget.SetItemCount(len(todos))

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("Gitea ToDos (%d)", len(widget.todos))

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if widget.todos == nil {
		return title, "No ToDos to display", false
	}

	str := widget.contentFrom(widget.todos)

	return title, str, false
}

func (widget *Widget) getTodos() ([]*gitea.Issue, error) {
	lo := gitea.ListOptions { PageSize:30 }
	opts := gitea.ListIssueOption{ Type: gitea.IssueTypeIssue, State: gitea.StateOpen , ListOptions: lo  } 

	todos, _, err := widget.giteaClient.ListIssues(opts)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// trim the todo body so it fits on a single line
func (widget *Widget) trimTodoBody(body string) string {
	r := []rune(body)

	// Cut at first occurence of a newline
	for i, a := range r {
		if a == '\n' {
			return string(r[:i])
		}
	}

	return body
}

func (widget *Widget) contentFrom(todos []*gitea.Issue) string {
	var str string

	for idx, todo := range todos {
		row := fmt.Sprintf(`[%s]%2d. `, widget.RowColor(idx), idx+1)
		//		if widget.settings.showProject {
		//	row = fmt.Sprintf(`%s%s `, row, todo.Project.Path)
		//}
		row = fmt.Sprintf(`%s[mediumpurple](%s)[%s] %s - %s`,
			row,
			todo.Repository.Name,
			widget.RowColor(idx),
			todo.Title,
			widget.trimTodoBody(todo.Body),
			
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(todo.Body))
	}

	return str
}

/*
func (widget *Widget) markAsDone() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.todos != nil && sel < len(widget.todos) {
		todo := widget.todos[sel]
		_, err := widget.giteaClient.Todos.MarkTodoAsDone(todo.ID)
		if err == nil {
			widget.Refresh()
		}
	}
}

func (widget *Widget) openTodo() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.todos != nil && sel < len(widget.todos) {
		todo := widget.todos[sel]
		utils.OpenFile(todo.TargetURL)
	}
}
*/
func (widget *Widget) markAsDone() {
	_ = "noop"
}
func (widget *Widget) openTodo() {
_ = "ok"
}
