#include "qpixmap.h"
#include <QString>
#include <QRegion>
#include <QColor>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QIODevice>
#include <QRect>
#include <QBitmap>
#include <QImage>
#include <QPixmap>
#include "_cgo_export.h"

class MyQPixmap: public QPixmap {
public:
};

int QPixmap_Depth(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->depth();
}

int QPixmap_Height(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->height();
}

int QPixmap_IsNull(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->isNull();
}

int QPixmap_IsQBitmap(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->isQBitmap();
}

int QPixmap_Width(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->width();
}

int QPixmap_ConvertFromImage(QtObjectPtr ptr, QtObjectPtr image, int flags){
	return static_cast<QPixmap*>(ptr)->convertFromImage(*static_cast<QImage*>(image), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_QPixmap_DefaultDepth(){
	return QPixmap::defaultDepth();
}

void QPixmap_Detach(QtObjectPtr ptr){
	static_cast<QPixmap*>(ptr)->detach();
}

void QPixmap_Fill(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QPixmap*>(ptr)->fill(*static_cast<QColor*>(color));
}

int QPixmap_HasAlpha(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->hasAlpha();
}

int QPixmap_HasAlphaChannel(QtObjectPtr ptr){
	return static_cast<QPixmap*>(ptr)->hasAlphaChannel();
}

int QPixmap_Load(QtObjectPtr ptr, char* fileName, char* format, int flags){
	return static_cast<QPixmap*>(ptr)->load(QString(fileName), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_LoadFromData2(QtObjectPtr ptr, QtObjectPtr data, char* format, int flags){
	return static_cast<QPixmap*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_Save2(QtObjectPtr ptr, QtObjectPtr device, char* format, int quality){
	return static_cast<QPixmap*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QPixmap_Save(QtObjectPtr ptr, char* fileName, char* format, int quality){
	return static_cast<QPixmap*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void QPixmap_Scroll2(QtObjectPtr ptr, int dx, int dy, QtObjectPtr rect, QtObjectPtr exposed){
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(rect), static_cast<QRegion*>(exposed));
}

void QPixmap_Scroll(QtObjectPtr ptr, int dx, int dy, int x, int y, int width, int height, QtObjectPtr exposed){
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, x, y, width, height, static_cast<QRegion*>(exposed));
}

void QPixmap_SetMask(QtObjectPtr ptr, QtObjectPtr mask){
	static_cast<QPixmap*>(ptr)->setMask(*static_cast<QBitmap*>(mask));
}

void QPixmap_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPixmap*>(ptr)->swap(*static_cast<QPixmap*>(other));
}

void QPixmap_DestroyQPixmap(QtObjectPtr ptr){
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

