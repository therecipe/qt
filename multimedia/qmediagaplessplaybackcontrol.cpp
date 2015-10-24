#include "qmediagaplessplaybackcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMediaContent>
#include <QMediaGaplessPlaybackControl>
#include "_cgo_export.h"

class MyQMediaGaplessPlaybackControl: public QMediaGaplessPlaybackControl {
public:
void Signal_AdvancedToNextMedia(){callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(this->objectName().toUtf8().data());};
};

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

int QMediaGaplessPlaybackControl_IsCrossfadeSupported(QtObjectPtr ptr){
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->isCrossfadeSupported();
}

void QMediaGaplessPlaybackControl_SetNextMedia(QtObjectPtr ptr, QtObjectPtr media){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setNextMedia(*static_cast<QMediaContent*>(media));
}

void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(QtObjectPtr ptr){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->~QMediaGaplessPlaybackControl();
}

