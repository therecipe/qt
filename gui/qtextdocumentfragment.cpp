#include "qtextdocumentfragment.h"
#include <QByteArray>
#include <QTextCursor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QTextDocumentFragment>
#include "_cgo_export.h"

class MyQTextDocumentFragment: public QTextDocumentFragment {
public:
};

void* QTextDocumentFragment_NewQTextDocumentFragment4(void* other){
	return new QTextDocumentFragment(*static_cast<QTextDocumentFragment*>(other));
}

void* QTextDocumentFragment_NewQTextDocumentFragment(){
	return new QTextDocumentFragment();
}

void* QTextDocumentFragment_NewQTextDocumentFragment3(void* cursor){
	return new QTextDocumentFragment(*static_cast<QTextCursor*>(cursor));
}

void* QTextDocumentFragment_NewQTextDocumentFragment2(void* document){
	return new QTextDocumentFragment(static_cast<QTextDocument*>(document));
}

int QTextDocumentFragment_IsEmpty(void* ptr){
	return static_cast<QTextDocumentFragment*>(ptr)->isEmpty();
}

char* QTextDocumentFragment_ToHtml(void* ptr, void* encoding){
	return static_cast<QTextDocumentFragment*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8().data();
}

char* QTextDocumentFragment_ToPlainText(void* ptr){
	return static_cast<QTextDocumentFragment*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocumentFragment_DestroyQTextDocumentFragment(void* ptr){
	static_cast<QTextDocumentFragment*>(ptr)->~QTextDocumentFragment();
}

