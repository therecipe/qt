#include "qrubberband.h"
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QSize>
#include <QRubberBand>
#include "_cgo_export.h"

class MyQRubberBand: public QRubberBand {
public:
};

void QRubberBand_SetGeometry(void* ptr, void* rect){
	static_cast<QRubberBand*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QRubberBand_NewQRubberBand(int s, void* p){
	return new QRubberBand(static_cast<QRubberBand::Shape>(s), static_cast<QWidget*>(p));
}

void QRubberBand_Move2(void* ptr, void* p){
	static_cast<QRubberBand*>(ptr)->move(*static_cast<QPoint*>(p));
}

void QRubberBand_Move(void* ptr, int x, int y){
	static_cast<QRubberBand*>(ptr)->move(x, y);
}

void QRubberBand_Resize2(void* ptr, void* size){
	static_cast<QRubberBand*>(ptr)->resize(*static_cast<QSize*>(size));
}

void QRubberBand_Resize(void* ptr, int width, int height){
	static_cast<QRubberBand*>(ptr)->resize(width, height);
}

void QRubberBand_SetGeometry2(void* ptr, int x, int y, int width, int height){
	static_cast<QRubberBand*>(ptr)->setGeometry(x, y, width, height);
}

int QRubberBand_Shape(void* ptr){
	return static_cast<QRubberBand*>(ptr)->shape();
}

void QRubberBand_DestroyQRubberBand(void* ptr){
	static_cast<QRubberBand*>(ptr)->~QRubberBand();
}

