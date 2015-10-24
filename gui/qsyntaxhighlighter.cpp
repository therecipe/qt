#include "qsyntaxhighlighter.h"
#include <QTextDocument>
#include <QMetaObject>
#include <QTextBlock>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSyntaxHighlighter>
#include "_cgo_export.h"

class MyQSyntaxHighlighter: public QSyntaxHighlighter {
public:
};

QtObjectPtr QSyntaxHighlighter_Document(QtObjectPtr ptr){
	return static_cast<QSyntaxHighlighter*>(ptr)->document();
}

void QSyntaxHighlighter_Rehighlight(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlight");
}

void QSyntaxHighlighter_RehighlightBlock(QtObjectPtr ptr, QtObjectPtr block){
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlightBlock", Q_ARG(QTextBlock, *static_cast<QTextBlock*>(block)));
}

void QSyntaxHighlighter_SetDocument(QtObjectPtr ptr, QtObjectPtr doc){
	static_cast<QSyntaxHighlighter*>(ptr)->setDocument(static_cast<QTextDocument*>(doc));
}

void QSyntaxHighlighter_DestroyQSyntaxHighlighter(QtObjectPtr ptr){
	static_cast<QSyntaxHighlighter*>(ptr)->~QSyntaxHighlighter();
}

