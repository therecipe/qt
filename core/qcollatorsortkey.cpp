#include "qcollatorsortkey.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCollator>
#include <QCollatorSortKey>
#include "_cgo_export.h"

class MyQCollatorSortKey: public QCollatorSortKey {
public:
};

void* QCollatorSortKey_NewQCollatorSortKey(void* other){
	return new QCollatorSortKey(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_Swap(void* ptr, void* other){
	static_cast<QCollatorSortKey*>(ptr)->swap(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_DestroyQCollatorSortKey(void* ptr){
	static_cast<QCollatorSortKey*>(ptr)->~QCollatorSortKey();
}

int QCollatorSortKey_Compare(void* ptr, void* otherKey){
	return static_cast<QCollatorSortKey*>(ptr)->compare(*static_cast<QCollatorSortKey*>(otherKey));
}

