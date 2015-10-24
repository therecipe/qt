#include "qimageencodercontrol.h"
#include <QImage>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QImageEncoderSettings>
#include <QImageEncoderControl>
#include "_cgo_export.h"

class MyQImageEncoderControl: public QImageEncoderControl {
public:
};

char* QImageEncoderControl_ImageCodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QImageEncoderControl*>(ptr)->imageCodecDescription(QString(codec)).toUtf8().data();
}

void QImageEncoderControl_SetImageSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QImageEncoderControl*>(ptr)->setImageSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char* QImageEncoderControl_SupportedImageCodecs(QtObjectPtr ptr){
	return static_cast<QImageEncoderControl*>(ptr)->supportedImageCodecs().join("|").toUtf8().data();
}

void QImageEncoderControl_DestroyQImageEncoderControl(QtObjectPtr ptr){
	static_cast<QImageEncoderControl*>(ptr)->~QImageEncoderControl();
}

