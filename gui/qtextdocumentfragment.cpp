#include "qtextdocumentfragment.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QTextCursor>
#include <QByteArray>
#include <QString>
#include <QTextDocumentFragment>
#include "_cgo_export.h"

class MyQTextDocumentFragment: public QTextDocumentFragment {
public:
};

QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment4(QtObjectPtr other){
	return new QTextDocumentFragment(*static_cast<QTextDocumentFragment*>(other));
}

QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment(){
	return new QTextDocumentFragment();
}

QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment3(QtObjectPtr cursor){
	return new QTextDocumentFragment(*static_cast<QTextCursor*>(cursor));
}

QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment2(QtObjectPtr document){
	return new QTextDocumentFragment(static_cast<QTextDocument*>(document));
}

int QTextDocumentFragment_IsEmpty(QtObjectPtr ptr){
	return static_cast<QTextDocumentFragment*>(ptr)->isEmpty();
}

char* QTextDocumentFragment_ToHtml(QtObjectPtr ptr, QtObjectPtr encoding){
	return static_cast<QTextDocumentFragment*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8().data();
}

char* QTextDocumentFragment_ToPlainText(QtObjectPtr ptr){
	return static_cast<QTextDocumentFragment*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocumentFragment_DestroyQTextDocumentFragment(QtObjectPtr ptr){
	static_cast<QTextDocumentFragment*>(ptr)->~QTextDocumentFragment();
}

