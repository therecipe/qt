#define protected public

#include "xmlpatterns.h"
#include "_cgo_export.h"

#include <QAbstractMessageHandler>
#include <QAbstractUriResolver>
#include <QAbstractXmlNodeModel>
#include <QAbstractXmlReceiver>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QNetworkAccessManager>
#include <QSimpleXmlNodeModel>
#include <QSourceLocation>
#include <QString>
#include <QStringRef>
#include <QTextCodec>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QXmlFormatter>
#include <QXmlItem>
#include <QXmlName>
#include <QXmlNamePool>
#include <QXmlNodeModelIndex>
#include <QXmlQuery>
#include <QXmlResultItems>
#include <QXmlSchema>
#include <QXmlSchemaValidator>
#include <QXmlSerializer>

class MyQAbstractMessageHandler: public QAbstractMessageHandler {
public:
	void timerEvent(QTimerEvent * event) { callbackQAbstractMessageHandlerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractMessageHandlerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractMessageHandlerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QAbstractMessageHandler_DestroyQAbstractMessageHandler(void* ptr){
	static_cast<QAbstractMessageHandler*>(ptr)->~QAbstractMessageHandler();
}

void QAbstractMessageHandler_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractMessageHandler*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractMessageHandler_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractMessageHandler_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractMessageHandler*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractMessageHandler_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractMessageHandler_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractMessageHandler*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractMessageHandler_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractMessageHandler*>(ptr)->QAbstractMessageHandler::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractUriResolver: public QAbstractUriResolver {
public:
	void timerEvent(QTimerEvent * event) { callbackQAbstractUriResolverTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractUriResolverChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAbstractUriResolverCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAbstractUriResolver_Resolve(void* ptr, void* relative, void* baseURI){
	return new QUrl(static_cast<QAbstractUriResolver*>(ptr)->resolve(*static_cast<QUrl*>(relative), *static_cast<QUrl*>(baseURI)));
}

void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr){
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

void QAbstractUriResolver_TimerEvent(void* ptr, void* event){
	static_cast<MyQAbstractUriResolver*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractUriResolver_TimerEventDefault(void* ptr, void* event){
	static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractUriResolver_ChildEvent(void* ptr, void* event){
	static_cast<MyQAbstractUriResolver*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractUriResolver_ChildEventDefault(void* ptr, void* event){
	static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractUriResolver_CustomEvent(void* ptr, void* event){
	static_cast<MyQAbstractUriResolver*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractUriResolver_CustomEventDefault(void* ptr, void* event){
	static_cast<QAbstractUriResolver*>(ptr)->QAbstractUriResolver::customEvent(static_cast<QEvent*>(event));
}

class MyQAbstractXmlNodeModel: public QAbstractXmlNodeModel {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void* QAbstractXmlNodeModel_BaseUri(void* ptr, void* n){
	return new QUrl(static_cast<QAbstractXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(n)));
}

int QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

void* QAbstractXmlNodeModel_DocumentUri(void* ptr, void* n){
	return new QUrl(static_cast<QAbstractXmlNodeModel*>(ptr)->documentUri(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QAbstractXmlNodeModel_ElementById(void* ptr, void* id){
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->elementById(*static_cast<QXmlName*>(id)));
}

int QAbstractXmlNodeModel_Kind(void* ptr, void* ni){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

void* QAbstractXmlNodeModel_Root(void* ptr, void* n){
	return new QXmlNodeModelIndex(static_cast<QAbstractXmlNodeModel*>(ptr)->root(*static_cast<QXmlNodeModelIndex*>(n)));
}

void* QAbstractXmlNodeModel_SourceLocation(void* ptr, void* index){
	return new QSourceLocation(static_cast<QAbstractXmlNodeModel*>(ptr)->sourceLocation(*static_cast<QXmlNodeModelIndex*>(index)));
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

char* QAbstractXmlNodeModel_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQAbstractXmlNodeModel*>(static_cast<QAbstractXmlNodeModel*>(ptr))) {
		return static_cast<MyQAbstractXmlNodeModel*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractXmlNodeModel_BASE").toUtf8().data();
}

void QAbstractXmlNodeModel_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQAbstractXmlNodeModel*>(static_cast<QAbstractXmlNodeModel*>(ptr))) {
		static_cast<MyQAbstractXmlNodeModel*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAbstractXmlReceiver: public QAbstractXmlReceiver {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
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

char* QAbstractXmlReceiver_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQAbstractXmlReceiver*>(static_cast<QAbstractXmlReceiver*>(ptr))) {
		return static_cast<MyQAbstractXmlReceiver*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAbstractXmlReceiver_BASE").toUtf8().data();
}

void QAbstractXmlReceiver_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQAbstractXmlReceiver*>(static_cast<QAbstractXmlReceiver*>(ptr))) {
		static_cast<MyQAbstractXmlReceiver*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSimpleXmlNodeModel: public QSimpleXmlNodeModel {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void* QSimpleXmlNodeModel_BaseUri(void* ptr, void* node){
	return new QUrl(static_cast<QSimpleXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(node)));
}

void* QSimpleXmlNodeModel_ElementById(void* ptr, void* id){
	return new QXmlNodeModelIndex(static_cast<QSimpleXmlNodeModel*>(ptr)->elementById(*static_cast<QXmlName*>(id)));
}

void* QSimpleXmlNodeModel_NamePool(void* ptr){
	return new QXmlNamePool(static_cast<QSimpleXmlNodeModel*>(ptr)->namePool());
}

char* QSimpleXmlNodeModel_StringValue(void* ptr, void* node){
	return static_cast<QSimpleXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8().data();
}

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr){
	static_cast<QSimpleXmlNodeModel*>(ptr)->~QSimpleXmlNodeModel();
}

char* QSimpleXmlNodeModel_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSimpleXmlNodeModel*>(static_cast<QSimpleXmlNodeModel*>(ptr))) {
		return static_cast<MyQSimpleXmlNodeModel*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSimpleXmlNodeModel_BASE").toUtf8().data();
}

void QSimpleXmlNodeModel_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSimpleXmlNodeModel*>(static_cast<QSimpleXmlNodeModel*>(ptr))) {
		static_cast<MyQSimpleXmlNodeModel*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSourceLocation_NewQSourceLocation(){
	return new QSourceLocation();
}

void* QSourceLocation_NewQSourceLocation2(void* other){
	return new QSourceLocation(*static_cast<QSourceLocation*>(other));
}

void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c){
	return new QSourceLocation(*static_cast<QUrl*>(u), l, c);
}

long long QSourceLocation_Column(void* ptr){
	return static_cast<long long>(static_cast<QSourceLocation*>(ptr)->column());
}

int QSourceLocation_IsNull(void* ptr){
	return static_cast<QSourceLocation*>(ptr)->isNull();
}

long long QSourceLocation_Line(void* ptr){
	return static_cast<long long>(static_cast<QSourceLocation*>(ptr)->line());
}

void QSourceLocation_SetColumn(void* ptr, long long newColumn){
	static_cast<QSourceLocation*>(ptr)->setColumn(static_cast<long long>(newColumn));
}

void QSourceLocation_SetLine(void* ptr, long long newLine){
	static_cast<QSourceLocation*>(ptr)->setLine(static_cast<long long>(newLine));
}

void QSourceLocation_SetUri(void* ptr, void* newUri){
	static_cast<QSourceLocation*>(ptr)->setUri(*static_cast<QUrl*>(newUri));
}

void* QSourceLocation_Uri(void* ptr){
	return new QUrl(static_cast<QSourceLocation*>(ptr)->uri());
}

void QSourceLocation_DestroyQSourceLocation(void* ptr){
	static_cast<QSourceLocation*>(ptr)->~QSourceLocation();
}

class MyQXmlFormatter: public QXmlFormatter {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlFormatter(const QXmlQuery &query, QIODevice *outputDevice) : QXmlFormatter(query, outputDevice) {};
	void characters(const QStringRef & value) { callbackQXmlFormatterCharacters(this, this->objectNameAbs().toUtf8().data(), new QStringRef(value)); };
	void comment(const QString & value) { callbackQXmlFormatterComment(this, this->objectNameAbs().toUtf8().data(), value.toUtf8().data()); };
	void endDocument() { callbackQXmlFormatterEndDocument(this, this->objectNameAbs().toUtf8().data()); };
	void endElement() { callbackQXmlFormatterEndElement(this, this->objectNameAbs().toUtf8().data()); };
	void endOfSequence() { callbackQXmlFormatterEndOfSequence(this, this->objectNameAbs().toUtf8().data()); };
	void startDocument() { callbackQXmlFormatterStartDocument(this, this->objectNameAbs().toUtf8().data()); };
	void startOfSequence() { callbackQXmlFormatterStartOfSequence(this, this->objectNameAbs().toUtf8().data()); };
};

void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice){
	return new MyQXmlFormatter(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlFormatter_Characters(void* ptr, void* value){
	static_cast<MyQXmlFormatter*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlFormatter_CharactersDefault(void* ptr, void* value){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::characters(*static_cast<QStringRef*>(value));
}

void QXmlFormatter_Comment(void* ptr, char* value){
	static_cast<MyQXmlFormatter*>(ptr)->comment(QString(value));
}

void QXmlFormatter_CommentDefault(void* ptr, char* value){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::comment(QString(value));
}

void QXmlFormatter_EndDocument(void* ptr){
	static_cast<MyQXmlFormatter*>(ptr)->endDocument();
}

void QXmlFormatter_EndDocumentDefault(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endDocument();
}

void QXmlFormatter_EndElement(void* ptr){
	static_cast<MyQXmlFormatter*>(ptr)->endElement();
}

void QXmlFormatter_EndElementDefault(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endElement();
}

void QXmlFormatter_EndOfSequence(void* ptr){
	static_cast<MyQXmlFormatter*>(ptr)->endOfSequence();
}

void QXmlFormatter_EndOfSequenceDefault(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::endOfSequence();
}

int QXmlFormatter_IndentationDepth(void* ptr){
	return static_cast<QXmlFormatter*>(ptr)->indentationDepth();
}

void QXmlFormatter_SetIndentationDepth(void* ptr, int depth){
	static_cast<QXmlFormatter*>(ptr)->setIndentationDepth(depth);
}

void QXmlFormatter_StartDocument(void* ptr){
	static_cast<MyQXmlFormatter*>(ptr)->startDocument();
}

void QXmlFormatter_StartDocumentDefault(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::startDocument();
}

void QXmlFormatter_StartOfSequence(void* ptr){
	static_cast<MyQXmlFormatter*>(ptr)->startOfSequence();
}

void QXmlFormatter_StartOfSequenceDefault(void* ptr){
	static_cast<QXmlFormatter*>(ptr)->QXmlFormatter::startOfSequence();
}

char* QXmlFormatter_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlFormatter*>(static_cast<QXmlFormatter*>(ptr))) {
		return static_cast<MyQXmlFormatter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlFormatter_BASE").toUtf8().data();
}

void QXmlFormatter_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlFormatter*>(static_cast<QXmlFormatter*>(ptr))) {
		static_cast<MyQXmlFormatter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

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

void* QXmlItem_ToNodeModelIndex(void* ptr){
	return new QXmlNodeModelIndex(static_cast<QXmlItem*>(ptr)->toNodeModelIndex());
}

void QXmlItem_DestroyQXmlItem(void* ptr){
	static_cast<QXmlItem*>(ptr)->~QXmlItem();
}

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

void* QXmlNamePool_NewQXmlNamePool(){
	return new QXmlNamePool();
}

void* QXmlNamePool_NewQXmlNamePool2(void* other){
	return new QXmlNamePool(*static_cast<QXmlNamePool*>(other));
}

void QXmlNamePool_DestroyQXmlNamePool(void* ptr){
	static_cast<QXmlNamePool*>(ptr)->~QXmlNamePool();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex(){
	return new QXmlNodeModelIndex();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other){
	return new QXmlNodeModelIndex(*static_cast<QXmlNodeModelIndex*>(other));
}

long long QXmlNodeModelIndex_AdditionalData(void* ptr){
	return static_cast<long long>(static_cast<QXmlNodeModelIndex*>(ptr)->additionalData());
}

long long QXmlNodeModelIndex_Data(void* ptr){
	return static_cast<long long>(static_cast<QXmlNodeModelIndex*>(ptr)->data());
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

void* QXmlQuery_NamePool(void* ptr){
	return new QXmlNamePool(static_cast<QXmlQuery*>(ptr)->namePool());
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

class MyQXmlResultItems: public QXmlResultItems {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlResultItems() : QXmlResultItems() {};
};

void* QXmlResultItems_NewQXmlResultItems(){
	return new MyQXmlResultItems();
}

void* QXmlResultItems_Current(void* ptr){
	return new QXmlItem(static_cast<QXmlResultItems*>(ptr)->current());
}

int QXmlResultItems_HasError(void* ptr){
	return static_cast<QXmlResultItems*>(ptr)->hasError();
}

void* QXmlResultItems_Next(void* ptr){
	return new QXmlItem(static_cast<QXmlResultItems*>(ptr)->next());
}

void QXmlResultItems_DestroyQXmlResultItems(void* ptr){
	static_cast<QXmlResultItems*>(ptr)->~QXmlResultItems();
}

char* QXmlResultItems_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlResultItems*>(static_cast<QXmlResultItems*>(ptr))) {
		return static_cast<MyQXmlResultItems*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlResultItems_BASE").toUtf8().data();
}

void QXmlResultItems_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlResultItems*>(static_cast<QXmlResultItems*>(ptr))) {
		static_cast<MyQXmlResultItems*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QXmlSchema_NewQXmlSchema(){
	return new QXmlSchema();
}

void* QXmlSchema_NewQXmlSchema2(void* other){
	return new QXmlSchema(*static_cast<QXmlSchema*>(other));
}

void* QXmlSchema_DocumentUri(void* ptr){
	return new QUrl(static_cast<QXmlSchema*>(ptr)->documentUri());
}

int QXmlSchema_IsValid(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->isValid();
}

int QXmlSchema_Load2(void* ptr, void* source, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(static_cast<QIODevice*>(source), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load3(void* ptr, char* data, void* documentUri){
	return static_cast<QXmlSchema*>(ptr)->load(QByteArray(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchema_Load(void* ptr, void* source){
	return static_cast<QXmlSchema*>(ptr)->load(*static_cast<QUrl*>(source));
}

void* QXmlSchema_MessageHandler(void* ptr){
	return static_cast<QXmlSchema*>(ptr)->messageHandler();
}

void* QXmlSchema_NamePool(void* ptr){
	return new QXmlNamePool(static_cast<QXmlSchema*>(ptr)->namePool());
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

void* QXmlSchemaValidator_NewQXmlSchemaValidator(){
	return new QXmlSchemaValidator();
}

void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema){
	return new QXmlSchemaValidator(*static_cast<QXmlSchema*>(schema));
}

void* QXmlSchemaValidator_MessageHandler(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->messageHandler();
}

void* QXmlSchemaValidator_NamePool(void* ptr){
	return new QXmlNamePool(static_cast<QXmlSchemaValidator*>(ptr)->namePool());
}

void* QXmlSchemaValidator_NetworkAccessManager(void* ptr){
	return static_cast<QXmlSchemaValidator*>(ptr)->networkAccessManager();
}

void* QXmlSchemaValidator_Schema(void* ptr){
	return new QXmlSchema(static_cast<QXmlSchemaValidator*>(ptr)->schema());
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

int QXmlSchemaValidator_Validate3(void* ptr, char* data, void* documentUri){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(QByteArray(data), *static_cast<QUrl*>(documentUri));
}

int QXmlSchemaValidator_Validate(void* ptr, void* source){
	return static_cast<QXmlSchemaValidator*>(ptr)->validate(*static_cast<QUrl*>(source));
}

void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr){
	static_cast<QXmlSchemaValidator*>(ptr)->~QXmlSchemaValidator();
}

class MyQXmlSerializer: public QXmlSerializer {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlSerializer(const QXmlQuery &query, QIODevice *outputDevice) : QXmlSerializer(query, outputDevice) {};
	void characters(const QStringRef & value) { callbackQXmlSerializerCharacters(this, this->objectNameAbs().toUtf8().data(), new QStringRef(value)); };
	void comment(const QString & value) { callbackQXmlSerializerComment(this, this->objectNameAbs().toUtf8().data(), value.toUtf8().data()); };
	void endDocument() { callbackQXmlSerializerEndDocument(this, this->objectNameAbs().toUtf8().data()); };
	void endElement() { callbackQXmlSerializerEndElement(this, this->objectNameAbs().toUtf8().data()); };
	void endOfSequence() { callbackQXmlSerializerEndOfSequence(this, this->objectNameAbs().toUtf8().data()); };
	void startDocument() { callbackQXmlSerializerStartDocument(this, this->objectNameAbs().toUtf8().data()); };
	void startOfSequence() { callbackQXmlSerializerStartOfSequence(this, this->objectNameAbs().toUtf8().data()); };
};

void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice){
	return new MyQXmlSerializer(*static_cast<QXmlQuery*>(query), static_cast<QIODevice*>(outputDevice));
}

void QXmlSerializer_Characters(void* ptr, void* value){
	static_cast<MyQXmlSerializer*>(ptr)->characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_CharactersDefault(void* ptr, void* value){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::characters(*static_cast<QStringRef*>(value));
}

void QXmlSerializer_Comment(void* ptr, char* value){
	static_cast<MyQXmlSerializer*>(ptr)->comment(QString(value));
}

void QXmlSerializer_CommentDefault(void* ptr, char* value){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::comment(QString(value));
}

void QXmlSerializer_EndDocument(void* ptr){
	static_cast<MyQXmlSerializer*>(ptr)->endDocument();
}

void QXmlSerializer_EndDocumentDefault(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endDocument();
}

void QXmlSerializer_EndElement(void* ptr){
	static_cast<MyQXmlSerializer*>(ptr)->endElement();
}

void QXmlSerializer_EndElementDefault(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endElement();
}

void* QXmlSerializer_Codec(void* ptr){
	return const_cast<QTextCodec*>(static_cast<QXmlSerializer*>(ptr)->codec());
}

void QXmlSerializer_EndOfSequence(void* ptr){
	static_cast<MyQXmlSerializer*>(ptr)->endOfSequence();
}

void QXmlSerializer_EndOfSequenceDefault(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::endOfSequence();
}

void* QXmlSerializer_OutputDevice(void* ptr){
	return static_cast<QXmlSerializer*>(ptr)->outputDevice();
}

void QXmlSerializer_SetCodec(void* ptr, void* outputCodec){
	static_cast<QXmlSerializer*>(ptr)->setCodec(static_cast<QTextCodec*>(outputCodec));
}

void QXmlSerializer_StartDocument(void* ptr){
	static_cast<MyQXmlSerializer*>(ptr)->startDocument();
}

void QXmlSerializer_StartDocumentDefault(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::startDocument();
}

void QXmlSerializer_StartOfSequence(void* ptr){
	static_cast<MyQXmlSerializer*>(ptr)->startOfSequence();
}

void QXmlSerializer_StartOfSequenceDefault(void* ptr){
	static_cast<QXmlSerializer*>(ptr)->QXmlSerializer::startOfSequence();
}

char* QXmlSerializer_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlSerializer*>(static_cast<QXmlSerializer*>(ptr))) {
		return static_cast<MyQXmlSerializer*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlSerializer_BASE").toUtf8().data();
}

void QXmlSerializer_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlSerializer*>(static_cast<QXmlSerializer*>(ptr))) {
		static_cast<MyQXmlSerializer*>(ptr)->setObjectNameAbs(QString(name));
	}
}

