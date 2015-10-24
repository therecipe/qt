#include "qmediavideoprobecontrol.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaVideoProbeControl>
#include "_cgo_export.h"

class MyQMediaVideoProbeControl: public QMediaVideoProbeControl {
public:
void Signal_Flush(){callbackQMediaVideoProbeControlFlush(this->objectName().toUtf8().data());};
};

void QMediaVideoProbeControl_ConnectFlush(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DisconnectFlush(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(QtObjectPtr ptr){
	static_cast<QMediaVideoProbeControl*>(ptr)->~QMediaVideoProbeControl();
}

