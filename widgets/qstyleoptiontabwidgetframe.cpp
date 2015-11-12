#include "qstyleoptiontabwidgetframe.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionTab>
#include <QStyleOptionTabWidgetFrame>
#include "_cgo_export.h"

class MyQStyleOptionTabWidgetFrame: public QStyleOptionTabWidgetFrame {
public:
};

void* QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame(){
	return new QStyleOptionTabWidgetFrame();
}

void* QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame2(void* other){
	return new QStyleOptionTabWidgetFrame(*static_cast<QStyleOptionTabWidgetFrame*>(other));
}

