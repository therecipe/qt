#include "qsequentialiterable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSequentialIterable>
#include "_cgo_export.h"

class MyQSequentialIterable: public QSequentialIterable {
public:
};

void* QSequentialIterable_At(void* ptr, int idx){
	return new QVariant(static_cast<QSequentialIterable*>(ptr)->at(idx));
}

int QSequentialIterable_CanReverseIterate(void* ptr){
	return static_cast<QSequentialIterable*>(ptr)->canReverseIterate();
}

int QSequentialIterable_Size(void* ptr){
	return static_cast<QSequentialIterable*>(ptr)->size();
}

