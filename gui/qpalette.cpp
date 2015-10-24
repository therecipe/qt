#include "qpalette.h"
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QPalette>
#include "_cgo_export.h"

class MyQPalette: public QPalette {
public:
};

int QPalette_NColorRoles_Type(){
	return QPalette::NColorRoles;
}

int QPalette_IsEqual(QtObjectPtr ptr, int cg1, int cg2){
	return static_cast<QPalette*>(ptr)->isEqual(static_cast<QPalette::ColorGroup>(cg1), static_cast<QPalette::ColorGroup>(cg2));
}

void QPalette_SetBrush2(QtObjectPtr ptr, int group, int role, QtObjectPtr brush){
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

QtObjectPtr QPalette_NewQPalette(){
	return new QPalette();
}

QtObjectPtr QPalette_NewQPalette8(QtObjectPtr other){
	return new QPalette(*static_cast<QPalette*>(other));
}

QtObjectPtr QPalette_NewQPalette3(int button){
	return new QPalette(static_cast<Qt::GlobalColor>(button));
}

QtObjectPtr QPalette_NewQPalette5(QtObjectPtr windowText, QtObjectPtr button, QtObjectPtr light, QtObjectPtr dark, QtObjectPtr mid, QtObjectPtr text, QtObjectPtr bright_text, QtObjectPtr base, QtObjectPtr window){
	return new QPalette(*static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

QtObjectPtr QPalette_NewQPalette2(QtObjectPtr button){
	return new QPalette(*static_cast<QColor*>(button));
}

QtObjectPtr QPalette_NewQPalette4(QtObjectPtr button, QtObjectPtr window){
	return new QPalette(*static_cast<QColor*>(button), *static_cast<QColor*>(window));
}

QtObjectPtr QPalette_NewQPalette7(QtObjectPtr p){
	return new QPalette(*static_cast<QPalette*>(p));
}

int QPalette_CurrentColorGroup(QtObjectPtr ptr){
	return static_cast<QPalette*>(ptr)->currentColorGroup();
}

int QPalette_IsBrushSet(QtObjectPtr ptr, int cg, int cr){
	return static_cast<QPalette*>(ptr)->isBrushSet(static_cast<QPalette::ColorGroup>(cg), static_cast<QPalette::ColorRole>(cr));
}

int QPalette_IsCopyOf(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QPalette*>(ptr)->isCopyOf(*static_cast<QPalette*>(p));
}

void QPalette_SetBrush(QtObjectPtr ptr, int role, QtObjectPtr brush){
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

void QPalette_SetColor(QtObjectPtr ptr, int group, int role, QtObjectPtr color){
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColor2(QtObjectPtr ptr, int role, QtObjectPtr color){
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColorGroup(QtObjectPtr ptr, int cg, QtObjectPtr windowText, QtObjectPtr button, QtObjectPtr light, QtObjectPtr dark, QtObjectPtr mid, QtObjectPtr text, QtObjectPtr bright_text, QtObjectPtr base, QtObjectPtr window){
	static_cast<QPalette*>(ptr)->setColorGroup(static_cast<QPalette::ColorGroup>(cg), *static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

void QPalette_SetCurrentColorGroup(QtObjectPtr ptr, int cg){
	static_cast<QPalette*>(ptr)->setCurrentColorGroup(static_cast<QPalette::ColorGroup>(cg));
}

void QPalette_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPalette*>(ptr)->swap(*static_cast<QPalette*>(other));
}

void QPalette_DestroyQPalette(QtObjectPtr ptr){
	static_cast<QPalette*>(ptr)->~QPalette();
}

