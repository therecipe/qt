#include "qquicktextdocument.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQuickItem>
#include <QQuickTextDocument>
#include "_cgo_export.h"

class MyQQuickTextDocument: public QQuickTextDocument {
public:
};

void* QQuickTextDocument_NewQQuickTextDocument(void* parent){
	return new QQuickTextDocument(static_cast<QQuickItem*>(parent));
}

void* QQuickTextDocument_TextDocument(void* ptr){
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

