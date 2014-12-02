#include "qlineedit.h"
#include <QLineEdit>
#include "cgoexport.h"

//MyQLineEdit
class MyQLineEdit: public QLineEdit {
Q_OBJECT
public:
void Signal_CursorPositionChanged() { callbackQLineEdit(this, QString("cursorPositionChanged").toUtf8().data()); };
void Signal_EditingFinished() { callbackQLineEdit(this, QString("editingFinished").toUtf8().data()); };
void Signal_ReturnPressed() { callbackQLineEdit(this, QString("returnPressed").toUtf8().data()); };
void Signal_SelectionChanged() { callbackQLineEdit(this, QString("selectionChanged").toUtf8().data()); };
void Signal_TextChanged() { callbackQLineEdit(this, QString("textChanged").toUtf8().data()); };
void Signal_TextEdited() { callbackQLineEdit(this, QString("textEdited").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_Copy();
void Slot_Cut();
void Slot_Paste();
void Slot_Redo();
void Slot_SelectAll();
void Slot_SetText(QString text);
void Slot_Undo();

};
#include "qlineedit.moc"

//Public Types
int QLineEdit_LeadingPosition() { return QLineEdit::LeadingPosition; }
int QLineEdit_TrailingPosition() { return QLineEdit::TrailingPosition; }
int QLineEdit_Normal() { return QLineEdit::Normal; }
int QLineEdit_NoEcho() { return QLineEdit::NoEcho; }
int QLineEdit_Password() { return QLineEdit::Password; }
int QLineEdit_PasswordEchoOnEdit() { return QLineEdit::PasswordEchoOnEdit; }

//Public Functions
QtObjectPtr QLineEdit_New_QWidget(QtObjectPtr parent)
{
	return new QLineEdit(((QWidget*)(parent)));
}

QtObjectPtr QLineEdit_New_String_QWidget(char* contents, QtObjectPtr parent)
{
	return new QLineEdit(QString(contents), ((QWidget*)(parent)));
}

void QLineEdit_Destroy(QtObjectPtr ptr)
{
	((QLineEdit*)(ptr))->~QLineEdit();
}

int QLineEdit_Alignment(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->alignment();
}

void QLineEdit_Backspace(QtObjectPtr ptr)
{
	((QLineEdit*)(ptr))->backspace();
}

QtObjectPtr QLineEdit_CreateStandardContextMenu(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->createStandardContextMenu();
}

void QLineEdit_CursorBackward_Bool_Int(QtObjectPtr ptr, int mark, int steps)
{
	((QLineEdit*)(ptr))->cursorBackward(mark != 0, steps);
}

void QLineEdit_CursorForward_Bool_Int(QtObjectPtr ptr, int mark, int steps)
{
	((QLineEdit*)(ptr))->cursorForward(mark != 0, steps);
}

int QLineEdit_CursorMoveStyle(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->cursorMoveStyle();
}

int QLineEdit_CursorPosition(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->cursorPosition();
}

void QLineEdit_CursorWordBackward_Bool(QtObjectPtr ptr, int mark)
{
	((QLineEdit*)(ptr))->cursorWordBackward(mark != 0);
}

void QLineEdit_CursorWordForward_Bool(QtObjectPtr ptr, int mark)
{
	((QLineEdit*)(ptr))->cursorWordForward(mark != 0);
}

void QLineEdit_Del(QtObjectPtr ptr)
{
	((QLineEdit*)(ptr))->del();
}

void QLineEdit_Deselect(QtObjectPtr ptr)
{
	((QLineEdit*)(ptr))->deselect();
}

char* QLineEdit_DisplayText(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->displayText().toUtf8().data();
}

int QLineEdit_DragEnabled(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->dragEnabled();
}

int QLineEdit_EchoMode(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->echoMode();
}

void QLineEdit_End_Bool(QtObjectPtr ptr, int mark)
{
	((QLineEdit*)(ptr))->end(mark != 0);
}

int QLineEdit_HasAcceptableInput(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->hasAcceptableInput();
}

int QLineEdit_HasFrame(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->hasFrame();
}

int QLineEdit_HasSelectedText(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->hasSelectedText();
}

void QLineEdit_Home_Bool(QtObjectPtr ptr, int mark)
{
	((QLineEdit*)(ptr))->home(mark != 0);
}

char* QLineEdit_InputMask(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->inputMask().toUtf8().data();
}

void QLineEdit_Insert_String(QtObjectPtr ptr, char* newText)
{
	((QLineEdit*)(ptr))->insert(QString(newText));
}

int QLineEdit_IsClearButtonEnabled(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->isClearButtonEnabled();
}

int QLineEdit_IsModified(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->isModified();
}

int QLineEdit_IsReadOnly(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->isReadOnly();
}

int QLineEdit_IsRedoAvailable(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->isRedoAvailable();
}

int QLineEdit_IsUndoAvailable(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->isUndoAvailable();
}

int QLineEdit_MaxLength(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->maxLength();
}

char* QLineEdit_PlaceholderText(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->placeholderText().toUtf8().data();
}

char* QLineEdit_SelectedText(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->selectedText().toUtf8().data();
}

int QLineEdit_SelectionStart(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->selectionStart();
}

void QLineEdit_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int flag)
{
	((QLineEdit*)(ptr))->setAlignment(((Qt::AlignmentFlag)(flag)));
}

void QLineEdit_SetClearButtonEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QLineEdit*)(ptr))->setClearButtonEnabled(enable != 0);
}

void QLineEdit_SetCursorMoveStyle_CursorMoveStyle(QtObjectPtr ptr, int style)
{
	((QLineEdit*)(ptr))->setCursorMoveStyle(((Qt::CursorMoveStyle)(style)));
}

void QLineEdit_SetCursorPosition_Int(QtObjectPtr ptr, int cursorPosition)
{
	((QLineEdit*)(ptr))->setCursorPosition(cursorPosition);
}

void QLineEdit_SetDragEnabled_Bool(QtObjectPtr ptr, int dragEnabled)
{
	((QLineEdit*)(ptr))->setDragEnabled(dragEnabled != 0);
}

void QLineEdit_SetEchoMode_EchoMode(QtObjectPtr ptr, int Ec)
{
	((QLineEdit*)(ptr))->setEchoMode(((QLineEdit::EchoMode)(Ec)));
}

void QLineEdit_SetFrame_Bool(QtObjectPtr ptr, int frame)
{
	((QLineEdit*)(ptr))->setFrame(frame != 0);
}

void QLineEdit_SetInputMask_String(QtObjectPtr ptr, char* inputMask)
{
	((QLineEdit*)(ptr))->setInputMask(QString(inputMask));
}

void QLineEdit_SetMaxLength_Int(QtObjectPtr ptr, int maxLength)
{
	((QLineEdit*)(ptr))->setMaxLength(maxLength);
}

void QLineEdit_SetModified_Bool(QtObjectPtr ptr, int modified)
{
	((QLineEdit*)(ptr))->setModified(modified != 0);
}

void QLineEdit_SetPlaceholderText_String(QtObjectPtr ptr, char* placeholderText)
{
	((QLineEdit*)(ptr))->setPlaceholderText(QString(placeholderText));
}

void QLineEdit_SetReadOnly_Bool(QtObjectPtr ptr, int readOnly)
{
	((QLineEdit*)(ptr))->setReadOnly(readOnly != 0);
}

void QLineEdit_SetSelection_Int_Int(QtObjectPtr ptr, int start, int length)
{
	((QLineEdit*)(ptr))->setSelection(start, length);
}

void QLineEdit_SetTextMargins_Int_Int_Int_Int(QtObjectPtr ptr, int left, int top, int right, int bottom)
{
	((QLineEdit*)(ptr))->setTextMargins(left, top, right, bottom);
}

void QLineEdit_SetValidator_QValidator(QtObjectPtr ptr, QtObjectPtr v)
{
	((QLineEdit*)(ptr))->setValidator(((QValidator*)(v)));
}

char* QLineEdit_Text(QtObjectPtr ptr)
{
	return ((QLineEdit*)(ptr))->text().toUtf8().data();
}

//Public Slots
void QLineEdit_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Clear, ((QLineEdit*)(ptr)), &QLineEdit::clear, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Clear, ((QLineEdit*)(ptr)), &QLineEdit::clear);
}

void QLineEdit_Clear(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Clear();
}

void QLineEdit_ConnectSlotCopy(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Copy, ((QLineEdit*)(ptr)), &QLineEdit::copy, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotCopy(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Copy, ((QLineEdit*)(ptr)), &QLineEdit::copy);
}

void QLineEdit_Copy(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Copy();
}

void QLineEdit_ConnectSlotCut(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Cut, ((QLineEdit*)(ptr)), &QLineEdit::cut, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotCut(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Cut, ((QLineEdit*)(ptr)), &QLineEdit::cut);
}

void QLineEdit_Cut(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Cut();
}

void QLineEdit_ConnectSlotPaste(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Paste, ((QLineEdit*)(ptr)), &QLineEdit::paste, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotPaste(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Paste, ((QLineEdit*)(ptr)), &QLineEdit::paste);
}

void QLineEdit_Paste(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Paste();
}

void QLineEdit_ConnectSlotRedo(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Redo, ((QLineEdit*)(ptr)), &QLineEdit::redo, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotRedo(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Redo, ((QLineEdit*)(ptr)), &QLineEdit::redo);
}

void QLineEdit_Redo(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Redo();
}

void QLineEdit_ConnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_SelectAll, ((QLineEdit*)(ptr)), &QLineEdit::selectAll, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_SelectAll, ((QLineEdit*)(ptr)), &QLineEdit::selectAll);
}

void QLineEdit_SelectAll(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_SelectAll();
}

void QLineEdit_ConnectSlotSetText(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_SetText, ((QLineEdit*)(ptr)), &QLineEdit::setText, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotSetText(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_SetText, ((QLineEdit*)(ptr)), &QLineEdit::setText);
}

void QLineEdit_SetText_String(QtObjectPtr ptr, char* text)
{
	((MyQLineEdit*)(ptr))->Slot_SetText(QString(text));
}

void QLineEdit_ConnectSlotUndo(QtObjectPtr ptr)
{
	QObject::connect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Undo, ((QLineEdit*)(ptr)), &QLineEdit::undo, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSlotUndo(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLineEdit*)(ptr)), &MyQLineEdit::Slot_Undo, ((QLineEdit*)(ptr)), &QLineEdit::undo);
}

void QLineEdit_Undo(QtObjectPtr ptr)
{
	((MyQLineEdit*)(ptr))->Slot_Undo();
}

//Signals
void QLineEdit_ConnectSignalCursorPositionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::cursorPositionChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_CursorPositionChanged, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalCursorPositionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::cursorPositionChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_CursorPositionChanged);
}

void QLineEdit_ConnectSignalEditingFinished(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::editingFinished, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_EditingFinished, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalEditingFinished(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::editingFinished, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_EditingFinished);
}

void QLineEdit_ConnectSignalReturnPressed(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::returnPressed, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_ReturnPressed, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalReturnPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::returnPressed, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_ReturnPressed);
}

void QLineEdit_ConnectSignalSelectionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::selectionChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_SelectionChanged, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalSelectionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::selectionChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_SelectionChanged);
}

void QLineEdit_ConnectSignalTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::textChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_TextChanged, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::textChanged, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_TextChanged);
}

void QLineEdit_ConnectSignalTextEdited(QtObjectPtr ptr)
{
	QObject::connect(((QLineEdit*)(ptr)), &QLineEdit::textEdited, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_TextEdited, Qt::QueuedConnection);
}

void QLineEdit_DisconnectSignalTextEdited(QtObjectPtr ptr)
{
	QObject::disconnect(((QLineEdit*)(ptr)), &QLineEdit::textEdited, ((MyQLineEdit*)(ptr)), &MyQLineEdit::Signal_TextEdited);
}

