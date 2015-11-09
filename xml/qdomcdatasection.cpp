#include "qdomcdatasection.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomCDATASection>
#include "_cgo_export.h"

class MyQDomCDATASection: public QDomCDATASection {
public:
};

void* QDomCDATASection_NewQDomCDATASection(){
	return new QDomCDATASection();
}

void* QDomCDATASection_NewQDomCDATASection2(void* x){
	return new QDomCDATASection(*static_cast<QDomCDATASection*>(x));
}

int QDomCDATASection_NodeType(void* ptr){
	return static_cast<QDomCDATASection*>(ptr)->nodeType();
}

