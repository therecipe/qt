#include "qassociativeiterable.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAssociativeIterable>
#include "_cgo_export.h"

class MyQAssociativeIterable: public QAssociativeIterable {
public:
};

int QAssociativeIterable_Size(void* ptr){
	return static_cast<QAssociativeIterable*>(ptr)->size();
}

void* QAssociativeIterable_Value(void* ptr, void* key){
	return new QVariant(static_cast<QAssociativeIterable*>(ptr)->value(*static_cast<QVariant*>(key)));
}

