#include "qmediaaudioprobecontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QMediaAudioProbeControl>
#include "_cgo_export.h"

class MyQMediaAudioProbeControl: public QMediaAudioProbeControl {
public:
void Signal_Flush(){callbackQMediaAudioProbeControlFlush(this->objectName().toUtf8().data());};
};

void QMediaAudioProbeControl_ConnectFlush(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DisconnectFlush(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(QtObjectPtr ptr){
	static_cast<QMediaAudioProbeControl*>(ptr)->~QMediaAudioProbeControl();
}

