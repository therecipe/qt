#include "qmediaplayercontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaPlayer>
#include <QObject>
#include <QMediaContent>
#include <QIODevice>
#include <QString>
#include <QMediaPlayerControl>
#include "_cgo_export.h"

class MyQMediaPlayerControl: public QMediaPlayerControl {
public:
void Signal_AudioAvailableChanged(bool audio){callbackQMediaPlayerControlAudioAvailableChanged(this->objectName().toUtf8().data(), audio);};
void Signal_BufferStatusChanged(int progress){callbackQMediaPlayerControlBufferStatusChanged(this->objectName().toUtf8().data(), progress);};
void Signal_Error(int error, const QString & errorString){callbackQMediaPlayerControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data());};
void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status){callbackQMediaPlayerControlMediaStatusChanged(this->objectName().toUtf8().data(), status);};
void Signal_MutedChanged(bool mute){callbackQMediaPlayerControlMutedChanged(this->objectName().toUtf8().data(), mute);};
void Signal_SeekableChanged(bool seekable){callbackQMediaPlayerControlSeekableChanged(this->objectName().toUtf8().data(), seekable);};
void Signal_StateChanged(QMediaPlayer::State state){callbackQMediaPlayerControlStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_VideoAvailableChanged(bool video){callbackQMediaPlayerControlVideoAvailableChanged(this->objectName().toUtf8().data(), video);};
void Signal_VolumeChanged(int volume){callbackQMediaPlayerControlVolumeChanged(this->objectName().toUtf8().data(), volume);};
};

void QMediaPlayerControl_ConnectAudioAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));;
}

void QMediaPlayerControl_DisconnectAudioAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));;
}

int QMediaPlayerControl_BufferStatus(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->bufferStatus();
}

void QMediaPlayerControl_ConnectBufferStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));;
}

void QMediaPlayerControl_DisconnectBufferStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));;
}

void QMediaPlayerControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));;
}

void QMediaPlayerControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));;
}

int QMediaPlayerControl_IsAudioAvailable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isAudioAvailable();
}

int QMediaPlayerControl_IsMuted(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isMuted();
}

int QMediaPlayerControl_IsSeekable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isSeekable();
}

int QMediaPlayerControl_IsVideoAvailable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isVideoAvailable();
}

int QMediaPlayerControl_MediaStatus(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->mediaStatus();
}

void QMediaPlayerControl_ConnectMediaStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));;
}

void QMediaPlayerControl_DisconnectMediaStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));;
}

void* QMediaPlayerControl_MediaStream(void* ptr){
	return const_cast<QIODevice*>(static_cast<QMediaPlayerControl*>(ptr)->mediaStream());
}

void QMediaPlayerControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));;
}

void QMediaPlayerControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));;
}

void QMediaPlayerControl_Pause(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->pause();
}

void QMediaPlayerControl_Play(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->play();
}

double QMediaPlayerControl_PlaybackRate(void* ptr){
	return static_cast<double>(static_cast<QMediaPlayerControl*>(ptr)->playbackRate());
}

void QMediaPlayerControl_ConnectSeekableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));;
}

void QMediaPlayerControl_DisconnectSeekableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));;
}

void QMediaPlayerControl_SetMedia(void* ptr, void* media, void* stream){
	static_cast<QMediaPlayerControl*>(ptr)->setMedia(*static_cast<QMediaContent*>(media), static_cast<QIODevice*>(stream));
}

void QMediaPlayerControl_SetMuted(void* ptr, int mute){
	static_cast<QMediaPlayerControl*>(ptr)->setMuted(mute != 0);
}

void QMediaPlayerControl_SetPlaybackRate(void* ptr, double rate){
	static_cast<QMediaPlayerControl*>(ptr)->setPlaybackRate(static_cast<qreal>(rate));
}

void QMediaPlayerControl_SetVolume(void* ptr, int volume){
	static_cast<QMediaPlayerControl*>(ptr)->setVolume(volume);
}

void QMediaPlayerControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));;
}

void QMediaPlayerControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));;
}

void QMediaPlayerControl_Stop(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->stop();
}

void QMediaPlayerControl_ConnectVideoAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));;
}

void QMediaPlayerControl_DisconnectVideoAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));;
}

int QMediaPlayerControl_Volume(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->volume();
}

void QMediaPlayerControl_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));;
}

void QMediaPlayerControl_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));;
}

void QMediaPlayerControl_DestroyQMediaPlayerControl(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->~QMediaPlayerControl();
}

