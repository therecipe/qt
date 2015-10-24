#include "qlineedit.h"
#include <QMargins>
#include <QObject>
#include <QAction>
#include <QLine>
#include <QMetaObject>
#include <QVariant>
#include <QPoint>
#include <QValidator>
#include <QString>
#include <QIcon>
#include <QCompleter>
#include <QUrl>
#include <QEvent>
#include <QWidget>
#include <QModelIndex>
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

int QLineEdit_Alignment(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->alignment();
}

int QLineEdit_CursorMoveStyle(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->cursorMoveStyle();
}

int QLineEdit_CursorPosition(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->cursorPosition();
}

char* QLineEdit_DisplayText(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->displayText().toUtf8().data();
}

int QLineEdit_DragEnabled(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->dragEnabled();
}

int QLineEdit_EchoMode(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->echoMode();
}

int QLineEdit_HasAcceptableInput(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->hasAcceptableInput();
}

int QLineEdit_HasFrame(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->hasFrame();
}

int QLineEdit_HasSelectedText(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->hasSelectedText();
}

char* QLineEdit_InputMask(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->inputMask().toUtf8().data();
}

int QLineEdit_IsClearButtonEnabled(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->isClearButtonEnabled();
}

int QLineEdit_IsModified(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->isModified();
}

int QLineEdit_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->isReadOnly();
}

int QLineEdit_IsRedoAvailable(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->isRedoAvailable();
}

int QLineEdit_IsUndoAvailable(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->isUndoAvailable();
}

int QLineEdit_MaxLength(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->maxLength();
}

char* QLineEdit_PlaceholderText(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->placeholderText().toUtf8().data();
}

char* QLineEdit_SelectedText(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->selectedText().toUtf8().data();
}

void QLineEdit_SetAlignment(QtObjectPtr ptr, int flag){
	static_cast<QLineEdit*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QLineEdit_SetClearButtonEnabled(QtObjectPtr ptr, int enable){
	static_cast<QLineEdit*>(ptr)->setClearButtonEnabled(enable != 0);
}

void QLineEdit_SetCursorMoveStyle(QtObjectPtr ptr, int style){
	static_cast<QLineEdit*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QLineEdit_SetCursorPosition(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setCursorPosition(v);
}

void QLineEdit_SetDragEnabled(QtObjectPtr ptr, int b){
	static_cast<QLineEdit*>(ptr)->setDragEnabled(b != 0);
}

void QLineEdit_SetEchoMode(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setEchoMode(static_cast<QLineEdit::EchoMode>(v));
}

void QLineEdit_SetFrame(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setFrame(v != 0);
}

void QLineEdit_SetInputMask(QtObjectPtr ptr, char* inputMask){
	static_cast<QLineEdit*>(ptr)->setInputMask(QString(inputMask));
}

void QLineEdit_SetMaxLength(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setMaxLength(v);
}

void QLineEdit_SetModified(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setModified(v != 0);
}

void QLineEdit_SetPlaceholderText(QtObjectPtr ptr, char* v){
	static_cast<QLineEdit*>(ptr)->setPlaceholderText(QString(v));
}

void QLineEdit_SetReadOnly(QtObjectPtr ptr, int v){
	static_cast<QLineEdit*>(ptr)->setReadOnly(v != 0);
}

void QLineEdit_SetText(QtObjectPtr ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

char* QLineEdit_Text(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->text().toUtf8().data();
}

QtObjectPtr QLineEdit_NewQLineEdit(QtObjectPtr parent){
	return new QLineEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QLineEdit_NewQLineEdit2(char* contents, QtObjectPtr parent){
	return new QLineEdit(QString(contents), static_cast<QWidget*>(parent));
}

QtObjectPtr QLineEdit_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, int position){
	return static_cast<QLineEdit*>(ptr)->addAction(*static_cast<QIcon*>(icon), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_AddAction(QtObjectPtr ptr, QtObjectPtr action, int position){
	static_cast<QLineEdit*>(ptr)->addAction(static_cast<QAction*>(action), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_Backspace(QtObjectPtr ptr){
	static_cast<QLineEdit*>(ptr)->backspace();
}

void QLineEdit_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "clear");
}

QtObjectPtr QLineEdit_Completer(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->completer();
}

void QLineEdit_Copy(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "copy");
}

QtObjectPtr QLineEdit_CreateStandardContextMenu(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->createStandardContextMenu();
}

void QLineEdit_CursorBackward(QtObjectPtr ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorBackward(mark != 0, steps);
}

void QLineEdit_CursorForward(QtObjectPtr ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorForward(mark != 0, steps);
}

int QLineEdit_CursorPositionAt(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QLineEdit*>(ptr)->cursorPositionAt(*static_cast<QPoint*>(pos));
}

void QLineEdit_ConnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_CursorWordBackward(QtObjectPtr ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordBackward(mark != 0);
}

void QLineEdit_CursorWordForward(QtObjectPtr ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordForward(mark != 0);
}

void QLineEdit_Cut(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "cut");
}

void QLineEdit_Del(QtObjectPtr ptr){
	static_cast<QLineEdit*>(ptr)->del();
}

void QLineEdit_Deselect(QtObjectPtr ptr){
	static_cast<QLineEdit*>(ptr)->deselect();
}

void QLineEdit_ConnectEditingFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_DisconnectEditingFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_End(QtObjectPtr ptr, int mark){
	static_cast<QLineEdit*>(ptr)->end(mark != 0);
}

int QLineEdit_Event(QtObjectPtr ptr, QtObjectPtr e){
	return static_cast<QLineEdit*>(ptr)->event(static_cast<QEvent*>(e));
}

void QLineEdit_GetTextMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->getTextMargins(&left, &top, &right, &bottom);
}

void QLineEdit_Home(QtObjectPtr ptr, int mark){
	static_cast<QLineEdit*>(ptr)->home(mark != 0);
}

char* QLineEdit_InputMethodQuery(QtObjectPtr ptr, int property){
	return static_cast<QLineEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)).toString().toUtf8().data();
}

void QLineEdit_Paste(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "paste");
}

void QLineEdit_Redo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "redo");
}

void QLineEdit_ConnectReturnPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_DisconnectReturnPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_SelectAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "selectAll");
}

void QLineEdit_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

void QLineEdit_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

int QLineEdit_SelectionStart(QtObjectPtr ptr){
	return static_cast<QLineEdit*>(ptr)->selectionStart();
}

void QLineEdit_SetCompleter(QtObjectPtr ptr, QtObjectPtr c){
	static_cast<QLineEdit*>(ptr)->setCompleter(static_cast<QCompleter*>(c));
}

void QLineEdit_SetSelection(QtObjectPtr ptr, int start, int length){
	static_cast<QLineEdit*>(ptr)->setSelection(start, length);
}

void QLineEdit_SetTextMargins2(QtObjectPtr ptr, QtObjectPtr margins){
	static_cast<QLineEdit*>(ptr)->setTextMargins(*static_cast<QMargins*>(margins));
}

void QLineEdit_SetTextMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->setTextMargins(left, top, right, bottom);
}

void QLineEdit_SetValidator(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QLineEdit*>(ptr)->setValidator(static_cast<QValidator*>(v));
}

void QLineEdit_ConnectTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_DisconnectTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_ConnectTextEdited(QtObjectPtr ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_DisconnectTextEdited(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_Undo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "undo");
}

QtObjectPtr QLineEdit_Validator(QtObjectPtr ptr){
	return const_cast<QValidator*>(static_cast<QLineEdit*>(ptr)->validator());
}

void QLineEdit_DestroyQLineEdit(QtObjectPtr ptr){
	static_cast<QLineEdit*>(ptr)->~QLineEdit();
}

