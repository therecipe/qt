#include "qmediaaudioprobecontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QMediaAudioProbeControl>
#include "_cgo_export.h"

class MyQMediaAudioProbeControl: public QMediaAudioProbeControl {
public:
void Signal_Flush(){callbackQMediaAudioProbeControlFlush(this->objectName().toUtf8().data());};
};

void QMediaAudioProbeControl_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(void* ptr){
	static_cast<QMediaAudioProbeControl*>(ptr)->~QMediaAudioProbeControl();
}

