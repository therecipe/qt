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

char* QSequentialIterable_At(QtObjectPtr ptr, int idx){
	return static_cast<QSequentialIterable*>(ptr)->at(idx).toString().toUtf8().data();
}

int QSequentialIterable_CanReverseIterate(QtObjectPtr ptr){
	return static_cast<QSequentialIterable*>(ptr)->canReverseIterate();
}

int QSequentialIterable_Size(QtObjectPtr ptr){
	return static_cast<QSequentialIterable*>(ptr)->size();
}

