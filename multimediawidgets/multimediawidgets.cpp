#define protected public
#define private public

#include "multimediawidgets.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
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
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsVideoItem>
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMediaObject>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVideoWidget>
#include <QVideoWidgetControl>
#include <QWheelEvent>
#include <QWidget>

class MyQCameraViewfinder: public QCameraViewfinder
{
public:
	MyQCameraViewfinder(QWidget *parent) : QCameraViewfinder(parent) {};
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQCameraViewfinder_MediaObject(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data())); };
	bool setMediaObject(QMediaObject * object) { return callbackQCameraViewfinder_SetMediaObject(this, this->objectName().toUtf8().data(), object) != 0; };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQCameraViewfinder_SetAspectRatioMode(this, this->objectName().toUtf8().data(), mode); };
	void setBrightness(int brightness) { callbackQCameraViewfinder_SetBrightness(this, this->objectName().toUtf8().data(), brightness); };
	void setContrast(int contrast) { callbackQCameraViewfinder_SetContrast(this, this->objectName().toUtf8().data(), contrast); };
	void setFullScreen(bool fullScreen) { callbackQCameraViewfinder_SetFullScreen(this, this->objectName().toUtf8().data(), fullScreen); };
	void setHue(int hue) { callbackQCameraViewfinder_SetHue(this, this->objectName().toUtf8().data(), hue); };
	void setSaturation(int saturation) { callbackQCameraViewfinder_SetSaturation(this, this->objectName().toUtf8().data(), saturation); };
	bool event(QEvent * event) { return callbackQCameraViewfinder_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	void hideEvent(QHideEvent * event) { callbackQCameraViewfinder_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQCameraViewfinder_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQCameraViewfinder_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQCameraViewfinder_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void showEvent(QShowEvent * event) { callbackQCameraViewfinder_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQCameraViewfinder_SizeHint(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data())); };
	void actionEvent(QActionEvent * event) { callbackQCameraViewfinder_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQCameraViewfinder_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQCameraViewfinder_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQCameraViewfinder_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQCameraViewfinder_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQCameraViewfinder_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQCameraViewfinder_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQCameraViewfinder_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQCameraViewfinder_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQCameraViewfinder_Metric(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQCameraViewfinder_MinimumSizeHint(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data())); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQCameraViewfinder_PaintEngine(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data())); };
	void setEnabled(bool vbo) { callbackQCameraViewfinder_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQCameraViewfinder_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQCameraViewfinder_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQCameraViewfinder_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQCameraViewfinder_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void changeEvent(QEvent * event) { callbackQCameraViewfinder_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQCameraViewfinder_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQCameraViewfinder_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQCameraViewfinder_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQCameraViewfinder_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQCameraViewfinder_HasHeightForWidth(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQCameraViewfinder_HeightForWidth(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQCameraViewfinder_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQCameraViewfinder_InitPainter(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQCameraViewfinder_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQCameraViewfinder_InputMethodQuery(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQCameraViewfinder_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQCameraViewfinder_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQCameraViewfinder_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQCameraViewfinder_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQCameraViewfinder_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQCameraViewfinder_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQCameraViewfinder_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQCameraViewfinder_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQCameraViewfinder_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQCameraViewfinder_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQCameraViewfinder_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQCameraViewfinder_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQCameraViewfinder_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQCameraViewfinder_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQCameraViewfinder_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQCameraViewfinder_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQCameraViewfinder_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQCameraViewfinder_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQCameraViewfinder_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQCameraViewfinder_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCameraViewfinder_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraViewfinder_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQCameraViewfinder_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQCameraViewfinder_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraViewfinder_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraViewfinder_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraViewfinder_MetaObject(const_cast<MyQCameraViewfinder*>(this), this->objectName().toUtf8().data())); };
};

void* QCameraViewfinder_NewQCameraViewfinder(void* parent)
{
	return new MyQCameraViewfinder(static_cast<QWidget*>(parent));
}

void* QCameraViewfinder_MediaObject(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

void* QCameraViewfinder_MediaObjectDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mediaObject();
}

int QCameraViewfinder_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QCameraViewfinder*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

int QCameraViewfinder_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setMediaObject(static_cast<QMediaObject*>(object));
}

void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

void QCameraViewfinder_SetAspectRatioMode(void* ptr, int mode)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QCameraViewfinder_SetBrightness(void* ptr, int brightness)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QCameraViewfinder_SetContrast(void* ptr, int contrast)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QCameraViewfinder_SetFullScreen(void* ptr, int fullScreen)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QCameraViewfinder_SetHue(void* ptr, int hue)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setHue", Q_ARG(int, hue));
}

void QCameraViewfinder_SetSaturation(void* ptr, int saturation)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

int QCameraViewfinder_Event(void* ptr, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->event(static_cast<QEvent*>(event));
}

int QCameraViewfinder_EventDefault(void* ptr, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::event(static_cast<QEvent*>(event));
}

void QCameraViewfinder_HideEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QCameraViewfinder_HideEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::hideEvent(static_cast<QHideEvent*>(event));
}

void QCameraViewfinder_MoveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QCameraViewfinder_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::moveEvent(static_cast<QMoveEvent*>(event));
}

void QCameraViewfinder_PaintEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QCameraViewfinder_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::paintEvent(static_cast<QPaintEvent*>(event));
}

void QCameraViewfinder_ResizeEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QCameraViewfinder_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QCameraViewfinder_ShowEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QCameraViewfinder_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showEvent(static_cast<QShowEvent*>(event));
}

void* QCameraViewfinder_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->sizeHint()).height());
}

void* QCameraViewfinder_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::sizeHint()).width(), static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::sizeHint()).height());
}

void QCameraViewfinder_ActionEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DropEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_DropEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_EnterEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_FocusInEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_LeaveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::leaveEvent(static_cast<QEvent*>(event));
}

int QCameraViewfinder_Metric(void* ptr, int m)
{
	return static_cast<QCameraViewfinder*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QCameraViewfinder_MetricDefault(void* ptr, int m)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QCameraViewfinder_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->minimumSizeHint()).height());
}

void* QCameraViewfinder_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::minimumSizeHint()).height());
}

void* QCameraViewfinder_PaintEngine(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->paintEngine();
}

void* QCameraViewfinder_PaintEngineDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::paintEngine();
}

void QCameraViewfinder_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QCameraViewfinder_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QCameraViewfinder_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QCameraViewfinder_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setVisible(visible != 0);
}

void QCameraViewfinder_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QCameraViewfinder_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QCameraViewfinder_ChangeEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::changeEvent(static_cast<QEvent*>(event));
}

int QCameraViewfinder_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

void QCameraViewfinder_CloseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QCameraViewfinder_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QCameraViewfinder_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QCameraViewfinder*>(ptr)->focusNextPrevChild(next != 0);
}

int QCameraViewfinder_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusNextPrevChild(next != 0);
}

int QCameraViewfinder_HasHeightForWidth(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->hasHeightForWidth();
}

int QCameraViewfinder_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::hasHeightForWidth();
}

int QCameraViewfinder_HeightForWidth(void* ptr, int w)
{
	return static_cast<QCameraViewfinder*>(ptr)->heightForWidth(w);
}

int QCameraViewfinder_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::heightForWidth(w);
}

void QCameraViewfinder_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "hide");
}

void QCameraViewfinder_InitPainter(void* ptr, void* painter)
{
	static_cast<QCameraViewfinder*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QCameraViewfinder_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::initPainter(static_cast<QPainter*>(painter));
}

void QCameraViewfinder_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QCameraViewfinder_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QCameraViewfinder_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QCameraViewfinder*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QCameraViewfinder_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QCameraViewfinder_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "lower");
}

void QCameraViewfinder_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QCameraViewfinder_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QCameraViewfinder*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QCameraViewfinder_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::nativeEvent(QByteArray(eventType), message, &result);
}

void QCameraViewfinder_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "raise");
}

void QCameraViewfinder_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "repaint");
}

void QCameraViewfinder_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QCameraViewfinder_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setFocus");
}

void QCameraViewfinder_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QCameraViewfinder_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "show");
}

void QCameraViewfinder_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showFullScreen");
}

void QCameraViewfinder_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showMaximized");
}

void QCameraViewfinder_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showMinimized");
}

void QCameraViewfinder_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showNormal");
}

void QCameraViewfinder_TabletEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "update");
}

void QCameraViewfinder_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "updateMicroFocus");
}

void QCameraViewfinder_WheelEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "deleteLater");
}

void QCameraViewfinder_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QCameraViewfinder_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QCameraViewfinder_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraViewfinder_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinder*>(ptr)->metaObject());
}

void* QCameraViewfinder_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::metaObject());
}

class MyQGraphicsVideoItem: public QGraphicsVideoItem
{
public:
	MyQGraphicsVideoItem(QGraphicsItem *parent) : QGraphicsVideoItem(parent) {};
	void Signal_NativeSizeChanged(const QSizeF & size) { callbackQGraphicsVideoItem_NativeSizeChanged(this, this->objectName().toUtf8().data(), new QSizeF(static_cast<QSizeF>(size).width(), static_cast<QSizeF>(size).height())); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsVideoItem_BoundingRect(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data())); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQGraphicsVideoItem_MediaObject(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data())); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsVideoItem_Paint(this, this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	bool setMediaObject(QMediaObject * object) { return callbackQGraphicsVideoItem_SetMediaObject(this, this->objectName().toUtf8().data(), object) != 0; };
	bool event(QEvent * ev) { return callbackQGraphicsVideoItem_Event(this, this->objectName().toUtf8().data(), ev) != 0; };
	void updateMicroFocus() { callbackQGraphicsVideoItem_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void advance(int phase) { callbackQGraphicsVideoItem_Advance(this, this->objectName().toUtf8().data(), phase); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsVideoItem_CollidesWithItem(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data(), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsVideoItem_CollidesWithPath(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data(), new QPainterPath(path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsVideoItem_Contains(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y())) != 0; };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsVideoItem_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQGraphicsVideoItem_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQGraphicsVideoItem_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQGraphicsVideoItem_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsVideoItem_InputMethodQuery(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data(), query)); };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsVideoItem_IsObscuredBy(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data(), const_cast<QGraphicsItem*>(item)) != 0; };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsVideoItem_ItemChange(this, this->objectName().toUtf8().data(), change, new QVariant(value))); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsVideoItem_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsVideoItem_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsVideoItem_OpaqueArea(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data())); };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsVideoItem_SceneEvent(this, this->objectName().toUtf8().data(), event) != 0; };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsVideoItem_SceneEventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsVideoItem_Shape(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data())); };
	int type() const { return callbackQGraphicsVideoItem_Type(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data()); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQGraphicsVideoItem_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsVideoItem_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsVideoItem_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsVideoItem_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsVideoItem_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGraphicsVideoItem_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsVideoItem_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsVideoItem_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsVideoItem_MetaObject(const_cast<MyQGraphicsVideoItem*>(this), this->objectName().toUtf8().data())); };
};

void QGraphicsVideoItem_ConnectNativeSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsVideoItem*>(ptr), static_cast<void (QGraphicsVideoItem::*)(const QSizeF &)>(&QGraphicsVideoItem::nativeSizeChanged), static_cast<MyQGraphicsVideoItem*>(ptr), static_cast<void (MyQGraphicsVideoItem::*)(const QSizeF &)>(&MyQGraphicsVideoItem::Signal_NativeSizeChanged));
}

void QGraphicsVideoItem_DisconnectNativeSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsVideoItem*>(ptr), static_cast<void (QGraphicsVideoItem::*)(const QSizeF &)>(&QGraphicsVideoItem::nativeSizeChanged), static_cast<MyQGraphicsVideoItem*>(ptr), static_cast<void (MyQGraphicsVideoItem::*)(const QSizeF &)>(&MyQGraphicsVideoItem::Signal_NativeSizeChanged));
}

void QGraphicsVideoItem_NativeSizeChanged(void* ptr, void* size)
{
	static_cast<QGraphicsVideoItem*>(ptr)->nativeSizeChanged(*static_cast<QSizeF*>(size));
}

void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent)
{
	return new MyQGraphicsVideoItem(static_cast<QGraphicsItem*>(parent));
}

int QGraphicsVideoItem_AspectRatioMode(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->aspectRatioMode();
}

void* QGraphicsVideoItem_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->boundingRect()).height());
}

void* QGraphicsVideoItem_BoundingRectDefault(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::boundingRect()).x(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::boundingRect()).y(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::boundingRect()).width(), static_cast<QRectF>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::boundingRect()).height());
}

void* QGraphicsVideoItem_MediaObject(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->mediaObject();
}

void* QGraphicsVideoItem_MediaObjectDefault(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mediaObject();
}

void* QGraphicsVideoItem_NativeSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QGraphicsVideoItem*>(ptr)->nativeSize()).width(), static_cast<QSizeF>(static_cast<QGraphicsVideoItem*>(ptr)->nativeSize()).height());
}

void* QGraphicsVideoItem_Offset(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QGraphicsVideoItem*>(ptr)->offset()).x(), static_cast<QPointF>(static_cast<QGraphicsVideoItem*>(ptr)->offset()).y());
}

void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsVideoItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, int mode)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QGraphicsVideoItem_SetOffset(void* ptr, void* offset)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsVideoItem_SetSize(void* ptr, void* size)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void* QGraphicsVideoItem_Size(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QGraphicsVideoItem*>(ptr)->size()).width(), static_cast<QSizeF>(static_cast<QGraphicsVideoItem*>(ptr)->size()).height());
}

void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr)
{
	static_cast<QGraphicsVideoItem*>(ptr)->~QGraphicsVideoItem();
}

int QGraphicsVideoItem_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

int QGraphicsVideoItem_Event(void* ptr, void* ev)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QGraphicsVideoItem_EventDefault(void* ptr, void* ev)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::event(static_cast<QEvent*>(ev));
}

void QGraphicsVideoItem_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsVideoItem*>(ptr), "updateMicroFocus");
}

void QGraphicsVideoItem_Advance(void* ptr, int phase)
{
	static_cast<QGraphicsVideoItem*>(ptr)->advance(phase);
}

void QGraphicsVideoItem_AdvanceDefault(void* ptr, int phase)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::advance(phase);
}

int QGraphicsVideoItem_CollidesWithItem(void* ptr, void* other, int mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsVideoItem_CollidesWithItemDefault(void* ptr, void* other, int mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsVideoItem_CollidesWithPath(void* ptr, void* path, int mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsVideoItem_CollidesWithPathDefault(void* ptr, void* path, int mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsVideoItem_Contains(void* ptr, void* point)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsVideoItem_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::contains(*static_cast<QPointF*>(point));
}

void QGraphicsVideoItem_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsVideoItem_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsVideoItem_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DropEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DropEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_FocusInEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsVideoItem_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QGraphicsVideoItem_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsVideoItem_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int QGraphicsVideoItem_IsObscuredBy(void* ptr, void* item)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

int QGraphicsVideoItem_IsObscuredByDefault(void* ptr, void* item)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsVideoItem_ItemChange(void* ptr, int change, void* value)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void* QGraphicsVideoItem_ItemChangeDefault(void* ptr, int change, void* value)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void QGraphicsVideoItem_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MousePressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void* QGraphicsVideoItem_OpaqueArea(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->opaqueArea());
}

void* QGraphicsVideoItem_OpaqueAreaDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::opaqueArea());
}

int QGraphicsVideoItem_SceneEvent(void* ptr, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->sceneEvent(static_cast<QEvent*>(event));
}

int QGraphicsVideoItem_SceneEventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::sceneEvent(static_cast<QEvent*>(event));
}

int QGraphicsVideoItem_SceneEventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

int QGraphicsVideoItem_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsVideoItem_Shape(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->shape());
}

void* QGraphicsVideoItem_ShapeDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::shape());
}

int QGraphicsVideoItem_Type(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->type();
}

int QGraphicsVideoItem_TypeDefault(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::type();
}

void QGraphicsVideoItem_WheelEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsVideoItem_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsVideoItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsVideoItem_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsVideoItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsVideoItem*>(ptr), "deleteLater");
}

void QGraphicsVideoItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGraphicsVideoItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGraphicsVideoItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsVideoItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsVideoItem*>(ptr)->metaObject());
}

void* QGraphicsVideoItem_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::metaObject());
}

class MyQVideoWidget: public QVideoWidget
{
public:
	MyQVideoWidget(QWidget *parent) : QVideoWidget(parent) {};
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQVideoWidget_MediaObject(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data())); };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQVideoWidget_SetAspectRatioMode(this, this->objectName().toUtf8().data(), mode); };
	void setBrightness(int brightness) { callbackQVideoWidget_SetBrightness(this, this->objectName().toUtf8().data(), brightness); };
	void setContrast(int contrast) { callbackQVideoWidget_SetContrast(this, this->objectName().toUtf8().data(), contrast); };
	void setFullScreen(bool fullScreen) { callbackQVideoWidget_SetFullScreen(this, this->objectName().toUtf8().data(), fullScreen); };
	void setHue(int hue) { callbackQVideoWidget_SetHue(this, this->objectName().toUtf8().data(), hue); };
	void setSaturation(int saturation) { callbackQVideoWidget_SetSaturation(this, this->objectName().toUtf8().data(), saturation); };
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidget_BrightnessChanged(this, this->objectName().toUtf8().data(), brightness); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidget_ContrastChanged(this, this->objectName().toUtf8().data(), contrast); };
	bool event(QEvent * event) { return callbackQVideoWidget_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidget_FullScreenChanged(this, this->objectName().toUtf8().data(), fullScreen); };
	void hideEvent(QHideEvent * event) { callbackQVideoWidget_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_HueChanged(int hue) { callbackQVideoWidget_HueChanged(this, this->objectName().toUtf8().data(), hue); };
	void moveEvent(QMoveEvent * event) { callbackQVideoWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQVideoWidget_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQVideoWidget_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidget_SaturationChanged(this, this->objectName().toUtf8().data(), saturation); };
	void showEvent(QShowEvent * event) { callbackQVideoWidget_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQVideoWidget_SizeHint(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data())); };
	bool setMediaObject(QMediaObject * object) { return callbackQVideoWidget_SetMediaObject(this, this->objectName().toUtf8().data(), object) != 0; };
	void actionEvent(QActionEvent * event) { callbackQVideoWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQVideoWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQVideoWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQVideoWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQVideoWidget_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQVideoWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQVideoWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQVideoWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQVideoWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQVideoWidget_Metric(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQVideoWidget_MinimumSizeHint(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data())); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQVideoWidget_PaintEngine(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data())); };
	void setEnabled(bool vbo) { callbackQVideoWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQVideoWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQVideoWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQVideoWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQVideoWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void changeEvent(QEvent * event) { callbackQVideoWidget_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQVideoWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQVideoWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQVideoWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQVideoWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQVideoWidget_HasHeightForWidth(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQVideoWidget_HeightForWidth(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQVideoWidget_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQVideoWidget_InitPainter(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQVideoWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQVideoWidget_InputMethodQuery(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQVideoWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQVideoWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQVideoWidget_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQVideoWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQVideoWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQVideoWidget_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQVideoWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQVideoWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQVideoWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQVideoWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQVideoWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQVideoWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQVideoWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQVideoWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQVideoWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQVideoWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQVideoWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQVideoWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQVideoWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQVideoWidget_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQVideoWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQVideoWidget_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidget_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoWidget_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQVideoWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQVideoWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoWidget_MetaObject(const_cast<MyQVideoWidget*>(this), this->objectName().toUtf8().data())); };
};

int QVideoWidget_AspectRatioMode(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->aspectRatioMode();
}

int QVideoWidget_Brightness(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->brightness();
}

int QVideoWidget_Contrast(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->contrast();
}

int QVideoWidget_Hue(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->hue();
}

void* QVideoWidget_MediaObject(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->mediaObject();
}

void* QVideoWidget_MediaObjectDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::mediaObject();
}

int QVideoWidget_Saturation(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->saturation();
}

void QVideoWidget_SetAspectRatioMode(void* ptr, int mode)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QVideoWidget_SetBrightness(void* ptr, int brightness)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QVideoWidget_SetContrast(void* ptr, int contrast)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QVideoWidget_SetFullScreen(void* ptr, int fullScreen)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QVideoWidget_SetHue(void* ptr, int hue)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHue", Q_ARG(int, hue));
}

void QVideoWidget_SetSaturation(void* ptr, int saturation)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

void* QVideoWidget_NewQVideoWidget(void* parent)
{
	return new MyQVideoWidget(static_cast<QWidget*>(parent));
}

void QVideoWidget_ConnectBrightnessChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));
}

void QVideoWidget_DisconnectBrightnessChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));
}

void QVideoWidget_BrightnessChanged(void* ptr, int brightness)
{
	static_cast<QVideoWidget*>(ptr)->brightnessChanged(brightness);
}

void QVideoWidget_ConnectContrastChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));
}

void QVideoWidget_DisconnectContrastChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));
}

void QVideoWidget_ContrastChanged(void* ptr, int contrast)
{
	static_cast<QVideoWidget*>(ptr)->contrastChanged(contrast);
}

int QVideoWidget_Event(void* ptr, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->event(static_cast<QEvent*>(event));
}

int QVideoWidget_EventDefault(void* ptr, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::event(static_cast<QEvent*>(event));
}

void QVideoWidget_ConnectFullScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));
}

void QVideoWidget_DisconnectFullScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));
}

void QVideoWidget_FullScreenChanged(void* ptr, int fullScreen)
{
	static_cast<QVideoWidget*>(ptr)->fullScreenChanged(fullScreen != 0);
}

void QVideoWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QVideoWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QVideoWidget_ConnectHueChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));
}

void QVideoWidget_DisconnectHueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));
}

void QVideoWidget_HueChanged(void* ptr, int hue)
{
	static_cast<QVideoWidget*>(ptr)->hueChanged(hue);
}

int QVideoWidget_IsFullScreen(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->isFullScreen();
}

void QVideoWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QVideoWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QVideoWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QVideoWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QVideoWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QVideoWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QVideoWidget_ConnectSaturationChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));
}

void QVideoWidget_DisconnectSaturationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));
}

void QVideoWidget_SaturationChanged(void* ptr, int saturation)
{
	static_cast<QVideoWidget*>(ptr)->saturationChanged(saturation);
}

void QVideoWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QVideoWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QVideoWidget_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->sizeHint()).height());
}

void* QVideoWidget_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::sizeHint()).width(), static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::sizeHint()).height());
}

void QVideoWidget_DestroyQVideoWidget(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->~QVideoWidget();
}

int QVideoWidget_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QVideoWidget*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

void QVideoWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::leaveEvent(static_cast<QEvent*>(event));
}

int QVideoWidget_Metric(void* ptr, int m)
{
	return static_cast<QVideoWidget*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QVideoWidget_MetricDefault(void* ptr, int m)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QVideoWidget_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->minimumSizeHint()).height());
}

void* QVideoWidget_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::minimumSizeHint()).height());
}

void* QVideoWidget_PaintEngine(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->paintEngine();
}

void* QVideoWidget_PaintEngineDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::paintEngine();
}

void QVideoWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QVideoWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QVideoWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QVideoWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setVisible(visible != 0);
}

void QVideoWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QVideoWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QVideoWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::changeEvent(static_cast<QEvent*>(event));
}

int QVideoWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

void QVideoWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QVideoWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QVideoWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QVideoWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QVideoWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusNextPrevChild(next != 0);
}

int QVideoWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->hasHeightForWidth();
}

int QVideoWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::hasHeightForWidth();
}

int QVideoWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QVideoWidget*>(ptr)->heightForWidth(w);
}

int QVideoWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::heightForWidth(w);
}

void QVideoWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "hide");
}

void QVideoWidget_InitPainter(void* ptr, void* painter)
{
	static_cast<QVideoWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QVideoWidget_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::initPainter(static_cast<QPainter*>(painter));
}

void QVideoWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QVideoWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QVideoWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QVideoWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QVideoWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QVideoWidget*>(ptr)->QVideoWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QVideoWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "lower");
}

void QVideoWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QVideoWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QVideoWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QVideoWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QVideoWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "raise");
}

void QVideoWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "repaint");
}

void QVideoWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QVideoWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFocus");
}

void QVideoWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QVideoWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "show");
}

void QVideoWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showFullScreen");
}

void QVideoWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showMaximized");
}

void QVideoWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showMinimized");
}

void QVideoWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showNormal");
}

void QVideoWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "update");
}

void QVideoWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "updateMicroFocus");
}

void QVideoWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::customEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "deleteLater");
}

void QVideoWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QVideoWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QVideoWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidget*>(ptr)->metaObject());
}

void* QVideoWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::metaObject());
}

class MyQVideoWidgetControl: public QVideoWidgetControl
{
public:
	MyQVideoWidgetControl(QObject *parent) : QVideoWidgetControl(parent) {};
	Qt::AspectRatioMode aspectRatioMode() const { return static_cast<Qt::AspectRatioMode>(callbackQVideoWidgetControl_AspectRatioMode(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data())); };
	int brightness() const { return callbackQVideoWidgetControl_Brightness(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data()); };
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidgetControl_BrightnessChanged(this, this->objectName().toUtf8().data(), brightness); };
	int contrast() const { return callbackQVideoWidgetControl_Contrast(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data()); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidgetControl_ContrastChanged(this, this->objectName().toUtf8().data(), contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidgetControl_FullScreenChanged(this, this->objectName().toUtf8().data(), fullScreen); };
	int hue() const { return callbackQVideoWidgetControl_Hue(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data()); };
	void Signal_HueChanged(int hue) { callbackQVideoWidgetControl_HueChanged(this, this->objectName().toUtf8().data(), hue); };
	bool isFullScreen() const { return callbackQVideoWidgetControl_IsFullScreen(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data()) != 0; };
	int saturation() const { return callbackQVideoWidgetControl_Saturation(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data()); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidgetControl_SaturationChanged(this, this->objectName().toUtf8().data(), saturation); };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQVideoWidgetControl_SetAspectRatioMode(this, this->objectName().toUtf8().data(), mode); };
	void setBrightness(int brightness) { callbackQVideoWidgetControl_SetBrightness(this, this->objectName().toUtf8().data(), brightness); };
	void setContrast(int contrast) { callbackQVideoWidgetControl_SetContrast(this, this->objectName().toUtf8().data(), contrast); };
	void setFullScreen(bool fullScreen) { callbackQVideoWidgetControl_SetFullScreen(this, this->objectName().toUtf8().data(), fullScreen); };
	void setHue(int hue) { callbackQVideoWidgetControl_SetHue(this, this->objectName().toUtf8().data(), hue); };
	void setSaturation(int saturation) { callbackQVideoWidgetControl_SetSaturation(this, this->objectName().toUtf8().data(), saturation); };
	QWidget * videoWidget() { return static_cast<QWidget*>(callbackQVideoWidgetControl_VideoWidget(this, this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidgetControl_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidgetControl_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoWidgetControl_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQVideoWidgetControl_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQVideoWidgetControl_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoWidgetControl_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQVideoWidgetControl_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoWidgetControl_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoWidgetControl_MetaObject(const_cast<MyQVideoWidgetControl*>(this), this->objectName().toUtf8().data())); };
};

void* QVideoWidgetControl_NewQVideoWidgetControl(void* parent)
{
	return new MyQVideoWidgetControl(static_cast<QObject*>(parent));
}

int QVideoWidgetControl_AspectRatioMode(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->aspectRatioMode();
}

int QVideoWidgetControl_Brightness(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->brightness();
}

void QVideoWidgetControl_ConnectBrightnessChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));
}

void QVideoWidgetControl_DisconnectBrightnessChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));
}

void QVideoWidgetControl_BrightnessChanged(void* ptr, int brightness)
{
	static_cast<QVideoWidgetControl*>(ptr)->brightnessChanged(brightness);
}

int QVideoWidgetControl_Contrast(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->contrast();
}

void QVideoWidgetControl_ConnectContrastChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));
}

void QVideoWidgetControl_DisconnectContrastChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));
}

void QVideoWidgetControl_ContrastChanged(void* ptr, int contrast)
{
	static_cast<QVideoWidgetControl*>(ptr)->contrastChanged(contrast);
}

void QVideoWidgetControl_ConnectFullScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));
}

void QVideoWidgetControl_DisconnectFullScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));
}

void QVideoWidgetControl_FullScreenChanged(void* ptr, int fullScreen)
{
	static_cast<QVideoWidgetControl*>(ptr)->fullScreenChanged(fullScreen != 0);
}

int QVideoWidgetControl_Hue(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->hue();
}

void QVideoWidgetControl_ConnectHueChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));
}

void QVideoWidgetControl_DisconnectHueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));
}

void QVideoWidgetControl_HueChanged(void* ptr, int hue)
{
	static_cast<QVideoWidgetControl*>(ptr)->hueChanged(hue);
}

int QVideoWidgetControl_IsFullScreen(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->isFullScreen();
}

int QVideoWidgetControl_Saturation(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->saturation();
}

void QVideoWidgetControl_ConnectSaturationChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));
}

void QVideoWidgetControl_DisconnectSaturationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));
}

void QVideoWidgetControl_SaturationChanged(void* ptr, int saturation)
{
	static_cast<QVideoWidgetControl*>(ptr)->saturationChanged(saturation);
}

void QVideoWidgetControl_SetAspectRatioMode(void* ptr, int mode)
{
	static_cast<QVideoWidgetControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWidgetControl_SetBrightness(void* ptr, int brightness)
{
	static_cast<QVideoWidgetControl*>(ptr)->setBrightness(brightness);
}

void QVideoWidgetControl_SetContrast(void* ptr, int contrast)
{
	static_cast<QVideoWidgetControl*>(ptr)->setContrast(contrast);
}

void QVideoWidgetControl_SetFullScreen(void* ptr, int fullScreen)
{
	static_cast<QVideoWidgetControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWidgetControl_SetHue(void* ptr, int hue)
{
	static_cast<QVideoWidgetControl*>(ptr)->setHue(hue);
}

void QVideoWidgetControl_SetSaturation(void* ptr, int saturation)
{
	static_cast<QVideoWidgetControl*>(ptr)->setSaturation(saturation);
}

void* QVideoWidgetControl_VideoWidget(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->videoWidget();
}

void QVideoWidgetControl_DestroyQVideoWidgetControl(void* ptr)
{
	static_cast<QVideoWidgetControl*>(ptr)->~QVideoWidgetControl();
}

void QVideoWidgetControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidgetControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoWidgetControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidgetControl*>(ptr), "deleteLater");
}

void QVideoWidgetControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QVideoWidgetControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoWidgetControl*>(ptr)->event(static_cast<QEvent*>(e));
}

int QVideoWidgetControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::event(static_cast<QEvent*>(e));
}

int QVideoWidgetControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidgetControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QVideoWidgetControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoWidgetControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidgetControl*>(ptr)->metaObject());
}

void* QVideoWidgetControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::metaObject());
}

