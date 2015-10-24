#include "qquicktextdocument.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQuickItem>
#include <QString>
#include <QQuickTextDocument>
#include "_cgo_export.h"

class MyQQuickTextDocument: public QQuickTextDocument {
public:
};

QtObjectPtr QQuickTextDocument_NewQQuickTextDocument(QtObjectPtr parent){
	return new QQuickTextDocument(static_cast<QQuickItem*>(parent));
}

QtObjectPtr QQuickTextDocument_TextDocument(QtObjectPtr ptr){
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

