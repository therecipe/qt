#include "qassociativeiterable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAssociativeIterable>
#include "_cgo_export.h"

class MyQAssociativeIterable: public QAssociativeIterable {
public:
};

int QAssociativeIterable_Size(QtObjectPtr ptr){
	return static_cast<QAssociativeIterable*>(ptr)->size();
}

char* QAssociativeIterable_Value(QtObjectPtr ptr, char* key){
	return static_cast<QAssociativeIterable*>(ptr)->value(QVariant(key)).toString().toUtf8().data();
}

