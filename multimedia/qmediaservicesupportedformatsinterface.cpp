#include "qmediaservicesupportedformatsinterface.h"
#include <QModelIndex>
#include <QMediaService>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMediaServiceSupportedFormatsInterface>
#include "_cgo_export.h"

class MyQMediaServiceSupportedFormatsInterface: public QMediaServiceSupportedFormatsInterface {
public:
};

char* QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(QtObjectPtr ptr){
	return static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->supportedMimeTypes().join("|").toUtf8().data();
}

void QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(QtObjectPtr ptr){
	static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->~QMediaServiceSupportedFormatsInterface();
}

