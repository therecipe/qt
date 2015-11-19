#include "qabstractmessagehandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractMessageHandler>
#include "_cgo_export.h"

class MyQAbstractMessageHandler: public QAbstractMessageHandler {
public:
};

void QAbstractMessageHandler_DestroyQAbstractMessageHandler(void* ptr){
	static_cast<QAbstractMessageHandler*>(ptr)->~QAbstractMessageHandler();
}

#include "qsimplexmlnodemodel.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QString>
#include <QVariant>
#include <QSimpleXmlNodeModel>
#include "_cgo_export.h"

class MyQSimpleXmlNodeModel: public QSimpleXmlNodeModel {
public:
};

char* QSimpleXmlNodeModel_StringValue(void* ptr, void* node){
	return static_cast<QSimpleXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8().data();
}

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr){
	static_cast<QSimpleXmlNodeModel*>(ptr)->~QSimpleXmlNodeModel();
}

#include "qxmlname.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNamePool>
#include <QString>
#include <QXmlName>
#include "_cgo_export.h"

class MyQXmlName: public QXmlName {
public:
};

void* QXmlName_NewQXmlName(){
	return new QXmlName();
}

void* QXmlName_NewQXmlName2(void* namePool, char* localName, char* namespaceURI, char* prefix){
	return new QXmlName(*static_cast<QXmlNamePool*>(namePool), QString(localName), QString(namespaceURI), QString(prefix));
}

int QXmlName_QXmlName_IsNCName(char* candidate){
	return QXmlName::isNCName(QString(candidate));
}

int QXmlName_IsNull(void* ptr){
	return static_cast<QXmlName*>(ptr)->isNull();
}

char* QXmlName_LocalName(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->localName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_NamespaceUri(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->namespaceUri(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_Prefix(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->prefix(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

char* QXmlName_ToClarkName(void* ptr, void* namePool){
	return static_cast<QXmlName*>(ptr)->toClarkName(*static_cast<QXmlNamePool*>(namePool)).toUtf8().data();
}

#include "qabstractxmlreceiver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QXmlName>
#include <QAbstractXmlReceiver>
#include "_cgo_export.h"

class MyQAbstractXmlReceiver: public QAbstractXmlReceiver {
public:
};

void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Characters(void* ptr, void* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QAbstractXmlReceiver_Comment(void* ptr, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->comment(QString(value));
}

void QAbstractXmlReceiver_EndDocument(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endDocument();
}

void QAbstractXmlReceiver_EndElement(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endElement();
}

void QAbstractXmlReceiver_EndOfSequence(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->endOfSequence();
}

void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name){
	static_cast<QAbstractXmlReceiver*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, char* value){
	static_cast<QAbstractXmlReceiver*>(ptr)->processingInstruction(*static_cast<QXmlName*>(target), QString(value));
}

void QAbstractXmlReceiver_StartDocument(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startDocument();
}

void QAbstractXmlReceiver_StartElement(void* ptr, void* name){
	static_cast<QAbstractXmlReceiver*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QAbstractXmlReceiver_StartOfSequence(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->startOfSequence();
}

void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr){
	static_cast<QAbstractXmlReceiver*>(ptr)->~QAbstractXmlReceiver();
}

#include "qxmlnodemodelindex.h"
#include <QModelIndex>
#include <QAbstractXmlNodeModel>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlNodeModelIndex>
#include "_cgo_export.h"

class MyQXmlNodeModelIndex: public QXmlNodeModelIndex {
public:
};

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex(){
	return new QXmlNodeModelIndex();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other){
	return new QXmlNodeModelIndex(*static_cast<QXmlNodeModelIndex*>(other));
}

void* QXmlNodeModelIndex_InternalPointer(void* ptr){
	return static_cast<QXmlNodeModelIndex*>(ptr)->internalPointer();
}

int QXmlNodeModelIndex_IsNull(void* ptr){
	return static_cast<QXmlNodeModelIndex*>(ptr)->isNull();
}

void* QXmlNodeModelIndex_Model(void* ptr){
	return const_cast<QAbstractXmlNodeModel*>(static_cast<QXmlNodeModelIndex*>(ptr)->model());
}

#include "qxmlserializer.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QStringRef>
#include <QString>
#include <QVariant>
#include <QIODevice>
#include <QXmlQuery>
#include <QTextCodec>
#include <QXmlSerializer>
#include "_cgo_export.h"

class MyQXmlSerializer: public QXmlSerializer {
public:
};

void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice){
	return new QXmlSerializer(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlSerializer_Attribute(void* ptr, void* name, void* value){
	static_cast<QXmlSerializer*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlSerializer_Characters(void* ptr, void* value){
	static_cast<QXmlSerializer*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_Comment(void* ptr, char* value){
	static_cast<QXmlSerializer*>(ptr)->comment(QString(value));
}

void QXmlSerializer_EndDocument(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endDocument();
}

void QXmlSerializer_EndElement(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endElement();
}

void* QXmlSerializer_Codec(void* ptr){
	return const_cast<QTextCodec*>(static_cast<QXmlSerializer*>(ptr)->codec());
}

void QXmlSerializer_EndOfSequence(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->endOfSequence();
}

void QXmlSerializer_NamespaceBinding(void* ptr, void* nb){
	static_cast<QXmlSerializer*>(ptr)->namespaceBinding(*static_cast<QXmlName*>(nb));
}

void* QXmlSerializer_OutputDevice(void* ptr){
	return static_cast<QXmlSerializer*>(ptr)->outputDevice();
}

void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, char* value){
	static_cast<QXmlSerializer*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlSerializer_SetCodec(void* ptr, void* outputCodec){
	static_cast<QXmlSerializer*>(ptr)->setCodec(static_cast<QTextCodec*>(outputCodec));
}

void QXmlSerializer_StartDocument(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->startDocument();
}

void QXmlSerializer_StartElement(void* ptr, void* name){
	static_cast<QXmlSerializer*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlSerializer_StartOfSequence(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->startOfSequence();
}

#include "qsourcelocation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSourceLocation>
#include "_cgo_export.h"

class MyQSourceLocation: public QSourceLocation {
public:
};

void* QSourceLocation_NewQSourceLocation(){
	return new QSourceLocation();
}

void* QSourceLocation_NewQSourceLocation2(void* other){
	return new QSourceLocation(*static_cast<QSourceLocation*>(other));
}

void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c){
	return new QSourceLocation(*static_cast<QUrl*>(u), l, c);
}

int QSourceLocation_IsNull(void* ptr){
	return static_cast<QSourceLocation*>(ptr)->isNull();
}

void QSourceLocation_SetUri(void* ptr, void* newUri){
	static_cast<QSourceLocation*>(ptr)->setUri(*static_cast<QUrl*>(newUri));
}

void QSourceLocation_DestroyQSourceLocation(void* ptr){
	static_cast<QSourceLocation*>(ptr)->~QSourceLocation();
}

#include "qxmlschemavalidator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkAccessManager>
#include <QAbstractUriResolver>
#include <QString>
#include <QByteArray>
#include <QIODevice>
#include <QXmlSchema>
#include <QAbstractMessageHandler>
#include <QXmlSchemaValidator>
#include "_cgo_export.h"

class MyQXmlSchemaValidator: public QXmlSchemaValidator {
public:
};

void* QXmlSchemaValidator_NewQXmlSchemaValidator(){
	return new QXmlSchemaValidator();
}

void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema){
	return new QXmlSchemaValidator(*static_cast<QXmlSchema*>(schema));
}

void* QXmlSchemaValidator_MessageHandler(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->messageHandler();
}

void* QXmlSchemaValidator_NetworkAccessManager(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->networkAccessManager();
}

void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler){
	static_cast<QXmlSchemaValidator*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager){
	static_cast<QXmlSchemaValidator*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchemaValidator_SetSchema(void* ptr, void* schema){
	static_cast<QXmlSchemaValidator*>(ptr)->setSchema(*static_cast<QXmlSchema*>(schema));
}

void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlSchemaValidator*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlSchemaValidator_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchemaValidator*>(ptr)->uriResolver());
}

int QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

int QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchemaValidator_Validate(void* ptr, void* source){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QUrl*>(source));
}

void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr){
	static_cast<QXmlSchemaValidator*>(ptr)->~QXmlSchemaValidator();
}

#include "qxmlformatter.h"
#include <QStringRef>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QIODevice>
#include <QXmlQuery>
#include <QXmlFormatter>
#include "_cgo_export.h"

class MyQXmlFormatter: public QXmlFormatter {
public:
};

void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice){
	return new QXmlFormatter(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlFormatter_Attribute(void* ptr, void* name, void* value){
	static_cast<QXmlFormatter*>(ptr)->attribute(*static_cast<QXmlName*>(name), *static_cast<QStringRef*>(value));
}

void QXmlFormatter_Characters(void* ptr, void* value){
	static_cast<QXmlFormatter*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlFormatter_Comment(void* ptr, char* value){
	static_cast<QXmlFormatter*>(ptr)->comment(QString(value));
}

void QXmlFormatter_EndDocument(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endDocument();
}

void QXmlFormatter_EndElement(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endElement();
}

void QXmlFormatter_EndOfSequence(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->endOfSequence();
}

int QXmlFormatter_IndentationDepth(void* ptr){
	return static_cast<QXmlFormatter*>(ptr)->indentationDepth();
}

void QXmlFormatter_ProcessingInstruction(void* ptr, void* name, char* value){
	static_cast<QXmlFormatter*>(ptr)->processingInstruction(*static_cast<QXmlName*>(name), QString(value));
}

void QXmlFormatter_SetIndentationDepth(void* ptr, int depth){
	static_cast<QXmlFormatter*>(ptr)->setIndentationDepth(depth);
}

void QXmlFormatter_StartDocument(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->startDocument();
}

void QXmlFormatter_StartElement(void* ptr, void* name){
	static_cast<QXmlFormatter*>(ptr)->startElement(*static_cast<QXmlName*>(name));
}

void QXmlFormatter_StartOfSequence(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->startOfSequence();
}

#include "qxmlschema.h"
#include <QNetworkAccessManager>
#include <QAbstractUriResolver>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractMessageHandler>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QXmlSchema>
#include "_cgo_export.h"

class MyQXmlSchema: public QXmlSchema {
public:
};

void* QXmlSchema_NewQXmlSchema(){
	return new QXmlSchema();
}

void* QXmlSchema_NewQXmlSchema2(void* other){
	return new QXmlSchema(*static_cast<QXmlSchema*>(other));
}

int QXmlSchema_IsValid(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->isValid();
}

int QXmlSchema_Load2(void* ptr, void* source, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load3(void* ptr, void* data, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load(void* ptr, void* source){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QUrl*>(source));
}

void* QXmlSchema_MessageHandler(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->messageHandler();
}

void* QXmlSchema_NetworkAccessManager(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->networkAccessManager();
}

void QXmlSchema_SetMessageHandler(void* ptr, void* handler){
	static_cast<QXmlSchema*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(handler));
}

void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager){
	static_cast<QXmlSchema*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QXmlSchema_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlSchema*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlSchema_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlSchema*>(ptr)->uriResolver());
}

void QXmlSchema_DestroyQXmlSchema(void* ptr){
	static_cast<QXmlSchema*>(ptr)->~QXmlSchema();
}

#include "qxmlnamepool.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QString>
#include <QXmlNamePool>
#include "_cgo_export.h"

class MyQXmlNamePool: public QXmlNamePool {
public:
};

void* QXmlNamePool_NewQXmlNamePool(){
	return new QXmlNamePool();
}

void* QXmlNamePool_NewQXmlNamePool2(void* other){
	return new QXmlNamePool(*static_cast<QXmlNamePool*>(other));
}

void QXmlNamePool_DestroyQXmlNamePool(void* ptr){
	static_cast<QXmlNamePool*>(ptr)->~QXmlNamePool();
}

#include "qabstractxmlnodemodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QAbstractXmlNodeModel>
#include "_cgo_export.h"

class MyQAbstractXmlNodeModel: public QAbstractXmlNodeModel {
public:
};

int QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

int QAbstractXmlNodeModel_Kind(void* ptr, void* ni){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

char* QAbstractXmlNodeModel_StringValue(void* ptr, void* n){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(n)).toUtf8().data();
}

void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node){
	return new QVariant(static_cast<QAbstractXmlNodeModel*>(ptr)->typedValue(*static_cast<QXmlNodeModelIndex*>(node)));
}

void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr){
	static_cast<QAbstractXmlNodeModel*>(ptr)->~QAbstractXmlNodeModel();
}

#include "qxmlitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QXmlItem>
#include "_cgo_export.h"

class MyQXmlItem: public QXmlItem {
public:
};

void* QXmlItem_NewQXmlItem(){
	return new QXmlItem();
}

void* QXmlItem_NewQXmlItem4(void* atomicValue){
	return new QXmlItem(*static_cast<QVariant*>(atomicValue));
}

void* QXmlItem_NewQXmlItem2(void* other){
	return new QXmlItem(*static_cast<QXmlItem*>(other));
}

void* QXmlItem_NewQXmlItem3(void* node){
	return new QXmlItem(*static_cast<QXmlNodeModelIndex*>(node));
}

int QXmlItem_IsNode(void* ptr){
	return static_cast<QXmlItem*>(ptr)->isNode();
}

int QXmlItem_IsNull(void* ptr){
	return static_cast<QXmlItem*>(ptr)->isNull();
}

void QXmlItem_DestroyQXmlItem(void* ptr){
	static_cast<QXmlItem*>(ptr)->~QXmlItem();
}

#include "qxmlquery.h"
#include <QModelIndex>
#include <QNetworkAccessManager>
#include <QIODevice>
#include <QXmlItem>
#include <QAbstractUriResolver>
#include <QXmlNamePool>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractMessageHandler>
#include <QXmlName>
#include <QXmlQuery>
#include "_cgo_export.h"

class MyQXmlQuery: public QXmlQuery {
public:
};

void* QXmlQuery_NewQXmlQuery(){
	return new QXmlQuery();
}

void* QXmlQuery_NewQXmlQuery4(int queryLanguage, void* np){
	return new QXmlQuery(static_cast<QXmlQuery::QueryLanguage>(queryLanguage), *static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery3(void* np){
	return new QXmlQuery(*static_cast<QXmlNamePool*>(np));
}

void* QXmlQuery_NewQXmlQuery2(void* other){
	return new QXmlQuery(*static_cast<QXmlQuery*>(other));
}

void QXmlQuery_BindVariable5(void* ptr, char* localName, void* device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable4(void* ptr, char* localName, void* value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable6(void* ptr, char* localName, void* query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(QString(localName), *static_cast<QXmlQuery*>(query));
}

void QXmlQuery_BindVariable2(void* ptr, void* name, void* device){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), static_cast<QIODevice*>(device));
}

void QXmlQuery_BindVariable(void* ptr, void* name, void* value){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlItem*>(value));
}

void QXmlQuery_BindVariable3(void* ptr, void* name, void* query){
	static_cast<QXmlQuery*>(ptr)->bindVariable(*static_cast<QXmlName*>(name), *static_cast<QXmlQuery*>(query));
}

int QXmlQuery_IsValid(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->isValid();
}

void* QXmlQuery_MessageHandler(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->messageHandler();
}

void* QXmlQuery_NetworkAccessManager(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->networkAccessManager();
}

int QXmlQuery_QueryLanguage(void* ptr){
	return static_cast<QXmlQuery*>(ptr)->queryLanguage();
}

int QXmlQuery_SetFocus3(void* ptr, void* document){
	return static_cast<QXmlQuery*>(ptr)->setFocus(static_cast<QIODevice*>(document));
}

int QXmlQuery_SetFocus4(void* ptr, char* focus){
	return static_cast<QXmlQuery*>(ptr)->setFocus(QString(focus));
}

int QXmlQuery_SetFocus2(void* ptr, void* documentURI){
	return static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetFocus(void* ptr, void* item){
	static_cast<QXmlQuery*>(ptr)->setFocus(*static_cast<QXmlItem*>(item));
}

void QXmlQuery_SetInitialTemplateName2(void* ptr, char* localName){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(QString(localName));
}

void QXmlQuery_SetInitialTemplateName(void* ptr, void* name){
	static_cast<QXmlQuery*>(ptr)->setInitialTemplateName(*static_cast<QXmlName*>(name));
}

void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler){
	static_cast<QXmlQuery*>(ptr)->setMessageHandler(static_cast<QAbstractMessageHandler*>(aMessageHandler));
}

void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager){
	static_cast<QXmlQuery*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(newManager));
}

void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(static_cast<QIODevice*>(sourceCode), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery3(void* ptr, char* sourceCode, void* documentURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(QString(sourceCode), *static_cast<QUrl*>(documentURI));
}

void QXmlQuery_SetQuery2(void* ptr, void* queryURI, void* baseURI){
	static_cast<QXmlQuery*>(ptr)->setQuery(*static_cast<QUrl*>(queryURI), *static_cast<QUrl*>(baseURI));
}

void QXmlQuery_SetUriResolver(void* ptr, void* resolver){
	static_cast<QXmlQuery*>(ptr)->setUriResolver(static_cast<QAbstractUriResolver*>(resolver));
}

void* QXmlQuery_UriResolver(void* ptr){
	return const_cast<QAbstractUriResolver*>(static_cast<QXmlQuery*>(ptr)->uriResolver());
}

void QXmlQuery_DestroyQXmlQuery(void* ptr){
	static_cast<QXmlQuery*>(ptr)->~QXmlQuery();
}

#include "qabstracturiresolver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractUriResolver>
#include "_cgo_export.h"

class MyQAbstractUriResolver: public QAbstractUriResolver {
public:
};

void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr){
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

#include "qxmlresultitems.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlResultItems>
#include "_cgo_export.h"

class MyQXmlResultItems: public QXmlResultItems {
public:
};

void* QXmlResultItems_NewQXmlResultItems(){
	return new QXmlResultItems();
}

int QXmlResultItems_HasError(void* ptr){
	return static_cast<QXmlResultItems*>(ptr)->hasError();
}

void QXmlResultItems_DestroyQXmlResultItems(void* ptr){
	static_cast<QXmlResultItems*>(ptr)->~QXmlResultItems();
}

