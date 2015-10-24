#include "qimage.h"
#include <QModelIndex>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QColor>
#include <QByteArray>
#include <QIODevice>
#include <QPixelFormat>
#include <QImage>
#include "_cgo_export.h"

class MyQImage: public QImage {
public:
};

int QImage_ColorCount(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->colorCount();
}

void QImage_Fill2(QtObjectPtr ptr, int color){
	static_cast<QImage*>(ptr)->fill(static_cast<Qt::GlobalColor>(color));
}

void QImage_Fill3(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QImage*>(ptr)->fill(*static_cast<QColor*>(color));
}

int QImage_Height(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->height();
}

int QImage_IsNull(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->isNull();
}

void QImage_SetOffset(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QImage*>(ptr)->setOffset(*static_cast<QPoint*>(offset));
}

void QImage_SetText(QtObjectPtr ptr, char* key, char* text){
	static_cast<QImage*>(ptr)->setText(QString(key), QString(text));
}

int QImage_Width(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->width();
}

int QImage_AllGray(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->allGray();
}

int QImage_BitPlaneCount(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->bitPlaneCount();
}

int QImage_ByteCount(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->byteCount();
}

int QImage_BytesPerLine(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->bytesPerLine();
}

int QImage_Depth(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->depth();
}

int QImage_DotsPerMeterX(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->dotsPerMeterX();
}

int QImage_DotsPerMeterY(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->dotsPerMeterY();
}

int QImage_Format(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->format();
}

int QImage_HasAlphaChannel(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->hasAlphaChannel();
}

void QImage_InvertPixels(QtObjectPtr ptr, int mode){
	static_cast<QImage*>(ptr)->invertPixels(static_cast<QImage::InvertMode>(mode));
}

int QImage_IsGrayscale(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->isGrayscale();
}

int QImage_Load2(QtObjectPtr ptr, QtObjectPtr device, char* format){
	return static_cast<QImage*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QImage_Load(QtObjectPtr ptr, char* fileName, char* format){
	return static_cast<QImage*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QImage_LoadFromData2(QtObjectPtr ptr, QtObjectPtr data, char* format){
	return static_cast<QImage*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format));
}

int QImage_PixelIndex(QtObjectPtr ptr, QtObjectPtr position){
	return static_cast<QImage*>(ptr)->pixelIndex(*static_cast<QPoint*>(position));
}

int QImage_PixelIndex2(QtObjectPtr ptr, int x, int y){
	return static_cast<QImage*>(ptr)->pixelIndex(x, y);
}

int QImage_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format, int quality){
	return static_cast<QImage*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QImage_Save(QtObjectPtr ptr, char* fileName, char* format, int quality){
	return static_cast<QImage*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void QImage_SetColorCount(QtObjectPtr ptr, int colorCount){
	static_cast<QImage*>(ptr)->setColorCount(colorCount);
}

void QImage_SetDotsPerMeterX(QtObjectPtr ptr, int x){
	static_cast<QImage*>(ptr)->setDotsPerMeterX(x);
}

void QImage_SetDotsPerMeterY(QtObjectPtr ptr, int y){
	static_cast<QImage*>(ptr)->setDotsPerMeterY(y);
}

void QImage_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QImage*>(ptr)->swap(*static_cast<QImage*>(other));
}

char* QImage_Text(QtObjectPtr ptr, char* key){
	return static_cast<QImage*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImage_TextKeys(QtObjectPtr ptr){
	return static_cast<QImage*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImage_QImage_ToImageFormat(QtObjectPtr format){
	return QImage::toImageFormat(*static_cast<QPixelFormat*>(format));
}

int QImage_Valid(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QImage*>(ptr)->valid(*static_cast<QPoint*>(pos));
}

int QImage_Valid2(QtObjectPtr ptr, int x, int y){
	return static_cast<QImage*>(ptr)->valid(x, y);
}

void QImage_DestroyQImage(QtObjectPtr ptr){
	static_cast<QImage*>(ptr)->~QImage();
}

