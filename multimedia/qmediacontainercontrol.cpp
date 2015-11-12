#include "qmediacontainercontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QMediaContainerControl>
#include "_cgo_export.h"

class MyQMediaContainerControl: public QMediaContainerControl {
public:
};

char* QMediaContainerControl_ContainerDescription(void* ptr, char* format){
	return static_cast<QMediaContainerControl*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaContainerControl_ContainerFormat(void* ptr){
	return static_cast<QMediaContainerControl*>(ptr)->containerFormat().toUtf8().data();
}

void QMediaContainerControl_SetContainerFormat(void* ptr, char* format){
	static_cast<QMediaContainerControl*>(ptr)->setContainerFormat(QString(format));
}

char* QMediaContainerControl_SupportedContainers(void* ptr){
	return static_cast<QMediaContainerControl*>(ptr)->supportedContainers().join("|").toUtf8().data();
}

void QMediaContainerControl_DestroyQMediaContainerControl(void* ptr){
	static_cast<QMediaContainerControl*>(ptr)->~QMediaContainerControl();
}

