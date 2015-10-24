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

QtObjectPtr QCollatorSortKey_NewQCollatorSortKey(QtObjectPtr other){
	return new QCollatorSortKey(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QCollatorSortKey*>(ptr)->swap(*static_cast<QCollatorSortKey*>(other));
}

void QCollatorSortKey_DestroyQCollatorSortKey(QtObjectPtr ptr){
	static_cast<QCollatorSortKey*>(ptr)->~QCollatorSortKey();
}

int QCollatorSortKey_Compare(QtObjectPtr ptr, QtObjectPtr otherKey){
	return static_cast<QCollatorSortKey*>(ptr)->compare(*static_cast<QCollatorSortKey*>(otherKey));
}

