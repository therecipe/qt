#include "qcursor.h"
#include <QUrl>
#include <QModelIndex>
#include <QBitmap>
#include <QScreen>
#include <QPixmap>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QCursor>
#include "_cgo_export.h"

class MyQCursor: public QCursor {
public:
};

void QCursor_QCursor_SetPos2(void* screen, int x, int y){
	QCursor::setPos(static_cast<QScreen*>(screen), x, y);
}

void QCursor_QCursor_SetPos(int x, int y){
	QCursor::setPos(x, y);
}

void* QCursor_NewQCursor(){
	return new QCursor();
}

void* QCursor_NewQCursor6(void* other){
	return new QCursor(*static_cast<QCursor*>(other));
}

void* QCursor_NewQCursor2(int shape){
	return new QCursor(static_cast<Qt::CursorShape>(shape));
}

void* QCursor_NewQCursor3(void* bitmap, void* mask, int hotX, int hotY){
	return new QCursor(*static_cast<QBitmap*>(bitmap), *static_cast<QBitmap*>(mask), hotX, hotY);
}

void* QCursor_NewQCursor5(void* c){
	return new QCursor(*static_cast<QCursor*>(c));
}

void* QCursor_NewQCursor4(void* pixmap, int hotX, int hotY){
	return new QCursor(*static_cast<QPixmap*>(pixmap), hotX, hotY);
}

void* QCursor_Bitmap(void* ptr){
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->bitmap());
}

void* QCursor_Mask(void* ptr){
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->mask());
}

void QCursor_QCursor_SetPos4(void* screen, void* p){
	QCursor::setPos(static_cast<QScreen*>(screen), *static_cast<QPoint*>(p));
}

void QCursor_QCursor_SetPos3(void* p){
	QCursor::setPos(*static_cast<QPoint*>(p));
}

void QCursor_SetShape(void* ptr, int shape){
	static_cast<QCursor*>(ptr)->setShape(static_cast<Qt::CursorShape>(shape));
}

int QCursor_Shape(void* ptr){
	return static_cast<QCursor*>(ptr)->shape();
}

void QCursor_DestroyQCursor(void* ptr){
	static_cast<QCursor*>(ptr)->~QCursor();
}

