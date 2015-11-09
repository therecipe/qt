#include "qdomnode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextStream>
#include <QDomNode>
#include "_cgo_export.h"

class MyQDomNode: public QDomNode {
public:
};

void* QDomNode_NewQDomNode(){
	return new QDomNode();
}

void* QDomNode_NewQDomNode2(void* n){
	return new QDomNode(*static_cast<QDomNode*>(n));
}

void QDomNode_Clear(void* ptr){
	static_cast<QDomNode*>(ptr)->clear();
}

int QDomNode_ColumnNumber(void* ptr){
	return static_cast<QDomNode*>(ptr)->columnNumber();
}

int QDomNode_HasAttributes(void* ptr){
	return static_cast<QDomNode*>(ptr)->hasAttributes();
}

int QDomNode_HasChildNodes(void* ptr){
	return static_cast<QDomNode*>(ptr)->hasChildNodes();
}

int QDomNode_IsAttr(void* ptr){
	return static_cast<QDomNode*>(ptr)->isAttr();
}

int QDomNode_IsCDATASection(void* ptr){
	return static_cast<QDomNode*>(ptr)->isCDATASection();
}

int QDomNode_IsCharacterData(void* ptr){
	return static_cast<QDomNode*>(ptr)->isCharacterData();
}

int QDomNode_IsComment(void* ptr){
	return static_cast<QDomNode*>(ptr)->isComment();
}

int QDomNode_IsDocument(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocument();
}

int QDomNode_IsDocumentFragment(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentFragment();
}

int QDomNode_IsDocumentType(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentType();
}

int QDomNode_IsElement(void* ptr){
	return static_cast<QDomNode*>(ptr)->isElement();
}

int QDomNode_IsEntity(void* ptr){
	return static_cast<QDomNode*>(ptr)->isEntity();
}

int QDomNode_IsEntityReference(void* ptr){
	return static_cast<QDomNode*>(ptr)->isEntityReference();
}

int QDomNode_IsNotation(void* ptr){
	return static_cast<QDomNode*>(ptr)->isNotation();
}

int QDomNode_IsNull(void* ptr){
	return static_cast<QDomNode*>(ptr)->isNull();
}

int QDomNode_IsProcessingInstruction(void* ptr){
	return static_cast<QDomNode*>(ptr)->isProcessingInstruction();
}

int QDomNode_IsSupported(void* ptr, char* feature, char* version){
	return static_cast<QDomNode*>(ptr)->isSupported(QString(feature), QString(version));
}

int QDomNode_IsText(void* ptr){
	return static_cast<QDomNode*>(ptr)->isText();
}

int QDomNode_LineNumber(void* ptr){
	return static_cast<QDomNode*>(ptr)->lineNumber();
}

char* QDomNode_LocalName(void* ptr){
	return static_cast<QDomNode*>(ptr)->localName().toUtf8().data();
}

char* QDomNode_NamespaceURI(void* ptr){
	return static_cast<QDomNode*>(ptr)->namespaceURI().toUtf8().data();
}

char* QDomNode_NodeName(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeName().toUtf8().data();
}

int QDomNode_NodeType(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeType();
}

char* QDomNode_NodeValue(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeValue().toUtf8().data();
}

void QDomNode_Normalize(void* ptr){
	static_cast<QDomNode*>(ptr)->normalize();
}

char* QDomNode_Prefix(void* ptr){
	return static_cast<QDomNode*>(ptr)->prefix().toUtf8().data();
}

void QDomNode_Save(void* ptr, void* stream, int indent, int encodingPolicy){
	static_cast<QDomNode*>(ptr)->save(*static_cast<QTextStream*>(stream), indent, static_cast<QDomNode::EncodingPolicy>(encodingPolicy));
}

void QDomNode_SetNodeValue(void* ptr, char* v){
	static_cast<QDomNode*>(ptr)->setNodeValue(QString(v));
}

void QDomNode_SetPrefix(void* ptr, char* pre){
	static_cast<QDomNode*>(ptr)->setPrefix(QString(pre));
}

void QDomNode_DestroyQDomNode(void* ptr){
	static_cast<QDomNode*>(ptr)->~QDomNode();
}

