#include "qdomtext.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDomText>
#include "_cgo_export.h"

class MyQDomText: public QDomText {
public:
};

void* QDomText_NewQDomText(){
	return new QDomText();
}

void* QDomText_NewQDomText2(void* x){
	return new QDomText(*static_cast<QDomText*>(x));
}

int QDomText_NodeType(void* ptr){
	return static_cast<QDomText*>(ptr)->nodeType();
}

