#include "qmediavideoprobecontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMediaVideoProbeControl>
#include "_cgo_export.h"

class MyQMediaVideoProbeControl: public QMediaVideoProbeControl {
public:
void Signal_Flush(){callbackQMediaVideoProbeControlFlush(this->objectName().toUtf8().data());};
};

void QMediaVideoProbeControl_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(void* ptr){
	static_cast<QMediaVideoProbeControl*>(ptr)->~QMediaVideoProbeControl();
}

