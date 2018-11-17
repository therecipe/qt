package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/printsupport"
	"github.com/therecipe/qt/widgets"
)

var rsrcPath string

func init() {
	if runtime.GOOS == "darwin" {
		rsrcPath = ":/qml/images/mac"
	} else {
		rsrcPath = ":/qml/images/win"
	}
}

type TextEdit struct {
	widgets.QMainWindow

	actionSave          *widgets.QAction
	actionTextBold      *widgets.QAction
	actionTextUnderline *widgets.QAction
	actionTextItalic    *widgets.QAction
	actionTextColor     *widgets.QAction
	actionAlignLeft     *widgets.QAction
	actionAlignCenter   *widgets.QAction
	actionAlignRight    *widgets.QAction
	actionAlignJustify  *widgets.QAction
	actionUndo          *widgets.QAction
	actionRedo          *widgets.QAction

	actionCut   *widgets.QAction
	actionCopy  *widgets.QAction
	actionPaste *widgets.QAction

	comboStyle *widgets.QComboBox
	comboFont  *widgets.QFontComboBox
	comboSize  *widgets.QComboBox

	tb       *widgets.QToolBar
	fileName string
	textEdit *widgets.QTextEdit
}

func initTextEdit() *TextEdit {
	var this = NewTextEdit(nil, 0)

	if runtime.GOOS == "darwin" {
		this.SetUnifiedTitleAndToolBarOnMac(true)
	}
	this.SetWindowTitle(core.QCoreApplication_ApplicationName())

	this.textEdit = widgets.NewQTextEdit(this)
	var textEdit = this.textEdit
	textEdit.ConnectCurrentCharFormatChanged(this.currentCharFormatChanged)
	textEdit.ConnectCursorPositionChanged(this.cursorPositionChanged)
	this.SetCentralWidget(textEdit)

	this.SetToolButtonStyle(core.Qt__ToolButtonFollowStyle)
	this.setupFileActions()
	this.setupEditActions()
	this.setupTextActions()

	var (
		helpMenu = this.MenuBar().AddMenu2("Help")
		action1  = helpMenu.AddAction("About")
		action2  = helpMenu.AddAction("About Qt")
	)
	action1.ConnectTriggered(func(checked bool) { this.about() })
	action2.ConnectTriggered(func(checked bool) { qApp.AboutQt() })

	var textFont = gui.NewQFont2("Helvetica", -1, -1, false)
	textFont.SetStyleHint(gui.QFont__SansSerif, gui.QFont__PreferDefault)
	textEdit.SetFont(textFont)
	this.fontChanged(textEdit.Font())
	this.colorChanged(textEdit.TextColor())
	this.alignmentChanged(textEdit.Alignment())

	textEdit.Document().ConnectModificationChanged(func(changed bool) {
		this.actionSave.SetEnabled(changed)
		this.SetWindowModified(changed)
	})

	textEdit.Document().ConnectUndoAvailable(this.actionUndo.SetEnabled)
	textEdit.Document().ConnectRedoAvailable(this.actionRedo.SetEnabled)

	this.SetWindowModified(textEdit.Document().IsModified())
	this.actionSave.SetEnabled(textEdit.Document().IsModified())
	this.actionUndo.SetEnabled(textEdit.Document().IsUndoAvailable())
	this.actionRedo.SetEnabled(textEdit.Document().IsRedoAvailable())

	if len(os.Getenv("QT_NO_CLIPBOARD")) != 0 {
		this.actionCut.SetEnabled(false)
		this.actionCopy.SetEnabled(false)
		qApp.Clipboard().ConnectDataChanged(this.clipboardDataChanged)
	}

	textEdit.SetFocus2()
	this.setCurrentFileName("")

	return this
}

func (t *TextEdit) currentCharFormatChanged(format *gui.QTextCharFormat) {
	t.fontChanged(format.Font())
	t.colorChanged(format.Foreground().Color())
}

func (t *TextEdit) cursorPositionChanged() {
	t.alignmentChanged(t.textEdit.Alignment())
}

func (t *TextEdit) setupFileActions() {
	var tb = t.AddToolBar3("File Actions")
	var menu = t.MenuBar().AddMenu2("&File")

	var newIcon = gui.QIcon_FromTheme2("document-new", gui.NewQIcon5(rsrcPath+"/filenew.png"))
	var a = menu.AddAction2(newIcon, "&New")
	a.ConnectTriggered(func(checked bool) { t.fileNew() })
	tb.QWidget.AddAction(a)
	a.SetPriority(widgets.QAction__LowPriority)
	a.SetShortcuts2(gui.QKeySequence__New)

	var openIcon = gui.QIcon_FromTheme2("document-open", gui.NewQIcon5(rsrcPath+"/fileopen.png"))
	a = menu.AddAction2(openIcon, "&Open...")
	a.ConnectTriggered(func(checked bool) { t.fileOpen() })
	a.SetShortcuts2(gui.QKeySequence__Open)
	tb.QWidget.AddAction(a)

	menu.AddSeparator()

	var saveIcon = gui.QIcon_FromTheme2("document-save", gui.NewQIcon5(rsrcPath+"/filesave.png"))
	t.actionSave = menu.AddAction2(saveIcon, "&Save")
	t.actionSave.ConnectTriggered(func(checked bool) { t.fileSave() })
	t.actionSave.SetShortcuts2(gui.QKeySequence__Save)
	t.actionSave.SetEnabled(false)
	tb.QWidget.AddAction(t.actionSave)

	a = menu.AddAction("Save &As...")
	a.ConnectTriggered(func(checked bool) { t.fileSaveAs() })
	a.SetPriority(widgets.QAction__LowPriority)
	menu.AddSeparator()

	if len(os.Getenv("QT_NO_PRINTER")) == 0 {
		var printIcon = gui.QIcon_FromTheme2("document-print", gui.NewQIcon5(rsrcPath+"/fileprint.png"))
		a = menu.AddAction2(printIcon, "&Print...")
		a.ConnectTriggered(func(checked bool) { t.filePrint() })
		a.SetPriority(widgets.QAction__LowPriority)
		a.SetShortcuts2(gui.QKeySequence__Print)
		tb.QWidget.AddAction(a)

		var filePrintIcon = gui.QIcon_FromTheme2("fileprint", gui.NewQIcon5(rsrcPath+"/fileprint.png"))
		a = menu.AddAction2(filePrintIcon, "Print Preview...")
		a.ConnectTriggered(func(checked bool) { t.filePrintPreview() })

		var exportPdfIcon = gui.QIcon_FromTheme2("exportpdf", gui.NewQIcon5(rsrcPath+"/exportpdf.png"))
		a = menu.AddAction2(exportPdfIcon, "&Export PDF...")
		a.ConnectTriggered(func(checked bool) { t.filePrintPdf() })
		a.SetPriority(widgets.QAction__LowPriority)
		a.SetShortcut(gui.NewQKeySequence2("Ctrl+D", gui.QKeySequence__NativeText))
		tb.QWidget.AddAction(a)

		menu.AddSeparator()
	}

	a = menu.AddAction("&Quit")
	a.ConnectTriggered(func(checked bool) { t.Close() })
	a.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
}

func (t *TextEdit) setupEditActions() {
	var tb = t.AddToolBar3("Edit Actions")
	var menu = t.MenuBar().AddMenu2("&Edit")

	var undoIcon = gui.QIcon_FromTheme2("edit-undo", gui.NewQIcon5(rsrcPath+"/editundo.png"))
	t.actionUndo = menu.AddAction2(undoIcon, "&Undo")
	t.actionUndo.ConnectTriggered(func(checked bool) { t.textEdit.Undo() })
	t.actionUndo.SetShortcuts2(gui.QKeySequence__Undo)
	tb.QWidget.AddAction(t.actionUndo)

	var redoIcon = gui.QIcon_FromTheme2("edit-redo", gui.NewQIcon5(rsrcPath+"/editredo.png"))
	t.actionRedo = menu.AddAction2(redoIcon, "&Redo")
	t.actionRedo.ConnectTriggered(func(checked bool) { t.textEdit.Redo() })
	t.actionRedo.SetPriority(widgets.QAction__LowPriority)
	t.actionRedo.SetShortcuts2(gui.QKeySequence__Redo)
	tb.QWidget.AddAction(t.actionRedo)
	menu.AddSeparator()

	if len(os.Getenv("QT_NO_CLIPBOARD")) == 0 {
		var cutIcon = gui.QIcon_FromTheme2("edit-cut", gui.NewQIcon5(rsrcPath+"/editcut.png"))
		t.actionCut = menu.AddAction2(cutIcon, "Cu&t")
		t.actionCut.ConnectTriggered(func(checked bool) { t.textEdit.Cut() })
		t.actionCut.SetPriority(widgets.QAction__LowPriority)
		t.actionCut.SetShortcuts2(gui.QKeySequence__Cut)
		tb.QWidget.AddAction(t.actionCut)

		var copyIcon = gui.QIcon_FromTheme2("edit-copy", gui.NewQIcon5(rsrcPath+"/editcopy.png"))
		t.actionCopy = menu.AddAction2(copyIcon, "&Copy")
		t.actionCopy.ConnectTriggered(func(checked bool) { t.textEdit.Copy() })
		t.actionCopy.SetPriority(widgets.QAction__LowPriority)
		t.actionCopy.SetShortcuts2(gui.QKeySequence__Copy)
		tb.QWidget.AddAction(t.actionCopy)

		var pasteIcon = gui.QIcon_FromTheme2("edit-paste", gui.NewQIcon5(rsrcPath+"/editpaste.png"))
		t.actionPaste = menu.AddAction2(pasteIcon, "&Paste")
		t.actionPaste.ConnectTriggered(func(checked bool) { t.textEdit.Paste() })
		t.actionPaste.SetPriority(widgets.QAction__LowPriority)
		t.actionPaste.SetShortcuts2(gui.QKeySequence__Paste)
		tb.QWidget.AddAction(t.actionPaste)

		var md = qApp.Clipboard().MimeData(gui.QClipboard__Clipboard)
		if md.Pointer() != nil {
			t.actionPaste.SetEnabled(md.HasText())
		}
	}
}

func (t *TextEdit) setupTextActions() {
	var tb = t.AddToolBar3("Format Actions")
	var menu = t.MenuBar().AddMenu2("F&ormat")

	var boldIcon = gui.QIcon_FromTheme2("format-text-bold", gui.NewQIcon5(rsrcPath+"/textbold.png"))
	t.actionTextBold = menu.AddAction2(boldIcon, "&Bold")
	t.actionTextBold.ConnectTriggered(func(checked bool) { t.textBold() })
	t.actionTextBold.SetShortcut(gui.NewQKeySequence2("Ctrl+B", gui.QKeySequence__NativeText))
	t.actionTextBold.SetPriority(widgets.QAction__LowPriority)
	var bold = gui.NewQFont()
	bold.SetBold(true)
	t.actionTextBold.SetFont(bold)
	tb.QWidget.AddAction(t.actionTextBold)
	t.actionTextBold.SetCheckable(true)

	var italicIcon = gui.QIcon_FromTheme2("format-text-italic", gui.NewQIcon5(rsrcPath+"/textitalic.png"))
	t.actionTextItalic = menu.AddAction2(italicIcon, "&Italic")
	t.actionTextItalic.ConnectTriggered(func(checked bool) { t.textItalic() })
	t.actionTextItalic.SetShortcut(gui.NewQKeySequence2("Ctrl+I", gui.QKeySequence__NativeText))
	t.actionTextItalic.SetPriority(widgets.QAction__LowPriority)
	var italic = gui.NewQFont()
	italic.SetItalic(true)
	t.actionTextItalic.SetFont(italic)
	tb.QWidget.AddAction(t.actionTextItalic)
	t.actionTextItalic.SetCheckable(true)

	var underlineIcon = gui.QIcon_FromTheme2("format-text-underline", gui.NewQIcon5(rsrcPath+"/textunder.png"))
	t.actionTextUnderline = menu.AddAction2(underlineIcon, "&Underline")
	t.actionTextUnderline.ConnectTriggered(func(checked bool) { t.textUnderline() })
	t.actionTextUnderline.SetShortcut(gui.NewQKeySequence2("Ctrl+U", gui.QKeySequence__NativeText))
	t.actionTextUnderline.SetPriority(widgets.QAction__LowPriority)
	var underline = gui.NewQFont()
	underline.SetUnderline(true)
	t.actionTextUnderline.SetFont(underline)
	tb.QWidget.AddAction(t.actionTextUnderline)
	t.actionTextUnderline.SetCheckable(true)

	menu.AddSeparator()

	var leftIcon = gui.QIcon_FromTheme2("format-justify-left", gui.NewQIcon5(rsrcPath+"/textleft.png"))
	t.actionAlignLeft = menu.AddAction2(leftIcon, "&Left")
	t.actionAlignLeft.SetShortcut(gui.NewQKeySequence2("Ctrl+L", gui.QKeySequence__NativeText))
	t.actionAlignLeft.SetCheckable(true)
	t.actionAlignLeft.SetPriority(widgets.QAction__LowPriority)

	var centerIcon = gui.QIcon_FromTheme2("format-justify-center", gui.NewQIcon5(rsrcPath+"/textcenter.png"))
	t.actionAlignCenter = menu.AddAction2(centerIcon, "C&enter")
	t.actionAlignCenter.SetShortcut(gui.NewQKeySequence2("Ctrl+E", gui.QKeySequence__NativeText))
	t.actionAlignCenter.SetCheckable(true)
	t.actionAlignCenter.SetPriority(widgets.QAction__LowPriority)

	var rightIcon = gui.QIcon_FromTheme2("format-justify-right", gui.NewQIcon5(rsrcPath+"/textright.png"))
	t.actionAlignRight = menu.AddAction2(rightIcon, "&Right")
	t.actionAlignRight.SetShortcut(gui.NewQKeySequence2("Ctrl+R", gui.QKeySequence__NativeText))
	t.actionAlignRight.SetCheckable(true)
	t.actionAlignRight.SetPriority(widgets.QAction__LowPriority)

	var fillIcon = gui.QIcon_FromTheme2("format-justify-fill", gui.NewQIcon5(rsrcPath+"/textjustify.png"))
	t.actionAlignJustify = menu.AddAction2(fillIcon, "&Justify")
	t.actionAlignJustify.SetShortcut(gui.NewQKeySequence2("Ctrl+J", gui.QKeySequence__NativeText))
	t.actionAlignJustify.SetCheckable(true)
	t.actionAlignJustify.SetPriority(widgets.QAction__LowPriority)

	// Make sure the alignLeft  is always left of the alignRight
	var alignGroup = widgets.NewQActionGroup(t)
	alignGroup.ConnectTriggered(t.textAlign)

	if qApp.IsLeftToRight() {
		alignGroup.AddAction(t.actionAlignLeft)
		alignGroup.AddAction(t.actionAlignCenter)
		alignGroup.AddAction(t.actionAlignRight)
	} else {
		alignGroup.AddAction(t.actionAlignRight)
		alignGroup.AddAction(t.actionAlignCenter)
		alignGroup.AddAction(t.actionAlignLeft)
	}
	alignGroup.AddAction(t.actionAlignJustify)

	tb.AddActions(alignGroup.Actions())
	menu.AddActions(alignGroup.Actions())

	menu.AddSeparator()

	var pix = gui.NewQPixmap3(16, 16)
	pix.Fill(gui.NewQColor2(core.Qt__black))
	t.actionTextColor = menu.AddAction2(gui.NewQIcon2(pix), "&Color...")
	t.actionTextColor.ConnectTriggered(func(checked bool) { t.textColor() })
	tb.QWidget.AddAction(t.actionTextColor)

	tb = t.AddToolBar3("Format Actions")
	tb.SetAllowedAreas(core.Qt__TopToolBarArea | core.Qt__BottomToolBarArea)
	t.AddToolBarBreak(core.Qt__TopToolBarArea)
	t.AddToolBar2(tb)

	t.comboStyle = widgets.NewQComboBox(tb)
	tb.AddWidget(t.comboStyle)
	t.comboStyle.AddItems([]string{
		"Standard",
		"Bullet List (Disc)",
		"Bullet List (Circle)",
		"Bullet List (Square)",
		"Ordered List (Decimal)",
		"Ordered List (Alpha lower)",
		"Ordered List (Alpha upper)",
		"Ordered List (Roman lower)",
		"Ordered List (Roman upper)",
	})
	t.comboStyle.ConnectActivated(t.textStyle)

	t.comboFont = widgets.NewQFontComboBox(tb)
	tb.AddWidget(t.comboFont)
	t.comboFont.ConnectActivated2(t.textFamily)

	t.comboSize = widgets.NewQComboBox(tb)
	t.comboSize.SetObjectName("comboSize")
	tb.AddWidget(t.comboSize)
	t.comboSize.SetEditable(true)

	var standardSizes = gui.QFontDatabase_StandardSizes()
	for _, size := range standardSizes {
		t.comboSize.AddItem(strconv.Itoa(size), core.NewQVariant())
	}
	t.comboSize.ConnectActivated2(t.textSize)
	t.comboSize.SetCurrentIndex(func() int {
		for index, size := range standardSizes {
			if size == widgets.QApplication_Font().PointSize() {
				return index
			}
		}
		return standardSizes[len(standardSizes)-1]
	}())
}

func (t *TextEdit) fontChanged(f *gui.QFont) {
	t.comboFont.SetCurrentIndex(t.comboFont.FindText(gui.NewQFontInfo(f).Family(), core.Qt__MatchExactly|core.Qt__MatchCaseSensitive))
	t.comboSize.SetCurrentIndex(t.comboSize.FindText(strconv.Itoa(f.PointSize()), core.Qt__MatchExactly|core.Qt__MatchCaseSensitive))
	t.actionTextBold.SetChecked(f.Bold())
	t.actionTextItalic.SetChecked(f.Italic())
	t.actionTextUnderline.SetChecked(f.Underline())
}

func (t *TextEdit) colorChanged(c *gui.QColor) {
	var pix = gui.NewQPixmap3(16, 16)
	pix.Fill(c)
	t.actionTextColor.SetIcon(gui.NewQIcon2(pix))
}

func (t *TextEdit) alignmentChanged(a core.Qt__AlignmentFlag) {
	switch a {
	case core.Qt__AlignLeft:
		{
			t.actionAlignLeft.SetChecked(true)
		}

	case core.Qt__AlignHCenter:
		{
			t.actionAlignCenter.SetChecked(true)
		}

	case core.Qt__AlignRight:
		{
			t.actionAlignRight.SetChecked(true)
		}

	case core.Qt__AlignJustify:
		{
			t.actionAlignJustify.SetChecked(true)
		}
	}
}

func (t *TextEdit) mergeFormatOnWordOrSelection(format *gui.QTextCharFormat) {
	var cursor = t.textEdit.TextCursor()
	if !cursor.HasSelection() {
		cursor.Select(gui.QTextCursor__WordUnderCursor)
	}
	cursor.MergeCharFormat(format)
	t.textEdit.MergeCurrentCharFormat(format)
}

func (t *TextEdit) about() {
	widgets.QMessageBox_About(t, "About", `This example demonstrates Qt's
rich text editing facilities in action, providing an example
document for you to experiment with.`)
}

func (t *TextEdit) clipboardDataChanged() {
	if len(os.Getenv("QT_NO_CLIPBOARD")) == 0 {
		var md = qApp.Clipboard().MimeData(gui.QClipboard__Clipboard)
		t.actionPaste.SetEnabled(md.HasText())
	}
}

func (t *TextEdit) load(f string) bool {
	if !(core.QFile_Exists(f)) {
		return false
	}
	var file = core.NewQFile2(f)
	if !file.Open(core.QIODevice__ReadOnly) {
		return false
	}
	defer file.Close()

	var (
		data  = file.ReadAll()
		codec = core.QTextCodec_CodecForHtml2(data)
		str   = codec.ToUnicode(data)
	)

	if core.Qt_MightBeRichText(str) {
		t.textEdit.SetHtml(str)
	} else {
		t.textEdit.SetPlainText(str)
	}

	t.setCurrentFileName(f)
	return true
}

func (t *TextEdit) maybeSave() bool {
	if !t.textEdit.Document().IsModified() {
		return true
	}

	var ret = widgets.QMessageBox_Warning(t, core.QCoreApplication_ApplicationName(), "The document has been modified.\nDo you want to save your changes?", widgets.QMessageBox__Save|widgets.QMessageBox__Discard|widgets.QMessageBox__Cancel, widgets.QMessageBox__NoButton)
	if ret == widgets.QMessageBox__Save {
		return t.fileSave()
	} else if ret == widgets.QMessageBox__Close {
		return false
	}

	return true
}

func (t *TextEdit) setCurrentFileName(fn string) {
	t.fileName = fn
	t.textEdit.Document().SetModified(false)

	var showName string
	if t.fileName == "" {
		showName = "untitled.txt"
	} else {
		showName = core.NewQFileInfo3(t.fileName).FileName()
	}

	t.SetWindowTitle(fmt.Sprintf("%v[*] - %v", showName, core.QCoreApplication_ApplicationName()))
	t.SetWindowModified(false)
}

func (t *TextEdit) fileNew() {
	if t.maybeSave() {
		t.textEdit.Clear()
		t.setCurrentFileName("")
	}
}

func (t *TextEdit) fileOpen() {
	var fileDialog = widgets.NewQFileDialog2(t, "Open File...", "", "")
	fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptOpen)
	fileDialog.SetFileMode(widgets.QFileDialog__ExistingFile)
	var mimeTypes = []string{"text/html", "text/plain"}
	fileDialog.SetMimeTypeFilters(mimeTypes)
	if fileDialog.Exec() != int(widgets.QDialog__Accepted) {
		return
	}
	var fn = fileDialog.SelectedFiles()[0]
	if t.load(fn) {
		t.StatusBar().ShowMessage(fmt.Sprintf("Opened %v", core.QDir_ToNativeSeparators(fn)), 0)
	} else {
		t.StatusBar().ShowMessage(fmt.Sprintf("Could not open %v", core.QDir_ToNativeSeparators(fn)), 0)
	}
}

func (t *TextEdit) fileSave() bool {
	if t.fileName == "" {
		return t.fileSaveAs()
	}
	if strings.HasPrefix(t.fileName, ":/") {
		return t.fileSaveAs()
	}

	var (
		writer  = gui.NewQTextDocumentWriter3(t.fileName, core.NewQByteArray())
		success = writer.Write(t.textEdit.Document())
	)

	if success {
		t.textEdit.Document().SetModified(false)
		t.StatusBar().ShowMessage(fmt.Sprintf("Wrote %v", core.QDir_ToNativeSeparators(t.fileName)), 0)
	} else {
		t.StatusBar().ShowMessage(fmt.Sprintf("Could not write to file %v", core.QDir_ToNativeSeparators(t.fileName)), 0)
	}

	return success
}

func (t *TextEdit) fileSaveAs() bool {
	var fileDialog = widgets.NewQFileDialog2(t, "Save as...", "", "")
	fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptSave)
	var mimeTypes = []string{"application/vnd.oasis.opendocument.text", "text/html", "text/plain"}
	fileDialog.SetMimeTypeFilters(mimeTypes)
	fileDialog.SetDefaultSuffix("odt")
	if fileDialog.Exec() != int(widgets.QDialog__Accepted) {
		return false
	}
	var fn = fileDialog.SelectedFiles()[0]
	t.setCurrentFileName(fn)
	return t.fileSave()
}

func (t *TextEdit) filePrint() {
	if len(os.Getenv("QT_NO_PRINTER")) == 0 && len(os.Getenv("QT_NO_PRINTDIALOG")) == 0 {
		var (
			printer = printsupport.NewQPrinter(printsupport.QPrinter__HighResolution)
			dlg     = printsupport.NewQPrintDialog(printer, t)
		)
		if t.textEdit.TextCursor().HasSelection() {
			dlg.SetOption(printsupport.QAbstractPrintDialog__PrintSelection, true)
		}
		dlg.SetWindowTitle("Print Document")
		if dlg.Exec() == int(widgets.QDialog__Accepted) {
			t.textEdit.Print(printer)
		}
		printer.DestroyQPrinter()
	}
}

func (t *TextEdit) filePrintPreview() {
	if len(os.Getenv("QT_NO_PRINTER")) == 0 && len(os.Getenv("QT_NO_PRINTDIALOG")) == 0 {
		var (
			printer = printsupport.NewQPrinter(printsupport.QPrinter__HighResolution)
			preview = printsupport.NewQPrintPreviewDialog(printer, t, 0)
		)
		preview.ConnectPaintRequested(t.printPreview)
		preview.Exec()
	}
}

func (t *TextEdit) printPreview(printer *printsupport.QPrinter) {
	if len(os.Getenv("QT_NO_PRINTER")) == 0 {
		t.textEdit.Print(printer)
	}
}

func (t *TextEdit) filePrintPdf() {
	if len(os.Getenv("QT_NO_PRINTER")) == 0 {
		var fileDialog = widgets.NewQFileDialog2(t, "Export PDF", "", "")
		fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptSave)
		fileDialog.SetMimeTypeFilters([]string{"application/pdf"})
		fileDialog.SetDefaultSuffix("pdf")
		if fileDialog.Exec() != int(widgets.QDialog__Accepted) {
			return
		}

		var (
			fileName = fileDialog.SelectedFiles()[0]
			printer  = printsupport.NewQPrinter(printsupport.QPrinter__HighResolution)
		)
		printer.SetOutputFormat(printsupport.QPrinter__PdfFormat)
		printer.SetOutputFileName(fileName)
		t.textEdit.Document().Print(printer)
		t.StatusBar().ShowMessage(fmt.Sprintf("Exported %v", core.QDir_ToNativeSeparators(fileName)), 0)
	}
}

func (t *TextEdit) textBold() {
	var fmt = gui.NewQTextCharFormat()
	var fw = gui.QFont__Normal
	if t.actionTextBold.IsChecked() {
		fw = gui.QFont__Bold
	}
	fmt.SetFontWeight(int(fw))
	t.mergeFormatOnWordOrSelection(fmt)
}

func (t *TextEdit) textUnderline() {
	var fmt = gui.NewQTextCharFormat()
	fmt.SetFontUnderline(t.actionTextUnderline.IsChecked())
	t.mergeFormatOnWordOrSelection(fmt)
}

func (t *TextEdit) textItalic() {
	var fmt = gui.NewQTextCharFormat()
	fmt.SetFontItalic(t.actionTextItalic.IsChecked())
	t.mergeFormatOnWordOrSelection(fmt)
}

func (t *TextEdit) textFamily(f string) {
	var fmt = gui.NewQTextCharFormat()
	fmt.SetFontFamily(f)
	t.mergeFormatOnWordOrSelection(fmt)
}

func (t *TextEdit) textSize(p string) {
	var pointSize, err = strconv.Atoi(p)
	if err != nil {
		return
	}

	if pointSize > 0 {
		var fmt = gui.NewQTextCharFormat()
		fmt.SetFontPointSize(float64(pointSize))
		t.mergeFormatOnWordOrSelection(fmt)
	}
}

func (t *TextEdit) textStyle(styleIndex int) {
	var cursor = t.textEdit.TextCursor()

	if styleIndex != 0 {

		var style = gui.QTextListFormat__ListDisc

		switch styleIndex {
		case 1:
			{
				style = gui.QTextListFormat__ListDisc
			}

		case 2:
			{
				style = gui.QTextListFormat__ListCircle
			}

		case 3:
			{
				style = gui.QTextListFormat__ListSquare
			}

		case 4:
			{
				style = gui.QTextListFormat__ListDecimal
			}

		case 5:
			{
				style = gui.QTextListFormat__ListLowerAlpha
			}

		case 6:
			{
				style = gui.QTextListFormat__ListUpperAlpha
			}

		case 7:
			{
				style = gui.QTextListFormat__ListLowerRoman
			}

		case 8:
			{
				style = gui.QTextListFormat__ListUpperRoman
			}
		}

		cursor.BeginEditBlock()

		var (
			blockFmt = cursor.BlockFormat()
			listFmt  = gui.NewQTextListFormat()
		)

		if cursor.CurrentList().Pointer() != nil {
			listFmt = gui.NewQTextListFormatFromPointer(cursor.CurrentList().Format().Pointer())
		} else {
			listFmt.SetIndent(blockFmt.Indent() + 1)
			blockFmt.SetIndent(0)
			cursor.SetBlockFormat(blockFmt)
		}

		listFmt.SetStyle(style)
		cursor.CreateList(listFmt)

		cursor.EndEditBlock()

	} else {
		var bfmt = gui.NewQTextBlockFormat()
		bfmt.SetObjectIndex(-1)
		cursor.MergeBlockFormat(bfmt)
	}
}

func (t *TextEdit) textColor() {
	var col = widgets.QColorDialog_GetColor(t.textEdit.TextColor(), t, "", 0)
	if !col.IsValid() {
		return
	}
	var fmt = gui.NewQTextCharFormat()
	fmt.SetForeground(gui.NewQBrush3(col, core.Qt__SolidPattern))
	t.mergeFormatOnWordOrSelection(fmt)
	t.colorChanged(col)
}

func (t *TextEdit) textAlign(a *widgets.QAction) {
	switch a.Pointer() {
	case t.actionAlignLeft.Pointer():
		{
			t.textEdit.SetAlignment(core.Qt__AlignLeft | core.Qt__AlignAbsolute)
		}

	case t.actionAlignCenter.Pointer():
		{
			t.textEdit.SetAlignment(core.Qt__AlignHCenter)
		}

	case t.actionAlignRight.Pointer():
		{
			t.textEdit.SetAlignment(core.Qt__AlignRight | core.Qt__AlignAbsolute)
		}

	case t.actionAlignJustify.Pointer():
		{
			t.textEdit.SetAlignment(core.Qt__AlignJustify)
		}
	}
}
