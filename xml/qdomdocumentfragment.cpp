#include "qdomdocumentfragment.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomDocument>
#include <QDomDocumentFragment>
#include "_cgo_export.h"

class MyQDomDocumentFragment: public QDomDocumentFragment {
public:
};

void* QDomDocumentFragment_NewQDomDocumentFragment(){
	return new QDomDocumentFragment();
}

void* QDomDocumentFragment_NewQDomDocumentFragment2(void* x){
	return new QDomDocumentFragment(*static_cast<QDomDocumentFragment*>(x));
}

int QDomDocumentFragment_NodeType(void* ptr){
	return static_cast<QDomDocumentFragment*>(ptr)->nodeType();
}

