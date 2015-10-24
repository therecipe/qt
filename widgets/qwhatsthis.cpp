#include "qwhatsthis.h"
#include <QModelIndex>
#include <QObject>
#include <QPoint>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWhatsThis>
#include "_cgo_export.h"

class MyQWhatsThis: public QWhatsThis {
public:
};

QtObjectPtr QWhatsThis_QWhatsThis_CreateAction(QtObjectPtr parent){
	return QWhatsThis::createAction(static_cast<QObject*>(parent));
}

void QWhatsThis_QWhatsThis_EnterWhatsThisMode(){
	QWhatsThis::enterWhatsThisMode();
}

void QWhatsThis_QWhatsThis_HideText(){
	QWhatsThis::hideText();
}

int QWhatsThis_QWhatsThis_InWhatsThisMode(){
	return QWhatsThis::inWhatsThisMode();
}

void QWhatsThis_QWhatsThis_LeaveWhatsThisMode(){
	QWhatsThis::leaveWhatsThisMode();
}

void QWhatsThis_QWhatsThis_ShowText(QtObjectPtr pos, char* text, QtObjectPtr w){
	QWhatsThis::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w));
}

