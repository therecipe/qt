#include "qquickitemgrabresult.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QQuickItem>
#include <QString>
#include <QVariant>
#include <QQuickItemGrabResult>
#include "_cgo_export.h"

class MyQQuickItemGrabResult: public QQuickItemGrabResult {
public:
void Signal_Ready(){callbackQQuickItemGrabResultReady(this->objectName().toUtf8().data());};
};

char* QQuickItemGrabResult_Url(QtObjectPtr ptr){
	return static_cast<QQuickItemGrabResult*>(ptr)->url().toString().toUtf8().data();
}

void QQuickItemGrabResult_ConnectReady(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

void QQuickItemGrabResult_DisconnectReady(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

int QQuickItemGrabResult_SaveToFile(QtObjectPtr ptr, char* fileName){
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString(fileName));
}

