#include "qsoundeffect.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QSound>
#include <QString>
#include <QVariant>
#include <QSoundEffect>
#include "_cgo_export.h"

class MyQSoundEffect: public QSoundEffect {
public:
void Signal_CategoryChanged(){callbackQSoundEffectCategoryChanged(this->objectName().toUtf8().data());};
void Signal_LoadedChanged(){callbackQSoundEffectLoadedChanged(this->objectName().toUtf8().data());};
void Signal_LoopCountChanged(){callbackQSoundEffectLoopCountChanged(this->objectName().toUtf8().data());};
void Signal_LoopsRemainingChanged(){callbackQSoundEffectLoopsRemainingChanged(this->objectName().toUtf8().data());};
void Signal_MutedChanged(){callbackQSoundEffectMutedChanged(this->objectName().toUtf8().data());};
void Signal_PlayingChanged(){callbackQSoundEffectPlayingChanged(this->objectName().toUtf8().data());};
void Signal_SourceChanged(){callbackQSoundEffectSourceChanged(this->objectName().toUtf8().data());};
void Signal_StatusChanged(){callbackQSoundEffectStatusChanged(this->objectName().toUtf8().data());};
void Signal_VolumeChanged(){callbackQSoundEffectVolumeChanged(this->objectName().toUtf8().data());};
};

int QSoundEffect_IsLoaded(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isLoaded();
}

int QSoundEffect_LoopsRemaining(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->loopsRemaining();
}

void QSoundEffect_Play(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "play");
}

void QSoundEffect_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "stop");
}

char* QSoundEffect_QSoundEffect_SupportedMimeTypes(){
	return QSoundEffect::supportedMimeTypes().join("|").toUtf8().data();
}

void* QSoundEffect_NewQSoundEffect(void* parent){
	return new QSoundEffect(static_cast<QObject*>(parent));
}

char* QSoundEffect_Category(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->category().toUtf8().data();
}

void QSoundEffect_ConnectCategoryChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

void QSoundEffect_DisconnectCategoryChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

int QSoundEffect_IsMuted(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isMuted();
}

int QSoundEffect_IsPlaying(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isPlaying();
}

void QSoundEffect_ConnectLoadedChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

void QSoundEffect_DisconnectLoadedChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

int QSoundEffect_LoopCount(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->loopCount();
}

void QSoundEffect_ConnectLoopCountChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_DisconnectLoopCountChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_ConnectLoopsRemainingChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_DisconnectLoopsRemainingChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_ConnectPlayingChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_DisconnectPlayingChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_SetCategory(void* ptr, char* category){
	static_cast<QSoundEffect*>(ptr)->setCategory(QString(category));
}

void QSoundEffect_SetLoopCount(void* ptr, int loopCount){
	static_cast<QSoundEffect*>(ptr)->setLoopCount(loopCount);
}

void QSoundEffect_SetMuted(void* ptr, int muted){
	static_cast<QSoundEffect*>(ptr)->setMuted(muted != 0);
}

void QSoundEffect_SetSource(void* ptr, void* url){
	static_cast<QSoundEffect*>(ptr)->setSource(*static_cast<QUrl*>(url));
}

void QSoundEffect_SetVolume(void* ptr, double volume){
	static_cast<QSoundEffect*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void QSoundEffect_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

void QSoundEffect_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

int QSoundEffect_Status(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->status();
}

void QSoundEffect_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

void QSoundEffect_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

double QSoundEffect_Volume(void* ptr){
	return static_cast<double>(static_cast<QSoundEffect*>(ptr)->volume());
}

void QSoundEffect_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DestroyQSoundEffect(void* ptr){
	static_cast<QSoundEffect*>(ptr)->~QSoundEffect();
}

