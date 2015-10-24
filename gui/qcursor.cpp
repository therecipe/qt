#include "qcursor.h"
#include <QModelIndex>
#include <QScreen>
#include <QPixmap>
#include <QBitmap>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QCursor>
#include "_cgo_export.h"

class MyQCursor: public QCursor {
public:
};

void QCursor_QCursor_SetPos2(QtObjectPtr screen, int x, int y){
	QCursor::setPos(static_cast<QScreen*>(screen), x, y);
}

void QCursor_QCursor_SetPos(int x, int y){
	QCursor::setPos(x, y);
}

QtObjectPtr QCursor_NewQCursor(){
	return new QCursor();
}

QtObjectPtr QCursor_NewQCursor6(QtObjectPtr other){
	return new QCursor(*static_cast<QCursor*>(other));
}

QtObjectPtr QCursor_NewQCursor2(int shape){
	return new QCursor(static_cast<Qt::CursorShape>(shape));
}

QtObjectPtr QCursor_NewQCursor3(QtObjectPtr bitmap, QtObjectPtr mask, int hotX, int hotY){
	return new QCursor(*static_cast<QBitmap*>(bitmap), *static_cast<QBitmap*>(mask), hotX, hotY);
}

QtObjectPtr QCursor_NewQCursor5(QtObjectPtr c){
	return new QCursor(*static_cast<QCursor*>(c));
}

QtObjectPtr QCursor_NewQCursor4(QtObjectPtr pixmap, int hotX, int hotY){
	return new QCursor(*static_cast<QPixmap*>(pixmap), hotX, hotY);
}

QtObjectPtr QCursor_Bitmap(QtObjectPtr ptr){
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->bitmap());
}

QtObjectPtr QCursor_Mask(QtObjectPtr ptr){
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->mask());
}

void QCursor_QCursor_SetPos4(QtObjectPtr screen, QtObjectPtr p){
	QCursor::setPos(static_cast<QScreen*>(screen), *static_cast<QPoint*>(p));
}

void QCursor_QCursor_SetPos3(QtObjectPtr p){
	QCursor::setPos(*static_cast<QPoint*>(p));
}

void QCursor_SetShape(QtObjectPtr ptr, int shape){
	static_cast<QCursor*>(ptr)->setShape(static_cast<Qt::CursorShape>(shape));
}

int QCursor_Shape(QtObjectPtr ptr){
	return static_cast<QCursor*>(ptr)->shape();
}

void QCursor_DestroyQCursor(QtObjectPtr ptr){
	static_cast<QCursor*>(ptr)->~QCursor();
}

