#include "qdomprocessinginstruction.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomProcessingInstruction>
#include "_cgo_export.h"

class MyQDomProcessingInstruction: public QDomProcessingInstruction {
public:
};

void* QDomProcessingInstruction_NewQDomProcessingInstruction(){
	return new QDomProcessingInstruction();
}

void* QDomProcessingInstruction_NewQDomProcessingInstruction2(void* x){
	return new QDomProcessingInstruction(*static_cast<QDomProcessingInstruction*>(x));
}

char* QDomProcessingInstruction_Data(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->data().toUtf8().data();
}

int QDomProcessingInstruction_NodeType(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->nodeType();
}

void QDomProcessingInstruction_SetData(void* ptr, char* d){
	static_cast<QDomProcessingInstruction*>(ptr)->setData(QString(d));
}

char* QDomProcessingInstruction_Target(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->target().toUtf8().data();
}

