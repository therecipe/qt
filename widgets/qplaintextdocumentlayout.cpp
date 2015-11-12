#include "qplaintextdocumentlayout.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QTextBlock>
#include <QString>
#include <QVariant>
#include <QPlainTextDocumentLayout>
#include "_cgo_export.h"

class MyQPlainTextDocumentLayout: public QPlainTextDocumentLayout {
public:
};

int QPlainTextDocumentLayout_CursorWidth(void* ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->cursorWidth();
}

void QPlainTextDocumentLayout_SetCursorWidth(void* ptr, int width){
	static_cast<QPlainTextDocumentLayout*>(ptr)->setCursorWidth(width);
}

void* QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(void* document){
	return new QPlainTextDocumentLayout(static_cast<QTextDocument*>(document));
}

void QPlainTextDocumentLayout_EnsureBlockLayout(void* ptr, void* block){
	static_cast<QPlainTextDocumentLayout*>(ptr)->ensureBlockLayout(*static_cast<QTextBlock*>(block));
}

int QPlainTextDocumentLayout_PageCount(void* ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->pageCount();
}

void QPlainTextDocumentLayout_RequestUpdate(void* ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->requestUpdate();
}

void QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(void* ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->~QPlainTextDocumentLayout();
}

