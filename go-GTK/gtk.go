package main

import (
	//"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	//"os"
	"os/exec"
	//"path"
	"regexp"
	"sort"
	"strings"
	//"fmt"
)

func uniq(strings []string) (ret []string) {
	return
}

func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Yasuhiro Matumoto <mattn.jp@gmail.com>"}
}

func main() {
	//var menuitem *gtk.MenuItem
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//GtkEntry
	entry := gtk.NewEntry()
	entry.SetText("Hello world")

	//GtkButton
	button := gtk.NewButtonWithLabel("Button with label")
	button.Clicked(func() {
		messagedialog := gtk.NewMessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			entry.GetText())
		messagedialog.Response(func() {
			println("Dialog OK!")
			filechooserdialog := gtk.NewFileChooserDialog(
				"Choose File...",
				button.GetTopLevelAsWindow(),
				gtk.FILE_CHOOSER_ACTION_OPEN,
				gtk.STOCK_OK,
				gtk.RESPONSE_ACCEPT)
			filter := gtk.NewFileFilter()
			filter.AddPattern("*.go")
			filechooserdialog.Response(func() {
				println(filechooserdialog.GetFilename())
				filechooserdialog.Destroy()
			})
			filechooserdialog.Run()
			messagedialog.Destroy()
		})
		messagedialog.Run()
	})
	window.Add(button)

	window.ShowAll()
	gtk.Main()
}
