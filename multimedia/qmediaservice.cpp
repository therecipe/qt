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

void QMediaService_ReleaseControl(QtObjectPtr ptr, QtObjectPtr control){
	static_cast<QMediaService*>(ptr)->releaseControl(static_cast<QMediaControl*>(control));
}

QtObjectPtr QMediaService_RequestControl(QtObjectPtr ptr, char* interfa){
	return static_cast<QMediaService*>(ptr)->requestControl(const_cast<const char*>(interfa));
}

void QMediaService_DestroyQMediaService(QtObjectPtr ptr){
	static_cast<QMediaService*>(ptr)->~QMediaService();
}

