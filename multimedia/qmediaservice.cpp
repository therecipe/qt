#include "qmediaservice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaControl>
#include <QMediaService>
#include "_cgo_export.h"

class MyQMediaService: public QMediaService {
public:
};

void QMediaService_ReleaseControl(void* ptr, void* control){
	static_cast<QMediaService*>(ptr)->releaseControl(static_cast<QMediaControl*>(control));
}

void* QMediaService_RequestControl(void* ptr, char* interfa){
	return static_cast<QMediaService*>(ptr)->requestControl(const_cast<const char*>(interfa));
}

void* QMediaService_RequestControl2(void* ptr){
	return static_cast<QMediaService*>(ptr)->requestControl<QMediaControl*>();
}

void QMediaService_DestroyQMediaService(void* ptr){
	static_cast<QMediaService*>(ptr)->~QMediaService();
}

