#include "qtextdocument.h"
#include <QByteArray>
#include <QPagedPaintDevice>
#include <QRect>
#include <QObject>
#include <QFont>
#include <QModelIndex>
#include <QAbstractTextDocumentLayout>
#include <QTextFormat>
#include <QSizeF>
#include <QTextOption>
#include <QSize>
#include <QString>
#include <QTextCursor>
#include <QRectF>
#include <QPainter>
#include <QUrl>
#include <QMetaObject>
#include <QVariant>
#include <QTextDocument>
#include "_cgo_export.h"

class MyQTextDocument: public QTextDocument {
public:
void Signal_BaseUrlChanged(const QUrl & url){callbackQTextDocumentBaseUrlChanged(this->objectName().toUtf8().data(), url.toString().toUtf8().data());};
void Signal_BlockCountChanged(int newBlockCount){callbackQTextDocumentBlockCountChanged(this->objectName().toUtf8().data(), newBlockCount);};
void Signal_ContentsChange(int position, int charsRemoved, int charsAdded){callbackQTextDocumentContentsChange(this->objectName().toUtf8().data(), position, charsRemoved, charsAdded);};
void Signal_ContentsChanged(){callbackQTextDocumentContentsChanged(this->objectName().toUtf8().data());};
void Signal_DocumentLayoutChanged(){callbackQTextDocumentDocumentLayoutChanged(this->objectName().toUtf8().data());};
void Signal_ModificationChanged(bool changed){callbackQTextDocumentModificationChanged(this->objectName().toUtf8().data(), changed);};
void Signal_RedoAvailable(bool available){callbackQTextDocumentRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_UndoAvailable(bool available){callbackQTextDocumentUndoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_UndoCommandAdded(){callbackQTextDocumentUndoCommandAdded(this->objectName().toUtf8().data());};
};

char* QTextDocument_BaseUrl(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->baseUrl().toString().toUtf8().data();
}

int QTextDocument_BlockCount(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->blockCount();
}

char* QTextDocument_DefaultStyleSheet(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->defaultStyleSheet().toUtf8().data();
}

int QTextDocument_IsModified(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->isModified();
}

int QTextDocument_IsUndoRedoEnabled(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->isUndoRedoEnabled();
}

void QTextDocument_MarkContentsDirty(QtObjectPtr ptr, int position, int length){
	static_cast<QTextDocument*>(ptr)->markContentsDirty(position, length);
}

int QTextDocument_MaximumBlockCount(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->maximumBlockCount();
}

void QTextDocument_SetBaseUrl(QtObjectPtr ptr, char* url){
	static_cast<QTextDocument*>(ptr)->setBaseUrl(QUrl(QString(url)));
}

void QTextDocument_SetDefaultStyleSheet(QtObjectPtr ptr, char* sheet){
	static_cast<QTextDocument*>(ptr)->setDefaultStyleSheet(QString(sheet));
}

void QTextDocument_SetMaximumBlockCount(QtObjectPtr ptr, int maximum){
	static_cast<QTextDocument*>(ptr)->setMaximumBlockCount(maximum);
}

void QTextDocument_SetModified(QtObjectPtr ptr, int m){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "setModified", Q_ARG(bool, m != 0));
}

void QTextDocument_SetPageSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QTextDocument*>(ptr)->setPageSize(*static_cast<QSizeF*>(size));
}

void QTextDocument_SetUndoRedoEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTextDocument*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QTextDocument_SetUseDesignMetrics(QtObjectPtr ptr, int b){
	static_cast<QTextDocument*>(ptr)->setUseDesignMetrics(b != 0);
}

int QTextDocument_UseDesignMetrics(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->useDesignMetrics();
}

QtObjectPtr QTextDocument_NewQTextDocument(QtObjectPtr parent){
	return new QTextDocument(static_cast<QObject*>(parent));
}

QtObjectPtr QTextDocument_NewQTextDocument2(char* text, QtObjectPtr parent){
	return new QTextDocument(QString(text), static_cast<QObject*>(parent));
}

void QTextDocument_AddResource(QtObjectPtr ptr, int ty, char* name, char* resource){
	static_cast<QTextDocument*>(ptr)->addResource(ty, QUrl(QString(name)), QVariant(resource));
}

void QTextDocument_AdjustSize(QtObjectPtr ptr){
	static_cast<QTextDocument*>(ptr)->adjustSize();
}

int QTextDocument_AvailableRedoSteps(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->availableRedoSteps();
}

int QTextDocument_AvailableUndoSteps(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->availableUndoSteps();
}

void QTextDocument_ConnectBaseUrlChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QUrl &)>(&QTextDocument::baseUrlChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QUrl &)>(&MyQTextDocument::Signal_BaseUrlChanged));;
}

void QTextDocument_DisconnectBaseUrlChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QUrl &)>(&QTextDocument::baseUrlChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QUrl &)>(&MyQTextDocument::Signal_BaseUrlChanged));;
}

void QTextDocument_ConnectBlockCountChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));;
}

void QTextDocument_DisconnectBlockCountChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));;
}

int QTextDocument_CharacterCount(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->characterCount();
}

void QTextDocument_Clear(QtObjectPtr ptr){
	static_cast<QTextDocument*>(ptr)->clear();
}

void QTextDocument_ClearUndoRedoStacks(QtObjectPtr ptr, int stacksToClear){
	static_cast<QTextDocument*>(ptr)->clearUndoRedoStacks(static_cast<QTextDocument::Stacks>(stacksToClear));
}

QtObjectPtr QTextDocument_Clone(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QTextDocument*>(ptr)->clone(static_cast<QObject*>(parent));
}

void QTextDocument_ConnectContentsChange(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));;
}

void QTextDocument_DisconnectContentsChange(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));;
}

void QTextDocument_ConnectContentsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));;
}

void QTextDocument_DisconnectContentsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));;
}

int QTextDocument_DefaultCursorMoveStyle(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->defaultCursorMoveStyle();
}

QtObjectPtr QTextDocument_DocumentLayout(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->documentLayout();
}

void QTextDocument_ConnectDocumentLayoutChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));;
}

void QTextDocument_DisconnectDocumentLayoutChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));;
}

void QTextDocument_DrawContents(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr rect){
	static_cast<QTextDocument*>(ptr)->drawContents(static_cast<QPainter*>(p), *static_cast<QRectF*>(rect));
}

int QTextDocument_IsEmpty(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->isEmpty();
}

int QTextDocument_IsRedoAvailable(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->isRedoAvailable();
}

int QTextDocument_IsUndoAvailable(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->isUndoAvailable();
}

int QTextDocument_LineCount(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->lineCount();
}

char* QTextDocument_MetaInformation(QtObjectPtr ptr, int info){
	return static_cast<QTextDocument*>(ptr)->metaInformation(static_cast<QTextDocument::MetaInformation>(info)).toUtf8().data();
}

void QTextDocument_ConnectModificationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));;
}

void QTextDocument_DisconnectModificationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));;
}

QtObjectPtr QTextDocument_Object(QtObjectPtr ptr, int objectIndex){
	return static_cast<QTextDocument*>(ptr)->object(objectIndex);
}

QtObjectPtr QTextDocument_ObjectForFormat(QtObjectPtr ptr, QtObjectPtr f){
	return static_cast<QTextDocument*>(ptr)->objectForFormat(*static_cast<QTextFormat*>(f));
}

int QTextDocument_PageCount(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->pageCount();
}

void QTextDocument_Print(QtObjectPtr ptr, QtObjectPtr printer){
	static_cast<QTextDocument*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextDocument_Redo2(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "redo");
}

void QTextDocument_Redo(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QTextDocument*>(ptr)->redo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectRedoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));;
}

void QTextDocument_DisconnectRedoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));;
}

char* QTextDocument_Resource(QtObjectPtr ptr, int ty, char* name){
	return static_cast<QTextDocument*>(ptr)->resource(ty, QUrl(QString(name))).toString().toUtf8().data();
}

int QTextDocument_Revision(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->revision();
}

QtObjectPtr QTextDocument_RootFrame(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->rootFrame();
}

void QTextDocument_SetDefaultCursorMoveStyle(QtObjectPtr ptr, int style){
	static_cast<QTextDocument*>(ptr)->setDefaultCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextDocument_SetDefaultFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QTextDocument*>(ptr)->setDefaultFont(*static_cast<QFont*>(font));
}

void QTextDocument_SetDefaultTextOption(QtObjectPtr ptr, QtObjectPtr option){
	static_cast<QTextDocument*>(ptr)->setDefaultTextOption(*static_cast<QTextOption*>(option));
}

void QTextDocument_SetDocumentLayout(QtObjectPtr ptr, QtObjectPtr layout){
	static_cast<QTextDocument*>(ptr)->setDocumentLayout(static_cast<QAbstractTextDocumentLayout*>(layout));
}

void QTextDocument_SetHtml(QtObjectPtr ptr, char* html){
	static_cast<QTextDocument*>(ptr)->setHtml(QString(html));
}

void QTextDocument_SetMetaInformation(QtObjectPtr ptr, int info, char* stri){
	static_cast<QTextDocument*>(ptr)->setMetaInformation(static_cast<QTextDocument::MetaInformation>(info), QString(stri));
}

void QTextDocument_SetPlainText(QtObjectPtr ptr, char* text){
	static_cast<QTextDocument*>(ptr)->setPlainText(QString(text));
}

char* QTextDocument_ToHtml(QtObjectPtr ptr, QtObjectPtr encoding){
	return static_cast<QTextDocument*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8().data();
}

char* QTextDocument_ToPlainText(QtObjectPtr ptr){
	return static_cast<QTextDocument*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocument_Undo2(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "undo");
}

void QTextDocument_Undo(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QTextDocument*>(ptr)->undo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectUndoAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));;
}

void QTextDocument_DisconnectUndoAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));;
}

void QTextDocument_ConnectUndoCommandAdded(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));;
}

void QTextDocument_DisconnectUndoCommandAdded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));;
}

void QTextDocument_DestroyQTextDocument(QtObjectPtr ptr){
	static_cast<QTextDocument*>(ptr)->~QTextDocument();
}

