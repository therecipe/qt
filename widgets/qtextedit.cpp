#include "qtextedit.h"
#include <QUrl>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QTextCursor>
#include <QColor>
#include <QMetaObject>
#include <QPoint>
#include <QObject>
#include <QTextDocument>
#include <QWidget>
#include <QFont>
#include <QPagedPaintDevice>
#include <QTextCharFormat>
#include <QTextOption>
#include <QTextEdit>
#include "_cgo_export.h"

class MyQTextEdit: public QTextEdit {
public:
void Signal_CopyAvailable(bool yes){callbackQTextEditCopyAvailable(this->objectName().toUtf8().data(), yes);};
void Signal_CursorPositionChanged(){callbackQTextEditCursorPositionChanged(this->objectName().toUtf8().data());};
void Signal_RedoAvailable(bool available){callbackQTextEditRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_SelectionChanged(){callbackQTextEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(){callbackQTextEditTextChanged(this->objectName().toUtf8().data());};
void Signal_UndoAvailable(bool available){callbackQTextEditUndoAvailable(this->objectName().toUtf8().data(), available);};
};

int QTextEdit_AcceptRichText(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->acceptRichText();
}

int QTextEdit_AutoFormatting(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->autoFormatting();
}

int QTextEdit_CursorWidth(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->cursorWidth();
}

QtObjectPtr QTextEdit_Document(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->document();
}

int QTextEdit_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->isReadOnly();
}

int QTextEdit_LineWrapColumnOrWidth(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapColumnOrWidth();
}

int QTextEdit_LineWrapMode(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapMode();
}

int QTextEdit_OverwriteMode(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->overwriteMode();
}

char* QTextEdit_PlaceholderText(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QTextEdit_Redo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "redo");
}

void QTextEdit_SetAcceptRichText(QtObjectPtr ptr, int accept){
	static_cast<QTextEdit*>(ptr)->setAcceptRichText(accept != 0);
}

void QTextEdit_SetAutoFormatting(QtObjectPtr ptr, int features){
	static_cast<QTextEdit*>(ptr)->setAutoFormatting(static_cast<QTextEdit::AutoFormattingFlag>(features));
}

void QTextEdit_SetCursorWidth(QtObjectPtr ptr, int width){
	static_cast<QTextEdit*>(ptr)->setCursorWidth(width);
}

void QTextEdit_SetDocument(QtObjectPtr ptr, QtObjectPtr document){
	static_cast<QTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QTextEdit_SetFontWeight(QtObjectPtr ptr, int weight){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontWeight", Q_ARG(int, weight));
}

void QTextEdit_SetHtml(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetLineWrapColumnOrWidth(QtObjectPtr ptr, int w){
	static_cast<QTextEdit*>(ptr)->setLineWrapColumnOrWidth(w);
}

void QTextEdit_SetLineWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QTextEdit*>(ptr)->setLineWrapMode(static_cast<QTextEdit::LineWrapMode>(mode));
}

void QTextEdit_SetOverwriteMode(QtObjectPtr ptr, int overwrite){
	static_cast<QTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QTextEdit_SetPlaceholderText(QtObjectPtr ptr, char* placeholderText){
	static_cast<QTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QTextEdit_SetReadOnly(QtObjectPtr ptr, int ro){
	static_cast<QTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QTextEdit_SetTabChangesFocus(QtObjectPtr ptr, int b){
	static_cast<QTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QTextEdit_SetTabStopWidth(QtObjectPtr ptr, int width){
	static_cast<QTextEdit*>(ptr)->setTabStopWidth(width);
}

void QTextEdit_SetTextInteractionFlags(QtObjectPtr ptr, int flags){
	static_cast<QTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QTextEdit_SetWordWrapMode(QtObjectPtr ptr, int policy){
	static_cast<QTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QTextEdit_TabChangesFocus(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->tabChangesFocus();
}

int QTextEdit_TabStopWidth(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->tabStopWidth();
}

int QTextEdit_TextInteractionFlags(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->textInteractionFlags();
}

char* QTextEdit_ToHtml(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->toHtml().toUtf8().data();
}

int QTextEdit_WordWrapMode(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->wordWrapMode();
}

void QTextEdit_ZoomIn(QtObjectPtr ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QTextEdit_ZoomOut(QtObjectPtr ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

QtObjectPtr QTextEdit_NewQTextEdit(QtObjectPtr parent){
	return new QTextEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QTextEdit_NewQTextEdit2(char* text, QtObjectPtr parent){
	return new QTextEdit(QString(text), static_cast<QWidget*>(parent));
}

int QTextEdit_Alignment(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->alignment();
}

char* QTextEdit_AnchorAt(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QTextEdit_Append(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "append", Q_ARG(QString, QString(text)));
}

int QTextEdit_CanPaste(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->canPaste();
}

void QTextEdit_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "clear");
}

void QTextEdit_Copy(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "copy");
}

void QTextEdit_ConnectCopyAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

void QTextEdit_DisconnectCopyAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

QtObjectPtr QTextEdit_CreateStandardContextMenu(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu();
}

QtObjectPtr QTextEdit_CreateStandardContextMenu2(QtObjectPtr ptr, QtObjectPtr position){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QTextEdit_ConnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_DisconnectCursorPositionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_Cut(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "cut");
}

char* QTextEdit_DocumentTitle(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QTextEdit_EnsureCursorVisible(QtObjectPtr ptr){
	static_cast<QTextEdit*>(ptr)->ensureCursorVisible();
}

char* QTextEdit_FontFamily(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->fontFamily().toUtf8().data();
}

int QTextEdit_FontItalic(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->fontItalic();
}

int QTextEdit_FontUnderline(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->fontUnderline();
}

int QTextEdit_FontWeight(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->fontWeight();
}

char* QTextEdit_InputMethodQuery(QtObjectPtr ptr, int property){
	return static_cast<QTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)).toString().toUtf8().data();
}

void QTextEdit_InsertHtml(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_InsertPlainText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QTextEdit_IsUndoRedoEnabled(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->isUndoRedoEnabled();
}

char* QTextEdit_LoadResource(QtObjectPtr ptr, int ty, char* name){
	return static_cast<QTextEdit*>(ptr)->loadResource(ty, QUrl(QString(name))).toString().toUtf8().data();
}

void QTextEdit_MergeCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr modifier){
	static_cast<QTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextEdit_MoveCursor(QtObjectPtr ptr, int operation, int mode){
	static_cast<QTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QTextEdit_Paste(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "paste");
}

void QTextEdit_Print(QtObjectPtr ptr, QtObjectPtr printer){
	static_cast<QTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextEdit_ConnectRedoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_DisconnectRedoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_ScrollToAnchor(QtObjectPtr ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "scrollToAnchor", Q_ARG(QString, QString(name)));
}

void QTextEdit_SelectAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "selectAll");
}

void QTextEdit_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_SetAlignment(QtObjectPtr ptr, int a){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setAlignment", Q_ARG(Qt::AlignmentFlag, static_cast<Qt::AlignmentFlag>(a)));
}

void QTextEdit_SetCurrentCharFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextEdit_SetCurrentFont(QtObjectPtr ptr, QtObjectPtr f){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(f)));
}

void QTextEdit_SetDocumentTitle(QtObjectPtr ptr, char* title){
	static_cast<QTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QTextEdit_SetFontFamily(QtObjectPtr ptr, char* fontFamily){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontFamily", Q_ARG(QString, QString(fontFamily)));
}

void QTextEdit_SetFontItalic(QtObjectPtr ptr, int italic){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontItalic", Q_ARG(bool, italic != 0));
}

void QTextEdit_SetFontUnderline(QtObjectPtr ptr, int underline){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontUnderline", Q_ARG(bool, underline != 0));
}

void QTextEdit_SetPlainText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetTextBackgroundColor(QtObjectPtr ptr, QtObjectPtr c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextBackgroundColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextColor(QtObjectPtr ptr, QtObjectPtr c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QTextEdit_SetUndoRedoEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QTextEdit_ConnectTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

void QTextEdit_DisconnectTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

char* QTextEdit_ToPlainText(QtObjectPtr ptr){
	return static_cast<QTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QTextEdit_Undo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "undo");
}

void QTextEdit_ConnectUndoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DisconnectUndoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DestroyQTextEdit(QtObjectPtr ptr){
	static_cast<QTextEdit*>(ptr)->~QTextEdit();
}

