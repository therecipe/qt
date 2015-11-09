#include "qstyleoptiondockwidget.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionDockWidget>
#include "_cgo_export.h"

class MyQStyleOptionDockWidget: public QStyleOptionDockWidget {
public:
};

void* QStyleOptionDockWidget_NewQStyleOptionDockWidget(){
	return new QStyleOptionDockWidget();
}

void* QStyleOptionDockWidget_NewQStyleOptionDockWidget2(void* other){
	return new QStyleOptionDockWidget(*static_cast<QStyleOptionDockWidget*>(other));
}

