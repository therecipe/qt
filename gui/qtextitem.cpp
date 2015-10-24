#include "qtextitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextItem>
#include "_cgo_export.h"

class MyQTextItem: public QTextItem {
public:
};

int QTextItem_RenderFlags(QtObjectPtr ptr){
	return static_cast<QTextItem*>(ptr)->renderFlags();
}

char* QTextItem_Text(QtObjectPtr ptr){
	return static_cast<QTextItem*>(ptr)->text().toUtf8().data();
}

