#include "qbitmap.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBitmap>
#include "_cgo_export.h"

class MyQBitmap: public QBitmap {
public:
};

void QBitmap_Clear(QtObjectPtr ptr){
	static_cast<QBitmap*>(ptr)->clear();
}

void QBitmap_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QBitmap*>(ptr)->swap(*static_cast<QBitmap*>(other));
}

void QBitmap_DestroyQBitmap(QtObjectPtr ptr){
	static_cast<QBitmap*>(ptr)->~QBitmap();
}

