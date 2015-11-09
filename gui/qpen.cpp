#include "qpen.h"
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QPen>
#include "_cgo_export.h"

class MyQPen: public QPen {
public:
};

void* QPen_NewQPen4(void* brush, double width, int style, int cap, int join){
	return new QPen(*static_cast<QBrush*>(brush), static_cast<qreal>(width), static_cast<Qt::PenStyle>(style), static_cast<Qt::PenCapStyle>(cap), static_cast<Qt::PenJoinStyle>(join));
}

void* QPen_NewQPen5(void* pen){
	return new QPen(*static_cast<QPen*>(pen));
}

void* QPen_Color(void* ptr){
	return new QColor(static_cast<QPen*>(ptr)->color());
}

void QPen_SetCapStyle(void* ptr, int style){
	static_cast<QPen*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPen_SetColor(void* ptr, void* color){
	static_cast<QPen*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QPen_SetJoinStyle(void* ptr, int style){
	static_cast<QPen*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPen_SetStyle(void* ptr, int style){
	static_cast<QPen*>(ptr)->setStyle(static_cast<Qt::PenStyle>(style));
}

void QPen_SetWidth(void* ptr, int width){
	static_cast<QPen*>(ptr)->setWidth(width);
}

int QPen_Style(void* ptr){
	return static_cast<QPen*>(ptr)->style();
}

int QPen_Width(void* ptr){
	return static_cast<QPen*>(ptr)->width();
}

double QPen_WidthF(void* ptr){
	return static_cast<double>(static_cast<QPen*>(ptr)->widthF());
}

void* QPen_NewQPen(){
	return new QPen();
}

void* QPen_NewQPen6(void* pen){
	return new QPen(*static_cast<QPen*>(pen));
}

void* QPen_NewQPen2(int style){
	return new QPen(static_cast<Qt::PenStyle>(style));
}

void* QPen_NewQPen3(void* color){
	return new QPen(*static_cast<QColor*>(color));
}

void* QPen_Brush(void* ptr){
	return new QBrush(static_cast<QPen*>(ptr)->brush());
}

int QPen_CapStyle(void* ptr){
	return static_cast<QPen*>(ptr)->capStyle();
}

double QPen_DashOffset(void* ptr){
	return static_cast<double>(static_cast<QPen*>(ptr)->dashOffset());
}

int QPen_IsCosmetic(void* ptr){
	return static_cast<QPen*>(ptr)->isCosmetic();
}

int QPen_IsSolid(void* ptr){
	return static_cast<QPen*>(ptr)->isSolid();
}

int QPen_JoinStyle(void* ptr){
	return static_cast<QPen*>(ptr)->joinStyle();
}

double QPen_MiterLimit(void* ptr){
	return static_cast<double>(static_cast<QPen*>(ptr)->miterLimit());
}

void QPen_SetBrush(void* ptr, void* brush){
	static_cast<QPen*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPen_SetCosmetic(void* ptr, int cosmetic){
	static_cast<QPen*>(ptr)->setCosmetic(cosmetic != 0);
}

void QPen_SetDashOffset(void* ptr, double offset){
	static_cast<QPen*>(ptr)->setDashOffset(static_cast<qreal>(offset));
}

void QPen_SetMiterLimit(void* ptr, double limit){
	static_cast<QPen*>(ptr)->setMiterLimit(static_cast<qreal>(limit));
}

void QPen_SetWidthF(void* ptr, double width){
	static_cast<QPen*>(ptr)->setWidthF(static_cast<qreal>(width));
}

void QPen_Swap(void* ptr, void* other){
	static_cast<QPen*>(ptr)->swap(*static_cast<QPen*>(other));
}

void QPen_DestroyQPen(void* ptr){
	static_cast<QPen*>(ptr)->~QPen();
}

