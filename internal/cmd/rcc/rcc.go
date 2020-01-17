package rcc

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/therecipe/qt/internal/binding/templater"

	"github.com/therecipe/qt/internal/cmd"

	"github.com/therecipe/qt/internal/utils"
)

var (
	ResourceNames      = make(map[string]string)
	ResourceNamesMutex = new(sync.Mutex)
)

func Rcc(path, target, tagsCustom, output_dir string, useuic, quickcompiler, deploying bool, skipSetup bool) {
	if utils.UseGOMOD(path) && !skipSetup {
		if !utils.ExistsDir(filepath.Join(filepath.Dir(utils.GOMOD(path)), "vendor")) {
			cmd := exec.Command("go", "mod", "vendor")
			cmd.Dir = path
			utils.RunCmd(cmd, "go mod vendor")
		}
		if utils.QT_DOCKER() {
			cmd := exec.Command("go", "get", "-v", "-d", "github.com/therecipe/qt/internal/binding/files/docs/"+utils.QT_API(utils.QT_VERSION())+"@"+cmd.QtModVersion(filepath.Dir(utils.GOMOD(path)))) //TODO: needs to pull 5.8.0 if QT_WEBKIT
			cmd.Dir = path
			if !utils.QT_PKG_CONFIG() {
				utils.RunCmdOptional(cmd, "go get docs") //TODO: this can fail if QT_PKG_CONFIG
			}

			if strings.HasPrefix(target, "sailfish") || strings.HasPrefix(target, "android") { //TODO: generate android and sailfish minimal instead
				cmd := exec.Command(filepath.Join(utils.GOBIN(), "qtsetup"), "generate", target)
				cmd.Dir = path
				utils.RunCmd(cmd, "run setup")
			}
		}
	}

	rcc(path, target, tagsCustom, output_dir, quickcompiler, useuic, true)

	if !deploying && utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			if utils.UseGOMOD(path) {
				utils.RunCmd(exec.Command("chown", "-R", idug, filepath.Dir(utils.GOMOD(path))), "chown files to user")
			} else {
				utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
			}
		}
	}
}

func rcc(path, target, tagsCustom, output_dir string, quickcompiler bool, useuic bool, root bool) {
	utils.Log.WithField("path", path).WithField("target", target).Debug("start Rcc")

	//TODO: cache non go asset (*.qml, *.js, *.ui, ...) hashes in rcc.go files to indentify staled assets in cached go packages
	//TODO: pure go rcc file generation for wasm/js targets (needed for more convenient -fast mode)

	if root {
		wg := new(sync.WaitGroup)
		defer wg.Wait()
		allImports := cmd.GetImports(path, target, tagsCustom, 0, false)
		wg.Add(len(allImports))
		for _, path := range allImports {
			go func(path string) {
				rcc(path, target, tagsCustom, path, quickcompiler, useuic, false)
				wg.Done()
			}(path)
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Error("failed to read dir")
		return
	}
	var hasQRCFiles bool
	var hasUIFiles bool
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".qrc" {
			hasQRCFiles = true
		}

		if file.IsDir() && file.Name() == "qml" ||
			(!file.IsDir() && file.Name() == "qml" && !file.Mode().IsRegular()) {
			hasQRCFiles = true
		}

		if file.IsDir() && file.Name() == "ui" ||
			(!file.IsDir() && file.Name() == "ui" && !file.Mode().IsRegular()) {
			hasUIFiles = true
		}
	}

	if hasUIFiles {
		uic := utils.ToolPath("uic", target)
		if (target == "windows" && utils.ExistsFile(uic+".exe")) || utils.ExistsFile(uic) {

			path := filepath.Join(path, "ui")

			files, err = ioutil.ReadDir(filepath.Join(path))
			if err != nil {
				utils.Log.WithError(err).Fatal("failed to read ui dir")
			}

			var fileList []string
			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".ui" {
					//TODO: check for buildTags
					fileList = append(fileList, filepath.Join(path, file.Name()))
				}
			}

			for _, filePath := range fileList {

				utils.RunCmd(exec.Command(uic, "-o", filepath.Join(path, "uic_tmp.cpp"), filePath), fmt.Sprintf("execute uic for %v on %v", target, runtime.GOOS))
				file := strings.Replace(utils.Load(filepath.Join(path, "uic_tmp.cpp")), "\r\n", "\n", -1)
				defer utils.RemoveAll(filepath.Join(path, "uic_tmp.cpp"))

				name := strings.TrimSpace(strings.Split(strings.Split(file, "class Ui_")[1], "{")[0])
				typ := strings.TrimSpace(strings.Split(strings.Split(file, "setupUi(")[1], " ")[0])

				bb := new(bytes.Buffer)
				defer bb.Reset()

				fmt.Fprintf(bb, "type __%v struct {}\n", strings.ToLower(name))
				fmt.Fprintf(bb, "func (*__%v) init(){}\n", strings.ToLower(name))

				fmt.Fprintf(bb, "type %v struct {\n", name)
				fmt.Fprintf(bb, "*__%v\n", strings.ToLower(name))
				fmt.Fprintf(bb, "*widgets.%v\n", typ)

				namesType := make(map[string]string)
				for _, l := range strings.Split(strings.Split(strings.Split(file, "public:")[1], "void")[0], ";") {
					if !strings.Contains(l, "*") {
						continue
					}
					ls := strings.Split(strings.TrimSpace(l), " *")

					module := "widgets"
					switch ls[0] {
					case "QWebEngineView":
						module = "webengine"
					case "QCameraViewfinder":
						module = "multimedia"
					case "QQuickWidget":
						module = "quick"
					}
					fmt.Fprintf(bb, "%v%v *%v.%v\n", strings.Title(ls[1][:1]), ls[1][1:], module, ls[0])

					namesType[ls[1]] = ls[0]
				}
				fmt.Fprintln(bb, "}")

				fmt.Fprintf(bb, "func New%[1]v(p widgets.QWidget_ITF) *%[1]v {\n", name)
				fmt.Fprintln(bb, "var par *widgets.QWidget")
				fmt.Fprintln(bb, "if p != nil {")
				fmt.Fprintln(bb, "par = p.QWidget_PTR()")
				fmt.Fprintln(bb, "}")

				if useuic {
					if typ == "QGroupBox" || typ == "QMdiArea" || typ == "QScrollArea" || typ == "QStackedWidget" || typ == "QTabWidget" || typ == "QWizardPage" {
						fmt.Fprintf(bb, "w := &%[1]v{%[2]v: widgets.New%[2]v(par)}\n", name, typ)
					} else if typ == "QDockWidget" {
						fmt.Fprintf(bb, "w := &%[1]v{%[2]v: widgets.New%[2]v2(par, core.Qt__WindowType(0))}\n", name, typ)
					} else {
						fmt.Fprintf(bb, "w := &%[1]v{%[2]v: widgets.New%[2]v(par, 0)}\n", name, typ)
					}
				} else {
					fmt.Fprintf(bb, "file := core.NewQFile2(\":/%v/%v\")\n", filepath.Base(filepath.Dir(filePath)), filepath.Base(filePath))
					fmt.Fprintln(bb, "file.Open(core.QIODevice__ReadOnly)")
					fmt.Fprintf(bb, "w := &%[1]v{%[2]v: widgets.New%[2]vFromPointer(uitools.NewQUiLoader(nil).Load(file, par).Pointer())}\n", name, typ)
					fmt.Fprintln(bb, "file.Close()")
				}

				fmt.Fprintln(bb, "w.setupUI()")
				fmt.Fprintln(bb, "w.init()")
				fmt.Fprintln(bb, "return w")
				fmt.Fprintln(bb, "}")

				typeForName := func(l string) string {
					for n, t := range namesType {
						if strings.Contains(l, fmt.Sprintf("w.%v%v.", strings.Title(n[:1]), n[1:])) {
							return t
						}
					}
					return ""
				}

				rewrite := func(i int) []byte {
					bb := new(bytes.Buffer)
					defer bb.Reset()

					file = strings.Replace(file, "\\\"", "_QUOTE_", -1)

					var extractedStrings []string
					ls := strings.Split(strings.Split(strings.Split(file, "*"+name+")")[i], "} //")[0], "\"")
					for ii, l := range ls {
						if ii%2 != 1 {
							continue
						}
						extractedStrings = append(extractedStrings, l)
						ls[ii] = "_PLACEHOLDER_"
					}

					ii := 0
					for _, l := range strings.Split(strings.Join(ls, "\""), ";") {

						//

						if !(strings.Contains(l, " else ")) {
							l = strings.Replace(l, "{", "", -1)
						}

						l = strings.Replace(l, "->", ".", -1)

						l = strings.Replace(l, " \n", "\n", -1)

						if strings.Contains(l, "const bool") {
							l = strings.Replace(l, "const bool", "", -1)
							l = strings.Replace(l, " = ", " := ", -1)
						}

						if strings.Contains(l, "___") {
							l = strings.Replace(l, " = ", " := ", -1)
						}

						if strings.Contains(l, "Q_UNUSED") {
							l = "/*" + l + "*/"
						}

						if strings.Contains(l, "QString::fromUtf8(") {
							l = strings.Replace(strings.Replace(l, "QString::fromUtf8(\"", "\"", -1), "\")", "\"", -1)

							l = strings.Replace(l, " = \"", " := \"", -1)
							l = strings.Replace(l, "QString", "", -1)
						}

						//

						for pname := range namesType {
							l = strings.Replace(l, " "+pname, " w."+pname, -1)
							l = strings.Replace(l, "("+pname, "(w."+pname, -1)
						}

						ls := strings.Split(l, ".")
						for i := range ls {
							if i > 0 {
								ls[i] = strings.Title(ls[i][:1]) + ls[i][1:]
							}
						}
						l = strings.Join(ls, ".")

						if strings.Contains(l, "if (") && strings.Contains(l, ")\n") {
							l = strings.Replace(l, ")\n", "){\n", -1)
							if !strings.Contains(l, "QIcon") {
								l += "\n}\n"
							}
						}

						//

						for _, n := range []string{"QBrush", "QSizePolicy"} {
							if strings.Contains(l, fmt.Sprintf(" %v ", n)) {

								lt := strings.Split(l, "(")
								ls := strings.Split(strings.TrimSpace(lt[0]), " ")

								ls[0], ls[1] = ls[1], ls[0]
								ls[1] = "= new " + ls[1]

								lt[0] = strings.Join(ls, " ")
								l = strings.Join(lt, "(")
							}
						}

						l = strings.Replace(l, "const QIcon", "", -1)

						for _, n := range []string{"QPalette", "QFont", "QIcon"} {
							if strings.Contains(l, fmt.Sprintf(" %v ", n)) {
								ls := strings.Split(strings.TrimSpace(l), " ")
								ls[0], ls[1] = ls[1], ls[0]
								ls[1] = "= new " + ls[1] + "()"
								l = strings.Join(ls, " ")
							}
						}

						//

						for i, n := range []string{"QWebEngineView", "QCameraViewfinder", "QQuickWidget", "QTableWidgetItem", "QTreeWidgetItem", "QListWidgetItem", "QSizePolicy", "QPalette", "QFont", "QIcon", "QBrush"} {
							if !strings.Contains(l, n) {
								continue
							}

							module := "widgets"
							switch {
							case i == 0:
								module = "webengine"
							case i == 1:
								module = "multimedia"
							case i == 2:
								module = "quick"
							case i >= 7:
								module = "gui"
							}

							if i <= 2 {
								l = strings.Replace(l, fmt.Sprintf("= new %v(", n), fmt.Sprintf("= %v.New%v(", module, n), -1)
							} else {
								if n == "QIcon" {
									l = strings.Replace(l, fmt.Sprintf("= %v(", n), fmt.Sprintf(":= %v.New%v5(", module, n), -1)
								}
								l = strings.Replace(l, fmt.Sprintf("= new %v(", n), fmt.Sprintf(":= %v.New%v(", module, n), -1)
							}
						}
						l = strings.Replace(l, "new ", "widgets.New", -1)

						//

						for _, n := range []string{"Qt", "QLocale"} {
							l = strings.Replace(l, n+"::", "core."+n+"__", -1)
						}

						l = strings.Replace(l, "QIcon::hasThemeIcon", "gui.QIcon_HasThemeIcon", -1)
						l = strings.Replace(l, "QIcon::fromTheme", "gui.QIcon_FromTheme", -1)
						if utils.QT_API_NUM(utils.QT_VERSION()) <= 5063 && strings.Contains(l, "gui.QIcon_FromTheme") {
							l = strings.TrimSuffix(l, ")") + ", gui.NewQIcon())"
						}

						l = strings.Replace(l, "QTabWidget::North", "widgets.QTabWidget__North", -1)
						l = strings.Replace(l, "QTabWidget::Rounded", "widgets.QTabWidget__Rounded", -1)

						for _, n := range []string{"QPalette", "QFont", "QIcon"} {
							l = strings.Replace(l, n+"::", "gui."+n+"__", -1)
						}

						for _, n := range []string{"QDockWidget", "QLCDNumber", "QLayout", "QTextEdit", "QPlainTextEdit", "QSlider", "QFrame", "QFormLayout", "QDialogButtonBox", "QSizePolicy", "QLineEdit", "QListView", "QAbstractItemView"} {
							l = strings.Replace(l, n+"::", "widgets."+n+"__", -1)
						}

						for _, n := range []string{"QQuickWidget"} {
							l = strings.Replace(l, n+"::", "quick."+n+"__", -1)
						}

						//

						l = strings.Replace(l, "("+name+")", "(w)", -1)
						l = strings.Replace(l, "("+name+".", "(w.", -1)
						l = strings.Replace(l, " "+name+".", "w.", -1)

						l = strings.Replace(l, "(QRect(", "(core.NewQRect4(", -1)

						if strings.Contains(l, "(QUrl(") {
							l = strings.Replace(l, "(QUrl(", "(core.NewQUrl3(", -1)
							ls := strings.Split(l, ")")
							ls[0] += ", 0"
							l = strings.Join(ls, ")")
						}

						l = strings.Replace(l, ".Resize(", ".Resize2(", -1)

						//

						for _, n := range []string{"QWidget", "QLabel", "QSizePolicy", "QListWidgetItem", "QFrame", "QToolBox", "QOpenGLWidget"} {
							if strings.Contains(l, ".New"+n+"(") {
								l = strings.TrimSuffix(l, ")") + ", 0)"
							}
						}
						l = strings.Replace(l, "(, 0)", "()", -1)

						for i, n := range []string{"QTreeWidgetItem", "QTableWidgetItem", "QListWidgetItem"} {
							if strings.Contains(l, ".New"+n+"(") {
								if strings.Contains(l, ".New"+n+"(w.") {
									if i == 2 {
										l = strings.Replace(l, ".New"+n+"(", ".New"+n+"(", -1)
									} else {
										l = strings.Replace(l, ".New"+n+"(", ".New"+n+"3(", -1)
										l = strings.TrimSuffix(l, ")") + ", 0)"
									}
								} else if strings.Contains(l, ".New"+n+"()") {
									l = strings.Replace(l, ".New"+n+"(", ".New"+n+"(0", -1)
								} else {
									l = strings.Replace(l, ".New"+n+"(", ".New"+n+"6(", -1)
									l = strings.TrimSuffix(l, ")") + ", 0)"
								}
							}
						}

						for _, n := range []string{"QToolBar", "QSizePolicy"} {
							if strings.Contains(l, ".New"+n+"(") {
								l = strings.Replace(l, ".New"+n+"(", ".New"+n+"2(", -1)
							}
						}

						for _, n := range []string{"QVBoxLayout", "QHBoxLayout"} {
							if strings.Contains(l, ".New"+n+"(") && !strings.Contains(l, ".New"+n+"()") {
								l = strings.Replace(l, ".New"+n+"(", ".New"+n+"2(", -1)
							}
						}

						for _, n := range []string{"QGridLayout"} {
							if strings.Contains(l, ".New"+n+"(") && strings.Contains(l, ".New"+n+"()") {
								l = strings.Replace(l, ".New"+n+"(", ".New"+n+"2(", -1)
							}
						}

						for _, n := range []string{"QBrush"} {
							if strings.Contains(l, ".New"+n+"(") {
								l = strings.Replace(l, ".New"+n+"(QColor(", ".New"+n+"3(gui.NewQColor3(", -1)
								l = strings.TrimSuffix(l, ")") + ", 1)"
							}
						}

						for _, n := range []string{"QDockWidget"} {
							if strings.Contains(l, ".New"+n+"(") {
								l = strings.Replace(l, ".New"+n+"(", ".New"+n+"2(", -1)
								l = strings.TrimSuffix(l, ")") + ", 0)"
							}
						}

						for _, n := range []string{"QFormLayout", "QWizardPage"} {
							if strings.Contains(l, ".New"+n+"(") && strings.Contains(l, ".New"+n+"()") {
								l = strings.Replace(l, ".New"+n+"(", ".New"+n+"(nil", -1)
							}
						}

						//

						for _, n := range []string{"QTableWidgetItem", "QTreeWidgetItem", "QListWidgetItem"} {
							if strings.Contains(l, n) {
								l = strings.Replace(l, n+" *__", "__", -1)
							}
						}

						//

						if strings.Contains(l, ".AddWidget(") {
							if strings.Contains(l, ",") {
								l = strings.Replace(l, ".AddWidget(", ".AddWidget3(", -1)
								l = strings.TrimSuffix(l, ")") + ", 0)"
							} else {
								if strings.Contains(typeForName(l), "Layout") {
									l = strings.Replace(l, ".AddWidget(", ".QLayout.AddWidget(", -1)
								}
							}
						}

						if strings.Contains(l, ".AddLayout(") {
							if strings.Contains(l, ",") {
								l = strings.Replace(l, ".AddLayout(", ".AddLayout2(", -1)
								l = strings.TrimSuffix(l, ")") + ", 0)"
							} else if strings.Contains(typeForName(l), "Layout") {
								l = strings.TrimSuffix(l, ")") + ", 0)"
							}
						}

						if strings.Contains(l, ".AddItem(") {
							if strings.Contains(l, "QString()") {
								l = strings.TrimSuffix(l, ")") + ", core.NewQVariant1(0))"
							} else if !strings.Contains(l, "(w.") || strings.Count(l, ",") >= 4 {
								l = strings.TrimSuffix(l, ")") + ", 0)"
							}

							if strings.Contains(l, "(icon") || strings.Contains(typeForName(l), "QToolBox") {
								l = strings.Replace(l, ".AddItem(", ".AddItem2(", -1)
							}
						}

						if strings.Contains(l, ".AddButton(") {
							l = strings.TrimSuffix(l, ")") + ", -1)"
						}

						if strings.Contains(l, ".SetTextAlignment(") {
							l = strings.Replace(l, ".SetTextAlignment(", ".SetTextAlignment(int(", -1)
							l = strings.TrimSuffix(l, ")") + "))"
						}

						if strings.Contains(l, ".SetAlignment(") {
							if strings.Contains(typeForName(l), "QGroupBox") {
								l = strings.Replace(l, ".SetAlignment(", ".SetAlignment(int(", -1)
								l = strings.TrimSuffix(l, ")") + "))"
							}
						}

						if strings.Contains(l, ".AddAction(") {
							l = strings.Replace(l, ".AddAction(", ".AddActions([]*widgets.QAction{", -1)
							l = strings.TrimSuffix(l, ")") + "})"
						}

						if strings.Contains(l, ".SetShortcut(") {
							l = strings.Replace(l, ".SetShortcut(", ".SetShortcut(gui.NewQKeySequence2(", -1)
							l = strings.TrimSuffix(l, "))") + "), 0))"
						}

						//

						if strings.Contains(l, "QSize(") {
							l = strings.Replace(l, "QSize(", "core.NewQSize(", -1)
							if !strings.Contains(l, "QSize()") {
								l = strings.Replace(l, "QSize(", "QSize2(", -1)
							}
						}

						if strings.Contains(l, "QLocale(") {
							l = strings.Replace(l, "QLocale(", "core.NewQLocale3(", -1)
						}

						if strings.Contains(l, "QCursor(") {
							l = strings.Replace(l, "QCursor(", "gui.NewQCursor2(", -1)
						}

						if strings.Contains(l, "QVariant(") {
							l = strings.Replace(l, "QVariant(", "core.NewQVariant1(", -1)
						}

						if strings.Contains(l, "QDateTime(QDate(") {
							l = strings.Replace(l, "QDateTime(QDate(", "core.NewQDateTime3(core.NewQDate3(", -1)
							l = strings.Replace(l, "QTime(", "core.NewQTime3(", -1)
							l = strings.TrimSuffix(l, ")))") + ", 0), 0))"
						}

						l = strings.Replace(l, "QWidget()", "QWidget(nil, 0)", -1)

						l = strings.Replace(l, "SetBrush(", "SetBrush2(", -1)

						//

						if strings.Contains(l, "retranslateUi(w)") {
							l = strings.Replace(l, "retranslateUi(w)", "w.retranslateUi()", -1)
						}

						l = strings.Replace(l, "QMetaObject::connectSlotsByName(", "core.QMetaObject_ConnectSlotsByName(", -1)
						l = strings.Replace(l, "QWidget::setTabOrder(", "widgets.QWidget_SetTabOrder(", -1)
						l = strings.Replace(l, "QCoreApplication::translate(", "core.QCoreApplication_Translate(", -1)
						l = strings.Replace(l, "QApplication::translate(", "widgets.QApplication_Translate(", -1)

						//

						ls = strings.Split(l, "\"")
						for j := range ls {
							if j%2 != 1 {
								continue
							}
							ls[j] = extractedStrings[ii]
							ii++
						}
						l = strings.Join(ls, "\"")
						l = strings.Replace(l, "_QUOTE_", "\\\"", -1)

						l = strings.Replace(l, "\"\n\"", "\"+\"", -1)

						l = strings.Replace(l, "QString()", "\"\"", -1)
						l = strings.Replace(l, ", nullptr)", ", \"\", 0)", -1)

						l = strings.Replace(l, ".IsEmpty()", "== \"\"", -1)

						if strings.Contains(l, "QPixmap(") {
							if utils.QT_API_NUM(utils.QT_VERSION()) >= 5130 {
								l = strings.Replace(l, "QPixmap(", "gui.NewQPixmap3(", -1)
							} else {
								l = strings.Replace(l, "QPixmap(", "gui.NewQPixmap5(", -1)
							}
							ls := strings.Split(l, ")")
							ls[0] += ", \"\", 0"
							l = strings.Join(ls, ")")
						}

						//

						l = strings.Replace(l, "#if", "//#if", -1)
						l = strings.Replace(l, "#endif", "//#endif", -1)

						if strings.Contains(l, "#endif") {
							if strings.Contains(l, "#ifndef Q_OS_MAC") {
								l = "}\nif runtime.GOOS != \"darwin\" {" + l
							} else if strings.Contains(l, "#if QT_CONFIG") {
								l = "}\nif true {" + l
							} else {
								l = "}\n" + l
							}
						} else {
							if strings.Contains(l, "#ifndef Q_OS_MAC") {
								l = "if runtime.GOOS != \"darwin\" {\n" + l
							}
							if strings.Contains(l, "#if QT_CONFIG") {
								l = "if true {" + l
							}
						}

						//

						if strings.Contains(l, "QObject::connect(") {
							ls := strings.Split(l, ",")

							src := strings.Split(ls[0], "(")[1]
							sig := strings.Title(strings.Split(ls[1], "(")[1])
							dst := strings.Replace(strings.TrimSpace(ls[2]), "W.", "w.", -1)
							slo := strings.Replace(strings.Title(strings.Split(ls[3], "(")[1]), "W.", "w.", -1)

							dstSlo := fmt.Sprintf("%v.%v", dst, slo)
							dstSlo = strings.Replace(dstSlo, ".w.", ".", -1)
							if !strings.Contains(dst, "w.") {
								dstSlo = slo
							}
							if !strings.Contains(dstSlo, "w.") {
								dstSlo = "w." + dstSlo
							}

							if strings.Count(dstSlo, ".") == 1 {
								for n := range namesType {
									if strings.HasSuffix(dstSlo, "."+strings.Title(n[:1])+n[1:]) {
										dstSlo += "." + sig
										break
									}
								}
							}

							if sig == "Clicked" && !strings.HasSuffix(dstSlo, sig) {
								var b string
								if strings.HasSuffix(dstSlo, ".SetEnabled") || strings.HasSuffix(dstSlo, ".SetVisible") {
									b = "b"
								}
								l = fmt.Sprintf("%v.Connect%v(func(b bool){%v(%v)})\n", src, sig, dstSlo, b)
							} else if strings.TrimPrefix(slo, "w.") == "AnimateClick" && !strings.HasSuffix(sig, strings.TrimPrefix(slo, "w.")) {
								l = fmt.Sprintf("%v.Connect%v(func(){%v(0)})\n", src, sig, dstSlo)
							} else {
								l = fmt.Sprintf("%v.Connect%v(%v)\n", src, sig, dstSlo)
							}
						}

						ls = strings.Split(l, "\n")
						for i := len(ls) - 1; i >= 0; i-- {
							if strings.HasPrefix(strings.TrimSpace(ls[i]), "//#") ||
								strings.TrimSpace(ls[i]) == "" {
								ls = append(ls[:i], ls[i+1:]...)
							}
						}
						l = strings.Join(ls, "\n")

						fmt.Fprintln(bb, strings.TrimSpace(l))
					}

					return bb.Bytes()
				}

				if useuic {
					fmt.Fprintf(bb, "func (w *%v) setupUI() {\n", name)
					bb.Write(rewrite(1))
					fmt.Fprintln(bb, "}")

					fmt.Fprintf(bb, "func (w *%v) retranslateUi() {\n", name)
					bb.Write(rewrite(2))
					fmt.Fprintln(bb, "}")
				} else {
					fmt.Fprintf(bb, "func (w *%v) setupUI() {\n", name)
					for n, t := range namesType {
						module := "widgets"
						switch t {
						case "QWebEngineView":
							module = "webengine"
						case "QCameraViewfinder":
							module = "multimedia"
						}
						fmt.Fprintf(bb, "w.%v = %v.New%vFromPointer(w.FindChild(\"%v\", core.Qt__FindChildrenRecursively).Pointer())\n", strings.Title(n[:1])+n[1:], module, t, n)
					}
					fmt.Fprintln(bb, "}")
				}

				by := bb.Bytes()
				bb = new(bytes.Buffer)
				defer bb.Reset()

				main := os.Getenv("QT_UIC_MAIN") == "true"

				if main {
					fmt.Fprintln(bb, "package main")
				} else {
					fmt.Fprintln(bb, "package ui")
				}
				fmt.Fprintln(bb, "import (")

				for i, p := range []string{"runtime", "core", "gui", "widgets", "webengine", "multimedia", "uitools", "quick"} {
					if !bytes.Contains(by, []byte(p+".")) {
						continue
					}
					if i == 0 {
						fmt.Fprintln(bb, "\""+p+"\"")
					} else {
						fmt.Fprintln(bb, "\"github.com/therecipe/qt/"+p+"\"")
					}
				}
				fmt.Fprintln(bb, "\n)")

				if main {
					fmt.Fprintln(bb, "func main() { widgets.NewQApplication(0, nil); New"+name+"(nil).Show();widgets.QApplication_Exec(); }")
				}

				bb.Write(by)

				out, err := format.Source(bb.Bytes())
				if err != nil {
					utils.Log.WithError(err).Errorln("failed to format ui:", "uic_"+strings.TrimSuffix(filepath.Base(filePath), ".ui")+".go")
					out = bb.Bytes()
				}
				utils.SaveBytes(filepath.Join(path, "uic_"+strings.TrimSuffix(filepath.Base(filePath), ".ui")+".go"), out)

				if !useuic {
					if dir := filepath.Dir(path); utils.ExistsDir(dir) {

						bb := new(bytes.Buffer)
						defer bb.Reset()

						fmt.Fprintln(bb, "<!DOCTYPE RCC><RCC version=\"1.0\">")
						fmt.Fprintln(bb, "<qresource>")

						files, err = ioutil.ReadDir(filepath.Join(dir, "ui"))
						if err != nil {
							utils.Log.WithError(err).Fatal("failed to read dir")
						}

						for _, file := range files {
							if !file.IsDir() && filepath.Ext(file.Name()) == ".ui" {
								fmt.Fprintf(bb, "<file>ui/%v</file>\n", file.Name())
							}
						}

						fmt.Fprintln(bb, "</qresource>")
						fmt.Fprintln(bb, "</RCC>")

						utils.SaveBytes(filepath.Join(dir, "uic.qrc"), bb.Bytes())

						hasQRCFiles = true
					}
				}
			}
		}
	}

	if !hasQRCFiles {
		return
	}

	rccQrc := filepath.Join(path, "rcc.qrc")

	env, tags, _, _ := cmd.BuildEnv(target, "", "")
	if tagsCustom != "" {
		tags = append(tags, strings.Split(tagsCustom, " ")...)
	}

	pkgCmd := utils.GoList("{{.Name}}", "-find", utils.BuildTags(tags))
	if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
		pkgCmd.Dir = path
	} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
		pkgCmd.Dir = filepath.Dir(utils.GOMOD(path))
		vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
		pkgCmd.Args = append(pkgCmd.Args, vl[len(vl)-1])
	}

	for k, v := range env {
		pkgCmd.Env = append(pkgCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	pkg := strings.TrimSpace(utils.RunCmd(pkgCmd, "run go list"))
	if pkg == "" {
		pkg = filepath.Base(path)
	}

	libs := []string{"Core"}
	if quickcompiler {
		libs = append(libs, "Qml")
	}
	rccCpp := filepath.Join(path, "rcc.cpp")
	if output_dir != "" {
		rccCpp = filepath.Join(output_dir, "rcc.cpp")
		templater.CgoTemplateSafe(pkg, output_dir, target, templater.RCC, "", "", libs)
	} else {
		templater.CgoTemplateSafe(pkg, path, target, templater.RCC, "", "", libs)
	}

	if dir := filepath.Join(path, "qml"); utils.ExistsDir(dir) {
		rcc := exec.Command(utils.ToolPath("rcc", target), "-project", "-o", rccQrc)
		rcc.Dir = dir
		utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.qrc on %v for %v", runtime.GOOS, target))

		content := utils.Load(rccQrc)
		content = strings.Replace(content, "<file>./", "<file>qml/", -1)
		if utils.ExistsFile(filepath.Join(path, "qtquickcontrols2.conf")) {
			content = strings.Replace(content, "<qresource>", "<qresource>\n<file>qtquickcontrols2.conf</file>", -1)
		}

		//TODO: filter out duplicate assets

		utils.Save(rccQrc, content)
	}

	files, err = ioutil.ReadDir(path)
	if err != nil {
		utils.Log.WithError(err).Fatal("failed to read dir")
	}
	var fileList []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".qrc" && !strings.HasSuffix(file.Name(), "_qml_cache.qrc") {
			//TODO: check for buildTags
			fileList = append(fileList, filepath.Join(path, file.Name()))
		}
	}

	nameCmd := utils.GoList("{{.ImportPath}}", "-find", utils.BuildTags(tags))
	if !utils.UseGOMOD(path) || (utils.UseGOMOD(path) && !strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/")) {
		nameCmd.Dir = path
	} else if utils.UseGOMOD(path) && strings.Contains(strings.Replace(path, "\\", "/", -1), "/vendor/") {
		nameCmd.Dir = filepath.Dir(utils.GOMOD(path))
		vl := strings.Split(strings.Replace(path, "\\", "/", -1), "/vendor/")
		nameCmd.Args = append(nameCmd.Args, vl[len(vl)-1])
	}

	for k, v := range env {
		nameCmd.Env = append(nameCmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	name := "rcc_" + strings.TrimSpace(utils.RunCmd(nameCmd, "run go list"))
	for _, s := range []string{"/", ".", "-"} {
		name = strings.Replace(name, s, "_", -1)
	}

	cachgen := utils.ToolPath("qmlcachegen", target)
	if ((target == "windows" && utils.ExistsFile(cachgen+".exe")) || utils.ExistsFile(cachgen)) && quickcompiler {
		utils.RemoveAll(rccCpp)

		var filteredFiles []string
		var possibleMixedFiles []string
		var pureFiles []string

		for _, f := range fileList {
			newName := filepath.Join(filepath.Dir(f), name+"_"+strings.TrimSuffix(strings.Replace(filepath.Base(f), "-", "_", -1), ".qrc")+"_qml_cache.qrc")

			utils.RunCmd(exec.Command(cachgen, "--filter-resource-file", "-o", newName, f), fmt.Sprintf("execute qmlcachegen filter on %v for %v", runtime.GOOS, target))

			for _, tBC := range strings.Split(strings.TrimSpace(utils.RunCmd(exec.Command(utils.ToolPath("rcc", target), "-list", f), "execute rcc")), "\n") {
				tBC = strings.TrimSpace(tBC)
				if strings.HasSuffix(tBC, ".qml") || strings.HasSuffix(tBC, ".js") {
					tBCT := strings.Replace(strings.Replace(tBC, ".", "_", -1), "-", "_", -1)

					cmd := exec.Command(cachgen)

					/* TODO: re-enable to warn about duplicate assets?
					for _, fl := range fileList {
						cmd.Args = append(cmd.Args, []string{"--resource", fl}...)
					}
					*/
					cmd.Args = append(cmd.Args, []string{"--resource", f}...)

					of := filepath.Join(filepath.Dir(f), "rcc_"+strings.TrimSuffix(strings.Replace(filepath.Base(f), "-", "_", -1), ".qrc")+"_"+filepath.Base(strings.Replace(filepath.Dir(tBCT), "-", "_", -1))+"_"+filepath.Base(tBCT+"_qml_cache.cpp"))
					ResourceNamesMutex.Lock()
					ResourceNames[of] = ""
					ResourceNamesMutex.Unlock()
					cmd.Args = append(cmd.Args, []string{"-o", of, tBC}...)
					utils.RunCmd(cmd, fmt.Sprintf("execute qmlcachegen cache on %v for %v", runtime.GOOS, target))
				}
			}

			tmpName := filepath.Join(filepath.Dir(f), name+"_"+strings.Replace(filepath.Base(f), "-", "_", -1))

			if utils.ExistsFile(newName) {
				filteredFiles = append(filteredFiles, newName)

				utils.Save(tmpName, utils.Load(f))
				possibleMixedFiles = append(possibleMixedFiles, tmpName)

				defer func() { utils.RemoveAll(newName) }()
			} else {

				utils.Save(tmpName, utils.Load(f))
				pureFiles = append(pureFiles, tmpName)
			}

			defer func() { utils.RemoveAll(tmpName) }()
		}

		if len(possibleMixedFiles) > 0 || len(pureFiles) > 0 {
			cmd := exec.Command(utils.ToolPath("qmlcachegen", target))
			cmd.Dir = path

			for i, _ := range possibleMixedFiles {
				cmd.Args = append(cmd.Args, fmt.Sprintf("--resource-file-mapping=%v=%v", possibleMixedFiles[i], filteredFiles[i]))
			}
			for _, f := range pureFiles {
				cmd.Args = append(cmd.Args, fmt.Sprintf("--resource-file-mapping=%v", f))
			}

			cmd.Args = append(cmd.Args, []string{"-o", "rcc_qmlcache_loader.cpp"}...)
			cmd.Args = append(cmd.Args, possibleMixedFiles...)
			cmd.Args = append(cmd.Args, pureFiles...)

			var initNameList []string
			for _, f := range append(possibleMixedFiles, pureFiles...) {
				initNameList = append(initNameList, strings.TrimSuffix(strings.Replace(filepath.Base(f), "-", "_", -1), ".qrc"))
			}
			ResourceNamesMutex.Lock()
			ResourceNames[filepath.Join(path, "rcc_qmlcache_loader.cpp")] = strings.Join(initNameList, "|")
			ResourceNamesMutex.Unlock()

			utils.RunCmd(cmd, fmt.Sprintf("execute qmlcachegen loader on %v for %v", runtime.GOOS, target))
		}

		for _, f := range filteredFiles {
			rcc := exec.Command(utils.ToolPath("rcc", target), "-name", strings.TrimSuffix(strings.Replace(filepath.Base(f), "-", "_", -1), ".qrc"), "-o", strings.Replace(filepath.Base(strings.TrimSuffix(f, ".qrc")), name, "rcc", -1)+".cpp")
			rcc.Dir = path
			rcc.Args = append(rcc.Args, f)
			utils.RunCmd(rcc, fmt.Sprintf("execute per file rcc *.cpp on %v for %v", runtime.GOOS, target))

			ResourceNamesMutex.Lock()
			ResourceNames[strings.Replace(strings.TrimSuffix(f, ".qrc"), name, "rcc", -1)+".cpp"] = strings.TrimSuffix(strings.Replace(filepath.Base(f), "-", "_", -1), ".qrc")
			ResourceNamesMutex.Unlock()
		}

	} else {
		ResourceNamesMutex.Lock()
		ResourceNames[rccCpp] = name
		ResourceNamesMutex.Unlock()

		rcc := exec.Command(utils.ToolPath("rcc", target), "-name", name, "-o", rccCpp)
		rcc.Args = append(rcc.Args, fileList...)
		utils.RunCmd(rcc, fmt.Sprintf("execute rcc *.cpp on %v for %v", runtime.GOOS, target))
	}

	if utils.QT_DEBUG_QML() {
		utils.Save("debug.pro", fmt.Sprintf("RESOURCES += %v", strings.Join(fileList, " ")))
	}
}
