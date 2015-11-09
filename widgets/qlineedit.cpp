#include "qlineedit.h"
#include <QUrl>
#include <QLine>
#include <QModelIndex>
#include <QAction>
#include <QValidator>
#include <QEvent>
#include <QWidget>
#include <QVariant>
#include <QObject>
#include <QMargins>
#include <QCompleter>
#include <QPoint>
#include <QString>
#include <QMetaObject>
#include <QIcon>
#include <QLineEdit>
#include "_cgo_export.h"

class MyQLineEdit: public QLineEdit {
public:
void Signal_CursorPositionChanged(int old, int n){callbackQLineEditCursorPositionChanged(this->objectName().toUtf8().data(), old, n);};
void Signal_EditingFinished(){callbackQLineEditEditingFinished(this->objectName().toUtf8().data());};
void Signal_ReturnPressed(){callbackQLineEditReturnPressed(this->objectName().toUtf8().data());};
void Signal_SelectionChanged(){callbackQLineEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(const QString & text){callbackQLineEditTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_TextEdited(const QString & text){callbackQLineEditTextEdited(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

int QLineEdit_Alignment(void* ptr){
	return static_cast<QLineEdit*>(ptr)->alignment();
}

int QLineEdit_CursorMoveStyle(void* ptr){
	return static_cast<QLineEdit*>(ptr)->cursorMoveStyle();
}

int QLineEdit_CursorPosition(void* ptr){
	return static_cast<QLineEdit*>(ptr)->cursorPosition();
}

char* QLineEdit_DisplayText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->displayText().toUtf8().data();
}

int QLineEdit_DragEnabled(void* ptr){
	return static_cast<QLineEdit*>(ptr)->dragEnabled();
}

int QLineEdit_EchoMode(void* ptr){
	return static_cast<QLineEdit*>(ptr)->echoMode();
}

int QLineEdit_HasAcceptableInput(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasAcceptableInput();
}

int QLineEdit_HasFrame(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasFrame();
}

int QLineEdit_HasSelectedText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasSelectedText();
}

char* QLineEdit_InputMask(void* ptr){
	return static_cast<QLineEdit*>(ptr)->inputMask().toUtf8().data();
}

int QLineEdit_IsClearButtonEnabled(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isClearButtonEnabled();
}

int QLineEdit_IsModified(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isModified();
}

int QLineEdit_IsReadOnly(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isReadOnly();
}

int QLineEdit_IsRedoAvailable(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isRedoAvailable();
}

int QLineEdit_IsUndoAvailable(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isUndoAvailable();
}

int QLineEdit_MaxLength(void* ptr){
	return static_cast<QLineEdit*>(ptr)->maxLength();
}

char* QLineEdit_PlaceholderText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->placeholderText().toUtf8().data();
}

char* QLineEdit_SelectedText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->selectedText().toUtf8().data();
}

void QLineEdit_SetAlignment(void* ptr, int flag){
	static_cast<QLineEdit*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QLineEdit_SetClearButtonEnabled(void* ptr, int enable){
	static_cast<QLineEdit*>(ptr)->setClearButtonEnabled(enable != 0);
}

void QLineEdit_SetCursorMoveStyle(void* ptr, int style){
	static_cast<QLineEdit*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QLineEdit_SetCursorPosition(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setCursorPosition(v);
}

void QLineEdit_SetDragEnabled(void* ptr, int b){
	static_cast<QLineEdit*>(ptr)->setDragEnabled(b != 0);
}

void QLineEdit_SetEchoMode(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setEchoMode(static_cast<QLineEdit::EchoMode>(v));
}

void QLineEdit_SetFrame(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setFrame(v != 0);
}

void QLineEdit_SetInputMask(void* ptr, char* inputMask){
	static_cast<QLineEdit*>(ptr)->setInputMask(QString(inputMask));
}

void QLineEdit_SetMaxLength(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setMaxLength(v);
}

void QLineEdit_SetModified(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setModified(v != 0);
}

void QLineEdit_SetPlaceholderText(void* ptr, char* v){
	static_cast<QLineEdit*>(ptr)->setPlaceholderText(QString(v));
}

void QLineEdit_SetReadOnly(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setReadOnly(v != 0);
}

void QLineEdit_SetText(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

char* QLineEdit_Text(void* ptr){
	return static_cast<QLineEdit*>(ptr)->text().toUtf8().data();
}

void* QLineEdit_NewQLineEdit(void* parent){
	return new QLineEdit(static_cast<QWidget*>(parent));
}

void* QLineEdit_NewQLineEdit2(char* contents, void* parent){
	return new QLineEdit(QString(contents), static_cast<QWidget*>(parent));
}

void* QLineEdit_AddAction2(void* ptr, void* icon, int position){
	return static_cast<QLineEdit*>(ptr)->addAction(*static_cast<QIcon*>(icon), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_AddAction(void* ptr, void* action, int position){
	static_cast<QLineEdit*>(ptr)->addAction(static_cast<QAction*>(action), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_Backspace(void* ptr){
	static_cast<QLineEdit*>(ptr)->backspace();
}

void QLineEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "clear");
}

void* QLineEdit_Completer(void* ptr){
	return static_cast<QLineEdit*>(ptr)->completer();
}

void QLineEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "copy");
}

void* QLineEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QLineEdit*>(ptr)->createStandardContextMenu();
}

void QLineEdit_CursorBackward(void* ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorBackward(mark != 0, steps);
}

void QLineEdit_CursorForward(void* ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorForward(mark != 0, steps);
}

int QLineEdit_CursorPositionAt(void* ptr, void* pos){
	return static_cast<QLineEdit*>(ptr)->cursorPositionAt(*static_cast<QPoint*>(pos));
}

void QLineEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_CursorWordBackward(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordBackward(mark != 0);
}

void QLineEdit_CursorWordForward(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordForward(mark != 0);
}

void QLineEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "cut");
}

void QLineEdit_Del(void* ptr){
	static_cast<QLineEdit*>(ptr)->del();
}

void QLineEdit_Deselect(void* ptr){
	static_cast<QLineEdit*>(ptr)->deselect();
}

void QLineEdit_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_End(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->end(mark != 0);
}

int QLineEdit_Event(void* ptr, void* e){
	return static_cast<QLineEdit*>(ptr)->event(static_cast<QEvent*>(e));
}

void QLineEdit_GetTextMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->getTextMargins(&left, &top, &right, &bottom);
}

void QLineEdit_Home(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->home(mark != 0);
}

void* QLineEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QLineEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QLineEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "paste");
}

void QLineEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "redo");
}

void QLineEdit_ConnectReturnPressed(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_DisconnectReturnPressed(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "selectAll");
}

void QLineEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

void QLineEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

int QLineEdit_SelectionStart(void* ptr){
	return static_cast<QLineEdit*>(ptr)->selectionStart();
}

void QLineEdit_SetCompleter(void* ptr, void* c){
	static_cast<QLineEdit*>(ptr)->setCompleter(static_cast<QCompleter*>(c));
}

void QLineEdit_SetSelection(void* ptr, int start, int length){
	static_cast<QLineEdit*>(ptr)->setSelection(start, length);
}

void QLineEdit_SetTextMargins2(void* ptr, void* margins){
	static_cast<QLineEdit*>(ptr)->setTextMargins(*static_cast<QMargins*>(margins));
}

void QLineEdit_SetTextMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->setTextMargins(left, top, right, bottom);
}

void QLineEdit_SetValidator(void* ptr, void* v){
	static_cast<QLineEdit*>(ptr)->setValidator(static_cast<QValidator*>(v));
}

void QLineEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_ConnectTextEdited(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_DisconnectTextEdited(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "undo");
}

void* QLineEdit_Validator(void* ptr){
	return const_cast<QValidator*>(static_cast<QLineEdit*>(ptr)->validator());
}

void QLineEdit_DestroyQLineEdit(void* ptr){
	static_cast<QLineEdit*>(ptr)->~QLineEdit();
}

