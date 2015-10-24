#include "qrubberband.h"
#include <QSize>
#include <QPoint>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRubberBand>
#include "_cgo_export.h"

class MyQRubberBand: public QRubberBand {
public:
};

void QRubberBand_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QRubberBand*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

QtObjectPtr QRubberBand_NewQRubberBand(int s, QtObjectPtr p){
	return new QRubberBand(static_cast<QRubberBand::Shape>(s), static_cast<QWidget*>(p));
}

void QRubberBand_Move2(QtObjectPtr ptr, QtObjectPtr p){
	static_cast<QRubberBand*>(ptr)->move(*static_cast<QPoint*>(p));
}

void QRubberBand_Move(QtObjectPtr ptr, int x, int y){
	static_cast<QRubberBand*>(ptr)->move(x, y);
}

void QRubberBand_Resize2(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QRubberBand*>(ptr)->resize(*static_cast<QSize*>(size));
}

void QRubberBand_Resize(QtObjectPtr ptr, int width, int height){
	static_cast<QRubberBand*>(ptr)->resize(width, height);
}

void QRubberBand_SetGeometry2(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QRubberBand*>(ptr)->setGeometry(x, y, width, height);
}

int QRubberBand_Shape(QtObjectPtr ptr){
	return static_cast<QRubberBand*>(ptr)->shape();
}

void QRubberBand_DestroyQRubberBand(QtObjectPtr ptr){
	static_cast<QRubberBand*>(ptr)->~QRubberBand();
}

