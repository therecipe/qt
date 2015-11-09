#include "qmediaplaylist.h"
#include <QUrl>
#include <QObject>
#include <QMediaContent>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QMetaObject>
#include <QModelIndex>
#include <QNetworkRequest>
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

int QMediaPlaylist_PlaybackMode(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->playbackMode();
}

void QMediaPlaylist_SetPlaybackMode(void* ptr, int mode){
	static_cast<QMediaPlaylist*>(ptr)->setPlaybackMode(static_cast<QMediaPlaylist::PlaybackMode>(mode));
}

void* QMediaPlaylist_NewQMediaPlaylist(void* parent){
	return new QMediaPlaylist(static_cast<QObject*>(parent));
}

int QMediaPlaylist_AddMedia(void* ptr, void* content){
	return static_cast<QMediaPlaylist*>(ptr)->addMedia(*static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_Clear(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->clear();
}

int QMediaPlaylist_CurrentIndex(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->currentIndex();
}

void QMediaPlaylist_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

void QMediaPlaylist_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

int QMediaPlaylist_Error(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->error();
}

char* QMediaPlaylist_ErrorString(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlaylist_InsertMedia(void* ptr, int pos, void* content){
	return static_cast<QMediaPlaylist*>(ptr)->insertMedia(pos, *static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_IsEmpty(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isEmpty();
}

int QMediaPlaylist_IsReadOnly(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isReadOnly();
}

void QMediaPlaylist_Load3(void* ptr, void* device, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

void QMediaPlaylist_Load(void* ptr, void* request, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QNetworkRequest*>(request), const_cast<const char*>(format));
}

void QMediaPlaylist_Load2(void* ptr, void* location, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_ConnectLoadFailed(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_DisconnectLoadFailed(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_ConnectLoaded(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_DisconnectLoaded(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_ConnectMediaAboutToBeInserted(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeInserted(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_ConnectMediaAboutToBeRemoved(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_ConnectMediaChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

void QMediaPlaylist_DisconnectMediaChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

int QMediaPlaylist_MediaCount(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaCount();
}

void QMediaPlaylist_ConnectMediaInserted(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

void QMediaPlaylist_DisconnectMediaInserted(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

void* QMediaPlaylist_MediaObject(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaObject();
}

void QMediaPlaylist_ConnectMediaRemoved(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_DisconnectMediaRemoved(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_Next(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "next");
}

int QMediaPlaylist_NextIndex(void* ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->nextIndex(steps);
}

void QMediaPlaylist_ConnectPlaybackModeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_DisconnectPlaybackModeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_Previous(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "previous");
}

int QMediaPlaylist_PreviousIndex(void* ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->previousIndex(steps);
}

int QMediaPlaylist_RemoveMedia(void* ptr, int pos){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(pos);
}

int QMediaPlaylist_RemoveMedia2(void* ptr, int start, int end){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(start, end);
}

int QMediaPlaylist_Save2(void* ptr, void* device, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QMediaPlaylist_Save(void* ptr, void* location, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_SetCurrentIndex(void* ptr, int playlistPosition){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "setCurrentIndex", Q_ARG(int, playlistPosition));
}

void QMediaPlaylist_Shuffle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "shuffle");
}

void QMediaPlaylist_DestroyQMediaPlaylist(void* ptr){
	static_cast<QMediaPlaylist*>(ptr)->~QMediaPlaylist();
}

