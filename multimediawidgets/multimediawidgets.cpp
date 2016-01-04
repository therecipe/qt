#define protected public

#include "multimediawidgets.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QCamera>
#include <QCameraViewfinder>
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
#include <QGraphicsItem>
#include <QGraphicsVideoItem>
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMediaObject>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QPointF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVideoWidget>
#include <QVideoWidgetControl>
#include <QWheelEvent>
#include <QWidget>

class MyQCameraViewfinder: public QCameraViewfinder {
public:
	MyQCameraViewfinder(QWidget *parent) : QCameraViewfinder(parent) {};
	void hideEvent(QHideEvent * event) { callbackQCameraViewfinderHideEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQCameraViewfinderMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQCameraViewfinderPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQCameraViewfinderResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void showEvent(QShowEvent * event) { callbackQCameraViewfinderShowEvent(this, this->objectName().toUtf8().data(), event); };
	void actionEvent(QActionEvent * event) { callbackQCameraViewfinderActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQCameraViewfinderDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQCameraViewfinderDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQCameraViewfinderDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQCameraViewfinderDropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQCameraViewfinderEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQCameraViewfinderFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQCameraViewfinderFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQCameraViewfinderLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQCameraViewfinderSetVisible(this, this->objectName().toUtf8().data(), visible)) { QCameraViewfinder::setVisible(visible); }; };
	void changeEvent(QEvent * event) { callbackQCameraViewfinderChangeEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQCameraViewfinderCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQCameraViewfinderContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQCameraViewfinderInitPainter(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQCameraViewfinderInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQCameraViewfinderKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQCameraViewfinderKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQCameraViewfinderMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQCameraViewfinderMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQCameraViewfinderMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQCameraViewfinderMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQCameraViewfinderTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QWheelEvent * event) { callbackQCameraViewfinderWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQCameraViewfinderTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCameraViewfinderChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQCameraViewfinderCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QCameraViewfinder_NewQCameraViewfinder(void* parent){
	return new MyQCameraViewfinder(static_cast<QWidget*>(parent));
}

void* QCameraViewfinder_MediaObject(void* ptr){
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

int QCameraViewfinder_SetMediaObject(void* ptr, void* object){
	return static_cast<QCameraViewfinder*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr){
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

void QCameraViewfinder_HideEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QCameraViewfinder_HideEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::hideEvent(static_cast<QHideEvent*>(event));
}

void QCameraViewfinder_MoveEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QCameraViewfinder_MoveEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::moveEvent(static_cast<QMoveEvent*>(event));
}

void QCameraViewfinder_PaintEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QCameraViewfinder_PaintEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::paintEvent(static_cast<QPaintEvent*>(event));
}

void QCameraViewfinder_ResizeEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QCameraViewfinder_ResizeEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QCameraViewfinder_ShowEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QCameraViewfinder_ShowEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showEvent(static_cast<QShowEvent*>(event));
}

void QCameraViewfinder_ActionEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_ActionEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DropEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_DropEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_EnterEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_EnterEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_FocusInEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusInEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_LeaveEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_LeaveEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::leaveEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQCameraViewfinder*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QCameraViewfinder_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QCameraViewfinder_ChangeEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_ChangeEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::changeEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_CloseEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_CloseEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QCameraViewfinder_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QCameraViewfinder_InitPainter(void* ptr, void* painter){
	static_cast<MyQCameraViewfinder*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QCameraViewfinder_InitPainterDefault(void* ptr, void* painter){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::initPainter(static_cast<QPainter*>(painter));
}

void QCameraViewfinder_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QCameraViewfinder_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QCameraViewfinder_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_TabletEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_TabletEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_WheelEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_WheelEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_TimerEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_TimerEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_ChildEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_ChildEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_CustomEvent(void* ptr, void* event){
	static_cast<MyQCameraViewfinder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_CustomEventDefault(void* ptr, void* event){
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::customEvent(static_cast<QEvent*>(event));
}

class MyQGraphicsVideoItem: public QGraphicsVideoItem {
public:
	MyQGraphicsVideoItem(QGraphicsItem *parent) : QGraphicsVideoItem(parent) {};
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsVideoItemPaint(this, this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsVideoItemTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsVideoItemChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGraphicsVideoItemCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent){
	return new MyQGraphicsVideoItem(static_cast<QGraphicsItem*>(parent));
}

int QGraphicsVideoItem_AspectRatioMode(void* ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->aspectRatioMode();
}

void* QGraphicsVideoItem_MediaObject(void* ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->mediaObject();
}

void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<MyQGraphicsVideoItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_PaintDefault(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QGraphicsVideoItem*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QGraphicsVideoItem_SetOffset(void* ptr, void* offset){
	static_cast<QGraphicsVideoItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsVideoItem_SetSize(void* ptr, void* size){
	static_cast<QGraphicsVideoItem*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr){
	static_cast<QGraphicsVideoItem*>(ptr)->~QGraphicsVideoItem();
}

void QGraphicsVideoItem_TimerEvent(void* ptr, void* event){
	static_cast<MyQGraphicsVideoItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_TimerEventDefault(void* ptr, void* event){
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_ChildEvent(void* ptr, void* event){
	static_cast<MyQGraphicsVideoItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_ChildEventDefault(void* ptr, void* event){
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_CustomEvent(void* ptr, void* event){
	static_cast<MyQGraphicsVideoItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsVideoItem_CustomEventDefault(void* ptr, void* event){
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::customEvent(static_cast<QEvent*>(event));
}

class MyQVideoWidget: public QVideoWidget {
public:
	MyQVideoWidget(QWidget *parent) : QVideoWidget(parent) {};
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidgetBrightnessChanged(this, this->objectName().toUtf8().data(), brightness); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidgetContrastChanged(this, this->objectName().toUtf8().data(), contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidgetFullScreenChanged(this, this->objectName().toUtf8().data(), fullScreen); };
	void hideEvent(QHideEvent * event) { callbackQVideoWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_HueChanged(int hue) { callbackQVideoWidgetHueChanged(this, this->objectName().toUtf8().data(), hue); };
	void moveEvent(QMoveEvent * event) { callbackQVideoWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQVideoWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQVideoWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidgetSaturationChanged(this, this->objectName().toUtf8().data(), saturation); };
	void showEvent(QShowEvent * event) { callbackQVideoWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void actionEvent(QActionEvent * event) { callbackQVideoWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQVideoWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQVideoWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQVideoWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQVideoWidgetDropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQVideoWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQVideoWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQVideoWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQVideoWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQVideoWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QVideoWidget::setVisible(visible); }; };
	void changeEvent(QEvent * event) { callbackQVideoWidgetChangeEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQVideoWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQVideoWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQVideoWidgetInitPainter(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQVideoWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQVideoWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQVideoWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQVideoWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQVideoWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQVideoWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQVideoWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQVideoWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QWheelEvent * event) { callbackQVideoWidgetWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQVideoWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QVideoWidget_AspectRatioMode(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->aspectRatioMode();
}

int QVideoWidget_Brightness(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->brightness();
}

int QVideoWidget_Contrast(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->contrast();
}

int QVideoWidget_Hue(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->hue();
}

void* QVideoWidget_MediaObject(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->mediaObject();
}

int QVideoWidget_Saturation(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->saturation();
}

void QVideoWidget_SetAspectRatioMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QVideoWidget_SetBrightness(void* ptr, int brightness){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QVideoWidget_SetContrast(void* ptr, int contrast){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QVideoWidget_SetFullScreen(void* ptr, int fullScreen){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QVideoWidget_SetHue(void* ptr, int hue){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHue", Q_ARG(int, hue));
}

void QVideoWidget_SetSaturation(void* ptr, int saturation){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

void* QVideoWidget_NewQVideoWidget(void* parent){
	return new MyQVideoWidget(static_cast<QWidget*>(parent));
}

void QVideoWidget_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));;
}

void QVideoWidget_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));;
}

void QVideoWidget_BrightnessChanged(void* ptr, int brightness){
	static_cast<QVideoWidget*>(ptr)->brightnessChanged(brightness);
}

void QVideoWidget_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));;
}

void QVideoWidget_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));;
}

void QVideoWidget_ContrastChanged(void* ptr, int contrast){
	static_cast<QVideoWidget*>(ptr)->contrastChanged(contrast);
}

int QVideoWidget_Event(void* ptr, void* event){
	return static_cast<QVideoWidget*>(ptr)->event(static_cast<QEvent*>(event));
}

void QVideoWidget_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));;
}

void QVideoWidget_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));;
}

void QVideoWidget_FullScreenChanged(void* ptr, int fullScreen){
	static_cast<QVideoWidget*>(ptr)->fullScreenChanged(fullScreen != 0);
}

void QVideoWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QVideoWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QVideoWidget_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));;
}

void QVideoWidget_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));;
}

void QVideoWidget_HueChanged(void* ptr, int hue){
	static_cast<QVideoWidget*>(ptr)->hueChanged(hue);
}

int QVideoWidget_IsFullScreen(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->isFullScreen();
}

void QVideoWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QVideoWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QVideoWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QVideoWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QVideoWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QVideoWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QVideoWidget_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));;
}

void QVideoWidget_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));;
}

void QVideoWidget_SaturationChanged(void* ptr, int saturation){
	static_cast<QVideoWidget*>(ptr)->saturationChanged(saturation);
}

void QVideoWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QVideoWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QVideoWidget_SizeHint(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->sizeHint()).height());
}

void QVideoWidget_DestroyQVideoWidget(void* ptr){
	static_cast<QVideoWidget*>(ptr)->~QVideoWidget();
}

void QVideoWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DropEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_DropEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQVideoWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QVideoWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QVideoWidget_ChangeEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_ChangeEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::changeEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QVideoWidget_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QVideoWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQVideoWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QVideoWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::initPainter(static_cast<QPainter*>(painter));
}

void QVideoWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QVideoWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QVideoWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_WheelEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_WheelEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQVideoWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::customEvent(static_cast<QEvent*>(event));
}

class MyQVideoWidgetControl: public QVideoWidgetControl {
public:
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidgetControlBrightnessChanged(this, this->objectName().toUtf8().data(), brightness); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidgetControlContrastChanged(this, this->objectName().toUtf8().data(), contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidgetControlFullScreenChanged(this, this->objectName().toUtf8().data(), fullScreen); };
	void Signal_HueChanged(int hue) { callbackQVideoWidgetControlHueChanged(this, this->objectName().toUtf8().data(), hue); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidgetControlSaturationChanged(this, this->objectName().toUtf8().data(), saturation); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidgetControlTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidgetControlChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQVideoWidgetControlCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QVideoWidgetControl_AspectRatioMode(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->aspectRatioMode();
}

int QVideoWidgetControl_Brightness(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->brightness();
}

void QVideoWidgetControl_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

void QVideoWidgetControl_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

void QVideoWidgetControl_BrightnessChanged(void* ptr, int brightness){
	static_cast<QVideoWidgetControl*>(ptr)->brightnessChanged(brightness);
}

int QVideoWidgetControl_Contrast(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->contrast();
}

void QVideoWidgetControl_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_ContrastChanged(void* ptr, int contrast){
	static_cast<QVideoWidgetControl*>(ptr)->contrastChanged(contrast);
}

void QVideoWidgetControl_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

void QVideoWidgetControl_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

void QVideoWidgetControl_FullScreenChanged(void* ptr, int fullScreen){
	static_cast<QVideoWidgetControl*>(ptr)->fullScreenChanged(fullScreen != 0);
}

int QVideoWidgetControl_Hue(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->hue();
}

void QVideoWidgetControl_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

void QVideoWidgetControl_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

void QVideoWidgetControl_HueChanged(void* ptr, int hue){
	static_cast<QVideoWidgetControl*>(ptr)->hueChanged(hue);
}

int QVideoWidgetControl_IsFullScreen(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->isFullScreen();
}

int QVideoWidgetControl_Saturation(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->saturation();
}

void QVideoWidgetControl_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_SaturationChanged(void* ptr, int saturation){
	static_cast<QVideoWidgetControl*>(ptr)->saturationChanged(saturation);
}

void QVideoWidgetControl_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QVideoWidgetControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWidgetControl_SetBrightness(void* ptr, int brightness){
	static_cast<QVideoWidgetControl*>(ptr)->setBrightness(brightness);
}

void QVideoWidgetControl_SetContrast(void* ptr, int contrast){
	static_cast<QVideoWidgetControl*>(ptr)->setContrast(contrast);
}

void QVideoWidgetControl_SetFullScreen(void* ptr, int fullScreen){
	static_cast<QVideoWidgetControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWidgetControl_SetHue(void* ptr, int hue){
	static_cast<QVideoWidgetControl*>(ptr)->setHue(hue);
}

void QVideoWidgetControl_SetSaturation(void* ptr, int saturation){
	static_cast<QVideoWidgetControl*>(ptr)->setSaturation(saturation);
}

void* QVideoWidgetControl_VideoWidget(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->videoWidget();
}

void QVideoWidgetControl_DestroyQVideoWidgetControl(void* ptr){
	static_cast<QVideoWidgetControl*>(ptr)->~QVideoWidgetControl();
}

void QVideoWidgetControl_TimerEvent(void* ptr, void* event){
	static_cast<MyQVideoWidgetControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_TimerEventDefault(void* ptr, void* event){
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_ChildEvent(void* ptr, void* event){
	static_cast<MyQVideoWidgetControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_ChildEventDefault(void* ptr, void* event){
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_CustomEvent(void* ptr, void* event){
	static_cast<MyQVideoWidgetControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidgetControl_CustomEventDefault(void* ptr, void* event){
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::customEvent(static_cast<QEvent*>(event));
}

