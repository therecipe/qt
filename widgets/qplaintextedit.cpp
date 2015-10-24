#include "qplaintextedit.h"
#include <QVariant>
#include <QMetaObject>
#include <QTextCharFormat>
#include <QTextOption>
#include <QObject>
#include <QPoint>
#include <QWidget>
#include <QString>
#include <QPagedPaintDevice>
#include <QTextCursor>
#include <QModelIndex>
#include <QTextDocument>
#include <QUrl>
#include <QPlainTextEdit>
#include "_cgo_export.h"

class MyQPlainTextEdit: public QPlainTextEdit {
public:
void Signal_BlockCountChanged(int newBlockCount){callbackQPlainTextEditBlockCountChanged(this->objectName().toUtf8().data(), newBlockCount);};
void Signal_CopyAvailable(bool yes){callbackQPlainTextEditCopyAvailable(this->objectName().toUtf8().data(), yes);};
void Signal_CursorPositionChanged(){callbackQPlainTextEditCursorPositionChanged(this->objectName().toUtf8().data());};
void Signal_ModificationChanged(bool changed){callbackQPlainTextEditModificationChanged(this->objectName().toUtf8().data(), changed);};
void Signal_RedoAvailable(bool available){callbackQPlainTextEditRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_SelectionChanged(){callbackQPlainTextEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(){callbackQPlainTextEditTextChanged(this->objectName().toUtf8().data());};
void Signal_UndoAvailable(bool available){callbackQPlainTextEditUndoAvailable(this->objectName().toUtf8().data(), available);};
};

int QPlainTextEdit_BackgroundVisible(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->backgroundVisible();
}

int QPlainTextEdit_BlockCount(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->blockCount();
}

int QPlainTextEdit_CenterOnScroll(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->centerOnScroll();
}

int QPlainTextEdit_CursorWidth(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->cursorWidth();
}

int QPlainTextEdit_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isReadOnly();
}

int QPlainTextEdit_LineWrapMode(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->lineWrapMode();
}

int QPlainTextEdit_OverwriteMode(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->overwriteMode();
}

char* QPlainTextEdit_PlaceholderText(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QPlainTextEdit_Redo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "redo");
}

void QPlainTextEdit_SetBackgroundVisible(QtObjectPtr ptr, int visible){
	static_cast<QPlainTextEdit*>(ptr)->setBackgroundVisible(visible != 0);
}

void QPlainTextEdit_SetCenterOnScroll(QtObjectPtr ptr, int enabled){
	static_cast<QPlainTextEdit*>(ptr)->setCenterOnScroll(enabled != 0);
}

void QPlainTextEdit_SetCursorWidth(QtObjectPtr ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setCursorWidth(width);
}

void QPlainTextEdit_SetLineWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QPlainTextEdit*>(ptr)->setLineWrapMode(static_cast<QPlainTextEdit::LineWrapMode>(mode));
}

void QPlainTextEdit_SetOverwriteMode(QtObjectPtr ptr, int overwrite){
	static_cast<QPlainTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QPlainTextEdit_SetPlaceholderText(QtObjectPtr ptr, char* placeholderText){
	static_cast<QPlainTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QPlainTextEdit_SetReadOnly(QtObjectPtr ptr, int ro){
	static_cast<QPlainTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QPlainTextEdit_SetTabChangesFocus(QtObjectPtr ptr, int b){
	static_cast<QPlainTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QPlainTextEdit_SetTabStopWidth(QtObjectPtr ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setTabStopWidth(width);
}

void QPlainTextEdit_SetTextInteractionFlags(QtObjectPtr ptr, int flags){
	static_cast<QPlainTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QPlainTextEdit_SetWordWrapMode(QtObjectPtr ptr, int policy){
	static_cast<QPlainTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QPlainTextEdit_TabChangesFocus(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabChangesFocus();
}

int QPlainTextEdit_TabStopWidth(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabStopWidth();
}

int QPlainTextEdit_TextInteractionFlags(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->textInteractionFlags();
}

int QPlainTextEdit_WordWrapMode(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->wordWrapMode();
}

void QPlainTextEdit_ZoomIn(QtObjectPtr ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QPlainTextEdit_ZoomOut(QtObjectPtr ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

QtObjectPtr QPlainTextEdit_NewQPlainTextEdit(QtObjectPtr parent){
	return new QPlainTextEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QPlainTextEdit_NewQPlainTextEdit2(char* text, QtObjectPtr parent){
	return new QPlainTextEdit(QString(text), static_cast<QWidget*>(parent));
}

char* QPlainTextEdit_AnchorAt(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QPlainTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QPlainTextEdit_AppendPlainText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_CenterCursor(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "centerCursor");
}

void QPlainTextEdit_AppendHtml(QtObjectPtr ptr, char* html){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendHtml", Q_ARG(QString, QString(html)));
}

void QPlainTextEdit_ConnectBlockCountChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

void QPlainTextEdit_DisconnectBlockCountChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

int QPlainTextEdit_CanPaste(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->canPaste();
}

void QPlainTextEdit_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "clear");
}

void QPlainTextEdit_Copy(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "copy");
}

void QPlainTextEdit_ConnectCopyAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

void QPlainTextEdit_DisconnectCopyAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

QtObjectPtr QPlainTextEdit_CreateStandardContextMenu(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu();
}

QtObjectPtr QPlainTextEdit_CreateStandardContextMenu2(QtObjectPtr ptr, QtObjectPtr position){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QPlainTextEdit_ConnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_Cut(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "cut");
}

QtObjectPtr QPlainTextEdit_Document(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->document();
}

char* QPlainTextEdit_DocumentTitle(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QPlainTextEdit_EnsureCursorVisible(QtObjectPtr ptr){
	static_cast<QPlainTextEdit*>(ptr)->ensureCursorVisible();
}

char* QPlainTextEdit_InputMethodQuery(QtObjectPtr ptr, int property){
	return static_cast<QPlainTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)).toString().toUtf8().data();
}

void QPlainTextEdit_InsertPlainText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QPlainTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isUndoRedoEnabled();
}

char* QPlainTextEdit_LoadResource(QtObjectPtr ptr, int ty, char* name){
	return static_cast<QPlainTextEdit*>(ptr)->loadResource(ty, QUrl(QString(name))).toString().toUtf8().data();
}

int QPlainTextEdit_MaximumBlockCount(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->maximumBlockCount();
}

void QPlainTextEdit_MergeCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr modifier){
	static_cast<QPlainTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QPlainTextEdit_ConnectModificationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_DisconnectModificationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_MoveCursor(QtObjectPtr ptr, int operation, int mode){
	static_cast<QPlainTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QPlainTextEdit_Paste(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "paste");
}

void QPlainTextEdit_Print(QtObjectPtr ptr, QtObjectPtr printer){
	static_cast<QPlainTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QPlainTextEdit_ConnectRedoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_DisconnectRedoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_SelectAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "selectAll");
}

void QPlainTextEdit_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_SetCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QPlainTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QPlainTextEdit_SetDocument(QtObjectPtr ptr, QtObjectPtr document){
	static_cast<QPlainTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QPlainTextEdit_SetDocumentTitle(QtObjectPtr ptr, char* title){
	static_cast<QPlainTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QPlainTextEdit_SetMaximumBlockCount(QtObjectPtr ptr, int maximum){
	static_cast<QPlainTextEdit*>(ptr)->setMaximumBlockCount(maximum);
}

void QPlainTextEdit_SetPlainText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QPlainTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QPlainTextEdit_SetUndoRedoEnabled(QtObjectPtr ptr, int enable){
	static_cast<QPlainTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QPlainTextEdit_ConnectTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

void QPlainTextEdit_DisconnectTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

char* QPlainTextEdit_ToPlainText(QtObjectPtr ptr){
	return static_cast<QPlainTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QPlainTextEdit_Undo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "undo");
}

void QPlainTextEdit_ConnectUndoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DisconnectUndoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DestroyQPlainTextEdit(QtObjectPtr ptr){
	static_cast<QPlainTextEdit*>(ptr)->~QPlainTextEdit();
}

