#include "qmediaplaylist.h"
#include <QNetworkRequest>
#include <QIODevice>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QMediaContent>
#include <QString>
#include <QUrl>
#include <QMetaObject>
#include <QMediaPlaylist>
#include "_cgo_export.h"

class MyQMediaPlaylist: public QMediaPlaylist {
public:
void Signal_CurrentIndexChanged(int position){callbackQMediaPlaylistCurrentIndexChanged(this->objectName().toUtf8().data(), position);};
void Signal_LoadFailed(){callbackQMediaPlaylistLoadFailed(this->objectName().toUtf8().data());};
void Signal_Loaded(){callbackQMediaPlaylistLoaded(this->objectName().toUtf8().data());};
void Signal_MediaAboutToBeInserted(int start, int end){callbackQMediaPlaylistMediaAboutToBeInserted(this->objectName().toUtf8().data(), start, end);};
void Signal_MediaAboutToBeRemoved(int start, int end){callbackQMediaPlaylistMediaAboutToBeRemoved(this->objectName().toUtf8().data(), start, end);};
void Signal_MediaChanged(int start, int end){callbackQMediaPlaylistMediaChanged(this->objectName().toUtf8().data(), start, end);};
void Signal_MediaInserted(int start, int end){callbackQMediaPlaylistMediaInserted(this->objectName().toUtf8().data(), start, end);};
void Signal_MediaRemoved(int start, int end){callbackQMediaPlaylistMediaRemoved(this->objectName().toUtf8().data(), start, end);};
void Signal_PlaybackModeChanged(QMediaPlaylist::PlaybackMode mode){callbackQMediaPlaylistPlaybackModeChanged(this->objectName().toUtf8().data(), mode);};
};

int QMediaPlaylist_PlaybackMode(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->playbackMode();
}

void QMediaPlaylist_SetPlaybackMode(QtObjectPtr ptr, int mode){
	static_cast<QMediaPlaylist*>(ptr)->setPlaybackMode(static_cast<QMediaPlaylist::PlaybackMode>(mode));
}

QtObjectPtr QMediaPlaylist_NewQMediaPlaylist(QtObjectPtr parent){
	return new QMediaPlaylist(static_cast<QObject*>(parent));
}

int QMediaPlaylist_AddMedia(QtObjectPtr ptr, QtObjectPtr content){
	return static_cast<QMediaPlaylist*>(ptr)->addMedia(*static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_Clear(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->clear();
}

int QMediaPlaylist_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->currentIndex();
}

void QMediaPlaylist_ConnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

void QMediaPlaylist_DisconnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

int QMediaPlaylist_Error(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->error();
}

char* QMediaPlaylist_ErrorString(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlaylist_InsertMedia(QtObjectPtr ptr, int pos, QtObjectPtr content){
	return static_cast<QMediaPlaylist*>(ptr)->insertMedia(pos, *static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_IsEmpty(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isEmpty();
}

int QMediaPlaylist_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isReadOnly();
}

void QMediaPlaylist_Load3(QtObjectPtr ptr, QtObjectPtr device, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

void QMediaPlaylist_Load(QtObjectPtr ptr, QtObjectPtr request, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QNetworkRequest*>(request), const_cast<const char*>(format));
}

void QMediaPlaylist_Load2(QtObjectPtr ptr, char* location, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(QUrl(QString(location)), const_cast<const char*>(format));
}

void QMediaPlaylist_ConnectLoadFailed(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_DisconnectLoadFailed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_ConnectLoaded(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_DisconnectLoaded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_ConnectMediaAboutToBeInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_ConnectMediaAboutToBeRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_ConnectMediaChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

void QMediaPlaylist_DisconnectMediaChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

int QMediaPlaylist_MediaCount(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaCount();
}

void QMediaPlaylist_ConnectMediaInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

void QMediaPlaylist_DisconnectMediaInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

QtObjectPtr QMediaPlaylist_MediaObject(QtObjectPtr ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaObject();
}

void QMediaPlaylist_ConnectMediaRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_DisconnectMediaRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_Next(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "next");
}

int QMediaPlaylist_NextIndex(QtObjectPtr ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->nextIndex(steps);
}

void QMediaPlaylist_ConnectPlaybackModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_DisconnectPlaybackModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_Previous(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "previous");
}

int QMediaPlaylist_PreviousIndex(QtObjectPtr ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->previousIndex(steps);
}

int QMediaPlaylist_RemoveMedia(QtObjectPtr ptr, int pos){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(pos);
}

int QMediaPlaylist_RemoveMedia2(QtObjectPtr ptr, int start, int end){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(start, end);
}

int QMediaPlaylist_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QMediaPlaylist_Save(QtObjectPtr ptr, char* location, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(QUrl(QString(location)), const_cast<const char*>(format));
}

void QMediaPlaylist_SetCurrentIndex(QtObjectPtr ptr, int playlistPosition){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "setCurrentIndex", Q_ARG(int, playlistPosition));
}

void QMediaPlaylist_Shuffle(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "shuffle");
}

void QMediaPlaylist_DestroyQMediaPlaylist(QtObjectPtr ptr){
	static_cast<QMediaPlaylist*>(ptr)->~QMediaPlaylist();
}

