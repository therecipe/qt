#include "qtextedit.h"
#include <QColor>
#include <QTextDocument>
#include <QTextCharFormat>
#include <QObject>
#include <QUrl>
#include <QFont>
#include <QMetaObject>
#include <QWidget>
#include <QModelIndex>
#include <QTextCursor>
#include <QPoint>
#include <QPagedPaintDevice>
#include <QString>
#include <QVariant>
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

int QTextEdit_AcceptRichText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->acceptRichText();
}

int QTextEdit_AutoFormatting(void* ptr){
	return static_cast<QTextEdit*>(ptr)->autoFormatting();
}

int QTextEdit_CursorWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->cursorWidth();
}

void* QTextEdit_Document(void* ptr){
	return static_cast<QTextEdit*>(ptr)->document();
}

int QTextEdit_IsReadOnly(void* ptr){
	return static_cast<QTextEdit*>(ptr)->isReadOnly();
}

int QTextEdit_LineWrapColumnOrWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapColumnOrWidth();
}

int QTextEdit_LineWrapMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapMode();
}

int QTextEdit_OverwriteMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->overwriteMode();
}

char* QTextEdit_PlaceholderText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QTextEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "redo");
}

void QTextEdit_SetAcceptRichText(void* ptr, int accept){
	static_cast<QTextEdit*>(ptr)->setAcceptRichText(accept != 0);
}

void QTextEdit_SetAutoFormatting(void* ptr, int features){
	static_cast<QTextEdit*>(ptr)->setAutoFormatting(static_cast<QTextEdit::AutoFormattingFlag>(features));
}

void QTextEdit_SetCursorWidth(void* ptr, int width){
	static_cast<QTextEdit*>(ptr)->setCursorWidth(width);
}

void QTextEdit_SetDocument(void* ptr, void* document){
	static_cast<QTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QTextEdit_SetFontWeight(void* ptr, int weight){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontWeight", Q_ARG(int, weight));
}

void QTextEdit_SetHtml(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetLineWrapColumnOrWidth(void* ptr, int w){
	static_cast<QTextEdit*>(ptr)->setLineWrapColumnOrWidth(w);
}

void QTextEdit_SetLineWrapMode(void* ptr, int mode){
	static_cast<QTextEdit*>(ptr)->setLineWrapMode(static_cast<QTextEdit::LineWrapMode>(mode));
}

void QTextEdit_SetOverwriteMode(void* ptr, int overwrite){
	static_cast<QTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QTextEdit_SetPlaceholderText(void* ptr, char* placeholderText){
	static_cast<QTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QTextEdit_SetReadOnly(void* ptr, int ro){
	static_cast<QTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QTextEdit_SetTabChangesFocus(void* ptr, int b){
	static_cast<QTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QTextEdit_SetTabStopWidth(void* ptr, int width){
	static_cast<QTextEdit*>(ptr)->setTabStopWidth(width);
}

void QTextEdit_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QTextEdit_SetWordWrapMode(void* ptr, int policy){
	static_cast<QTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QTextEdit_TabChangesFocus(void* ptr){
	return static_cast<QTextEdit*>(ptr)->tabChangesFocus();
}

int QTextEdit_TabStopWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->tabStopWidth();
}

int QTextEdit_TextInteractionFlags(void* ptr){
	return static_cast<QTextEdit*>(ptr)->textInteractionFlags();
}

char* QTextEdit_ToHtml(void* ptr){
	return static_cast<QTextEdit*>(ptr)->toHtml().toUtf8().data();
}

int QTextEdit_WordWrapMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->wordWrapMode();
}

void QTextEdit_ZoomIn(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QTextEdit_ZoomOut(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

void* QTextEdit_NewQTextEdit(void* parent){
	return new QTextEdit(static_cast<QWidget*>(parent));
}

void* QTextEdit_NewQTextEdit2(char* text, void* parent){
	return new QTextEdit(QString(text), static_cast<QWidget*>(parent));
}

int QTextEdit_Alignment(void* ptr){
	return static_cast<QTextEdit*>(ptr)->alignment();
}

char* QTextEdit_AnchorAt(void* ptr, void* pos){
	return static_cast<QTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QTextEdit_Append(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "append", Q_ARG(QString, QString(text)));
}

int QTextEdit_CanPaste(void* ptr){
	return static_cast<QTextEdit*>(ptr)->canPaste();
}

void QTextEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "clear");
}

void QTextEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "copy");
}

void QTextEdit_ConnectCopyAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

void QTextEdit_DisconnectCopyAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

void* QTextEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu();
}

void* QTextEdit_CreateStandardContextMenu2(void* ptr, void* position){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QTextEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "cut");
}

char* QTextEdit_DocumentTitle(void* ptr){
	return static_cast<QTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QTextEdit_EnsureCursorVisible(void* ptr){
	static_cast<QTextEdit*>(ptr)->ensureCursorVisible();
}

char* QTextEdit_FontFamily(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontFamily().toUtf8().data();
}

int QTextEdit_FontItalic(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontItalic();
}

double QTextEdit_FontPointSize(void* ptr){
	return static_cast<double>(static_cast<QTextEdit*>(ptr)->fontPointSize());
}

int QTextEdit_FontUnderline(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontUnderline();
}

int QTextEdit_FontWeight(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontWeight();
}

void* QTextEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QTextEdit_InsertHtml(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_InsertPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QTextEdit_IsUndoRedoEnabled(void* ptr){
	return static_cast<QTextEdit*>(ptr)->isUndoRedoEnabled();
}

void* QTextEdit_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QTextEdit*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

void QTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier){
	static_cast<QTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextEdit_MoveCursor(void* ptr, int operation, int mode){
	static_cast<QTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QTextEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "paste");
}

void QTextEdit_Print(void* ptr, void* printer){
	static_cast<QTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextEdit_ConnectRedoAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_DisconnectRedoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_ScrollToAnchor(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "scrollToAnchor", Q_ARG(QString, QString(name)));
}

void QTextEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "selectAll");
}

void QTextEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_SetAlignment(void* ptr, int a){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setAlignment", Q_ARG(Qt::AlignmentFlag, static_cast<Qt::AlignmentFlag>(a)));
}

void QTextEdit_SetCurrentCharFormat(void* ptr, void* format){
	static_cast<QTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextEdit_SetCurrentFont(void* ptr, void* f){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(f)));
}

void QTextEdit_SetDocumentTitle(void* ptr, char* title){
	static_cast<QTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QTextEdit_SetFontFamily(void* ptr, char* fontFamily){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontFamily", Q_ARG(QString, QString(fontFamily)));
}

void QTextEdit_SetFontItalic(void* ptr, int italic){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontItalic", Q_ARG(bool, italic != 0));
}

void QTextEdit_SetFontPointSize(void* ptr, double s){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontPointSize", Q_ARG(qreal, static_cast<qreal>(s)));
}

void QTextEdit_SetFontUnderline(void* ptr, int underline){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontUnderline", Q_ARG(bool, underline != 0));
}

void QTextEdit_SetPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetTextBackgroundColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextBackgroundColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextCursor(void* ptr, void* cursor){
	static_cast<QTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QTextEdit_SetUndoRedoEnabled(void* ptr, int enable){
	static_cast<QTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void* QTextEdit_TextBackgroundColor(void* ptr){
	return new QColor(static_cast<QTextEdit*>(ptr)->textBackgroundColor());
}

void QTextEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

void QTextEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

void* QTextEdit_TextColor(void* ptr){
	return new QColor(static_cast<QTextEdit*>(ptr)->textColor());
}

char* QTextEdit_ToPlainText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QTextEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "undo");
}

void QTextEdit_ConnectUndoAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DisconnectUndoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DestroyQTextEdit(void* ptr){
	static_cast<QTextEdit*>(ptr)->~QTextEdit();
}

