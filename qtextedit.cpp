#include "qtextedit.h"
#include <QTextEdit>
#include "cgoexport.h"

//MyQTextEdit
class MyQTextEdit: public QTextEdit {
Q_OBJECT
public:
void Signal_CopyAvailable() { callbackQTextEdit(this, QString("copyAvailable").toUtf8().data()); };
void Signal_CursorPositionChanged() { callbackQTextEdit(this, QString("cursorPositionChanged").toUtf8().data()); };
void Signal_RedoAvailable() { callbackQTextEdit(this, QString("redoAvailable").toUtf8().data()); };
void Signal_SelectionChanged() { callbackQTextEdit(this, QString("selectionChanged").toUtf8().data()); };
void Signal_TextChanged() { callbackQTextEdit(this, QString("textChanged").toUtf8().data()); };
void Signal_UndoAvailable() { callbackQTextEdit(this, QString("undoAvailable").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_Copy();
void Slot_Cut();
void Slot_Paste();
void Slot_Redo();
void Slot_SelectAll();
void Slot_SetFontItalic(bool italic);
void Slot_SetFontUnderline(bool underline);
void Slot_SetFontWeight(int weight);
void Slot_Undo();
void Slot_ZoomIn(int rang);
void Slot_ZoomOut(int rang);

};
#include "qtextedit.moc"


//Public Functions
QtObjectPtr QTextEdit_New_QWidget(QtObjectPtr parent)
{
	return new QTextEdit(((QWidget*)(parent)));
}

QtObjectPtr QTextEdit_New_String_QWidget(char* text, QtObjectPtr parent)
{
	return new QTextEdit(QString(text), ((QWidget*)(parent)));
}

void QTextEdit_Destroy(QtObjectPtr ptr)
{
	((QTextEdit*)(ptr))->~QTextEdit();
}

int QTextEdit_AcceptRichText(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->acceptRichText();
}

int QTextEdit_Alignment(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->alignment();
}

int QTextEdit_CanPaste(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->canPaste();
}

QtObjectPtr QTextEdit_CreateStandardContextMenu(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->createStandardContextMenu();
}

int QTextEdit_CursorWidth(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->cursorWidth();
}

char* QTextEdit_DocumentTitle(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->documentTitle().toUtf8().data();
}

void QTextEdit_EnsureCursorVisible(QtObjectPtr ptr)
{
	((QTextEdit*)(ptr))->ensureCursorVisible();
}

char* QTextEdit_FontFamily(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->fontFamily().toUtf8().data();
}

int QTextEdit_FontItalic(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->fontItalic();
}

int QTextEdit_FontUnderline(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->fontUnderline();
}

int QTextEdit_FontWeight(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->fontWeight();
}

int QTextEdit_IsReadOnly(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->isReadOnly();
}

int QTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->isUndoRedoEnabled();
}

int QTextEdit_LineWrapColumnOrWidth(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->lineWrapColumnOrWidth();
}

int QTextEdit_OverwriteMode(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->overwriteMode();
}

char* QTextEdit_PlaceholderText(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->placeholderText().toUtf8().data();
}

void QTextEdit_SetAcceptRichText_Bool(QtObjectPtr ptr, int accept)
{
	((QTextEdit*)(ptr))->setAcceptRichText(accept != 0);
}

void QTextEdit_SetCursorWidth_Int(QtObjectPtr ptr, int width)
{
	((QTextEdit*)(ptr))->setCursorWidth(width);
}

void QTextEdit_SetDocumentTitle_String(QtObjectPtr ptr, char* title)
{
	((QTextEdit*)(ptr))->setDocumentTitle(QString(title));
}

void QTextEdit_SetLineWrapColumnOrWidth_Int(QtObjectPtr ptr, int w)
{
	((QTextEdit*)(ptr))->setLineWrapColumnOrWidth(w);
}

void QTextEdit_SetOverwriteMode_Bool(QtObjectPtr ptr, int overwrite)
{
	((QTextEdit*)(ptr))->setOverwriteMode(overwrite != 0);
}

void QTextEdit_SetPlaceholderText_String(QtObjectPtr ptr, char* placeholderText)
{
	((QTextEdit*)(ptr))->setPlaceholderText(QString(placeholderText));
}

void QTextEdit_SetReadOnly_Bool(QtObjectPtr ptr, int ro)
{
	((QTextEdit*)(ptr))->setReadOnly(ro != 0);
}

void QTextEdit_SetTabChangesFocus_Bool(QtObjectPtr ptr, int b)
{
	((QTextEdit*)(ptr))->setTabChangesFocus(b != 0);
}

void QTextEdit_SetTabStopWidth_Int(QtObjectPtr ptr, int width)
{
	((QTextEdit*)(ptr))->setTabStopWidth(width);
}

void QTextEdit_SetTextInteractionFlags_TextInteractionFlag(QtObjectPtr ptr, int flags)
{
	((QTextEdit*)(ptr))->setTextInteractionFlags(((Qt::TextInteractionFlag)(flags)));
}

void QTextEdit_SetUndoRedoEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QTextEdit*)(ptr))->setUndoRedoEnabled(enable != 0);
}

int QTextEdit_TabChangesFocus(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->tabChangesFocus();
}

int QTextEdit_TabStopWidth(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->tabStopWidth();
}

int QTextEdit_TextInteractionFlags(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->textInteractionFlags();
}

char* QTextEdit_ToHtml(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->toHtml().toUtf8().data();
}

char* QTextEdit_ToPlainText(QtObjectPtr ptr)
{
	return ((QTextEdit*)(ptr))->toPlainText().toUtf8().data();
}

//Public Slots
void QTextEdit_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Clear, ((QTextEdit*)(ptr)), &QTextEdit::clear, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Clear, ((QTextEdit*)(ptr)), &QTextEdit::clear);
}

void QTextEdit_Clear(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Clear();
}

void QTextEdit_ConnectSlotCopy(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Copy, ((QTextEdit*)(ptr)), &QTextEdit::copy, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotCopy(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Copy, ((QTextEdit*)(ptr)), &QTextEdit::copy);
}

void QTextEdit_Copy(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Copy();
}

void QTextEdit_ConnectSlotCut(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Cut, ((QTextEdit*)(ptr)), &QTextEdit::cut, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotCut(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Cut, ((QTextEdit*)(ptr)), &QTextEdit::cut);
}

void QTextEdit_Cut(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Cut();
}

void QTextEdit_ConnectSlotPaste(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Paste, ((QTextEdit*)(ptr)), &QTextEdit::paste, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotPaste(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Paste, ((QTextEdit*)(ptr)), &QTextEdit::paste);
}

void QTextEdit_Paste(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Paste();
}

void QTextEdit_ConnectSlotRedo(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Redo, ((QTextEdit*)(ptr)), &QTextEdit::redo, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotRedo(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Redo, ((QTextEdit*)(ptr)), &QTextEdit::redo);
}

void QTextEdit_Redo(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Redo();
}

void QTextEdit_ConnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SelectAll, ((QTextEdit*)(ptr)), &QTextEdit::selectAll, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SelectAll, ((QTextEdit*)(ptr)), &QTextEdit::selectAll);
}

void QTextEdit_SelectAll(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_SelectAll();
}

void QTextEdit_ConnectSlotSetFontItalic(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontItalic, ((QTextEdit*)(ptr)), &QTextEdit::setFontItalic, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotSetFontItalic(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontItalic, ((QTextEdit*)(ptr)), &QTextEdit::setFontItalic);
}

void QTextEdit_SetFontItalic_Bool(QtObjectPtr ptr, int italic)
{
	((MyQTextEdit*)(ptr))->Slot_SetFontItalic(italic != 0);
}

void QTextEdit_ConnectSlotSetFontUnderline(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontUnderline, ((QTextEdit*)(ptr)), &QTextEdit::setFontUnderline, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotSetFontUnderline(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontUnderline, ((QTextEdit*)(ptr)), &QTextEdit::setFontUnderline);
}

void QTextEdit_SetFontUnderline_Bool(QtObjectPtr ptr, int underline)
{
	((MyQTextEdit*)(ptr))->Slot_SetFontUnderline(underline != 0);
}

void QTextEdit_ConnectSlotSetFontWeight(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontWeight, ((QTextEdit*)(ptr)), &QTextEdit::setFontWeight, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotSetFontWeight(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_SetFontWeight, ((QTextEdit*)(ptr)), &QTextEdit::setFontWeight);
}

void QTextEdit_SetFontWeight_Int(QtObjectPtr ptr, int weight)
{
	((MyQTextEdit*)(ptr))->Slot_SetFontWeight(weight);
}

void QTextEdit_ConnectSlotUndo(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Undo, ((QTextEdit*)(ptr)), &QTextEdit::undo, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotUndo(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_Undo, ((QTextEdit*)(ptr)), &QTextEdit::undo);
}

void QTextEdit_Undo(QtObjectPtr ptr)
{
	((MyQTextEdit*)(ptr))->Slot_Undo();
}

void QTextEdit_ConnectSlotZoomIn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_ZoomIn, ((QTextEdit*)(ptr)), &QTextEdit::zoomIn, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotZoomIn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_ZoomIn, ((QTextEdit*)(ptr)), &QTextEdit::zoomIn);
}

void QTextEdit_ZoomIn_Int(QtObjectPtr ptr, int rang)
{
	((MyQTextEdit*)(ptr))->Slot_ZoomIn(rang);
}

void QTextEdit_ConnectSlotZoomOut(QtObjectPtr ptr)
{
	QObject::connect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_ZoomOut, ((QTextEdit*)(ptr)), &QTextEdit::zoomOut, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSlotZoomOut(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTextEdit*)(ptr)), &MyQTextEdit::Slot_ZoomOut, ((QTextEdit*)(ptr)), &QTextEdit::zoomOut);
}

void QTextEdit_ZoomOut_Int(QtObjectPtr ptr, int rang)
{
	((MyQTextEdit*)(ptr))->Slot_ZoomOut(rang);
}

//Signals
void QTextEdit_ConnectSignalCopyAvailable(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::copyAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_CopyAvailable, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalCopyAvailable(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::copyAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_CopyAvailable);
}

void QTextEdit_ConnectSignalCursorPositionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::cursorPositionChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_CursorPositionChanged, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalCursorPositionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::cursorPositionChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_CursorPositionChanged);
}

void QTextEdit_ConnectSignalRedoAvailable(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::redoAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_RedoAvailable, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalRedoAvailable(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::redoAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_RedoAvailable);
}

void QTextEdit_ConnectSignalSelectionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::selectionChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_SelectionChanged, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalSelectionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::selectionChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_SelectionChanged);
}

void QTextEdit_ConnectSignalTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::textChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_TextChanged, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::textChanged, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_TextChanged);
}

void QTextEdit_ConnectSignalUndoAvailable(QtObjectPtr ptr)
{
	QObject::connect(((QTextEdit*)(ptr)), &QTextEdit::undoAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_UndoAvailable, Qt::QueuedConnection);
}

void QTextEdit_DisconnectSignalUndoAvailable(QtObjectPtr ptr)
{
	QObject::disconnect(((QTextEdit*)(ptr)), &QTextEdit::undoAvailable, ((MyQTextEdit*)(ptr)), &MyQTextEdit::Signal_UndoAvailable);
}

