#include "qtextdocument.h"
#include <QAbstractTextDocumentLayout>
#include <QFont>
#include <QRectF>
#include <QSize>
#include <QVariant>
#include <QTextFormat>
#include <QPagedPaintDevice>
#include <QPainter>
#include <QString>
#include <QUrl>
#include <QTextCursor>
#include <QRect>
#include <QSizeF>
#include <QModelIndex>
#include <QByteArray>
#include <QMetaObject>
#include <QObject>
#include <QTextOption>
#include <QTextDocument>
#include "_cgo_export.h"

class MyQTextDocument: public QTextDocument {
public:
void Signal_BlockCountChanged(int newBlockCount){callbackQTextDocumentBlockCountChanged(this->objectName().toUtf8().data(), newBlockCount);};
void Signal_ContentsChange(int position, int charsRemoved, int charsAdded){callbackQTextDocumentContentsChange(this->objectName().toUtf8().data(), position, charsRemoved, charsAdded);};
void Signal_ContentsChanged(){callbackQTextDocumentContentsChanged(this->objectName().toUtf8().data());};
void Signal_DocumentLayoutChanged(){callbackQTextDocumentDocumentLayoutChanged(this->objectName().toUtf8().data());};
void Signal_ModificationChanged(bool changed){callbackQTextDocumentModificationChanged(this->objectName().toUtf8().data(), changed);};
void Signal_RedoAvailable(bool available){callbackQTextDocumentRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_UndoAvailable(bool available){callbackQTextDocumentUndoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_UndoCommandAdded(){callbackQTextDocumentUndoCommandAdded(this->objectName().toUtf8().data());};
};

int QTextDocument_BlockCount(void* ptr){
	return static_cast<QTextDocument*>(ptr)->blockCount();
}

char* QTextDocument_DefaultStyleSheet(void* ptr){
	return static_cast<QTextDocument*>(ptr)->defaultStyleSheet().toUtf8().data();
}

double QTextDocument_DocumentMargin(void* ptr){
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->documentMargin());
}

double QTextDocument_IndentWidth(void* ptr){
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->indentWidth());
}

int QTextDocument_IsModified(void* ptr){
	return static_cast<QTextDocument*>(ptr)->isModified();
}

int QTextDocument_IsUndoRedoEnabled(void* ptr){
	return static_cast<QTextDocument*>(ptr)->isUndoRedoEnabled();
}

void QTextDocument_MarkContentsDirty(void* ptr, int position, int length){
	static_cast<QTextDocument*>(ptr)->markContentsDirty(position, length);
}

int QTextDocument_MaximumBlockCount(void* ptr){
	return static_cast<QTextDocument*>(ptr)->maximumBlockCount();
}

void QTextDocument_SetBaseUrl(void* ptr, void* url){
	static_cast<QTextDocument*>(ptr)->setBaseUrl(*static_cast<QUrl*>(url));
}

void QTextDocument_SetDefaultStyleSheet(void* ptr, char* sheet){
	static_cast<QTextDocument*>(ptr)->setDefaultStyleSheet(QString(sheet));
}

void QTextDocument_SetDocumentMargin(void* ptr, double margin){
	static_cast<QTextDocument*>(ptr)->setDocumentMargin(static_cast<qreal>(margin));
}

void QTextDocument_SetMaximumBlockCount(void* ptr, int maximum){
	static_cast<QTextDocument*>(ptr)->setMaximumBlockCount(maximum);
}

void QTextDocument_SetModified(void* ptr, int m){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "setModified", Q_ARG(bool, m != 0));
}

void QTextDocument_SetPageSize(void* ptr, void* size){
	static_cast<QTextDocument*>(ptr)->setPageSize(*static_cast<QSizeF*>(size));
}

void QTextDocument_SetTextWidth(void* ptr, double width){
	static_cast<QTextDocument*>(ptr)->setTextWidth(static_cast<qreal>(width));
}

void QTextDocument_SetUndoRedoEnabled(void* ptr, int enable){
	static_cast<QTextDocument*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QTextDocument_SetUseDesignMetrics(void* ptr, int b){
	static_cast<QTextDocument*>(ptr)->setUseDesignMetrics(b != 0);
}

double QTextDocument_TextWidth(void* ptr){
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->textWidth());
}

int QTextDocument_UseDesignMetrics(void* ptr){
	return static_cast<QTextDocument*>(ptr)->useDesignMetrics();
}

void* QTextDocument_NewQTextDocument(void* parent){
	return new QTextDocument(static_cast<QObject*>(parent));
}

void* QTextDocument_NewQTextDocument2(char* text, void* parent){
	return new QTextDocument(QString(text), static_cast<QObject*>(parent));
}

void QTextDocument_AddResource(void* ptr, int ty, void* name, void* resource){
	static_cast<QTextDocument*>(ptr)->addResource(ty, *static_cast<QUrl*>(name), *static_cast<QVariant*>(resource));
}

void QTextDocument_AdjustSize(void* ptr){
	static_cast<QTextDocument*>(ptr)->adjustSize();
}

int QTextDocument_AvailableRedoSteps(void* ptr){
	return static_cast<QTextDocument*>(ptr)->availableRedoSteps();
}

int QTextDocument_AvailableUndoSteps(void* ptr){
	return static_cast<QTextDocument*>(ptr)->availableUndoSteps();
}

void QTextDocument_ConnectBlockCountChanged(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));;
}

void QTextDocument_DisconnectBlockCountChanged(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));;
}

int QTextDocument_CharacterCount(void* ptr){
	return static_cast<QTextDocument*>(ptr)->characterCount();
}

void QTextDocument_Clear(void* ptr){
	static_cast<QTextDocument*>(ptr)->clear();
}

void QTextDocument_ClearUndoRedoStacks(void* ptr, int stacksToClear){
	static_cast<QTextDocument*>(ptr)->clearUndoRedoStacks(static_cast<QTextDocument::Stacks>(stacksToClear));
}

void* QTextDocument_Clone(void* ptr, void* parent){
	return static_cast<QTextDocument*>(ptr)->clone(static_cast<QObject*>(parent));
}

void QTextDocument_ConnectContentsChange(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));;
}

void QTextDocument_DisconnectContentsChange(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));;
}

void QTextDocument_ConnectContentsChanged(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));;
}

void QTextDocument_DisconnectContentsChanged(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));;
}

int QTextDocument_DefaultCursorMoveStyle(void* ptr){
	return static_cast<QTextDocument*>(ptr)->defaultCursorMoveStyle();
}

void* QTextDocument_DocumentLayout(void* ptr){
	return static_cast<QTextDocument*>(ptr)->documentLayout();
}

void QTextDocument_ConnectDocumentLayoutChanged(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));;
}

void QTextDocument_DisconnectDocumentLayoutChanged(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));;
}

void QTextDocument_DrawContents(void* ptr, void* p, void* rect){
	static_cast<QTextDocument*>(ptr)->drawContents(static_cast<QPainter*>(p), *static_cast<QRectF*>(rect));
}

double QTextDocument_IdealWidth(void* ptr){
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->idealWidth());
}

int QTextDocument_IsEmpty(void* ptr){
	return static_cast<QTextDocument*>(ptr)->isEmpty();
}

int QTextDocument_IsRedoAvailable(void* ptr){
	return static_cast<QTextDocument*>(ptr)->isRedoAvailable();
}

int QTextDocument_IsUndoAvailable(void* ptr){
	return static_cast<QTextDocument*>(ptr)->isUndoAvailable();
}

int QTextDocument_LineCount(void* ptr){
	return static_cast<QTextDocument*>(ptr)->lineCount();
}

char* QTextDocument_MetaInformation(void* ptr, int info){
	return static_cast<QTextDocument*>(ptr)->metaInformation(static_cast<QTextDocument::MetaInformation>(info)).toUtf8().data();
}

void QTextDocument_ConnectModificationChanged(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));;
}

void QTextDocument_DisconnectModificationChanged(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));;
}

void* QTextDocument_Object(void* ptr, int objectIndex){
	return static_cast<QTextDocument*>(ptr)->object(objectIndex);
}

void* QTextDocument_ObjectForFormat(void* ptr, void* f){
	return static_cast<QTextDocument*>(ptr)->objectForFormat(*static_cast<QTextFormat*>(f));
}

int QTextDocument_PageCount(void* ptr){
	return static_cast<QTextDocument*>(ptr)->pageCount();
}

void QTextDocument_Print(void* ptr, void* printer){
	static_cast<QTextDocument*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextDocument_Redo2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "redo");
}

void QTextDocument_Redo(void* ptr, void* cursor){
	static_cast<QTextDocument*>(ptr)->redo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectRedoAvailable(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));;
}

void QTextDocument_DisconnectRedoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));;
}

void* QTextDocument_Resource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QTextDocument*>(ptr)->resource(ty, *static_cast<QUrl*>(name)));
}

int QTextDocument_Revision(void* ptr){
	return static_cast<QTextDocument*>(ptr)->revision();
}

void* QTextDocument_RootFrame(void* ptr){
	return static_cast<QTextDocument*>(ptr)->rootFrame();
}

void QTextDocument_SetDefaultCursorMoveStyle(void* ptr, int style){
	static_cast<QTextDocument*>(ptr)->setDefaultCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextDocument_SetDefaultFont(void* ptr, void* font){
	static_cast<QTextDocument*>(ptr)->setDefaultFont(*static_cast<QFont*>(font));
}

void QTextDocument_SetDefaultTextOption(void* ptr, void* option){
	static_cast<QTextDocument*>(ptr)->setDefaultTextOption(*static_cast<QTextOption*>(option));
}

void QTextDocument_SetDocumentLayout(void* ptr, void* layout){
	static_cast<QTextDocument*>(ptr)->setDocumentLayout(static_cast<QAbstractTextDocumentLayout*>(layout));
}

void QTextDocument_SetHtml(void* ptr, char* html){
	static_cast<QTextDocument*>(ptr)->setHtml(QString(html));
}

void QTextDocument_SetIndentWidth(void* ptr, double width){
	static_cast<QTextDocument*>(ptr)->setIndentWidth(static_cast<qreal>(width));
}

void QTextDocument_SetMetaInformation(void* ptr, int info, char* stri){
	static_cast<QTextDocument*>(ptr)->setMetaInformation(static_cast<QTextDocument::MetaInformation>(info), QString(stri));
}

void QTextDocument_SetPlainText(void* ptr, char* text){
	static_cast<QTextDocument*>(ptr)->setPlainText(QString(text));
}

char* QTextDocument_ToHtml(void* ptr, void* encoding){
	return static_cast<QTextDocument*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8().data();
}

char* QTextDocument_ToPlainText(void* ptr){
	return static_cast<QTextDocument*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocument_Undo2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "undo");
}

void QTextDocument_Undo(void* ptr, void* cursor){
	static_cast<QTextDocument*>(ptr)->undo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectUndoAvailable(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));;
}

void QTextDocument_DisconnectUndoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));;
}

void QTextDocument_ConnectUndoCommandAdded(void* ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));;
}

void QTextDocument_DisconnectUndoCommandAdded(void* ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));;
}

void QTextDocument_DestroyQTextDocument(void* ptr){
	static_cast<QTextDocument*>(ptr)->~QTextDocument();
}

