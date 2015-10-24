#include "qmediacontent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaPlaylist>
#include <QNetworkRequest>
#include <QMediaResource>
#include <QString>
#include <QMediaContent>
#include "_cgo_export.h"

class MyQMediaContent: public QMediaContent {
public:
};

QtObjectPtr QMediaContent_NewQMediaContent(){
	return new QMediaContent();
}

QtObjectPtr QMediaContent_NewQMediaContent7(QtObjectPtr playlist, char* contentUrl, int takeOwnership){
	return new QMediaContent(static_cast<QMediaPlaylist*>(playlist), QUrl(QString(contentUrl)), takeOwnership != 0);
}

QtObjectPtr QMediaContent_NewQMediaContent6(QtObjectPtr other){
	return new QMediaContent(*static_cast<QMediaContent*>(other));
}

QtObjectPtr QMediaContent_NewQMediaContent4(QtObjectPtr resource){
	return new QMediaContent(*static_cast<QMediaResource*>(resource));
}

QtObjectPtr QMediaContent_NewQMediaContent3(QtObjectPtr request){
	return new QMediaContent(*static_cast<QNetworkRequest*>(request));
}

QtObjectPtr QMediaContent_NewQMediaContent2(char* url){
	return new QMediaContent(QUrl(QString(url)));
}

char* QMediaContent_CanonicalUrl(QtObjectPtr ptr){
	return static_cast<QMediaContent*>(ptr)->canonicalUrl().toString().toUtf8().data();
}

int QMediaContent_IsNull(QtObjectPtr ptr){
	return static_cast<QMediaContent*>(ptr)->isNull();
}

QtObjectPtr QMediaContent_Playlist(QtObjectPtr ptr){
	return static_cast<QMediaContent*>(ptr)->playlist();
}

void QMediaContent_DestroyQMediaContent(QtObjectPtr ptr){
	static_cast<QMediaContent*>(ptr)->~QMediaContent();
}

