#include "qimage.h"
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QPixelFormat>
#include <QIODevice>
#include <QPoint>
#include <QUrl>
#include <QColor>
#include <QByteArray>
#include <QImage>
#include "_cgo_export.h"

class MyQImage: public QImage {
public:
};

int QImage_ColorCount(void* ptr){
	return static_cast<QImage*>(ptr)->colorCount();
}

void QImage_Fill2(void* ptr, int color){
	static_cast<QImage*>(ptr)->fill(static_cast<Qt::GlobalColor>(color));
}

void QImage_Fill3(void* ptr, void* color){
	static_cast<QImage*>(ptr)->fill(*static_cast<QColor*>(color));
}

int QImage_Height(void* ptr){
	return static_cast<QImage*>(ptr)->height();
}

int QImage_IsNull(void* ptr){
	return static_cast<QImage*>(ptr)->isNull();
}

void QImage_SetOffset(void* ptr, void* offset){
	static_cast<QImage*>(ptr)->setOffset(*static_cast<QPoint*>(offset));
}

void QImage_SetText(void* ptr, char* key, char* text){
	static_cast<QImage*>(ptr)->setText(QString(key), QString(text));
}

int QImage_Width(void* ptr){
	return static_cast<QImage*>(ptr)->width();
}

int QImage_AllGray(void* ptr){
	return static_cast<QImage*>(ptr)->allGray();
}

int QImage_BitPlaneCount(void* ptr){
	return static_cast<QImage*>(ptr)->bitPlaneCount();
}

int QImage_ByteCount(void* ptr){
	return static_cast<QImage*>(ptr)->byteCount();
}

int QImage_BytesPerLine(void* ptr){
	return static_cast<QImage*>(ptr)->bytesPerLine();
}

int QImage_Depth(void* ptr){
	return static_cast<QImage*>(ptr)->depth();
}

double QImage_DevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QImage*>(ptr)->devicePixelRatio());
}

int QImage_DotsPerMeterX(void* ptr){
	return static_cast<QImage*>(ptr)->dotsPerMeterX();
}

int QImage_DotsPerMeterY(void* ptr){
	return static_cast<QImage*>(ptr)->dotsPerMeterY();
}

int QImage_Format(void* ptr){
	return static_cast<QImage*>(ptr)->format();
}

int QImage_HasAlphaChannel(void* ptr){
	return static_cast<QImage*>(ptr)->hasAlphaChannel();
}

void QImage_InvertPixels(void* ptr, int mode){
	static_cast<QImage*>(ptr)->invertPixels(static_cast<QImage::InvertMode>(mode));
}

int QImage_IsGrayscale(void* ptr){
	return static_cast<QImage*>(ptr)->isGrayscale();
}

int QImage_Load2(void* ptr, void* device, char* format){
	return static_cast<QImage*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QImage_Load(void* ptr, char* fileName, char* format){
	return static_cast<QImage*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QImage_LoadFromData2(void* ptr, void* data, char* format){
	return static_cast<QImage*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format));
}

int QImage_PixelIndex(void* ptr, void* position){
	return static_cast<QImage*>(ptr)->pixelIndex(*static_cast<QPoint*>(position));
}

int QImage_PixelIndex2(void* ptr, int x, int y){
	return static_cast<QImage*>(ptr)->pixelIndex(x, y);
}

int QImage_Save2(void* ptr, void* device, char* format, int quality){
	return static_cast<QImage*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QImage_Save(void* ptr, char* fileName, char* format, int quality){
	return static_cast<QImage*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void QImage_SetColorCount(void* ptr, int colorCount){
	static_cast<QImage*>(ptr)->setColorCount(colorCount);
}

void QImage_SetDevicePixelRatio(void* ptr, double scaleFactor){
	static_cast<QImage*>(ptr)->setDevicePixelRatio(static_cast<qreal>(scaleFactor));
}

void QImage_SetDotsPerMeterX(void* ptr, int x){
	static_cast<QImage*>(ptr)->setDotsPerMeterX(x);
}

void QImage_SetDotsPerMeterY(void* ptr, int y){
	static_cast<QImage*>(ptr)->setDotsPerMeterY(y);
}

void QImage_Swap(void* ptr, void* other){
	static_cast<QImage*>(ptr)->swap(*static_cast<QImage*>(other));
}

char* QImage_Text(void* ptr, char* key){
	return static_cast<QImage*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImage_TextKeys(void* ptr){
	return static_cast<QImage*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImage_QImage_ToImageFormat(void* format){
	return QImage::toImageFormat(*static_cast<QPixelFormat*>(format));
}

int QImage_Valid(void* ptr, void* pos){
	return static_cast<QImage*>(ptr)->valid(*static_cast<QPoint*>(pos));
}

int QImage_Valid2(void* ptr, int x, int y){
	return static_cast<QImage*>(ptr)->valid(x, y);
}

void QImage_DestroyQImage(void* ptr){
	static_cast<QImage*>(ptr)->~QImage();
}

