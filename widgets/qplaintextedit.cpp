#include "qplaintextedit.h"
#include <QTextDocument>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QPoint>
#include <QTextOption>
#include <QTextCharFormat>
#include <QString>
#include <QMetaObject>
#include <QObject>
#include <QPagedPaintDevice>
#include <QTextCursor>
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

int QPlainTextEdit_BackgroundVisible(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->backgroundVisible();
}

int QPlainTextEdit_BlockCount(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->blockCount();
}

int QPlainTextEdit_CenterOnScroll(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->centerOnScroll();
}

int QPlainTextEdit_CursorWidth(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->cursorWidth();
}

int QPlainTextEdit_IsReadOnly(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isReadOnly();
}

int QPlainTextEdit_LineWrapMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->lineWrapMode();
}

int QPlainTextEdit_OverwriteMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->overwriteMode();
}

char* QPlainTextEdit_PlaceholderText(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QPlainTextEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "redo");
}

void QPlainTextEdit_SetBackgroundVisible(void* ptr, int visible){
	static_cast<QPlainTextEdit*>(ptr)->setBackgroundVisible(visible != 0);
}

void QPlainTextEdit_SetCenterOnScroll(void* ptr, int enabled){
	static_cast<QPlainTextEdit*>(ptr)->setCenterOnScroll(enabled != 0);
}

void QPlainTextEdit_SetCursorWidth(void* ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setCursorWidth(width);
}

void QPlainTextEdit_SetLineWrapMode(void* ptr, int mode){
	static_cast<QPlainTextEdit*>(ptr)->setLineWrapMode(static_cast<QPlainTextEdit::LineWrapMode>(mode));
}

void QPlainTextEdit_SetOverwriteMode(void* ptr, int overwrite){
	static_cast<QPlainTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QPlainTextEdit_SetPlaceholderText(void* ptr, char* placeholderText){
	static_cast<QPlainTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QPlainTextEdit_SetReadOnly(void* ptr, int ro){
	static_cast<QPlainTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QPlainTextEdit_SetTabChangesFocus(void* ptr, int b){
	static_cast<QPlainTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QPlainTextEdit_SetTabStopWidth(void* ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setTabStopWidth(width);
}

void QPlainTextEdit_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QPlainTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QPlainTextEdit_SetWordWrapMode(void* ptr, int policy){
	static_cast<QPlainTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QPlainTextEdit_TabChangesFocus(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabChangesFocus();
}

int QPlainTextEdit_TabStopWidth(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabStopWidth();
}

int QPlainTextEdit_TextInteractionFlags(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->textInteractionFlags();
}

int QPlainTextEdit_WordWrapMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->wordWrapMode();
}

void QPlainTextEdit_ZoomIn(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QPlainTextEdit_ZoomOut(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

void* QPlainTextEdit_NewQPlainTextEdit(void* parent){
	return new QPlainTextEdit(static_cast<QWidget*>(parent));
}

void* QPlainTextEdit_NewQPlainTextEdit2(char* text, void* parent){
	return new QPlainTextEdit(QString(text), static_cast<QWidget*>(parent));
}

char* QPlainTextEdit_AnchorAt(void* ptr, void* pos){
	return static_cast<QPlainTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QPlainTextEdit_AppendPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_CenterCursor(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "centerCursor");
}

void QPlainTextEdit_AppendHtml(void* ptr, char* html){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendHtml", Q_ARG(QString, QString(html)));
}

void QPlainTextEdit_ConnectBlockCountChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

void QPlainTextEdit_DisconnectBlockCountChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

int QPlainTextEdit_CanPaste(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->canPaste();
}

void QPlainTextEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "clear");
}

void QPlainTextEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "copy");
}

void QPlainTextEdit_ConnectCopyAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

void QPlainTextEdit_DisconnectCopyAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

void* QPlainTextEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu();
}

void* QPlainTextEdit_CreateStandardContextMenu2(void* ptr, void* position){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QPlainTextEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "cut");
}

void* QPlainTextEdit_Document(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->document();
}

char* QPlainTextEdit_DocumentTitle(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QPlainTextEdit_EnsureCursorVisible(void* ptr){
	static_cast<QPlainTextEdit*>(ptr)->ensureCursorVisible();
}

void* QPlainTextEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QPlainTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QPlainTextEdit_InsertPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QPlainTextEdit_IsUndoRedoEnabled(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isUndoRedoEnabled();
}

void* QPlainTextEdit_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QPlainTextEdit*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

int QPlainTextEdit_MaximumBlockCount(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->maximumBlockCount();
}

void QPlainTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier){
	static_cast<QPlainTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QPlainTextEdit_ConnectModificationChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_DisconnectModificationChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_MoveCursor(void* ptr, int operation, int mode){
	static_cast<QPlainTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QPlainTextEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "paste");
}

void QPlainTextEdit_Print(void* ptr, void* printer){
	static_cast<QPlainTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QPlainTextEdit_ConnectRedoAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_DisconnectRedoAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "selectAll");
}

void QPlainTextEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_SetCurrentCharFormat(void* ptr, void* format){
	static_cast<QPlainTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QPlainTextEdit_SetDocument(void* ptr, void* document){
	static_cast<QPlainTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QPlainTextEdit_SetDocumentTitle(void* ptr, char* title){
	static_cast<QPlainTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QPlainTextEdit_SetMaximumBlockCount(void* ptr, int maximum){
	static_cast<QPlainTextEdit*>(ptr)->setMaximumBlockCount(maximum);
}

void QPlainTextEdit_SetPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_SetTextCursor(void* ptr, void* cursor){
	static_cast<QPlainTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QPlainTextEdit_SetUndoRedoEnabled(void* ptr, int enable){
	static_cast<QPlainTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QPlainTextEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

void QPlainTextEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

char* QPlainTextEdit_ToPlainText(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QPlainTextEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "undo");
}

void QPlainTextEdit_ConnectUndoAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DisconnectUndoAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DestroyQPlainTextEdit(void* ptr){
	static_cast<QPlainTextEdit*>(ptr)->~QPlainTextEdit();
}

