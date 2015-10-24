#include "qbrush.h"
#include <QTransform>
#include <QPixmap>
#include <QImage>
#include <QColor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGradient>
#include <QBrush>
#include "_cgo_export.h"

class MyQBrush: public QBrush {
public:
};

QtObjectPtr QBrush_NewQBrush4(int color, int style){
	return new QBrush(static_cast<Qt::GlobalColor>(color), static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QBrush*>(ptr)->setColor(*static_cast<QColor*>(color));
}

QtObjectPtr QBrush_NewQBrush(){
	return new QBrush();
}

QtObjectPtr QBrush_NewQBrush2(int style){
	return new QBrush(static_cast<Qt::BrushStyle>(style));
}

QtObjectPtr QBrush_NewQBrush6(int color, QtObjectPtr pixmap){
	return new QBrush(static_cast<Qt::GlobalColor>(color), *static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QBrush_NewQBrush9(QtObjectPtr other){
	return new QBrush(*static_cast<QBrush*>(other));
}

QtObjectPtr QBrush_NewQBrush3(QtObjectPtr color, int style){
	return new QBrush(*static_cast<QColor*>(color), static_cast<Qt::BrushStyle>(style));
}

QtObjectPtr QBrush_NewQBrush5(QtObjectPtr color, QtObjectPtr pixmap){
	return new QBrush(*static_cast<QColor*>(color), *static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QBrush_NewQBrush10(QtObjectPtr gradient){
	return new QBrush(*static_cast<QGradient*>(gradient));
}

QtObjectPtr QBrush_NewQBrush8(QtObjectPtr image){
	return new QBrush(*static_cast<QImage*>(image));
}

QtObjectPtr QBrush_NewQBrush7(QtObjectPtr pixmap){
	return new QBrush(*static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QBrush_Gradient(QtObjectPtr ptr){
	return const_cast<QGradient*>(static_cast<QBrush*>(ptr)->gradient());
}

int QBrush_IsOpaque(QtObjectPtr ptr){
	return static_cast<QBrush*>(ptr)->isOpaque();
}

void QBrush_SetColor2(QtObjectPtr ptr, int color){
	static_cast<QBrush*>(ptr)->setColor(static_cast<Qt::GlobalColor>(color));
}

void QBrush_SetStyle(QtObjectPtr ptr, int style){
	static_cast<QBrush*>(ptr)->setStyle(static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetTexture(QtObjectPtr ptr, QtObjectPtr pixmap){
	static_cast<QBrush*>(ptr)->setTexture(*static_cast<QPixmap*>(pixmap));
}

void QBrush_SetTextureImage(QtObjectPtr ptr, QtObjectPtr image){
	static_cast<QBrush*>(ptr)->setTextureImage(*static_cast<QImage*>(image));
}

void QBrush_SetTransform(QtObjectPtr ptr, QtObjectPtr matrix){
	static_cast<QBrush*>(ptr)->setTransform(*static_cast<QTransform*>(matrix));
}

int QBrush_Style(QtObjectPtr ptr){
	return static_cast<QBrush*>(ptr)->style();
}

void QBrush_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QBrush*>(ptr)->swap(*static_cast<QBrush*>(other));
}

void QBrush_DestroyQBrush(QtObjectPtr ptr){
	static_cast<QBrush*>(ptr)->~QBrush();
}

