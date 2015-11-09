#include "qabstractvideofilter.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QAbstractVideoFilter>
#include "_cgo_export.h"

class MyQAbstractVideoFilter: public QAbstractVideoFilter {
public:
void Signal_ActiveChanged(){callbackQAbstractVideoFilterActiveChanged(this->objectName().toUtf8().data());};
};

int QAbstractVideoFilter_IsActive(void* ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->isActive();
}

void QAbstractVideoFilter_SetActive(void* ptr, int v){
	static_cast<QAbstractVideoFilter*>(ptr)->setActive(v != 0);
}

void QAbstractVideoFilter_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

void QAbstractVideoFilter_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

void* QAbstractVideoFilter_CreateFilterRunnable(void* ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->createFilterRunnable();
}

