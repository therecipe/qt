#include "qstylefactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleFactory>
#include "_cgo_export.h"

class MyQStyleFactory: public QStyleFactory {
public:
};

void* QStyleFactory_QStyleFactory_Create(char* key){
	return QStyleFactory::create(QString(key));
}

char* QStyleFactory_QStyleFactory_Keys(){
	return QStyleFactory::keys().join("|").toUtf8().data();
}

