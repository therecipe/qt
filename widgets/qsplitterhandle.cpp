#include "qsplitterhandle.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSplitter>
#include <QSplitterHandle>
#include "_cgo_export.h"

class MyQSplitterHandle: public QSplitterHandle {
public:
};

void* QSplitterHandle_NewQSplitterHandle(int orientation, void* parent){
	return new QSplitterHandle(static_cast<Qt::Orientation>(orientation), static_cast<QSplitter*>(parent));
}

int QSplitterHandle_OpaqueResize(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->opaqueResize();
}

int QSplitterHandle_Orientation(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->orientation();
}

void QSplitterHandle_SetOrientation(void* ptr, int orientation){
	static_cast<QSplitterHandle*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void* QSplitterHandle_Splitter(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->splitter();
}

void QSplitterHandle_DestroyQSplitterHandle(void* ptr){
	static_cast<QSplitterHandle*>(ptr)->~QSplitterHandle();
}

