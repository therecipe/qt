#include "qmediagaplessplaybackcontrol.h"
#include <QModelIndex>
#include <QMediaContent>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMediaGaplessPlaybackControl>
#include "_cgo_export.h"

class MyQMediaGaplessPlaybackControl: public QMediaGaplessPlaybackControl {
public:
void Signal_AdvancedToNextMedia(){callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(this->objectName().toUtf8().data());};
};

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(void* ptr){
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(void* ptr){
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

double QMediaGaplessPlaybackControl_CrossfadeTime(void* ptr){
	return static_cast<double>(static_cast<QMediaGaplessPlaybackControl*>(ptr)->crossfadeTime());
}

int QMediaGaplessPlaybackControl_IsCrossfadeSupported(void* ptr){
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->isCrossfadeSupported();
}

void QMediaGaplessPlaybackControl_SetCrossfadeTime(void* ptr, double crossfadeTime){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setCrossfadeTime(static_cast<qreal>(crossfadeTime));
}

void QMediaGaplessPlaybackControl_SetNextMedia(void* ptr, void* media){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setNextMedia(*static_cast<QMediaContent*>(media));
}

void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(void* ptr){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->~QMediaGaplessPlaybackControl();
}

