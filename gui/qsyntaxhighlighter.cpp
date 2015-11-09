#include "qsyntaxhighlighter.h"
#include <QTextBlock>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QSyntaxHighlighter>
#include "_cgo_export.h"

class MyQSyntaxHighlighter: public QSyntaxHighlighter {
public:
};

void* QSyntaxHighlighter_Document(void* ptr){
	return static_cast<QSyntaxHighlighter*>(ptr)->document();
}

void QSyntaxHighlighter_Rehighlight(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlight");
}

void QSyntaxHighlighter_RehighlightBlock(void* ptr, void* block){
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlightBlock", Q_ARG(QTextBlock, *static_cast<QTextBlock*>(block)));
}

void QSyntaxHighlighter_SetDocument(void* ptr, void* doc){
	static_cast<QSyntaxHighlighter*>(ptr)->setDocument(static_cast<QTextDocument*>(doc));
}

void QSyntaxHighlighter_DestroyQSyntaxHighlighter(void* ptr){
	static_cast<QSyntaxHighlighter*>(ptr)->~QSyntaxHighlighter();
}

