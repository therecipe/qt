#include "qabstractvideofilter.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractVideoFilter>
#include "_cgo_export.h"

class MyQAbstractVideoFilter: public QAbstractVideoFilter {
public:
void Signal_ActiveChanged(){callbackQAbstractVideoFilterActiveChanged(this->objectName().toUtf8().data());};
};

int QAbstractVideoFilter_IsActive(QtObjectPtr ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->isActive();
}

void QAbstractVideoFilter_SetActive(QtObjectPtr ptr, int v){
	static_cast<QAbstractVideoFilter*>(ptr)->setActive(v != 0);
}

void QAbstractVideoFilter_ConnectActiveChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

void QAbstractVideoFilter_DisconnectActiveChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

QtObjectPtr QAbstractVideoFilter_CreateFilterRunnable(QtObjectPtr ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->createFilterRunnable();
}

