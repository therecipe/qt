// +build !minimal

#define protected public
#define private public

#include "printsupport.h"
#include "_cgo_export.h"

#include <QAbstractPrintDialog>
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
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QList>
#include <QMargins>
#include <QMarginsF>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPageLayout>
#include <QPageSetupDialog>
#include <QPageSize>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPrintDialog>
#include <QPrintEngine>
#include <QPrintPreviewDialog>
#include <QPrintPreviewWidget>
#include <QPrinter>
#include <QPrinterInfo>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

class MyQAbstractPrintDialog: public QAbstractPrintDialog
{
public:
	MyQAbstractPrintDialog(QPrinter *printer, QWidget *parent) : QAbstractPrintDialog(printer, parent) {};
	int exec() { return callbackQAbstractPrintDialog_Exec(this); };
	void accept() { callbackQAbstractPrintDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQAbstractPrintDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQAbstractPrintDialog_ContextMenuEvent(this, e); };
	void done(int r) { callbackQAbstractPrintDialog_Done(this, r); };
	void keyPressEvent(QKeyEvent * e) { callbackQAbstractPrintDialog_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_MinimumSizeHint(const_cast<MyQAbstractPrintDialog*>(this))); };
	void open() { callbackQAbstractPrintDialog_Open(this); };
	void reject() { callbackQAbstractPrintDialog_Reject(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQAbstractPrintDialog_ResizeEvent(this, vqr); };
	void setVisible(bool visible) { callbackQAbstractPrintDialog_SetVisible(this, visible); };
	void showEvent(QShowEvent * event) { callbackQAbstractPrintDialog_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_SizeHint(const_cast<MyQAbstractPrintDialog*>(this))); };
	void actionEvent(QActionEvent * event) { callbackQAbstractPrintDialog_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQAbstractPrintDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQAbstractPrintDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQAbstractPrintDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQAbstractPrintDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQAbstractPrintDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQAbstractPrintDialog_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQAbstractPrintDialog_LeaveEvent(this, event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQAbstractPrintDialog_Metric(const_cast<MyQAbstractPrintDialog*>(this), m); };
	void moveEvent(QMoveEvent * event) { callbackQAbstractPrintDialog_MoveEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQAbstractPrintDialog_PaintEngine(const_cast<MyQAbstractPrintDialog*>(this))); };
	void paintEvent(QPaintEvent * event) { callbackQAbstractPrintDialog_PaintEvent(this, event); };
	void setEnabled(bool vbo) { callbackQAbstractPrintDialog_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQAbstractPrintDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQAbstractPrintDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQAbstractPrintDialog_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQAbstractPrintDialog_ChangeEvent(this, event); };
	bool close() { return callbackQAbstractPrintDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQAbstractPrintDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQAbstractPrintDialog_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQAbstractPrintDialog_HasHeightForWidth(const_cast<MyQAbstractPrintDialog*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQAbstractPrintDialog_HeightForWidth(const_cast<MyQAbstractPrintDialog*>(this), w); };
	void hide() { callbackQAbstractPrintDialog_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQAbstractPrintDialog_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQAbstractPrintDialog_InputMethodQuery(const_cast<MyQAbstractPrintDialog*>(this), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQAbstractPrintDialog_KeyReleaseEvent(this, event); };
	void lower() { callbackQAbstractPrintDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQAbstractPrintDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQAbstractPrintDialog_Raise(this); };
	void repaint() { callbackQAbstractPrintDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQAbstractPrintDialog_SetDisabled(this, disable); };
	void setFocus() { callbackQAbstractPrintDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQAbstractPrintDialog_SetHidden(this, hidden); };
	void show() { callbackQAbstractPrintDialog_Show(this); };
	void showFullScreen() { callbackQAbstractPrintDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQAbstractPrintDialog_ShowMaximized(this); };
	void showMinimized() { callbackQAbstractPrintDialog_ShowMinimized(this); };
	void showNormal() { callbackQAbstractPrintDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQAbstractPrintDialog_TabletEvent(this, event); };
	void update() { callbackQAbstractPrintDialog_Update(this); };
	void updateMicroFocus() { callbackQAbstractPrintDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQAbstractPrintDialog_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractPrintDialog_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAbstractPrintDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractPrintDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractPrintDialog_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractPrintDialog_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractPrintDialog_MetaObject(const_cast<MyQAbstractPrintDialog*>(this))); };
};

void* QAbstractPrintDialog_NewQAbstractPrintDialog(void* printer, void* parent)
{
	return new MyQAbstractPrintDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent));
}

int QAbstractPrintDialog_Exec(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->exec();
}

int QAbstractPrintDialog_FromPage(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->fromPage();
}

int QAbstractPrintDialog_MaxPage(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->maxPage();
}

int QAbstractPrintDialog_MinPage(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->minPage();
}

long long QAbstractPrintDialog_PrintRange(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->printRange();
}

void* QAbstractPrintDialog_Printer(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->printer();
}

void QAbstractPrintDialog_SetFromTo(void* ptr, int from, int to)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setFromTo(from, to);
}

void QAbstractPrintDialog_SetMinMax(void* ptr, int min, int max)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setMinMax(min, max);
}

void QAbstractPrintDialog_SetPrintRange(void* ptr, long long ran)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setPrintRange(static_cast<QAbstractPrintDialog::PrintRange>(ran));
}

int QAbstractPrintDialog_ToPage(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->toPage();
}

void QAbstractPrintDialog_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "accept");
}

void QAbstractPrintDialog_AcceptDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::accept();
}

void QAbstractPrintDialog_CloseEvent(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->closeEvent(static_cast<QCloseEvent*>(e));
}

void QAbstractPrintDialog_CloseEventDefault(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QAbstractPrintDialog_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QAbstractPrintDialog_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QAbstractPrintDialog_Done(void* ptr, int r)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "done", Q_ARG(int, r));
}

void QAbstractPrintDialog_DoneDefault(void* ptr, int r)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::done(r);
}

void QAbstractPrintDialog_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QAbstractPrintDialog_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* QAbstractPrintDialog_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QAbstractPrintDialog_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QAbstractPrintDialog_Open(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "open");
}

void QAbstractPrintDialog_OpenDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::open();
}

void QAbstractPrintDialog_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "reject");
}

void QAbstractPrintDialog_RejectDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::reject();
}

void QAbstractPrintDialog_ResizeEvent(void* ptr, void* vqr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QAbstractPrintDialog_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QAbstractPrintDialog_SetVisible(void* ptr, char visible)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setVisible(visible != 0);
}

void QAbstractPrintDialog_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setVisible(visible != 0);
}

void QAbstractPrintDialog_ShowEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QAbstractPrintDialog_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QAbstractPrintDialog_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QAbstractPrintDialog_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QAbstractPrintDialog_ActionEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QAbstractPrintDialog_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QAbstractPrintDialog_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QAbstractPrintDialog_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QAbstractPrintDialog_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QAbstractPrintDialog_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QAbstractPrintDialog_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QAbstractPrintDialog_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QAbstractPrintDialog_DropEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QAbstractPrintDialog_DropEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QAbstractPrintDialog_EnterEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::enterEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_FocusInEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QAbstractPrintDialog_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QAbstractPrintDialog_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QAbstractPrintDialog_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QAbstractPrintDialog_HideEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QAbstractPrintDialog_HideEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QAbstractPrintDialog_LeaveEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::leaveEvent(static_cast<QEvent*>(event));
}

int QAbstractPrintDialog_Metric(void* ptr, long long m)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QAbstractPrintDialog_MetricDefault(void* ptr, long long m)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QAbstractPrintDialog_MoveEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QAbstractPrintDialog_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QAbstractPrintDialog_PaintEngine(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->paintEngine();
}

void* QAbstractPrintDialog_PaintEngineDefault(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::paintEngine();
}

void QAbstractPrintDialog_PaintEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QAbstractPrintDialog_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QAbstractPrintDialog_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QAbstractPrintDialog_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setEnabled(vbo != 0);
}

void QAbstractPrintDialog_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QAbstractPrintDialog_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setStyleSheet(QString(styleSheet));
}

void QAbstractPrintDialog_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QAbstractPrintDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setWindowModified(vbo != 0);
}

void QAbstractPrintDialog_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QAbstractPrintDialog_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setWindowTitle(QString(vqs));
}

void QAbstractPrintDialog_ChangeEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::changeEvent(static_cast<QEvent*>(event));
}

char QAbstractPrintDialog_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QAbstractPrintDialog_CloseDefault(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::close();
}

char QAbstractPrintDialog_Event(void* ptr, void* event)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->event(static_cast<QEvent*>(event));
}

char QAbstractPrintDialog_EventDefault(void* ptr, void* event)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::event(static_cast<QEvent*>(event));
}

char QAbstractPrintDialog_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->focusNextPrevChild(next != 0);
}

char QAbstractPrintDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusNextPrevChild(next != 0);
}

char QAbstractPrintDialog_HasHeightForWidth(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->hasHeightForWidth();
}

char QAbstractPrintDialog_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hasHeightForWidth();
}

int QAbstractPrintDialog_HeightForWidth(void* ptr, int w)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->heightForWidth(w);
}

int QAbstractPrintDialog_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::heightForWidth(w);
}

void QAbstractPrintDialog_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "hide");
}

void QAbstractPrintDialog_HideDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hide();
}

void QAbstractPrintDialog_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QAbstractPrintDialog_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QAbstractPrintDialog_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QAbstractPrintDialog*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QAbstractPrintDialog_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QAbstractPrintDialog_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QAbstractPrintDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QAbstractPrintDialog_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "lower");
}

void QAbstractPrintDialog_LowerDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::lower();
}

void QAbstractPrintDialog_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MousePressEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QAbstractPrintDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QAbstractPrintDialog_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QAbstractPrintDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QAbstractPrintDialog_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "raise");
}

void QAbstractPrintDialog_RaiseDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::raise();
}

void QAbstractPrintDialog_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "repaint");
}

void QAbstractPrintDialog_RepaintDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::repaint();
}

void QAbstractPrintDialog_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QAbstractPrintDialog_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setDisabled(disable != 0);
}

void QAbstractPrintDialog_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setFocus");
}

void QAbstractPrintDialog_SetFocus2Default(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setFocus();
}

void QAbstractPrintDialog_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QAbstractPrintDialog_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setHidden(hidden != 0);
}

void QAbstractPrintDialog_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "show");
}

void QAbstractPrintDialog_ShowDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::show();
}

void QAbstractPrintDialog_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "showFullScreen");
}

void QAbstractPrintDialog_ShowFullScreenDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showFullScreen();
}

void QAbstractPrintDialog_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "showMaximized");
}

void QAbstractPrintDialog_ShowMaximizedDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showMaximized();
}

void QAbstractPrintDialog_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "showMinimized");
}

void QAbstractPrintDialog_ShowMinimizedDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showMinimized();
}

void QAbstractPrintDialog_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "showNormal");
}

void QAbstractPrintDialog_ShowNormalDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showNormal();
}

void QAbstractPrintDialog_TabletEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QAbstractPrintDialog_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QAbstractPrintDialog_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "update");
}

void QAbstractPrintDialog_UpdateDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::update();
}

void QAbstractPrintDialog_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "updateMicroFocus");
}

void QAbstractPrintDialog_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::updateMicroFocus();
}

void QAbstractPrintDialog_WheelEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QAbstractPrintDialog_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QAbstractPrintDialog_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractPrintDialog_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractPrintDialog_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractPrintDialog_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractPrintDialog_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractPrintDialog*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractPrintDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractPrintDialog_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::customEvent(static_cast<QEvent*>(event));
}

void QAbstractPrintDialog_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractPrintDialog*>(ptr), "deleteLater");
}

void QAbstractPrintDialog_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::deleteLater();
}

void QAbstractPrintDialog_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractPrintDialog*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractPrintDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAbstractPrintDialog_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAbstractPrintDialog_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractPrintDialog_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractPrintDialog*>(ptr)->metaObject());
}

void* QAbstractPrintDialog_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::metaObject());
}

class MyQPageSetupDialog: public QPageSetupDialog
{
public:
	MyQPageSetupDialog(QPrinter *printer, QWidget *parent) : QPageSetupDialog(printer, parent) {};
	MyQPageSetupDialog(QWidget *parent) : QPageSetupDialog(parent) {};
	void done(int result) { callbackQPageSetupDialog_Done(this, result); };
	int exec() { return callbackQPageSetupDialog_Exec(this); };
	void setVisible(bool visible) { callbackQPageSetupDialog_SetVisible(this, visible); };
	void accept() { callbackQPageSetupDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQPageSetupDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQPageSetupDialog_ContextMenuEvent(this, e); };
	void keyPressEvent(QKeyEvent * e) { callbackQPageSetupDialog_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPageSetupDialog_MinimumSizeHint(const_cast<MyQPageSetupDialog*>(this))); };
	void reject() { callbackQPageSetupDialog_Reject(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQPageSetupDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQPageSetupDialog_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPageSetupDialog_SizeHint(const_cast<MyQPageSetupDialog*>(this))); };
	void actionEvent(QActionEvent * event) { callbackQPageSetupDialog_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPageSetupDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPageSetupDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPageSetupDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPageSetupDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPageSetupDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPageSetupDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPageSetupDialog_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQPageSetupDialog_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPageSetupDialog_LeaveEvent(this, event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPageSetupDialog_Metric(const_cast<MyQPageSetupDialog*>(this), m); };
	void moveEvent(QMoveEvent * event) { callbackQPageSetupDialog_MoveEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPageSetupDialog_PaintEngine(const_cast<MyQPageSetupDialog*>(this))); };
	void paintEvent(QPaintEvent * event) { callbackQPageSetupDialog_PaintEvent(this, event); };
	void setEnabled(bool vbo) { callbackQPageSetupDialog_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPageSetupDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPageSetupDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPageSetupDialog_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQPageSetupDialog_ChangeEvent(this, event); };
	bool close() { return callbackQPageSetupDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPageSetupDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPageSetupDialog_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQPageSetupDialog_HasHeightForWidth(const_cast<MyQPageSetupDialog*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQPageSetupDialog_HeightForWidth(const_cast<MyQPageSetupDialog*>(this), w); };
	void hide() { callbackQPageSetupDialog_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPageSetupDialog_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPageSetupDialog_InputMethodQuery(const_cast<MyQPageSetupDialog*>(this), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPageSetupDialog_KeyReleaseEvent(this, event); };
	void lower() { callbackQPageSetupDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPageSetupDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPageSetupDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQPageSetupDialog_Raise(this); };
	void repaint() { callbackQPageSetupDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQPageSetupDialog_SetDisabled(this, disable); };
	void setFocus() { callbackQPageSetupDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPageSetupDialog_SetHidden(this, hidden); };
	void show() { callbackQPageSetupDialog_Show(this); };
	void showFullScreen() { callbackQPageSetupDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQPageSetupDialog_ShowMaximized(this); };
	void showMinimized() { callbackQPageSetupDialog_ShowMinimized(this); };
	void showNormal() { callbackQPageSetupDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPageSetupDialog_TabletEvent(this, event); };
	void update() { callbackQPageSetupDialog_Update(this); };
	void updateMicroFocus() { callbackQPageSetupDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPageSetupDialog_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQPageSetupDialog_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQPageSetupDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPageSetupDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPageSetupDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQPageSetupDialog_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPageSetupDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPageSetupDialog_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPageSetupDialog_MetaObject(const_cast<MyQPageSetupDialog*>(this))); };
};

void* QPageSetupDialog_Printer(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->printer();
}

void QPageSetupDialog_DestroyQPageSetupDialog(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->~QPageSetupDialog();
}

void QPageSetupDialog_Done(void* ptr, int result)
{
	static_cast<QPageSetupDialog*>(ptr)->done(result);
}

void QPageSetupDialog_DoneDefault(void* ptr, int result)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::done(result);
}

void QPageSetupDialog_Open(void* ptr, void* receiver, char* member)
{
	static_cast<QPageSetupDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QPageSetupDialog_NewQPageSetupDialog(void* printer, void* parent)
{
	return new MyQPageSetupDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent));
}

void* QPageSetupDialog_NewQPageSetupDialog2(void* parent)
{
	return new MyQPageSetupDialog(static_cast<QWidget*>(parent));
}

int QPageSetupDialog_Exec(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->exec();
}

int QPageSetupDialog_ExecDefault(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::exec();
}

void QPageSetupDialog_SetVisible(void* ptr, char visible)
{
	static_cast<QPageSetupDialog*>(ptr)->setVisible(visible != 0);
}

void QPageSetupDialog_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setVisible(visible != 0);
}

void QPageSetupDialog_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "accept");
}

void QPageSetupDialog_AcceptDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::accept();
}

void QPageSetupDialog_CloseEvent(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->closeEvent(static_cast<QCloseEvent*>(e));
}

void QPageSetupDialog_CloseEventDefault(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QPageSetupDialog_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPageSetupDialog_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPageSetupDialog_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QPageSetupDialog_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* QPageSetupDialog_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPageSetupDialog_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPageSetupDialog_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "reject");
}

void QPageSetupDialog_RejectDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::reject();
}

void QPageSetupDialog_ResizeEvent(void* ptr, void* vqr)
{
	static_cast<QPageSetupDialog*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPageSetupDialog_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPageSetupDialog_ShowEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QPageSetupDialog_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QPageSetupDialog_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPageSetupDialog_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPageSetupDialog_ActionEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QPageSetupDialog_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QPageSetupDialog_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPageSetupDialog_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPageSetupDialog_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPageSetupDialog_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPageSetupDialog_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPageSetupDialog_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPageSetupDialog_DropEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QPageSetupDialog_DropEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QPageSetupDialog_EnterEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::enterEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_FocusInEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_HideEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QPageSetupDialog_HideEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QPageSetupDialog_LeaveEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::leaveEvent(static_cast<QEvent*>(event));
}

int QPageSetupDialog_Metric(void* ptr, long long m)
{
	return static_cast<QPageSetupDialog*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QPageSetupDialog_MetricDefault(void* ptr, long long m)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QPageSetupDialog_MoveEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QPageSetupDialog_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QPageSetupDialog_PaintEngine(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->paintEngine();
}

void* QPageSetupDialog_PaintEngineDefault(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::paintEngine();
}

void QPageSetupDialog_PaintEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QPageSetupDialog_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPageSetupDialog_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QPageSetupDialog_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setEnabled(vbo != 0);
}

void QPageSetupDialog_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QPageSetupDialog_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setStyleSheet(QString(styleSheet));
}

void QPageSetupDialog_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QPageSetupDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setWindowModified(vbo != 0);
}

void QPageSetupDialog_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QPageSetupDialog_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setWindowTitle(QString(vqs));
}

void QPageSetupDialog_ChangeEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::changeEvent(static_cast<QEvent*>(event));
}

char QPageSetupDialog_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QPageSetupDialog_CloseDefault(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::close();
}

char QPageSetupDialog_Event(void* ptr, void* event)
{
	return static_cast<QPageSetupDialog*>(ptr)->event(static_cast<QEvent*>(event));
}

char QPageSetupDialog_EventDefault(void* ptr, void* event)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::event(static_cast<QEvent*>(event));
}

char QPageSetupDialog_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QPageSetupDialog*>(ptr)->focusNextPrevChild(next != 0);
}

char QPageSetupDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusNextPrevChild(next != 0);
}

char QPageSetupDialog_HasHeightForWidth(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->hasHeightForWidth();
}

char QPageSetupDialog_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hasHeightForWidth();
}

int QPageSetupDialog_HeightForWidth(void* ptr, int w)
{
	return static_cast<QPageSetupDialog*>(ptr)->heightForWidth(w);
}

int QPageSetupDialog_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::heightForWidth(w);
}

void QPageSetupDialog_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "hide");
}

void QPageSetupDialog_HideDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hide();
}

void QPageSetupDialog_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPageSetupDialog_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QPageSetupDialog_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QPageSetupDialog*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QPageSetupDialog_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QPageSetupDialog_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPageSetupDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPageSetupDialog_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "lower");
}

void QPageSetupDialog_LowerDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::lower();
}

void QPageSetupDialog_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MousePressEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QPageSetupDialog_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPageSetupDialog*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QPageSetupDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPageSetupDialog_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "raise");
}

void QPageSetupDialog_RaiseDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::raise();
}

void QPageSetupDialog_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "repaint");
}

void QPageSetupDialog_RepaintDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::repaint();
}

void QPageSetupDialog_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QPageSetupDialog_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setDisabled(disable != 0);
}

void QPageSetupDialog_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setFocus");
}

void QPageSetupDialog_SetFocus2Default(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setFocus();
}

void QPageSetupDialog_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QPageSetupDialog_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setHidden(hidden != 0);
}

void QPageSetupDialog_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "show");
}

void QPageSetupDialog_ShowDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::show();
}

void QPageSetupDialog_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "showFullScreen");
}

void QPageSetupDialog_ShowFullScreenDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showFullScreen();
}

void QPageSetupDialog_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "showMaximized");
}

void QPageSetupDialog_ShowMaximizedDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showMaximized();
}

void QPageSetupDialog_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "showMinimized");
}

void QPageSetupDialog_ShowMinimizedDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showMinimized();
}

void QPageSetupDialog_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "showNormal");
}

void QPageSetupDialog_ShowNormalDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showNormal();
}

void QPageSetupDialog_TabletEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPageSetupDialog_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPageSetupDialog_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "update");
}

void QPageSetupDialog_UpdateDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::update();
}

void QPageSetupDialog_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "updateMicroFocus");
}

void QPageSetupDialog_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::updateMicroFocus();
}

void QPageSetupDialog_WheelEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPageSetupDialog_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPageSetupDialog_TimerEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPageSetupDialog_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPageSetupDialog_ChildEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPageSetupDialog_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QPageSetupDialog_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPageSetupDialog*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPageSetupDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPageSetupDialog_CustomEvent(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::customEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPageSetupDialog*>(ptr), "deleteLater");
}

void QPageSetupDialog_DeleteLaterDefault(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::deleteLater();
}

void QPageSetupDialog_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPageSetupDialog*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPageSetupDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QPageSetupDialog_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPageSetupDialog*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QPageSetupDialog_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPageSetupDialog_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPageSetupDialog*>(ptr)->metaObject());
}

void* QPageSetupDialog_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::metaObject());
}

class MyQPrintDialog: public QPrintDialog
{
public:
	MyQPrintDialog(QPrinter *printer, QWidget *parent) : QPrintDialog(printer, parent) {};
	MyQPrintDialog(QWidget *parent) : QPrintDialog(parent) {};
	void Signal_Accepted(QPrinter * printer) { callbackQPrintDialog_Accepted(this, printer); };
	void done(int result) { callbackQPrintDialog_Done(this, result); };
	int exec() { return callbackQPrintDialog_Exec(this); };
	void setVisible(bool visible) { callbackQPrintDialog_SetVisible(this, visible); };
	void accept() { callbackQPrintDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQPrintDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQPrintDialog_ContextMenuEvent(this, e); };
	void keyPressEvent(QKeyEvent * e) { callbackQPrintDialog_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPrintDialog_MinimumSizeHint(const_cast<MyQPrintDialog*>(this))); };
	void reject() { callbackQPrintDialog_Reject(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQPrintDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQPrintDialog_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPrintDialog_SizeHint(const_cast<MyQPrintDialog*>(this))); };
	void actionEvent(QActionEvent * event) { callbackQPrintDialog_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPrintDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPrintDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPrintDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPrintDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPrintDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPrintDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPrintDialog_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQPrintDialog_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPrintDialog_LeaveEvent(this, event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPrintDialog_Metric(const_cast<MyQPrintDialog*>(this), m); };
	void moveEvent(QMoveEvent * event) { callbackQPrintDialog_MoveEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrintDialog_PaintEngine(const_cast<MyQPrintDialog*>(this))); };
	void paintEvent(QPaintEvent * event) { callbackQPrintDialog_PaintEvent(this, event); };
	void setEnabled(bool vbo) { callbackQPrintDialog_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPrintDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPrintDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPrintDialog_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQPrintDialog_ChangeEvent(this, event); };
	bool close() { return callbackQPrintDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPrintDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPrintDialog_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQPrintDialog_HasHeightForWidth(const_cast<MyQPrintDialog*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQPrintDialog_HeightForWidth(const_cast<MyQPrintDialog*>(this), w); };
	void hide() { callbackQPrintDialog_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPrintDialog_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPrintDialog_InputMethodQuery(const_cast<MyQPrintDialog*>(this), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPrintDialog_KeyReleaseEvent(this, event); };
	void lower() { callbackQPrintDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPrintDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPrintDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPrintDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPrintDialog_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPrintDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQPrintDialog_Raise(this); };
	void repaint() { callbackQPrintDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQPrintDialog_SetDisabled(this, disable); };
	void setFocus() { callbackQPrintDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPrintDialog_SetHidden(this, hidden); };
	void show() { callbackQPrintDialog_Show(this); };
	void showFullScreen() { callbackQPrintDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQPrintDialog_ShowMaximized(this); };
	void showMinimized() { callbackQPrintDialog_ShowMinimized(this); };
	void showNormal() { callbackQPrintDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPrintDialog_TabletEvent(this, event); };
	void update() { callbackQPrintDialog_Update(this); };
	void updateMicroFocus() { callbackQPrintDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPrintDialog_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQPrintDialog_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQPrintDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPrintDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPrintDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQPrintDialog_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPrintDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPrintDialog_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPrintDialog_MetaObject(const_cast<MyQPrintDialog*>(this))); };
};

long long QPrintDialog_Options(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->options();
}

void QPrintDialog_SetOptions(void* ptr, long long options)
{
	static_cast<QPrintDialog*>(ptr)->setOptions(static_cast<QAbstractPrintDialog::PrintDialogOption>(options));
}

void QPrintDialog_ConnectAccepted(void* ptr)
{
	QObject::connect(static_cast<QPrintDialog*>(ptr), static_cast<void (QPrintDialog::*)(QPrinter *)>(&QPrintDialog::accepted), static_cast<MyQPrintDialog*>(ptr), static_cast<void (MyQPrintDialog::*)(QPrinter *)>(&MyQPrintDialog::Signal_Accepted));
}

void QPrintDialog_DisconnectAccepted(void* ptr)
{
	QObject::disconnect(static_cast<QPrintDialog*>(ptr), static_cast<void (QPrintDialog::*)(QPrinter *)>(&QPrintDialog::accepted), static_cast<MyQPrintDialog*>(ptr), static_cast<void (MyQPrintDialog::*)(QPrinter *)>(&MyQPrintDialog::Signal_Accepted));
}

void QPrintDialog_Accepted(void* ptr, void* printer)
{
	static_cast<QPrintDialog*>(ptr)->accepted(static_cast<QPrinter*>(printer));
}

void QPrintDialog_Done(void* ptr, int result)
{
	static_cast<QPrintDialog*>(ptr)->done(result);
}

void QPrintDialog_DoneDefault(void* ptr, int result)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::done(result);
}

void QPrintDialog_Open(void* ptr, void* receiver, char* member)
{
	static_cast<QPrintDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QPrintDialog_Printer(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->printer();
}

void QPrintDialog_SetOption(void* ptr, long long option, char on)
{
	static_cast<QPrintDialog*>(ptr)->setOption(static_cast<QAbstractPrintDialog::PrintDialogOption>(option), on != 0);
}

char QPrintDialog_TestOption(void* ptr, long long option)
{
	return static_cast<QPrintDialog*>(ptr)->testOption(static_cast<QAbstractPrintDialog::PrintDialogOption>(option));
}

void* QPrintDialog_NewQPrintDialog(void* printer, void* parent)
{
	return new MyQPrintDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent));
}

void* QPrintDialog_NewQPrintDialog2(void* parent)
{
	return new MyQPrintDialog(static_cast<QWidget*>(parent));
}

int QPrintDialog_Exec(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->exec();
}

int QPrintDialog_ExecDefault(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::exec();
}

void QPrintDialog_SetVisible(void* ptr, char visible)
{
	static_cast<QPrintDialog*>(ptr)->setVisible(visible != 0);
}

void QPrintDialog_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setVisible(visible != 0);
}

void QPrintDialog_DestroyQPrintDialog(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->~QPrintDialog();
}

void QPrintDialog_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "accept");
}

void QPrintDialog_AcceptDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::accept();
}

void QPrintDialog_CloseEvent(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->closeEvent(static_cast<QCloseEvent*>(e));
}

void QPrintDialog_CloseEventDefault(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QPrintDialog_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPrintDialog_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPrintDialog_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QPrintDialog_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* QPrintDialog_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintDialog_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->QPrintDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintDialog_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "reject");
}

void QPrintDialog_RejectDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::reject();
}

void QPrintDialog_ResizeEvent(void* ptr, void* vqr)
{
	static_cast<QPrintDialog*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPrintDialog_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPrintDialog_ShowEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QPrintDialog_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QPrintDialog_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintDialog_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->QPrintDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintDialog_ActionEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintDialog_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintDialog_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintDialog_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintDialog_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintDialog_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintDialog_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintDialog_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintDialog_DropEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintDialog_DropEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintDialog_EnterEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::enterEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_FocusInEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintDialog_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintDialog_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintDialog_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintDialog_HideEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintDialog_HideEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintDialog_LeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::leaveEvent(static_cast<QEvent*>(event));
}

int QPrintDialog_Metric(void* ptr, long long m)
{
	return static_cast<QPrintDialog*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QPrintDialog_MetricDefault(void* ptr, long long m)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QPrintDialog_MoveEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QPrintDialog_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QPrintDialog_PaintEngine(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->paintEngine();
}

void* QPrintDialog_PaintEngineDefault(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::paintEngine();
}

void QPrintDialog_PaintEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintDialog_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintDialog_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QPrintDialog_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setEnabled(vbo != 0);
}

void QPrintDialog_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QPrintDialog_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setStyleSheet(QString(styleSheet));
}

void QPrintDialog_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QPrintDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setWindowModified(vbo != 0);
}

void QPrintDialog_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QPrintDialog_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setWindowTitle(QString(vqs));
}

void QPrintDialog_ChangeEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::changeEvent(static_cast<QEvent*>(event));
}

char QPrintDialog_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QPrintDialog_CloseDefault(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::close();
}

char QPrintDialog_Event(void* ptr, void* event)
{
	return static_cast<QPrintDialog*>(ptr)->event(static_cast<QEvent*>(event));
}

char QPrintDialog_EventDefault(void* ptr, void* event)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::event(static_cast<QEvent*>(event));
}

char QPrintDialog_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QPrintDialog*>(ptr)->focusNextPrevChild(next != 0);
}

char QPrintDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusNextPrevChild(next != 0);
}

char QPrintDialog_HasHeightForWidth(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->hasHeightForWidth();
}

char QPrintDialog_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::hasHeightForWidth();
}

int QPrintDialog_HeightForWidth(void* ptr, int w)
{
	return static_cast<QPrintDialog*>(ptr)->heightForWidth(w);
}

int QPrintDialog_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::heightForWidth(w);
}

void QPrintDialog_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "hide");
}

void QPrintDialog_HideDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::hide();
}

void QPrintDialog_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPrintDialog_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QPrintDialog_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintDialog*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QPrintDialog_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintDialog*>(ptr)->QPrintDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QPrintDialog_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintDialog_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "lower");
}

void QPrintDialog_LowerDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::lower();
}

void QPrintDialog_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MousePressEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPrintDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QPrintDialog_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintDialog*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QPrintDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPrintDialog_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "raise");
}

void QPrintDialog_RaiseDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::raise();
}

void QPrintDialog_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "repaint");
}

void QPrintDialog_RepaintDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::repaint();
}

void QPrintDialog_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QPrintDialog_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setDisabled(disable != 0);
}

void QPrintDialog_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setFocus");
}

void QPrintDialog_SetFocus2Default(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setFocus();
}

void QPrintDialog_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QPrintDialog_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::setHidden(hidden != 0);
}

void QPrintDialog_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "show");
}

void QPrintDialog_ShowDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::show();
}

void QPrintDialog_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "showFullScreen");
}

void QPrintDialog_ShowFullScreenDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::showFullScreen();
}

void QPrintDialog_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "showMaximized");
}

void QPrintDialog_ShowMaximizedDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::showMaximized();
}

void QPrintDialog_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "showMinimized");
}

void QPrintDialog_ShowMinimizedDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::showMinimized();
}

void QPrintDialog_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "showNormal");
}

void QPrintDialog_ShowNormalDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::showNormal();
}

void QPrintDialog_TabletEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintDialog_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintDialog_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "update");
}

void QPrintDialog_UpdateDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::update();
}

void QPrintDialog_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "updateMicroFocus");
}

void QPrintDialog_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::updateMicroFocus();
}

void QPrintDialog_WheelEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintDialog_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintDialog_TimerEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintDialog_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintDialog_ChildEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPrintDialog_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QPrintDialog_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintDialog*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintDialog_CustomEvent(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::customEvent(static_cast<QEvent*>(event));
}

void QPrintDialog_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintDialog*>(ptr), "deleteLater");
}

void QPrintDialog_DeleteLaterDefault(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::deleteLater();
}

void QPrintDialog_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintDialog*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintDialog*>(ptr)->QPrintDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QPrintDialog_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintDialog*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QPrintDialog_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintDialog*>(ptr)->QPrintDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPrintDialog_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintDialog*>(ptr)->metaObject());
}

void* QPrintDialog_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintDialog*>(ptr)->QPrintDialog::metaObject());
}

class MyQPrintEngine: public QPrintEngine
{
public:
	bool abort() { return callbackQPrintEngine_Abort(this) != 0; };
	int metric(QPaintDevice::PaintDeviceMetric id) const { return callbackQPrintEngine_Metric(const_cast<MyQPrintEngine*>(this), id); };
	bool newPage() { return callbackQPrintEngine_NewPage(this) != 0; };
	QPrinter::PrinterState printerState() const { return static_cast<QPrinter::PrinterState>(callbackQPrintEngine_PrinterState(const_cast<MyQPrintEngine*>(this))); };
	QVariant property(QPrintEngine::PrintEnginePropertyKey key) const { return *static_cast<QVariant*>(callbackQPrintEngine_Property(const_cast<MyQPrintEngine*>(this), key)); };
	void setProperty(QPrintEngine::PrintEnginePropertyKey key, const QVariant & value) { callbackQPrintEngine_SetProperty(this, key, const_cast<QVariant*>(&value)); };
	 ~MyQPrintEngine() { callbackQPrintEngine_DestroyQPrintEngine(this); };
};

char QPrintEngine_Abort(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->abort();
}

int QPrintEngine_Metric(void* ptr, long long id)
{
	return static_cast<QPrintEngine*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(id));
}

char QPrintEngine_NewPage(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->newPage();
}

long long QPrintEngine_PrinterState(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->printerState();
}

void* QPrintEngine_Property(void* ptr, long long key)
{
	return new QVariant(static_cast<QPrintEngine*>(ptr)->property(static_cast<QPrintEngine::PrintEnginePropertyKey>(key)));
}

void QPrintEngine_SetProperty(void* ptr, long long key, void* value)
{
	static_cast<QPrintEngine*>(ptr)->setProperty(static_cast<QPrintEngine::PrintEnginePropertyKey>(key), *static_cast<QVariant*>(value));
}

void QPrintEngine_DestroyQPrintEngine(void* ptr)
{
	static_cast<QPrintEngine*>(ptr)->~QPrintEngine();
}

void QPrintEngine_DestroyQPrintEngineDefault(void* ptr)
{

}

class MyQPrintPreviewDialog: public QPrintPreviewDialog
{
public:
	MyQPrintPreviewDialog(QWidget *parent, Qt::WindowFlags flags) : QPrintPreviewDialog(parent, flags) {};
	MyQPrintPreviewDialog(QPrinter *printer, QWidget *parent, Qt::WindowFlags flags) : QPrintPreviewDialog(printer, parent, flags) {};
	void done(int result) { callbackQPrintPreviewDialog_Done(this, result); };
	void Signal_PaintRequested(QPrinter * printer) { callbackQPrintPreviewDialog_PaintRequested(this, printer); };
	void setVisible(bool visible) { callbackQPrintPreviewDialog_SetVisible(this, visible); };
	void accept() { callbackQPrintPreviewDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQPrintPreviewDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQPrintPreviewDialog_ContextMenuEvent(this, e); };
	int exec() { return callbackQPrintPreviewDialog_Exec(this); };
	void keyPressEvent(QKeyEvent * e) { callbackQPrintPreviewDialog_KeyPressEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewDialog_MinimumSizeHint(const_cast<MyQPrintPreviewDialog*>(this))); };
	void reject() { callbackQPrintPreviewDialog_Reject(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQPrintPreviewDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQPrintPreviewDialog_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewDialog_SizeHint(const_cast<MyQPrintPreviewDialog*>(this))); };
	void actionEvent(QActionEvent * event) { callbackQPrintPreviewDialog_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPrintPreviewDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPrintPreviewDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPrintPreviewDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPrintPreviewDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPrintPreviewDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPrintPreviewDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPrintPreviewDialog_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQPrintPreviewDialog_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPrintPreviewDialog_LeaveEvent(this, event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPrintPreviewDialog_Metric(const_cast<MyQPrintPreviewDialog*>(this), m); };
	void moveEvent(QMoveEvent * event) { callbackQPrintPreviewDialog_MoveEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrintPreviewDialog_PaintEngine(const_cast<MyQPrintPreviewDialog*>(this))); };
	void paintEvent(QPaintEvent * event) { callbackQPrintPreviewDialog_PaintEvent(this, event); };
	void setEnabled(bool vbo) { callbackQPrintPreviewDialog_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPrintPreviewDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPrintPreviewDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPrintPreviewDialog_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQPrintPreviewDialog_ChangeEvent(this, event); };
	bool close() { return callbackQPrintPreviewDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPrintPreviewDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPrintPreviewDialog_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQPrintPreviewDialog_HasHeightForWidth(const_cast<MyQPrintPreviewDialog*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQPrintPreviewDialog_HeightForWidth(const_cast<MyQPrintPreviewDialog*>(this), w); };
	void hide() { callbackQPrintPreviewDialog_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPrintPreviewDialog_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPrintPreviewDialog_InputMethodQuery(const_cast<MyQPrintPreviewDialog*>(this), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPrintPreviewDialog_KeyReleaseEvent(this, event); };
	void lower() { callbackQPrintPreviewDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPrintPreviewDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQPrintPreviewDialog_Raise(this); };
	void repaint() { callbackQPrintPreviewDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQPrintPreviewDialog_SetDisabled(this, disable); };
	void setFocus() { callbackQPrintPreviewDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPrintPreviewDialog_SetHidden(this, hidden); };
	void show() { callbackQPrintPreviewDialog_Show(this); };
	void showFullScreen() { callbackQPrintPreviewDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQPrintPreviewDialog_ShowMaximized(this); };
	void showMinimized() { callbackQPrintPreviewDialog_ShowMinimized(this); };
	void showNormal() { callbackQPrintPreviewDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPrintPreviewDialog_TabletEvent(this, event); };
	void update() { callbackQPrintPreviewDialog_Update(this); };
	void updateMicroFocus() { callbackQPrintPreviewDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPrintPreviewDialog_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQPrintPreviewDialog_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQPrintPreviewDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPrintPreviewDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPrintPreviewDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQPrintPreviewDialog_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPrintPreviewDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPrintPreviewDialog_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPrintPreviewDialog_MetaObject(const_cast<MyQPrintPreviewDialog*>(this))); };
};

void* QPrintPreviewDialog_NewQPrintPreviewDialog2(void* parent, long long flags)
{
	return new MyQPrintPreviewDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QPrintPreviewDialog_NewQPrintPreviewDialog(void* printer, void* parent, long long flags)
{
	return new MyQPrintPreviewDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QPrintPreviewDialog_Done(void* ptr, int result)
{
	static_cast<QPrintPreviewDialog*>(ptr)->done(result);
}

void QPrintPreviewDialog_DoneDefault(void* ptr, int result)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::done(result);
}

void QPrintPreviewDialog_Open(void* ptr, void* receiver, char* member)
{
	static_cast<QPrintPreviewDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QPrintPreviewDialog_ConnectPaintRequested(void* ptr)
{
	QObject::connect(static_cast<QPrintPreviewDialog*>(ptr), static_cast<void (QPrintPreviewDialog::*)(QPrinter *)>(&QPrintPreviewDialog::paintRequested), static_cast<MyQPrintPreviewDialog*>(ptr), static_cast<void (MyQPrintPreviewDialog::*)(QPrinter *)>(&MyQPrintPreviewDialog::Signal_PaintRequested));
}

void QPrintPreviewDialog_DisconnectPaintRequested(void* ptr)
{
	QObject::disconnect(static_cast<QPrintPreviewDialog*>(ptr), static_cast<void (QPrintPreviewDialog::*)(QPrinter *)>(&QPrintPreviewDialog::paintRequested), static_cast<MyQPrintPreviewDialog*>(ptr), static_cast<void (MyQPrintPreviewDialog::*)(QPrinter *)>(&MyQPrintPreviewDialog::Signal_PaintRequested));
}

void QPrintPreviewDialog_PaintRequested(void* ptr, void* printer)
{
	static_cast<QPrintPreviewDialog*>(ptr)->paintRequested(static_cast<QPrinter*>(printer));
}

void* QPrintPreviewDialog_Printer(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->printer();
}

void QPrintPreviewDialog_SetVisible(void* ptr, char visible)
{
	static_cast<QPrintPreviewDialog*>(ptr)->setVisible(visible != 0);
}

void QPrintPreviewDialog_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setVisible(visible != 0);
}

void QPrintPreviewDialog_DestroyQPrintPreviewDialog(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->~QPrintPreviewDialog();
}

void QPrintPreviewDialog_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "accept");
}

void QPrintPreviewDialog_AcceptDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::accept();
}

void QPrintPreviewDialog_CloseEvent(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->closeEvent(static_cast<QCloseEvent*>(e));
}

void QPrintPreviewDialog_CloseEventDefault(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QPrintPreviewDialog_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPrintPreviewDialog_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

int QPrintPreviewDialog_Exec(void* ptr)
{
	int returnArg;
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "exec", Q_RETURN_ARG(int, returnArg));
	return returnArg;
}

int QPrintPreviewDialog_ExecDefault(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::exec();
}

void QPrintPreviewDialog_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QPrintPreviewDialog_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void* QPrintPreviewDialog_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewDialog_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintPreviewDialog_Reject(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "reject");
}

void QPrintPreviewDialog_RejectDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::reject();
}

void QPrintPreviewDialog_ResizeEvent(void* ptr, void* vqr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPrintPreviewDialog_ResizeEventDefault(void* ptr, void* vqr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPrintPreviewDialog_ShowEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QPrintPreviewDialog_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QPrintPreviewDialog_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewDialog_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintPreviewDialog_ActionEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewDialog_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewDialog_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewDialog_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewDialog_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewDialog_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewDialog_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewDialog_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewDialog_DropEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewDialog_DropEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewDialog_EnterEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_FocusInEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_HideEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewDialog_HideEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewDialog_LeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::leaveEvent(static_cast<QEvent*>(event));
}

int QPrintPreviewDialog_Metric(void* ptr, long long m)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QPrintPreviewDialog_MetricDefault(void* ptr, long long m)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QPrintPreviewDialog_MoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QPrintPreviewDialog_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QPrintPreviewDialog_PaintEngine(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->paintEngine();
}

void* QPrintPreviewDialog_PaintEngineDefault(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::paintEngine();
}

void QPrintPreviewDialog_PaintEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewDialog_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewDialog_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QPrintPreviewDialog_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setEnabled(vbo != 0);
}

void QPrintPreviewDialog_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QPrintPreviewDialog_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setStyleSheet(QString(styleSheet));
}

void QPrintPreviewDialog_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QPrintPreviewDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setWindowModified(vbo != 0);
}

void QPrintPreviewDialog_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QPrintPreviewDialog_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setWindowTitle(QString(vqs));
}

void QPrintPreviewDialog_ChangeEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::changeEvent(static_cast<QEvent*>(event));
}

char QPrintPreviewDialog_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QPrintPreviewDialog_CloseDefault(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::close();
}

char QPrintPreviewDialog_Event(void* ptr, void* event)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->event(static_cast<QEvent*>(event));
}

char QPrintPreviewDialog_EventDefault(void* ptr, void* event)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::event(static_cast<QEvent*>(event));
}

char QPrintPreviewDialog_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->focusNextPrevChild(next != 0);
}

char QPrintPreviewDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusNextPrevChild(next != 0);
}

char QPrintPreviewDialog_HasHeightForWidth(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->hasHeightForWidth();
}

char QPrintPreviewDialog_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hasHeightForWidth();
}

int QPrintPreviewDialog_HeightForWidth(void* ptr, int w)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->heightForWidth(w);
}

int QPrintPreviewDialog_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::heightForWidth(w);
}

void QPrintPreviewDialog_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "hide");
}

void QPrintPreviewDialog_HideDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hide();
}

void QPrintPreviewDialog_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPrintPreviewDialog_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QPrintPreviewDialog_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintPreviewDialog*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QPrintPreviewDialog_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QPrintPreviewDialog_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewDialog_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "lower");
}

void QPrintPreviewDialog_LowerDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::lower();
}

void QPrintPreviewDialog_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MousePressEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QPrintPreviewDialog_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QPrintPreviewDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPrintPreviewDialog_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "raise");
}

void QPrintPreviewDialog_RaiseDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::raise();
}

void QPrintPreviewDialog_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "repaint");
}

void QPrintPreviewDialog_RepaintDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::repaint();
}

void QPrintPreviewDialog_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QPrintPreviewDialog_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setDisabled(disable != 0);
}

void QPrintPreviewDialog_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setFocus");
}

void QPrintPreviewDialog_SetFocus2Default(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setFocus();
}

void QPrintPreviewDialog_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QPrintPreviewDialog_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setHidden(hidden != 0);
}

void QPrintPreviewDialog_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "show");
}

void QPrintPreviewDialog_ShowDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::show();
}

void QPrintPreviewDialog_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "showFullScreen");
}

void QPrintPreviewDialog_ShowFullScreenDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showFullScreen();
}

void QPrintPreviewDialog_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "showMaximized");
}

void QPrintPreviewDialog_ShowMaximizedDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showMaximized();
}

void QPrintPreviewDialog_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "showMinimized");
}

void QPrintPreviewDialog_ShowMinimizedDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showMinimized();
}

void QPrintPreviewDialog_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "showNormal");
}

void QPrintPreviewDialog_ShowNormalDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showNormal();
}

void QPrintPreviewDialog_TabletEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewDialog_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewDialog_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "update");
}

void QPrintPreviewDialog_UpdateDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::update();
}

void QPrintPreviewDialog_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "updateMicroFocus");
}

void QPrintPreviewDialog_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::updateMicroFocus();
}

void QPrintPreviewDialog_WheelEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintPreviewDialog_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintPreviewDialog_TimerEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintPreviewDialog_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintPreviewDialog_ChildEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewDialog_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewDialog_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintPreviewDialog*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewDialog_CustomEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewDialog*>(ptr), "deleteLater");
}

void QPrintPreviewDialog_DeleteLaterDefault(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::deleteLater();
}

void QPrintPreviewDialog_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintPreviewDialog*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QPrintPreviewDialog_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QPrintPreviewDialog_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPrintPreviewDialog_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintPreviewDialog*>(ptr)->metaObject());
}

void* QPrintPreviewDialog_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::metaObject());
}

class MyQPrintPreviewWidget: public QPrintPreviewWidget
{
public:
	MyQPrintPreviewWidget(QPrinter *printer, QWidget *parent, Qt::WindowFlags flags) : QPrintPreviewWidget(printer, parent, flags) {};
	MyQPrintPreviewWidget(QWidget *parent, Qt::WindowFlags flags) : QPrintPreviewWidget(parent, flags) {};
	void fitInView() { callbackQPrintPreviewWidget_FitInView(this); };
	void fitToWidth() { callbackQPrintPreviewWidget_FitToWidth(this); };
	void Signal_PaintRequested(QPrinter * printer) { callbackQPrintPreviewWidget_PaintRequested(this, printer); };
	void Signal_PreviewChanged() { callbackQPrintPreviewWidget_PreviewChanged(this); };
	void print() { callbackQPrintPreviewWidget_Print(this); };
	void setAllPagesViewMode() { callbackQPrintPreviewWidget_SetAllPagesViewMode(this); };
	void setCurrentPage(int page) { callbackQPrintPreviewWidget_SetCurrentPage(this, page); };
	void setFacingPagesViewMode() { callbackQPrintPreviewWidget_SetFacingPagesViewMode(this); };
	void setLandscapeOrientation() { callbackQPrintPreviewWidget_SetLandscapeOrientation(this); };
	void setOrientation(QPrinter::Orientation orientation) { callbackQPrintPreviewWidget_SetOrientation(this, orientation); };
	void setPortraitOrientation() { callbackQPrintPreviewWidget_SetPortraitOrientation(this); };
	void setSinglePageViewMode() { callbackQPrintPreviewWidget_SetSinglePageViewMode(this); };
	void setViewMode(QPrintPreviewWidget::ViewMode mode) { callbackQPrintPreviewWidget_SetViewMode(this, mode); };
	void setVisible(bool visible) { callbackQPrintPreviewWidget_SetVisible(this, visible); };
	void setZoomFactor(qreal factor) { callbackQPrintPreviewWidget_SetZoomFactor(this, factor); };
	void setZoomMode(QPrintPreviewWidget::ZoomMode zoomMode) { callbackQPrintPreviewWidget_SetZoomMode(this, zoomMode); };
	void updatePreview() { callbackQPrintPreviewWidget_UpdatePreview(this); };
	void zoomIn(qreal factor) { callbackQPrintPreviewWidget_ZoomIn(this, factor); };
	void zoomOut(qreal factor) { callbackQPrintPreviewWidget_ZoomOut(this, factor); };
	void actionEvent(QActionEvent * event) { callbackQPrintPreviewWidget_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPrintPreviewWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPrintPreviewWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPrintPreviewWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPrintPreviewWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPrintPreviewWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPrintPreviewWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPrintPreviewWidget_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQPrintPreviewWidget_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPrintPreviewWidget_LeaveEvent(this, event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPrintPreviewWidget_Metric(const_cast<MyQPrintPreviewWidget*>(this), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewWidget_MinimumSizeHint(const_cast<MyQPrintPreviewWidget*>(this))); };
	void moveEvent(QMoveEvent * event) { callbackQPrintPreviewWidget_MoveEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrintPreviewWidget_PaintEngine(const_cast<MyQPrintPreviewWidget*>(this))); };
	void paintEvent(QPaintEvent * event) { callbackQPrintPreviewWidget_PaintEvent(this, event); };
	void setEnabled(bool vbo) { callbackQPrintPreviewWidget_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPrintPreviewWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPrintPreviewWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPrintPreviewWidget_SetWindowTitle(this, vqsPacked); };
	void showEvent(QShowEvent * event) { callbackQPrintPreviewWidget_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewWidget_SizeHint(const_cast<MyQPrintPreviewWidget*>(this))); };
	void changeEvent(QEvent * event) { callbackQPrintPreviewWidget_ChangeEvent(this, event); };
	bool close() { return callbackQPrintPreviewWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQPrintPreviewWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQPrintPreviewWidget_ContextMenuEvent(this, event); };
	bool event(QEvent * event) { return callbackQPrintPreviewWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPrintPreviewWidget_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQPrintPreviewWidget_HasHeightForWidth(const_cast<MyQPrintPreviewWidget*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQPrintPreviewWidget_HeightForWidth(const_cast<MyQPrintPreviewWidget*>(this), w); };
	void hide() { callbackQPrintPreviewWidget_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPrintPreviewWidget_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPrintPreviewWidget_InputMethodQuery(const_cast<MyQPrintPreviewWidget*>(this), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQPrintPreviewWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPrintPreviewWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQPrintPreviewWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPrintPreviewWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQPrintPreviewWidget_Raise(this); };
	void repaint() { callbackQPrintPreviewWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQPrintPreviewWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQPrintPreviewWidget_SetDisabled(this, disable); };
	void setFocus() { callbackQPrintPreviewWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPrintPreviewWidget_SetHidden(this, hidden); };
	void show() { callbackQPrintPreviewWidget_Show(this); };
	void showFullScreen() { callbackQPrintPreviewWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQPrintPreviewWidget_ShowMaximized(this); };
	void showMinimized() { callbackQPrintPreviewWidget_ShowMinimized(this); };
	void showNormal() { callbackQPrintPreviewWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPrintPreviewWidget_TabletEvent(this, event); };
	void update() { callbackQPrintPreviewWidget_Update(this); };
	void updateMicroFocus() { callbackQPrintPreviewWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPrintPreviewWidget_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQPrintPreviewWidget_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQPrintPreviewWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPrintPreviewWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPrintPreviewWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQPrintPreviewWidget_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPrintPreviewWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPrintPreviewWidget_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPrintPreviewWidget_MetaObject(const_cast<MyQPrintPreviewWidget*>(this))); };
};

void* QPrintPreviewWidget_NewQPrintPreviewWidget(void* printer, void* parent, long long flags)
{
	return new MyQPrintPreviewWidget(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QPrintPreviewWidget_NewQPrintPreviewWidget2(void* parent, long long flags)
{
	return new MyQPrintPreviewWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QPrintPreviewWidget_CurrentPage(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->currentPage();
}

void QPrintPreviewWidget_FitInView(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "fitInView");
}

void QPrintPreviewWidget_FitInViewDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::fitInView();
}

void QPrintPreviewWidget_FitToWidth(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "fitToWidth");
}

void QPrintPreviewWidget_FitToWidthDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::fitToWidth();
}

long long QPrintPreviewWidget_Orientation(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->orientation();
}

int QPrintPreviewWidget_PageCount(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->pageCount();
}

void QPrintPreviewWidget_ConnectPaintRequested(void* ptr)
{
	QObject::connect(static_cast<QPrintPreviewWidget*>(ptr), static_cast<void (QPrintPreviewWidget::*)(QPrinter *)>(&QPrintPreviewWidget::paintRequested), static_cast<MyQPrintPreviewWidget*>(ptr), static_cast<void (MyQPrintPreviewWidget::*)(QPrinter *)>(&MyQPrintPreviewWidget::Signal_PaintRequested));
}

void QPrintPreviewWidget_DisconnectPaintRequested(void* ptr)
{
	QObject::disconnect(static_cast<QPrintPreviewWidget*>(ptr), static_cast<void (QPrintPreviewWidget::*)(QPrinter *)>(&QPrintPreviewWidget::paintRequested), static_cast<MyQPrintPreviewWidget*>(ptr), static_cast<void (MyQPrintPreviewWidget::*)(QPrinter *)>(&MyQPrintPreviewWidget::Signal_PaintRequested));
}

void QPrintPreviewWidget_PaintRequested(void* ptr, void* printer)
{
	static_cast<QPrintPreviewWidget*>(ptr)->paintRequested(static_cast<QPrinter*>(printer));
}

void QPrintPreviewWidget_ConnectPreviewChanged(void* ptr)
{
	QObject::connect(static_cast<QPrintPreviewWidget*>(ptr), static_cast<void (QPrintPreviewWidget::*)()>(&QPrintPreviewWidget::previewChanged), static_cast<MyQPrintPreviewWidget*>(ptr), static_cast<void (MyQPrintPreviewWidget::*)()>(&MyQPrintPreviewWidget::Signal_PreviewChanged));
}

void QPrintPreviewWidget_DisconnectPreviewChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPrintPreviewWidget*>(ptr), static_cast<void (QPrintPreviewWidget::*)()>(&QPrintPreviewWidget::previewChanged), static_cast<MyQPrintPreviewWidget*>(ptr), static_cast<void (MyQPrintPreviewWidget::*)()>(&MyQPrintPreviewWidget::Signal_PreviewChanged));
}

void QPrintPreviewWidget_PreviewChanged(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->previewChanged();
}

void QPrintPreviewWidget_Print(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "print");
}

void QPrintPreviewWidget_PrintDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::print();
}

void QPrintPreviewWidget_SetAllPagesViewMode(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setAllPagesViewMode");
}

void QPrintPreviewWidget_SetAllPagesViewModeDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setAllPagesViewMode();
}

void QPrintPreviewWidget_SetCurrentPage(void* ptr, int page)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setCurrentPage", Q_ARG(int, page));
}

void QPrintPreviewWidget_SetCurrentPageDefault(void* ptr, int page)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setCurrentPage(page);
}

void QPrintPreviewWidget_SetFacingPagesViewMode(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setFacingPagesViewMode");
}

void QPrintPreviewWidget_SetFacingPagesViewModeDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setFacingPagesViewMode();
}

void QPrintPreviewWidget_SetLandscapeOrientation(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setLandscapeOrientation");
}

void QPrintPreviewWidget_SetLandscapeOrientationDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setLandscapeOrientation();
}

void QPrintPreviewWidget_SetOrientation(void* ptr, long long orientation)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setOrientation", Q_ARG(QPrinter::Orientation, static_cast<QPrinter::Orientation>(orientation)));
}

void QPrintPreviewWidget_SetOrientationDefault(void* ptr, long long orientation)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setOrientation(static_cast<QPrinter::Orientation>(orientation));
}

void QPrintPreviewWidget_SetPortraitOrientation(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setPortraitOrientation");
}

void QPrintPreviewWidget_SetPortraitOrientationDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setPortraitOrientation();
}

void QPrintPreviewWidget_SetSinglePageViewMode(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setSinglePageViewMode");
}

void QPrintPreviewWidget_SetSinglePageViewModeDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setSinglePageViewMode();
}

void QPrintPreviewWidget_SetViewMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setViewMode", Q_ARG(QPrintPreviewWidget::ViewMode, static_cast<QPrintPreviewWidget::ViewMode>(mode)));
}

void QPrintPreviewWidget_SetViewModeDefault(void* ptr, long long mode)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setViewMode(static_cast<QPrintPreviewWidget::ViewMode>(mode));
}

void QPrintPreviewWidget_SetVisible(void* ptr, char visible)
{
	static_cast<QPrintPreviewWidget*>(ptr)->setVisible(visible != 0);
}

void QPrintPreviewWidget_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setVisible(visible != 0);
}

void QPrintPreviewWidget_SetZoomFactor(void* ptr, double factor)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setZoomFactor", Q_ARG(qreal, factor));
}

void QPrintPreviewWidget_SetZoomFactorDefault(void* ptr, double factor)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setZoomFactor(factor);
}

void QPrintPreviewWidget_SetZoomMode(void* ptr, long long zoomMode)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setZoomMode", Q_ARG(QPrintPreviewWidget::ZoomMode, static_cast<QPrintPreviewWidget::ZoomMode>(zoomMode)));
}

void QPrintPreviewWidget_SetZoomModeDefault(void* ptr, long long zoomMode)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setZoomMode(static_cast<QPrintPreviewWidget::ZoomMode>(zoomMode));
}

void QPrintPreviewWidget_UpdatePreview(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "updatePreview");
}

void QPrintPreviewWidget_UpdatePreviewDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::updatePreview();
}

long long QPrintPreviewWidget_ViewMode(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->viewMode();
}

double QPrintPreviewWidget_ZoomFactor(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->zoomFactor();
}

void QPrintPreviewWidget_ZoomIn(void* ptr, double factor)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "zoomIn", Q_ARG(qreal, factor));
}

void QPrintPreviewWidget_ZoomInDefault(void* ptr, double factor)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::zoomIn(factor);
}

long long QPrintPreviewWidget_ZoomMode(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->zoomMode();
}

void QPrintPreviewWidget_ZoomOut(void* ptr, double factor)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "zoomOut", Q_ARG(qreal, factor));
}

void QPrintPreviewWidget_ZoomOutDefault(void* ptr, double factor)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::zoomOut(factor);
}

void QPrintPreviewWidget_DestroyQPrintPreviewWidget(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->~QPrintPreviewWidget();
}

void QPrintPreviewWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::leaveEvent(static_cast<QEvent*>(event));
}

int QPrintPreviewWidget_Metric(void* ptr, long long m)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QPrintPreviewWidget_MetricDefault(void* ptr, long long m)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QPrintPreviewWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintPreviewWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QPrintPreviewWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QPrintPreviewWidget_PaintEngine(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->paintEngine();
}

void* QPrintPreviewWidget_PaintEngineDefault(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::paintEngine();
}

void QPrintPreviewWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewWidget_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QPrintPreviewWidget_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setEnabled(vbo != 0);
}

void QPrintPreviewWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QPrintPreviewWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setStyleSheet(QString(styleSheet));
}

void QPrintPreviewWidget_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QPrintPreviewWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setWindowModified(vbo != 0);
}

void QPrintPreviewWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QPrintPreviewWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setWindowTitle(QString(vqs));
}

void QPrintPreviewWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QPrintPreviewWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QPrintPreviewWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPrintPreviewWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::changeEvent(static_cast<QEvent*>(event));
}

char QPrintPreviewWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QPrintPreviewWidget_CloseDefault(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::close();
}

void QPrintPreviewWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QPrintPreviewWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QPrintPreviewWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QPrintPreviewWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

char QPrintPreviewWidget_Event(void* ptr, void* event)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->event(static_cast<QEvent*>(event));
}

char QPrintPreviewWidget_EventDefault(void* ptr, void* event)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::event(static_cast<QEvent*>(event));
}

char QPrintPreviewWidget_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->focusNextPrevChild(next != 0);
}

char QPrintPreviewWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusNextPrevChild(next != 0);
}

char QPrintPreviewWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->hasHeightForWidth();
}

char QPrintPreviewWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hasHeightForWidth();
}

int QPrintPreviewWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->heightForWidth(w);
}

int QPrintPreviewWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::heightForWidth(w);
}

void QPrintPreviewWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "hide");
}

void QPrintPreviewWidget_HideDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hide();
}

void QPrintPreviewWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPrintPreviewWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QPrintPreviewWidget_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintPreviewWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QPrintPreviewWidget_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QPrintPreviewWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "lower");
}

void QPrintPreviewWidget_LowerDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::lower();
}

void QPrintPreviewWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QPrintPreviewWidget_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QPrintPreviewWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPrintPreviewWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "raise");
}

void QPrintPreviewWidget_RaiseDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::raise();
}

void QPrintPreviewWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "repaint");
}

void QPrintPreviewWidget_RepaintDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::repaint();
}

void QPrintPreviewWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QPrintPreviewWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QPrintPreviewWidget_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QPrintPreviewWidget_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setDisabled(disable != 0);
}

void QPrintPreviewWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setFocus");
}

void QPrintPreviewWidget_SetFocus2Default(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setFocus();
}

void QPrintPreviewWidget_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QPrintPreviewWidget_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setHidden(hidden != 0);
}

void QPrintPreviewWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "show");
}

void QPrintPreviewWidget_ShowDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::show();
}

void QPrintPreviewWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "showFullScreen");
}

void QPrintPreviewWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showFullScreen();
}

void QPrintPreviewWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "showMaximized");
}

void QPrintPreviewWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showMaximized();
}

void QPrintPreviewWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "showMinimized");
}

void QPrintPreviewWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showMinimized();
}

void QPrintPreviewWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "showNormal");
}

void QPrintPreviewWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showNormal();
}

void QPrintPreviewWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "update");
}

void QPrintPreviewWidget_UpdateDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::update();
}

void QPrintPreviewWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "updateMicroFocus");
}

void QPrintPreviewWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::updateMicroFocus();
}

void QPrintPreviewWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintPreviewWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QPrintPreviewWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintPreviewWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPrintPreviewWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintPreviewWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "deleteLater");
}

void QPrintPreviewWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::deleteLater();
}

void QPrintPreviewWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPrintPreviewWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QPrintPreviewWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QPrintPreviewWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPrintPreviewWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintPreviewWidget*>(ptr)->metaObject());
}

void* QPrintPreviewWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::metaObject());
}

class MyQPrinter: public QPrinter
{
public:
	MyQPrinter(PrinterMode mode) : QPrinter(mode) {};
	MyQPrinter(const QPrinterInfo &printer, PrinterMode mode) : QPrinter(printer, mode) {};
	bool newPage() { return callbackQPrinter_NewPage(this) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrinter_PaintEngine(const_cast<MyQPrinter*>(this))); };
	bool setPageSize(const QPageSize & pageSize) { return callbackQPrinter_SetPageSize(this, const_cast<QPageSize*>(&pageSize)) != 0; };
	void setPageSize(QPagedPaintDevice::PageSize size) { callbackQPrinter_SetPageSize2(this, size); };
	void setPageSizeMM(const QSizeF & size) { callbackQPrinter_SetPageSizeMM(this, const_cast<QSizeF*>(&size)); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPrinter_Metric(const_cast<MyQPrinter*>(this), metric); };
};

int QPrinter_FromPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->fromPage();
}

struct QtPrintSupport_PackedString QPrinter_OutputFileName(void* ptr)
{
	return ({ QByteArray tde340f = static_cast<QPrinter*>(ptr)->outputFileName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tde340f.prepend("WHITESPACE").constData()+10), tde340f.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_PrinterSelectionOption(void* ptr)
{
	return ({ QByteArray te1c3f5 = static_cast<QPrinter*>(ptr)->printerSelectionOption().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(te1c3f5.prepend("WHITESPACE").constData()+10), te1c3f5.size()-10 }; });
}

void QPrinter_SetPrinterSelectionOption(void* ptr, char* option)
{
	static_cast<QPrinter*>(ptr)->setPrinterSelectionOption(QString(option));
}

void* QPrinter_NewQPrinter(long long mode)
{
	return new MyQPrinter(static_cast<QPrinter::PrinterMode>(mode));
}

void* QPrinter_NewQPrinter2(void* printer, long long mode)
{
	return new MyQPrinter(*static_cast<QPrinterInfo*>(printer), static_cast<QPrinter::PrinterMode>(mode));
}

char QPrinter_Abort(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->abort();
}

char QPrinter_CollateCopies(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->collateCopies();
}

long long QPrinter_ColorMode(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->colorMode();
}

int QPrinter_CopyCount(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->copyCount();
}

struct QtPrintSupport_PackedString QPrinter_Creator(void* ptr)
{
	return ({ QByteArray ta24181 = static_cast<QPrinter*>(ptr)->creator().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(ta24181.prepend("WHITESPACE").constData()+10), ta24181.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_DocName(void* ptr)
{
	return ({ QByteArray t72c43e = static_cast<QPrinter*>(ptr)->docName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t72c43e.prepend("WHITESPACE").constData()+10), t72c43e.size()-10 }; });
}

long long QPrinter_Duplex(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->duplex();
}

char QPrinter_FontEmbeddingEnabled(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->fontEmbeddingEnabled();
}

char QPrinter_FullPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->fullPage();
}

char QPrinter_IsValid(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->isValid();
}

char QPrinter_NewPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->newPage();
}

char QPrinter_NewPageDefault(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->QPrinter::newPage();
}

long long QPrinter_OutputFormat(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->outputFormat();
}

void* QPrinter_PageLayout(void* ptr)
{
	return new QPageLayout(static_cast<QPrinter*>(ptr)->pageLayout());
}

long long QPrinter_PageOrder(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->pageOrder();
}

void* QPrinter_PageRect(void* ptr, long long unit)
{
	return ({ QRectF tmpValue = static_cast<QPrinter*>(ptr)->pageRect(static_cast<QPrinter::Unit>(unit)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPrinter_PaintEngine(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->paintEngine();
}

void* QPrinter_PaintEngineDefault(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->QPrinter::paintEngine();
}

void* QPrinter_PaperRect(void* ptr, long long unit)
{
	return ({ QRectF tmpValue = static_cast<QPrinter*>(ptr)->paperRect(static_cast<QPrinter::Unit>(unit)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QPrinter_PaperSource(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->paperSource();
}

void* QPrinter_PrintEngine(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printEngine();
}

struct QtPrintSupport_PackedString QPrinter_PrintProgram(void* ptr)
{
	return ({ QByteArray tb89767 = static_cast<QPrinter*>(ptr)->printProgram().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tb89767.prepend("WHITESPACE").constData()+10), tb89767.size()-10 }; });
}

long long QPrinter_PrintRange(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printRange();
}

struct QtPrintSupport_PackedString QPrinter_PrinterName(void* ptr)
{
	return ({ QByteArray ta7ec2f = static_cast<QPrinter*>(ptr)->printerName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(ta7ec2f.prepend("WHITESPACE").constData()+10), ta7ec2f.size()-10 }; });
}

long long QPrinter_PrinterState(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printerState();
}

int QPrinter_Resolution(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->resolution();
}

void QPrinter_SetCollateCopies(void* ptr, char collate)
{
	static_cast<QPrinter*>(ptr)->setCollateCopies(collate != 0);
}

void QPrinter_SetColorMode(void* ptr, long long newColorMode)
{
	static_cast<QPrinter*>(ptr)->setColorMode(static_cast<QPrinter::ColorMode>(newColorMode));
}

void QPrinter_SetCopyCount(void* ptr, int count)
{
	static_cast<QPrinter*>(ptr)->setCopyCount(count);
}

void QPrinter_SetCreator(void* ptr, char* creator)
{
	static_cast<QPrinter*>(ptr)->setCreator(QString(creator));
}

void QPrinter_SetDocName(void* ptr, char* name)
{
	static_cast<QPrinter*>(ptr)->setDocName(QString(name));
}

void QPrinter_SetDuplex(void* ptr, long long duplex)
{
	static_cast<QPrinter*>(ptr)->setDuplex(static_cast<QPrinter::DuplexMode>(duplex));
}

void QPrinter_SetEngines(void* ptr, void* printEngine, void* paintEngine)
{
	static_cast<QPrinter*>(ptr)->setEngines(static_cast<QPrintEngine*>(printEngine), static_cast<QPaintEngine*>(paintEngine));
}

void QPrinter_SetFontEmbeddingEnabled(void* ptr, char enable)
{
	static_cast<QPrinter*>(ptr)->setFontEmbeddingEnabled(enable != 0);
}

void QPrinter_SetFromTo(void* ptr, int from, int to)
{
	static_cast<QPrinter*>(ptr)->setFromTo(from, to);
}

void QPrinter_SetFullPage(void* ptr, char fp)
{
	static_cast<QPrinter*>(ptr)->setFullPage(fp != 0);
}

void QPrinter_SetOutputFileName(void* ptr, char* fileName)
{
	static_cast<QPrinter*>(ptr)->setOutputFileName(QString(fileName));
}

void QPrinter_SetOutputFormat(void* ptr, long long format)
{
	static_cast<QPrinter*>(ptr)->setOutputFormat(static_cast<QPrinter::OutputFormat>(format));
}

char QPrinter_SetPageLayout(void* ptr, void* newLayout)
{
	return static_cast<QPrinter*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newLayout));
}

char QPrinter_SetPageMargins(void* ptr, void* margins)
{
	return static_cast<QPrinter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

char QPrinter_SetPageMargins2(void* ptr, void* margins, long long units)
{
	return static_cast<QPrinter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

void QPrinter_SetPageOrder(void* ptr, long long pageOrder)
{
	static_cast<QPrinter*>(ptr)->setPageOrder(static_cast<QPrinter::PageOrder>(pageOrder));
}

char QPrinter_SetPageOrientation(void* ptr, long long orientation)
{
	return static_cast<QPrinter*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

char QPrinter_SetPageSize(void* ptr, void* pageSize)
{
	return static_cast<QPrinter*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

char QPrinter_SetPageSizeDefault(void* ptr, void* pageSize)
{
	return static_cast<QPrinter*>(ptr)->QPrinter::setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPrinter_SetPaperSource(void* ptr, long long source)
{
	static_cast<QPrinter*>(ptr)->setPaperSource(static_cast<QPrinter::PaperSource>(source));
}

void QPrinter_SetPrintProgram(void* ptr, char* printProg)
{
	static_cast<QPrinter*>(ptr)->setPrintProgram(QString(printProg));
}

void QPrinter_SetPrintRange(void* ptr, long long ran)
{
	static_cast<QPrinter*>(ptr)->setPrintRange(static_cast<QPrinter::PrintRange>(ran));
}

void QPrinter_SetPrinterName(void* ptr, char* name)
{
	static_cast<QPrinter*>(ptr)->setPrinterName(QString(name));
}

void QPrinter_SetResolution(void* ptr, int dpi)
{
	static_cast<QPrinter*>(ptr)->setResolution(dpi);
}

char QPrinter_SupportsMultipleCopies(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->supportsMultipleCopies();
}

int QPrinter_ToPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->toPage();
}

void QPrinter_DestroyQPrinter(void* ptr)
{
	static_cast<QPrinter*>(ptr)->~QPrinter();
}

void QPrinter_SetPageSize2(void* ptr, long long size)
{
	static_cast<QPrinter*>(ptr)->setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPrinter_SetPageSize2Default(void* ptr, long long size)
{
	static_cast<QPrinter*>(ptr)->QPrinter::setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPrinter_SetPageSizeMM(void* ptr, void* size)
{
	static_cast<QPrinter*>(ptr)->setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPrinter_SetPageSizeMMDefault(void* ptr, void* size)
{
	static_cast<QPrinter*>(ptr)->QPrinter::setPageSizeMM(*static_cast<QSizeF*>(size));
}

int QPrinter_Metric(void* ptr, long long metric)
{
	return static_cast<QPrinter*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPrinter_MetricDefault(void* ptr, long long metric)
{
	return static_cast<QPrinter*>(ptr)->QPrinter::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPrinterInfo_NewQPrinterInfo()
{
	return new QPrinterInfo();
}

void* QPrinterInfo_NewQPrinterInfo3(void* printer)
{
	return new QPrinterInfo(*static_cast<QPrinter*>(printer));
}

void* QPrinterInfo_NewQPrinterInfo2(void* other)
{
	return new QPrinterInfo(*static_cast<QPrinterInfo*>(other));
}

struct QtPrintSupport_PackedString QPrinterInfo_QPrinterInfo_AvailablePrinterNames()
{
	return ({ QByteArray tb6ad9e = QPrinterInfo::availablePrinterNames().join("|").toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tb6ad9e.prepend("WHITESPACE").constData()+10), tb6ad9e.size()-10 }; });
}

struct QtPrintSupport_PackedList QPrinterInfo_QPrinterInfo_AvailablePrinters()
{
	return ({ QList<QPrinterInfo>* tmpValue = new QList<QPrinterInfo>(QPrinterInfo::availablePrinters()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

long long QPrinterInfo_DefaultDuplexMode(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->defaultDuplexMode();
}

void* QPrinterInfo_DefaultPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->defaultPageSize());
}

void* QPrinterInfo_QPrinterInfo_DefaultPrinter()
{
	return new QPrinterInfo(QPrinterInfo::defaultPrinter());
}

struct QtPrintSupport_PackedString QPrinterInfo_QPrinterInfo_DefaultPrinterName()
{
	return ({ QByteArray t8ce1a9 = QPrinterInfo::defaultPrinterName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t8ce1a9.prepend("WHITESPACE").constData()+10), t8ce1a9.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_Description(void* ptr)
{
	return ({ QByteArray t0d7900 = static_cast<QPrinterInfo*>(ptr)->description().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t0d7900.prepend("WHITESPACE").constData()+10), t0d7900.size()-10 }; });
}

char QPrinterInfo_IsDefault(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->isDefault();
}

char QPrinterInfo_IsNull(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->isNull();
}

char QPrinterInfo_IsRemote(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->isRemote();
}

struct QtPrintSupport_PackedString QPrinterInfo_Location(void* ptr)
{
	return ({ QByteArray t360fc3 = static_cast<QPrinterInfo*>(ptr)->location().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t360fc3.prepend("WHITESPACE").constData()+10), t360fc3.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_MakeAndModel(void* ptr)
{
	return ({ QByteArray t372e5c = static_cast<QPrinterInfo*>(ptr)->makeAndModel().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t372e5c.prepend("WHITESPACE").constData()+10), t372e5c.size()-10 }; });
}

void* QPrinterInfo_MaximumPhysicalPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->maximumPhysicalPageSize());
}

void* QPrinterInfo_MinimumPhysicalPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->minimumPhysicalPageSize());
}

void* QPrinterInfo_QPrinterInfo_PrinterInfo(char* printerName)
{
	return new QPrinterInfo(QPrinterInfo::printerInfo(QString(printerName)));
}

struct QtPrintSupport_PackedString QPrinterInfo_PrinterName(void* ptr)
{
	return ({ QByteArray te478b0 = static_cast<QPrinterInfo*>(ptr)->printerName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(te478b0.prepend("WHITESPACE").constData()+10), te478b0.size()-10 }; });
}

long long QPrinterInfo_State(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->state();
}

struct QtPrintSupport_PackedList QPrinterInfo_SupportedPageSizes(void* ptr)
{
	return ({ QList<QPageSize>* tmpValue = new QList<QPageSize>(static_cast<QPrinterInfo*>(ptr)->supportedPageSizes()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

char QPrinterInfo_SupportsCustomPageSizes(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->supportsCustomPageSizes();
}

void QPrinterInfo_DestroyQPrinterInfo(void* ptr)
{
	static_cast<QPrinterInfo*>(ptr)->~QPrinterInfo();
}

void* QPrinterInfo_availablePrinters_atList(void* ptr, int i)
{
	return new QPrinterInfo(static_cast<QList<QPrinterInfo>*>(ptr)->at(i));
}

void* QPrinterInfo_supportedPageSizes_atList(void* ptr, int i)
{
	return new QPageSize(static_cast<QList<QPageSize>*>(ptr)->at(i));
}

