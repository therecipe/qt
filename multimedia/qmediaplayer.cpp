#include "qmediaplayer.h"
#include <QUrl>
#include <QMetaObject>
#include <QMediaContent>
#include <QIODevice>
#include <QAbstractVideoSurface>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMediaPlaylist>
#include <QObject>
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

int QMediaPlayer_BufferStatus(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->bufferStatus();
}

char* QMediaPlayer_ErrorString(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlayer_IsAudioAvailable(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->isAudioAvailable();
}

int QMediaPlayer_IsMuted(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->isMuted();
}

int QMediaPlayer_IsSeekable(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->isSeekable();
}

int QMediaPlayer_IsVideoAvailable(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->isVideoAvailable();
}

int QMediaPlayer_MediaStatus(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->mediaStatus();
}

QtObjectPtr QMediaPlayer_Playlist(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->playlist();
}

void QMediaPlayer_SetMuted(QtObjectPtr ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QMediaPlayer_SetPlaylist(QtObjectPtr ptr, QtObjectPtr playlist){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaylist", Q_ARG(QMediaPlaylist*, static_cast<QMediaPlaylist*>(playlist)));
}

void QMediaPlayer_SetVolume(QtObjectPtr ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QMediaPlayer_Volume(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->volume();
}

QtObjectPtr QMediaPlayer_NewQMediaPlayer(QtObjectPtr parent, int flags){
	return new QMediaPlayer(static_cast<QObject*>(parent), static_cast<QMediaPlayer::Flag>(flags));
}

void QMediaPlayer_ConnectAudioAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_DisconnectAudioAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_ConnectBufferStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

void QMediaPlayer_DisconnectBufferStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

int QMediaPlayer_Error(QtObjectPtr ptr){
	return static_cast<QMediaPlayer*>(ptr)->error();
}

void QMediaPlayer_ConnectMediaStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

void QMediaPlayer_DisconnectMediaStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

QtObjectPtr QMediaPlayer_MediaStream(QtObjectPtr ptr){
	return const_cast<QIODevice*>(static_cast<QMediaPlayer*>(ptr)->mediaStream());
}

void QMediaPlayer_ConnectMutedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_DisconnectMutedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_Pause(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "pause");
}

void QMediaPlayer_Play(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "play");
}

void QMediaPlayer_ConnectSeekableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_DisconnectSeekableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_SetMedia(QtObjectPtr ptr, QtObjectPtr media, QtObjectPtr stream){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMedia", Q_ARG(QMediaContent, *static_cast<QMediaContent*>(media)), Q_ARG(QIODevice*, static_cast<QIODevice*>(stream)));
}

void QMediaPlayer_SetVideoOutput3(QtObjectPtr ptr, QtObjectPtr surface){
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QAbstractVideoSurface*>(surface));
}

void QMediaPlayer_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "stop");
}

void QMediaPlayer_ConnectVideoAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_DisconnectVideoAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_ConnectVolumeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DisconnectVolumeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DestroyQMediaPlayer(QtObjectPtr ptr){
	static_cast<QMediaPlayer*>(ptr)->~QMediaPlayer();
}

