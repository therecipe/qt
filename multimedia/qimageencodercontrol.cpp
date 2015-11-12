#include "qimageencodercontrol.h"
#include <QModelIndex>
#include <QImage>
#include <QImageEncoderSettings>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QImageEncoderControl>
#include "_cgo_export.h"

class MyQImageEncoderControl: public QImageEncoderControl {
public:
};

char* QImageEncoderControl_ImageCodecDescription(void* ptr, char* codec){
	return static_cast<QImageEncoderControl*>(ptr)->imageCodecDescription(QString(codec)).toUtf8().data();
}

void QImageEncoderControl_SetImageSettings(void* ptr, void* settings){
	static_cast<QImageEncoderControl*>(ptr)->setImageSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char* QImageEncoderControl_SupportedImageCodecs(void* ptr){
	return static_cast<QImageEncoderControl*>(ptr)->supportedImageCodecs().join("|").toUtf8().data();
}

void QImageEncoderControl_DestroyQImageEncoderControl(void* ptr){
	static_cast<QImageEncoderControl*>(ptr)->~QImageEncoderControl();
}

