#include "qmediaplayer.h"
#include <QModelIndex>
#include <QObject>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMediaPlaylist>
#include <QAbstractVideoSurface>
#include <QMediaContent>
#include <QMetaObject>
#include <QMediaPlayer>
#include "_cgo_export.h"

class MyQMediaPlayer: public QMediaPlayer {
public:
void Signal_AudioAvailableChanged(bool available){callbackQMediaPlayerAudioAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_BufferStatusChanged(int percentFilled){callbackQMediaPlayerBufferStatusChanged(this->objectName().toUtf8().data(), percentFilled);};
void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status){callbackQMediaPlayerMediaStatusChanged(this->objectName().toUtf8().data(), status);};
void Signal_MutedChanged(bool muted){callbackQMediaPlayerMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_SeekableChanged(bool seekable){callbackQMediaPlayerSeekableChanged(this->objectName().toUtf8().data(), seekable);};
void Signal_StateChanged(QMediaPlayer::State state){callbackQMediaPlayerStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_VideoAvailableChanged(bool videoAvailable){callbackQMediaPlayerVideoAvailableChanged(this->objectName().toUtf8().data(), videoAvailable);};
void Signal_VolumeChanged(int volume){callbackQMediaPlayerVolumeChanged(this->objectName().toUtf8().data(), volume);};
};

int QMediaPlayer_BufferStatus(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->bufferStatus();
}

char* QMediaPlayer_ErrorString(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlayer_IsAudioAvailable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isAudioAvailable();
}

int QMediaPlayer_IsMuted(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isMuted();
}

int QMediaPlayer_IsSeekable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isSeekable();
}

int QMediaPlayer_IsVideoAvailable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isVideoAvailable();
}

int QMediaPlayer_MediaStatus(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->mediaStatus();
}

double QMediaPlayer_PlaybackRate(void* ptr){
	return static_cast<double>(static_cast<QMediaPlayer*>(ptr)->playbackRate());
}

void* QMediaPlayer_Playlist(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->playlist();
}

void QMediaPlayer_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QMediaPlayer_SetPlaybackRate(void* ptr, double rate){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaybackRate", Q_ARG(qreal, static_cast<qreal>(rate)));
}

void QMediaPlayer_SetPlaylist(void* ptr, void* playlist){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaylist", Q_ARG(QMediaPlaylist*, static_cast<QMediaPlaylist*>(playlist)));
}

void QMediaPlayer_SetVolume(void* ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QMediaPlayer_Volume(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->volume();
}

void* QMediaPlayer_NewQMediaPlayer(void* parent, int flags){
	return new QMediaPlayer(static_cast<QObject*>(parent), static_cast<QMediaPlayer::Flag>(flags));
}

void QMediaPlayer_ConnectAudioAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_DisconnectAudioAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_ConnectBufferStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

void QMediaPlayer_DisconnectBufferStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

int QMediaPlayer_Error(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->error();
}

void QMediaPlayer_ConnectMediaStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

void QMediaPlayer_DisconnectMediaStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

void* QMediaPlayer_MediaStream(void* ptr){
	return const_cast<QIODevice*>(static_cast<QMediaPlayer*>(ptr)->mediaStream());
}

void QMediaPlayer_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "pause");
}

void QMediaPlayer_Play(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "play");
}

void QMediaPlayer_ConnectSeekableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_DisconnectSeekableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_SetMedia(void* ptr, void* media, void* stream){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMedia", Q_ARG(QMediaContent, *static_cast<QMediaContent*>(media)), Q_ARG(QIODevice*, static_cast<QIODevice*>(stream)));
}

void QMediaPlayer_SetVideoOutput3(void* ptr, void* surface){
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QAbstractVideoSurface*>(surface));
}

void QMediaPlayer_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "stop");
}

void QMediaPlayer_ConnectVideoAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_DisconnectVideoAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DestroyQMediaPlayer(void* ptr){
	static_cast<QMediaPlayer*>(ptr)->~QMediaPlayer();
}

