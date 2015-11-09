#include "qpalette.h"
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QPalette>
#include "_cgo_export.h"

class MyQPalette: public QPalette {
public:
};

int QPalette_NColorRoles_Type(){
	return QPalette::NColorRoles;
}

void* QPalette_Brush(void* ptr, int group, int role){
	return new QBrush(static_cast<QPalette*>(ptr)->brush(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role)));
}

int QPalette_IsEqual(void* ptr, int cg1, int cg2){
	return static_cast<QPalette*>(ptr)->isEqual(static_cast<QPalette::ColorGroup>(cg1), static_cast<QPalette::ColorGroup>(cg2));
}

void QPalette_SetBrush2(void* ptr, int group, int role, void* brush){
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

void* QPalette_NewQPalette(){
	return new QPalette();
}

void* QPalette_NewQPalette8(void* other){
	return new QPalette(*static_cast<QPalette*>(other));
}

void* QPalette_NewQPalette3(int button){
	return new QPalette(static_cast<Qt::GlobalColor>(button));
}

void* QPalette_NewQPalette5(void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window){
	return new QPalette(*static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

void* QPalette_NewQPalette2(void* button){
	return new QPalette(*static_cast<QColor*>(button));
}

void* QPalette_NewQPalette4(void* button, void* window){
	return new QPalette(*static_cast<QColor*>(button), *static_cast<QColor*>(window));
}

void* QPalette_NewQPalette7(void* p){
	return new QPalette(*static_cast<QPalette*>(p));
}

void* QPalette_AlternateBase(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->alternateBase());
}

void* QPalette_Base(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->base());
}

void* QPalette_BrightText(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->brightText());
}

void* QPalette_Brush2(void* ptr, int role){
	return new QBrush(static_cast<QPalette*>(ptr)->brush(static_cast<QPalette::ColorRole>(role)));
}

void* QPalette_Button(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->button());
}

void* QPalette_ButtonText(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->buttonText());
}

void* QPalette_Color(void* ptr, int group, int role){
	return new QColor(static_cast<QPalette*>(ptr)->color(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role)));
}

void* QPalette_Color2(void* ptr, int role){
	return new QColor(static_cast<QPalette*>(ptr)->color(static_cast<QPalette::ColorRole>(role)));
}

int QPalette_CurrentColorGroup(void* ptr){
	return static_cast<QPalette*>(ptr)->currentColorGroup();
}

void* QPalette_Dark(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->dark());
}

void* QPalette_Highlight(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->highlight());
}

void* QPalette_HighlightedText(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->highlightedText());
}

int QPalette_IsBrushSet(void* ptr, int cg, int cr){
	return static_cast<QPalette*>(ptr)->isBrushSet(static_cast<QPalette::ColorGroup>(cg), static_cast<QPalette::ColorRole>(cr));
}

int QPalette_IsCopyOf(void* ptr, void* p){
	return static_cast<QPalette*>(ptr)->isCopyOf(*static_cast<QPalette*>(p));
}

void* QPalette_Light(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->light());
}

void* QPalette_Link(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->link());
}

void* QPalette_LinkVisited(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->linkVisited());
}

void* QPalette_Mid(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->mid());
}

void* QPalette_Midlight(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->midlight());
}

void QPalette_SetBrush(void* ptr, int role, void* brush){
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

void QPalette_SetColor(void* ptr, int group, int role, void* color){
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColor2(void* ptr, int role, void* color){
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColorGroup(void* ptr, int cg, void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window){
	static_cast<QPalette*>(ptr)->setColorGroup(static_cast<QPalette::ColorGroup>(cg), *static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

void QPalette_SetCurrentColorGroup(void* ptr, int cg){
	static_cast<QPalette*>(ptr)->setCurrentColorGroup(static_cast<QPalette::ColorGroup>(cg));
}

void* QPalette_Shadow(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->shadow());
}

void QPalette_Swap(void* ptr, void* other){
	static_cast<QPalette*>(ptr)->swap(*static_cast<QPalette*>(other));
}

void* QPalette_Text(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->text());
}

void* QPalette_ToolTipBase(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->toolTipBase());
}

void* QPalette_ToolTipText(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->toolTipText());
}

void* QPalette_Window(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->window());
}

void* QPalette_WindowText(void* ptr){
	return new QBrush(static_cast<QPalette*>(ptr)->windowText());
}

void QPalette_DestroyQPalette(void* ptr){
	static_cast<QPalette*>(ptr)->~QPalette();
}

