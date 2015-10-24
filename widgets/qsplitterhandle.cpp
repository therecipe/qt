#include "qsplitterhandle.h"
#include <QUrl>
#include <QModelIndex>
#include <QSplitter>
#include <QString>
#include <QVariant>
#include <QSplitterHandle>
#include "_cgo_export.h"

class MyQSplitterHandle: public QSplitterHandle {
public:
};

QtObjectPtr QSplitterHandle_NewQSplitterHandle(int orientation, QtObjectPtr parent){
	return new QSplitterHandle(static_cast<Qt::Orientation>(orientation), static_cast<QSplitter*>(parent));
}

int QSplitterHandle_OpaqueResize(QtObjectPtr ptr){
	return static_cast<QSplitterHandle*>(ptr)->opaqueResize();
}

int QSplitterHandle_Orientation(QtObjectPtr ptr){
	return static_cast<QSplitterHandle*>(ptr)->orientation();
}

void QSplitterHandle_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QSplitterHandle*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

QtObjectPtr QSplitterHandle_Splitter(QtObjectPtr ptr){
	return static_cast<QSplitterHandle*>(ptr)->splitter();
}

void QSplitterHandle_DestroyQSplitterHandle(QtObjectPtr ptr){
	static_cast<QSplitterHandle*>(ptr)->~QSplitterHandle();
}

