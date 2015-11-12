#include "qbrush.h"
#include <QModelIndex>
#include <QTransform>
#include <QColor>
#include <QImage>
#include <QUrl>
#include <QVariant>
#include <QPixmap>
#include <QGradient>
#include <QString>
#include <QBrush>
#include "_cgo_export.h"

class MyQBrush: public QBrush {
public:
};

void* QBrush_NewQBrush4(int color, int style){
	return new QBrush(static_cast<Qt::GlobalColor>(color), static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetColor(void* ptr, void* color){
	static_cast<QBrush*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QBrush_NewQBrush(){
	return new QBrush();
}

void* QBrush_NewQBrush2(int style){
	return new QBrush(static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush6(int color, void* pixmap){
	return new QBrush(static_cast<Qt::GlobalColor>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush9(void* other){
	return new QBrush(*static_cast<QBrush*>(other));
}

void* QBrush_NewQBrush3(void* color, int style){
	return new QBrush(*static_cast<QColor*>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush5(void* color, void* pixmap){
	return new QBrush(*static_cast<QColor*>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush10(void* gradient){
	return new QBrush(*static_cast<QGradient*>(gradient));
}

void* QBrush_NewQBrush8(void* image){
	return new QBrush(*static_cast<QImage*>(image));
}

void* QBrush_NewQBrush7(void* pixmap){
	return new QBrush(*static_cast<QPixmap*>(pixmap));
}

void* QBrush_Color(void* ptr){
	return new QColor(static_cast<QBrush*>(ptr)->color());
}

void* QBrush_Gradient(void* ptr){
	return const_cast<QGradient*>(static_cast<QBrush*>(ptr)->gradient());
}

int QBrush_IsOpaque(void* ptr){
	return static_cast<QBrush*>(ptr)->isOpaque();
}

void QBrush_SetColor2(void* ptr, int color){
	static_cast<QBrush*>(ptr)->setColor(static_cast<Qt::GlobalColor>(color));
}

void QBrush_SetStyle(void* ptr, int style){
	static_cast<QBrush*>(ptr)->setStyle(static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetTexture(void* ptr, void* pixmap){
	static_cast<QBrush*>(ptr)->setTexture(*static_cast<QPixmap*>(pixmap));
}

void QBrush_SetTextureImage(void* ptr, void* image){
	static_cast<QBrush*>(ptr)->setTextureImage(*static_cast<QImage*>(image));
}

void QBrush_SetTransform(void* ptr, void* matrix){
	static_cast<QBrush*>(ptr)->setTransform(*static_cast<QTransform*>(matrix));
}

int QBrush_Style(void* ptr){
	return static_cast<QBrush*>(ptr)->style();
}

void QBrush_Swap(void* ptr, void* other){
	static_cast<QBrush*>(ptr)->swap(*static_cast<QBrush*>(other));
}

void QBrush_DestroyQBrush(void* ptr){
	static_cast<QBrush*>(ptr)->~QBrush();
}

