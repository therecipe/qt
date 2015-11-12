#include "qpixmap.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QImage>
#include <QByteArray>
#include <QRegion>
#include <QColor>
#include <QRect>
#include <QBitmap>
#include <QIODevice>
#include <QPixmap>
#include "_cgo_export.h"

class MyQPixmap: public QPixmap {
public:
};

int QPixmap_Depth(void* ptr){
	return static_cast<QPixmap*>(ptr)->depth();
}

int QPixmap_Height(void* ptr){
	return static_cast<QPixmap*>(ptr)->height();
}

int QPixmap_IsNull(void* ptr){
	return static_cast<QPixmap*>(ptr)->isNull();
}

int QPixmap_IsQBitmap(void* ptr){
	return static_cast<QPixmap*>(ptr)->isQBitmap();
}

int QPixmap_Width(void* ptr){
	return static_cast<QPixmap*>(ptr)->width();
}

int QPixmap_ConvertFromImage(void* ptr, void* image, int flags){
	return static_cast<QPixmap*>(ptr)->convertFromImage(*static_cast<QImage*>(image), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_QPixmap_DefaultDepth(){
	return QPixmap::defaultDepth();
}

void QPixmap_Detach(void* ptr){
	static_cast<QPixmap*>(ptr)->detach();
}

double QPixmap_DevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QPixmap*>(ptr)->devicePixelRatio());
}

void QPixmap_Fill(void* ptr, void* color){
	static_cast<QPixmap*>(ptr)->fill(*static_cast<QColor*>(color));
}

int QPixmap_HasAlpha(void* ptr){
	return static_cast<QPixmap*>(ptr)->hasAlpha();
}

int QPixmap_HasAlphaChannel(void* ptr){
	return static_cast<QPixmap*>(ptr)->hasAlphaChannel();
}

int QPixmap_Load(void* ptr, char* fileName, char* format, int flags){
	return static_cast<QPixmap*>(ptr)->load(QString(fileName), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_LoadFromData2(void* ptr, void* data, char* format, int flags){
	return static_cast<QPixmap*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_Save2(void* ptr, void* device, char* format, int quality){
	return static_cast<QPixmap*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QPixmap_Save(void* ptr, char* fileName, char* format, int quality){
	return static_cast<QPixmap*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void QPixmap_Scroll2(void* ptr, int dx, int dy, void* rect, void* exposed){
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(rect), static_cast<QRegion*>(exposed));
}

void QPixmap_Scroll(void* ptr, int dx, int dy, int x, int y, int width, int height, void* exposed){
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, x, y, width, height, static_cast<QRegion*>(exposed));
}

void QPixmap_SetDevicePixelRatio(void* ptr, double scaleFactor){
	static_cast<QPixmap*>(ptr)->setDevicePixelRatio(static_cast<qreal>(scaleFactor));
}

void QPixmap_SetMask(void* ptr, void* mask){
	static_cast<QPixmap*>(ptr)->setMask(*static_cast<QBitmap*>(mask));
}

void QPixmap_Swap(void* ptr, void* other){
	static_cast<QPixmap*>(ptr)->swap(*static_cast<QPixmap*>(other));
}

void QPixmap_DestroyQPixmap(void* ptr){
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

