#include "qquickitemgrabresult.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQuickItem>
#include <QQuickItemGrabResult>
#include "_cgo_export.h"

class MyQQuickItemGrabResult: public QQuickItemGrabResult {
public:
void Signal_Ready(){callbackQQuickItemGrabResultReady(this->objectName().toUtf8().data());};
};

void QQuickItemGrabResult_ConnectReady(void* ptr){
	QObject::connect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

void QQuickItemGrabResult_DisconnectReady(void* ptr){
	QObject::disconnect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

int QQuickItemGrabResult_SaveToFile(void* ptr, char* fileName){
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString(fileName));
}

