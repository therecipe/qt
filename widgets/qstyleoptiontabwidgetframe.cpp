#include "qstyleoptiontabwidgetframe.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionTab>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionTabWidgetFrame>
#include "_cgo_export.h"

class MyQStyleOptionTabWidgetFrame: public QStyleOptionTabWidgetFrame {
public:
};

QtObjectPtr QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame(){
	return new QStyleOptionTabWidgetFrame();
}

QtObjectPtr QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame2(QtObjectPtr other){
	return new QStyleOptionTabWidgetFrame(*static_cast<QStyleOptionTabWidgetFrame*>(other));
}

