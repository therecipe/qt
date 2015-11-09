#include "qmediaservicesupportedformatsinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaService>
#include <QMediaServiceSupportedFormatsInterface>
#include "_cgo_export.h"

class MyQMediaServiceSupportedFormatsInterface: public QMediaServiceSupportedFormatsInterface {
public:
};

char* QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(void* ptr){
	return static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->supportedMimeTypes().join("|").toUtf8().data();
}

void QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(void* ptr){
	static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->~QMediaServiceSupportedFormatsInterface();
}

