#include "qpen.h"
#include <QColor>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPen>
#include "_cgo_export.h"

class MyQPen: public QPen {
public:
};

QtObjectPtr QPen_NewQPen5(QtObjectPtr pen){
	return new QPen(*static_cast<QPen*>(pen));
}

void QPen_SetCapStyle(QtObjectPtr ptr, int style){
	static_cast<QPen*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPen_SetColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QPen*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QPen_SetJoinStyle(QtObjectPtr ptr, int style){
	static_cast<QPen*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPen_SetStyle(QtObjectPtr ptr, int style){
	static_cast<QPen*>(ptr)->setStyle(static_cast<Qt::PenStyle>(style));
}

void QPen_SetWidth(QtObjectPtr ptr, int width){
	static_cast<QPen*>(ptr)->setWidth(width);
}

int QPen_Style(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->style();
}

int QPen_Width(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->width();
}

QtObjectPtr QPen_NewQPen(){
	return new QPen();
}

QtObjectPtr QPen_NewQPen6(QtObjectPtr pen){
	return new QPen(*static_cast<QPen*>(pen));
}

QtObjectPtr QPen_NewQPen2(int style){
	return new QPen(static_cast<Qt::PenStyle>(style));
}

QtObjectPtr QPen_NewQPen3(QtObjectPtr color){
	return new QPen(*static_cast<QColor*>(color));
}

int QPen_CapStyle(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->capStyle();
}

int QPen_IsCosmetic(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->isCosmetic();
}

int QPen_IsSolid(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->isSolid();
}

int QPen_JoinStyle(QtObjectPtr ptr){
	return static_cast<QPen*>(ptr)->joinStyle();
}

void QPen_SetBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QPen*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPen_SetCosmetic(QtObjectPtr ptr, int cosmetic){
	static_cast<QPen*>(ptr)->setCosmetic(cosmetic != 0);
}

void QPen_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPen*>(ptr)->swap(*static_cast<QPen*>(other));
}

void QPen_DestroyQPen(QtObjectPtr ptr){
	static_cast<QPen*>(ptr)->~QPen();
}

