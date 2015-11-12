#include "qmediacontent.h"
#include <QUrl>
#include <QModelIndex>
#include <QNetworkRequest>
#include <QMediaResource>
#include <QMediaPlaylist>
#include <QString>
#include <QVariant>
#include <QMediaContent>
#include "_cgo_export.h"

class MyQMediaContent: public QMediaContent {
public:
};

void* QMediaContent_NewQMediaContent(){
	return new QMediaContent();
}

void* QMediaContent_NewQMediaContent7(void* playlist, void* contentUrl, int takeOwnership){
	return new QMediaContent(static_cast<QMediaPlaylist*>(playlist), *static_cast<QUrl*>(contentUrl), takeOwnership != 0);
}

void* QMediaContent_NewQMediaContent6(void* other){
	return new QMediaContent(*static_cast<QMediaContent*>(other));
}

void* QMediaContent_NewQMediaContent4(void* resource){
	return new QMediaContent(*static_cast<QMediaResource*>(resource));
}

void* QMediaContent_NewQMediaContent3(void* request){
	return new QMediaContent(*static_cast<QNetworkRequest*>(request));
}

void* QMediaContent_NewQMediaContent2(void* url){
	return new QMediaContent(*static_cast<QUrl*>(url));
}

int QMediaContent_IsNull(void* ptr){
	return static_cast<QMediaContent*>(ptr)->isNull();
}

void* QMediaContent_Playlist(void* ptr){
	return static_cast<QMediaContent*>(ptr)->playlist();
}

void QMediaContent_DestroyQMediaContent(void* ptr){
	static_cast<QMediaContent*>(ptr)->~QMediaContent();
}

