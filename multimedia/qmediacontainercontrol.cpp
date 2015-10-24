#include "qmediacontainercontrol.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMediaContainerControl>
#include "_cgo_export.h"

class MyQMediaContainerControl: public QMediaContainerControl {
public:
};

char* QMediaContainerControl_ContainerDescription(QtObjectPtr ptr, char* format){
	return static_cast<QMediaContainerControl*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaContainerControl_ContainerFormat(QtObjectPtr ptr){
	return static_cast<QMediaContainerControl*>(ptr)->containerFormat().toUtf8().data();
}

void QMediaContainerControl_SetContainerFormat(QtObjectPtr ptr, char* format){
	static_cast<QMediaContainerControl*>(ptr)->setContainerFormat(QString(format));
}

char* QMediaContainerControl_SupportedContainers(QtObjectPtr ptr){
	return static_cast<QMediaContainerControl*>(ptr)->supportedContainers().join("|").toUtf8().data();
}

void QMediaContainerControl_DestroyQMediaContainerControl(QtObjectPtr ptr){
	static_cast<QMediaContainerControl*>(ptr)->~QMediaContainerControl();
}

