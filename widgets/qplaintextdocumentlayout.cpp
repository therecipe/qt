#include "qplaintextdocumentlayout.h"
#include <QModelIndex>
#include <QTextDocument>
#include <QTextBlock>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPlainTextDocumentLayout>
#include "_cgo_export.h"

class MyQPlainTextDocumentLayout: public QPlainTextDocumentLayout {
public:
};

int QPlainTextDocumentLayout_CursorWidth(QtObjectPtr ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->cursorWidth();
}

void QPlainTextDocumentLayout_SetCursorWidth(QtObjectPtr ptr, int width){
	static_cast<QPlainTextDocumentLayout*>(ptr)->setCursorWidth(width);
}

QtObjectPtr QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(QtObjectPtr document){
	return new QPlainTextDocumentLayout(static_cast<QTextDocument*>(document));
}

void QPlainTextDocumentLayout_EnsureBlockLayout(QtObjectPtr ptr, QtObjectPtr block){
	static_cast<QPlainTextDocumentLayout*>(ptr)->ensureBlockLayout(*static_cast<QTextBlock*>(block));
}

int QPlainTextDocumentLayout_PageCount(QtObjectPtr ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->pageCount();
}

void QPlainTextDocumentLayout_RequestUpdate(QtObjectPtr ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->requestUpdate();
}

void QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(QtObjectPtr ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->~QPlainTextDocumentLayout();
}

