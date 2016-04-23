#define protected public

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
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneWheelEvent>
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
#include <QPaintDevice>
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
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsSvgItemPaint(this, this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void advance(int phase) { callbackQGraphicsSvgItemAdvance(this, this->objectName().toUtf8().data(), phase); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsSvgItemContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItemDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItemDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItemDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItemDropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQGraphicsSvgItemFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQGraphicsSvgItemFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItemHoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItemHoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItemHoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQGraphicsSvgItemInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsSvgItemKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsSvgItemKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItemMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItemMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItemMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItemMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQGraphicsSvgItemWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsSvgItemTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsSvgItemChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGraphicsSvgItemCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QGraphicsSvgItem_BoundingRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QGraphicsSvgItem*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QGraphicsSvgItem*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QGraphicsSvgItem*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QGraphicsSvgItem*>(ptr)->boundingRect()).height());
}

char* QGraphicsSvgItem_ElementId(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8().data();
}

void* QGraphicsSvgItem_MaximumCacheSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize()).width(), static_cast<QSize>(static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize()).height());
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<MyQGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSvgItem_PaintDefault(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
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

void QGraphicsSvgItem_Advance(void* ptr, int phase){
	static_cast<MyQGraphicsSvgItem*>(ptr)->advance(phase);
}

void QGraphicsSvgItem_AdvanceDefault(void* ptr, int phase){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::advance(phase);
}

void QGraphicsSvgItem_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsSvgItem_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsSvgItem_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DropEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DropEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_FocusInEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusInEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_HoverEnterEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverEnterEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverLeaveEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverLeaveEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverMoveEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverMoveEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsSvgItem_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsSvgItem_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MousePressEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MousePressEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_WheelEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsSvgItem_WheelEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsSvgItem_TimerEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsSvgItem_TimerEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsSvgItem_ChildEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsSvgItem_ChildEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsSvgItem_CustomEvent(void* ptr, void* event){
	static_cast<MyQGraphicsSvgItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_CustomEventDefault(void* ptr, void* event){
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::customEvent(static_cast<QEvent*>(event));
}

class MyQSvgGenerator: public QSvgGenerator {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQSvgGenerator() : QSvgGenerator() {};
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

void* QSvgGenerator_ViewBoxF(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSvgGenerator*>(ptr)->viewBoxF()).x(), static_cast<QRectF>(static_cast<QSvgGenerator*>(ptr)->viewBoxF()).y(), static_cast<QRectF>(static_cast<QSvgGenerator*>(ptr)->viewBoxF()).width(), static_cast<QRectF>(static_cast<QSvgGenerator*>(ptr)->viewBoxF()).height());
}

void* QSvgGenerator_NewQSvgGenerator(){
	return new MyQSvgGenerator();
}

int QSvgGenerator_Metric(void* ptr, int metric){
	return static_cast<QSvgGenerator*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QSvgGenerator_PaintEngine(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->paintEngine();
}

void* QSvgGenerator_ViewBox(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).x(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).y(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).width(), static_cast<QRect>(static_cast<QSvgGenerator*>(ptr)->viewBox()).height());
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr){
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

char* QSvgGenerator_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSvgGenerator*>(static_cast<QSvgGenerator*>(ptr))) {
		return static_cast<MyQSvgGenerator*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSvgGenerator_BASE").toUtf8().data();
}

void QSvgGenerator_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSvgGenerator*>(static_cast<QSvgGenerator*>(ptr))) {
		static_cast<MyQSvgGenerator*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSvgRenderer: public QSvgRenderer {
public:
	void Signal_RepaintNeeded() { callbackQSvgRendererRepaintNeeded(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSvgRendererTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSvgRendererChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSvgRendererCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void* QSvgRenderer_ViewBoxF(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->viewBoxF()).x(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->viewBoxF()).y(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->viewBoxF()).width(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->viewBoxF()).height());
}

void* QSvgRenderer_NewQSvgRenderer(void* parent){
	return new QSvgRenderer(static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent){
	return new QSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer3(char* contents, void* parent){
	return new QSvgRenderer(QByteArray(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent){
	return new QSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

int QSvgRenderer_Animated(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

void* QSvgRenderer_BoundsOnElement(void* ptr, char* id){
	return new QRectF(static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString(id))).x(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString(id))).y(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString(id))).width(), static_cast<QRectF>(static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString(id))).height());
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

int QSvgRenderer_Load2(void* ptr, char* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QByteArray, QByteArray(contents)));
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

void QSvgRenderer_RepaintNeeded(void* ptr){
	static_cast<QSvgRenderer*>(ptr)->repaintNeeded();
}

void* QSvgRenderer_ViewBox(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).x(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).y(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).width(), static_cast<QRect>(static_cast<QSvgRenderer*>(ptr)->viewBox()).height());
}

void QSvgRenderer_DestroyQSvgRenderer(void* ptr){
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

void QSvgRenderer_TimerEvent(void* ptr, void* event){
	static_cast<MyQSvgRenderer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgRenderer_TimerEventDefault(void* ptr, void* event){
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgRenderer_ChildEvent(void* ptr, void* event){
	static_cast<MyQSvgRenderer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSvgRenderer_ChildEventDefault(void* ptr, void* event){
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgRenderer_CustomEvent(void* ptr, void* event){
	static_cast<MyQSvgRenderer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSvgRenderer_CustomEventDefault(void* ptr, void* event){
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::customEvent(static_cast<QEvent*>(event));
}

class MyQSvgWidget: public QSvgWidget {
public:
	MyQSvgWidget(QWidget *parent) : QSvgWidget(parent) {};
	MyQSvgWidget(const QString &file, QWidget *parent) : QSvgWidget(file, parent) {};
	void paintEvent(QPaintEvent * event) { callbackQSvgWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void actionEvent(QActionEvent * event) { callbackQSvgWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQSvgWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQSvgWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQSvgWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQSvgWidgetDropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQSvgWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQSvgWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQSvgWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQSvgWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQSvgWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQSvgWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQSvgWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QSvgWidget::setVisible(visible); }; };
	void showEvent(QShowEvent * event) { callbackQSvgWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void changeEvent(QEvent * event) { callbackQSvgWidgetChangeEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQSvgWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQSvgWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQSvgWidgetInitPainter(const_cast<MyQSvgWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQSvgWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQSvgWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQSvgWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQSvgWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQSvgWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQSvgWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQSvgWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQSvgWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQSvgWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QWheelEvent * event) { callbackQSvgWidgetWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQSvgWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSvgWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSvgWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSvgWidget_NewQSvgWidget(void* parent){
	return new MyQSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(char* file, void* parent){
	return new MyQSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, char* contents){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, QByteArray(contents)));
}

void QSvgWidget_Load(void* ptr, char* file){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

void QSvgWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QSvgWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::paintEvent(static_cast<QPaintEvent*>(event));
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

void QSvgWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QSvgWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QSvgWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QSvgWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QSvgWidget_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QSvgWidget_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QSvgWidget_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QSvgWidget_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QSvgWidget_DropEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QSvgWidget_DropEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QSvgWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::enterEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QSvgWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QSvgWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QSvgWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QSvgWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQSvgWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QSvgWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QSvgWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QSvgWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QSvgWidget_ChangeEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_ChangeEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::changeEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QSvgWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QSvgWidget_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QSvgWidget_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QSvgWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQSvgWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QSvgWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::initPainter(static_cast<QPainter*>(painter));
}

void QSvgWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QSvgWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QSvgWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QSvgWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QSvgWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QSvgWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QSvgWidget_WheelEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QSvgWidget_WheelEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QSvgWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSvgWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQSvgWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::customEvent(static_cast<QEvent*>(event));
}

