#include "svg.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QGraphicsSvgItem>
#include <QHideEvent>
#include <QIODevice>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QSvgGenerator>
#include <QSvgRenderer>
#include <QSvgWidget>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWheelEvent>
#include <QWidget>
#include <QXmlStreamReader>

class MyQGraphicsSvgItem: public QGraphicsSvgItem {
public:
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { if (!callbackQGraphicsSvgItemPaint(this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget)) { QGraphicsSvgItem::paint(painter, option, widget); }; };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQGraphicsSvgItemTimerEvent(this->objectName().toUtf8().data(), event)) { QGraphicsSvgItem::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQGraphicsSvgItemChildEvent(this->objectName().toUtf8().data(), event)) { QGraphicsSvgItem::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQGraphicsSvgItemCustomEvent(this->objectName().toUtf8().data(), event)) { QGraphicsSvgItem::customEvent(event); }; };
};

char* QGraphicsSvgItem_ElementId(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8().data();
}

void* QGraphicsSvgItem_MaximumCacheSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize()).width(), static_cast<QSize>(static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize()).height());
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void* QGraphicsSvgItem_Renderer(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

void QGraphicsSvgItem_SetElementId(void* ptr, char* id){
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString(id));
}

void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size){
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer){
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

int QGraphicsSvgItem_Type(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->type();
}

class MyQSvgGenerator: public QSvgGenerator {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQSvgGenerator() : QSvgGenerator() {};
protected:
};

char* QSvgGenerator_Description(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->description().toUtf8().data();
}

char* QSvgGenerator_FileName(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8().data();
}

void* QSvgGenerator_OutputDevice(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

int QSvgGenerator_Resolution(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}

void QSvgGenerator_SetDescription(void* ptr, char* description){
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString(description));
}

void QSvgGenerator_SetFileName(void* ptr, char* fileName){
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString(fileName));
}

void QSvgGenerator_SetOutputDevice(void* ptr, void* outputDevice){
	static_cast<QSvgGenerator*>(ptr)->setOutputDevice(static_cast<QIODevice*>(outputDevice));
}

void QSvgGenerator_SetResolution(void* ptr, int dpi){
	static_cast<QSvgGenerator*>(ptr)->setResolution(dpi);
}

void QSvgGenerator_SetSize(void* ptr, void* size){
	static_cast<QSvgGenerator*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QSvgGenerator_SetTitle(void* ptr, char* title){
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString(title));
}

void QSvgGenerator_SetViewBox(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

void* QSvgGenerator_Size(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSvgGenerator*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QSvgGenerator*>(ptr)->size()).height());
}

char* QSvgGenerator_Title(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->title().toUtf8().data();
}

void* QSvgGenerator_NewQSvgGenerator(){
	return new MyQSvgGenerator();
}

void* QSvgGenerator_ViewBox(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).x(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).y(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).width(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).height());
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr){
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

char* QSvgGenerator_ObjectNameAbs(void* ptr){
	return static_cast<MyQSvgGenerator*>(ptr)->objectNameAbs().toUtf8().data();
}

void QSvgGenerator_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQSvgGenerator*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQSvgRenderer: public QSvgRenderer {
public:
	void Signal_RepaintNeeded() { callbackQSvgRendererRepaintNeeded(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQSvgRendererTimerEvent(this->objectName().toUtf8().data(), event)) { QSvgRenderer::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQSvgRendererChildEvent(this->objectName().toUtf8().data(), event)) { QSvgRenderer::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQSvgRendererCustomEvent(this->objectName().toUtf8().data(), event)) { QSvgRenderer::customEvent(event); }; };
};

int QSvgRenderer_FramesPerSecond(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(void* ptr, int num){
	static_cast<QSvgRenderer*>(ptr)->setFramesPerSecond(num);
}

void QSvgRenderer_SetViewBox(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRect*>(viewbox));
}

void QSvgRenderer_SetViewBox2(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRectF*>(viewbox));
}

void* QSvgRenderer_NewQSvgRenderer(void* parent){
	return new QSvgRenderer(static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent){
	return new QSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent){
	return new QSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent){
	return new QSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

int QSvgRenderer_Animated(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

void* QSvgRenderer_DefaultSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSvgRenderer*>(ptr)->defaultSize()).width(), static_cast<QSize>(static_cast<QSvgRenderer*>(ptr)->defaultSize()).height());
}

int QSvgRenderer_ElementExists(void* ptr, char* id){
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString(id));
}

int QSvgRenderer_IsValid(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

int QSvgRenderer_Load3(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
}

int QSvgRenderer_Load2(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

int QSvgRenderer_Load(void* ptr, char* filename){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QString, QString(filename)));
}

void QSvgRenderer_Render(void* ptr, void* painter){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3(void* ptr, void* painter, char* elementId, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QString, QString(elementId)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_ConnectRepaintNeeded(void* ptr){
	QObject::connect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DisconnectRepaintNeeded(void* ptr){
	QObject::disconnect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void* QSvgRenderer_ViewBox(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).x(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).y(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).width(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).height());
}

void QSvgRenderer_DestroyQSvgRenderer(void* ptr){
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

class MyQSvgWidget: public QSvgWidget {
public:
	MyQSvgWidget(QWidget *parent) : QSvgWidget(parent) {};
	MyQSvgWidget(const QString &file, QWidget *parent) : QSvgWidget(file, parent) {};
	void setVisible(bool visible) { if (!callbackQSvgWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QSvgWidget::setVisible(visible); }; };
protected:
	void paintEvent(QPaintEvent * event) { if (!callbackQSvgWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::paintEvent(event); }; };
	void actionEvent(QActionEvent * event) { if (!callbackQSvgWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::actionEvent(event); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQSvgWidgetDragEnterEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQSvgWidgetDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQSvgWidgetDragMoveEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQSvgWidgetDropEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::dropEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQSvgWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::enterEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQSvgWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::focusOutEvent(event); }; };
	void hideEvent(QHideEvent * event) { if (!callbackQSvgWidgetHideEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::hideEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQSvgWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQSvgWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::moveEvent(event); }; };
	void showEvent(QShowEvent * event) { if (!callbackQSvgWidgetShowEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::showEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQSvgWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::closeEvent(event); }; };
	void contextMenuEvent(QContextMenuEvent * event) { if (!callbackQSvgWidgetContextMenuEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::contextMenuEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQSvgWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QSvgWidget::initPainter(painter); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQSvgWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQSvgWidgetKeyPressEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQSvgWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQSvgWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQSvgWidgetMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQSvgWidgetMousePressEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQSvgWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::mouseReleaseEvent(event); }; };
	void resizeEvent(QResizeEvent * event) { if (!callbackQSvgWidgetResizeEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::resizeEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQSvgWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::tabletEvent(event); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQSvgWidgetWheelEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQSvgWidgetTimerEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQSvgWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQSvgWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::customEvent(event); }; };
};

void* QSvgWidget_NewQSvgWidget(void* parent){
	return new MyQSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(char* file, void* parent){
	return new MyQSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, void* contents){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load(void* ptr, char* file){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

void* QSvgWidget_Renderer(void* ptr){
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void* QSvgWidget_SizeHint(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSvgWidget*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QSvgWidget*>(ptr)->sizeHint()).height());
}

void QSvgWidget_DestroyQSvgWidget(void* ptr){
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

