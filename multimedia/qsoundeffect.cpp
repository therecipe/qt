#include "qsoundeffect.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
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

int QSoundEffect_IsLoaded(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->isLoaded();
}

int QSoundEffect_LoopsRemaining(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->loopsRemaining();
}

void QSoundEffect_Play(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "play");
}

void QSoundEffect_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "stop");
}

char* QSoundEffect_QSoundEffect_SupportedMimeTypes(){
	return QSoundEffect::supportedMimeTypes().join("|").toUtf8().data();
}

QtObjectPtr QSoundEffect_NewQSoundEffect(QtObjectPtr parent){
	return new QSoundEffect(static_cast<QObject*>(parent));
}

char* QSoundEffect_Category(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->category().toUtf8().data();
}

void QSoundEffect_ConnectCategoryChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

void QSoundEffect_DisconnectCategoryChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

int QSoundEffect_IsMuted(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->isMuted();
}

int QSoundEffect_IsPlaying(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->isPlaying();
}

void QSoundEffect_ConnectLoadedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

void QSoundEffect_DisconnectLoadedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

int QSoundEffect_LoopCount(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->loopCount();
}

void QSoundEffect_ConnectLoopCountChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_DisconnectLoopCountChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_ConnectLoopsRemainingChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_DisconnectLoopsRemainingChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_ConnectMutedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_DisconnectMutedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_ConnectPlayingChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_DisconnectPlayingChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_SetCategory(QtObjectPtr ptr, char* category){
	static_cast<QSoundEffect*>(ptr)->setCategory(QString(category));
}

void QSoundEffect_SetLoopCount(QtObjectPtr ptr, int loopCount){
	static_cast<QSoundEffect*>(ptr)->setLoopCount(loopCount);
}

void QSoundEffect_SetMuted(QtObjectPtr ptr, int muted){
	static_cast<QSoundEffect*>(ptr)->setMuted(muted != 0);
}

void QSoundEffect_SetSource(QtObjectPtr ptr, char* url){
	static_cast<QSoundEffect*>(ptr)->setSource(QUrl(QString(url)));
}

char* QSoundEffect_Source(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->source().toString().toUtf8().data();
}

void QSoundEffect_ConnectSourceChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

void QSoundEffect_DisconnectSourceChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

int QSoundEffect_Status(QtObjectPtr ptr){
	return static_cast<QSoundEffect*>(ptr)->status();
}

void QSoundEffect_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

void QSoundEffect_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

void QSoundEffect_ConnectVolumeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DisconnectVolumeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DestroyQSoundEffect(QtObjectPtr ptr){
	static_cast<QSoundEffect*>(ptr)->~QSoundEffect();
}

