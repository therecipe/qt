#include "qlinef.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QLine>
#include <QPointF>
#include <QLineF>
#include "_cgo_export.h"

class MyQLineF: public QLineF {
public:
};

int QLineF_Intersect(QtObjectPtr ptr, QtObjectPtr line, QtObjectPtr intersectionPoint){
	return static_cast<QLineF*>(ptr)->intersect(*static_cast<QLineF*>(line), static_cast<QPointF*>(intersectionPoint));
}

QtObjectPtr QLineF_NewQLineF(){
	return new QLineF();
}

QtObjectPtr QLineF_NewQLineF4(QtObjectPtr line){
	return new QLineF(*static_cast<QLine*>(line));
}

QtObjectPtr QLineF_NewQLineF2(QtObjectPtr p1, QtObjectPtr p2){
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

int QLineF_IsNull(QtObjectPtr ptr){
	return static_cast<QLineF*>(ptr)->isNull();
}

void QLineF_SetP1(QtObjectPtr ptr, QtObjectPtr p1){
	static_cast<QLineF*>(ptr)->setP1(*static_cast<QPointF*>(p1));
}

void QLineF_SetP2(QtObjectPtr ptr, QtObjectPtr p2){
	static_cast<QLineF*>(ptr)->setP2(*static_cast<QPointF*>(p2));
}

void QLineF_SetPoints(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2){
	static_cast<QLineF*>(ptr)->setPoints(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QLineF_Translate(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QLineF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

