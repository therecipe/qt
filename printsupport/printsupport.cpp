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
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPageSetupDialog>
#include <QPageSize>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPoint>
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
	MyQAbstractPrintDialog(QPrinter *printer, QWidget *parent = Q_NULLPTR) : QAbstractPrintDialog(printer, parent) {QAbstractPrintDialog_QAbstractPrintDialog_QRegisterMetaType();};
	int exec() { return callbackQAbstractPrintDialog_Exec(this); };
	void accept() { callbackQAbstractPrintDialog_Accept(this); };
	void Signal_Accepted() { callbackQAbstractPrintDialog_Accepted(this); };
	void closeEvent(QCloseEvent * e) { callbackQAbstractPrintDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQAbstractPrintDialog_ContextMenuEvent(this, e); };
	void done(int r) { callbackQAbstractPrintDialog_Done(this, r); };
	void Signal_Finished(int result) { callbackQAbstractPrintDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQAbstractPrintDialog_KeyPressEvent(this, e); };
	void open() { callbackQAbstractPrintDialog_Open(this); };
	void reject() { callbackQAbstractPrintDialog_Reject(this); };
	void Signal_Rejected() { callbackQAbstractPrintDialog_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQAbstractPrintDialog_ResizeEvent(this, vqr); };
	void setVisible(bool visible) { callbackQAbstractPrintDialog_SetVisible(this, visible); };
	void showEvent(QShowEvent * event) { callbackQAbstractPrintDialog_ShowEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQAbstractPrintDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQAbstractPrintDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQAbstractPrintDialog_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQAbstractPrintDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQAbstractPrintDialog_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQAbstractPrintDialog_ChangeEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQAbstractPrintDialog_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQAbstractPrintDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQAbstractPrintDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQAbstractPrintDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQAbstractPrintDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQAbstractPrintDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusOutEvent(this, event); };
	void hide() { callbackQAbstractPrintDialog_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQAbstractPrintDialog_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQAbstractPrintDialog_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQAbstractPrintDialog_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQAbstractPrintDialog_LeaveEvent(this, event); };
	void lower() { callbackQAbstractPrintDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQAbstractPrintDialog_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQAbstractPrintDialog_PaintEvent(this, event); };
	void raise() { callbackQAbstractPrintDialog_Raise(this); };
	void repaint() { callbackQAbstractPrintDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQAbstractPrintDialog_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQAbstractPrintDialog_SetEnabled(this, vbo); };
	void setFocus() { callbackQAbstractPrintDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQAbstractPrintDialog_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQAbstractPrintDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQAbstractPrintDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQAbstractPrintDialog_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQAbstractPrintDialog_Show(this); };
	void showFullScreen() { callbackQAbstractPrintDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQAbstractPrintDialog_ShowMaximized(this); };
	void showMinimized() { callbackQAbstractPrintDialog_ShowMinimized(this); };
	void showNormal() { callbackQAbstractPrintDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQAbstractPrintDialog_TabletEvent(this, event); };
	void update() { callbackQAbstractPrintDialog_Update(this); };
	void updateMicroFocus() { callbackQAbstractPrintDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQAbstractPrintDialog_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQAbstractPrintDialog_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtPrintSupport_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQAbstractPrintDialog_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQAbstractPrintDialog_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQAbstractPrintDialog_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQAbstractPrintDialog_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQAbstractPrintDialog_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQAbstractPrintDialog_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void childEvent(QChildEvent * event) { callbackQAbstractPrintDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractPrintDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractPrintDialog_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractPrintDialog_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPrintSupport_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractPrintDialog_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractPrintDialog_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractPrintDialog_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAbstractPrintDialog*)

int QAbstractPrintDialog_QAbstractPrintDialog_QRegisterMetaType(){qRegisterMetaType<QAbstractPrintDialog*>(); return qRegisterMetaType<MyQAbstractPrintDialog*>();}

void* QAbstractPrintDialog_NewQAbstractPrintDialog(void* printer, void* parent)
{
		return new MyQAbstractPrintDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent));
}

int QAbstractPrintDialog_Exec(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->exec();
}

void QAbstractPrintDialog_SetFromTo(void* ptr, int from, int to)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setFromTo(from, to);
}

void QAbstractPrintDialog_SetMinMax(void* ptr, int min, int max)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setMinMax(min, max);
}

void QAbstractPrintDialog_SetOptionTabs(void* ptr, void* tabs)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setOptionTabs(*static_cast<QList<QWidget *>*>(tabs));
}

void QAbstractPrintDialog_SetPrintRange(void* ptr, long long ran)
{
	static_cast<QAbstractPrintDialog*>(ptr)->setPrintRange(static_cast<QAbstractPrintDialog::PrintRange>(ran));
}

long long QAbstractPrintDialog_PrintRange(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->printRange();
}

void* QAbstractPrintDialog_Printer(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->printer();
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

int QAbstractPrintDialog_ToPage(void* ptr)
{
	return static_cast<QAbstractPrintDialog*>(ptr)->toPage();
}

void* QAbstractPrintDialog___setOptionTabs_tabs_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void QAbstractPrintDialog___setOptionTabs_tabs_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QAbstractPrintDialog___setOptionTabs_tabs_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>;
}

void* QAbstractPrintDialog___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QAbstractPrintDialog___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QAbstractPrintDialog___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QAbstractPrintDialog___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QAbstractPrintDialog___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QAbstractPrintDialog___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QAbstractPrintDialog___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QAbstractPrintDialog___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QAbstractPrintDialog___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QAbstractPrintDialog___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QAbstractPrintDialog___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractPrintDialog___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QAbstractPrintDialog___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractPrintDialog___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractPrintDialog___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractPrintDialog___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractPrintDialog___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractPrintDialog___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractPrintDialog___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QAbstractPrintDialog___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractPrintDialog___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QAbstractPrintDialog___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QAbstractPrintDialog___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractPrintDialog___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QAbstractPrintDialog_EventFilterDefault(void* ptr, void* o, void* e)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
	}
}

void QAbstractPrintDialog_AcceptDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::accept();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::accept();
	}
}

void QAbstractPrintDialog_CloseEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::closeEvent(static_cast<QCloseEvent*>(e));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::closeEvent(static_cast<QCloseEvent*>(e));
	}
}

void QAbstractPrintDialog_ContextMenuEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
	}
}

void QAbstractPrintDialog_DoneDefault(void* ptr, int r)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::done(r);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::done(r);
	}
}

void QAbstractPrintDialog_KeyPressEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
	}
}

void QAbstractPrintDialog_OpenDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::open();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::open();
	}
}

void QAbstractPrintDialog_RejectDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::reject();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::reject();
	}
}

void QAbstractPrintDialog_ResizeEventDefault(void* ptr, void* vqr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
	}
}

void QAbstractPrintDialog_SetVisibleDefault(void* ptr, char visible)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setVisible(visible != 0);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setVisible(visible != 0);
	}
}

void QAbstractPrintDialog_ShowEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::showEvent(static_cast<QShowEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showEvent(static_cast<QShowEvent*>(event));
	}
}

void* QAbstractPrintDialog_MinimumSizeHintDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->QPrintDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QAbstractPrintDialog_SizeHintDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QPrintDialog*>(ptr)->QPrintDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

char QAbstractPrintDialog_CloseDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::close();
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::close();
	}
}

char QAbstractPrintDialog_EventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::event(static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::event(static_cast<QEvent*>(event));
	}
}

char QAbstractPrintDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusNextPrevChild(next != 0);
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusNextPrevChild(next != 0);
	}
}

char QAbstractPrintDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
	}
}

void QAbstractPrintDialog_ActionEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::actionEvent(static_cast<QActionEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::actionEvent(static_cast<QActionEvent*>(event));
	}
}

void QAbstractPrintDialog_ChangeEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::changeEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::changeEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractPrintDialog_DragEnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
	}
}

void QAbstractPrintDialog_DragLeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
	}
}

void QAbstractPrintDialog_DragMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
	}
}

void QAbstractPrintDialog_DropEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::dropEvent(static_cast<QDropEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::dropEvent(static_cast<QDropEvent*>(event));
	}
}

void QAbstractPrintDialog_EnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::enterEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::enterEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractPrintDialog_FocusInEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusInEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusInEvent(static_cast<QFocusEvent*>(event));
	}
}

void QAbstractPrintDialog_FocusOutEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
	}
}

void QAbstractPrintDialog_HideDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::hide();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hide();
	}
}

void QAbstractPrintDialog_HideEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hideEvent(static_cast<QHideEvent*>(event));
	}
}

void QAbstractPrintDialog_InputMethodEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	}
}

void QAbstractPrintDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void QAbstractPrintDialog_LeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::leaveEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::leaveEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractPrintDialog_LowerDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::lower();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::lower();
	}
}

void QAbstractPrintDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	}
}

void QAbstractPrintDialog_MouseMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	}
}

void QAbstractPrintDialog_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	}
}

void QAbstractPrintDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	}
}

void QAbstractPrintDialog_MoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::moveEvent(static_cast<QMoveEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::moveEvent(static_cast<QMoveEvent*>(event));
	}
}

void QAbstractPrintDialog_PaintEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::paintEvent(static_cast<QPaintEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::paintEvent(static_cast<QPaintEvent*>(event));
	}
}

void QAbstractPrintDialog_RaiseDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::raise();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::raise();
	}
}

void QAbstractPrintDialog_RepaintDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::repaint();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::repaint();
	}
}

void QAbstractPrintDialog_SetDisabledDefault(void* ptr, char disable)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setDisabled(disable != 0);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setDisabled(disable != 0);
	}
}

void QAbstractPrintDialog_SetEnabledDefault(void* ptr, char vbo)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setEnabled(vbo != 0);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setEnabled(vbo != 0);
	}
}

void QAbstractPrintDialog_SetFocus2Default(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setFocus();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setFocus();
	}
}

void QAbstractPrintDialog_SetHiddenDefault(void* ptr, char hidden)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setHidden(hidden != 0);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setHidden(hidden != 0);
	}
}

void QAbstractPrintDialog_SetStyleSheetDefault(void* ptr, struct QtPrintSupport_PackedString styleSheet)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	}
}

void QAbstractPrintDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setWindowModified(vbo != 0);
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setWindowModified(vbo != 0);
	}
}

void QAbstractPrintDialog_SetWindowTitleDefault(void* ptr, struct QtPrintSupport_PackedString vqs)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	}
}

void QAbstractPrintDialog_ShowDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::show();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::show();
	}
}

void QAbstractPrintDialog_ShowFullScreenDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::showFullScreen();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showFullScreen();
	}
}

void QAbstractPrintDialog_ShowMaximizedDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::showMaximized();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showMaximized();
	}
}

void QAbstractPrintDialog_ShowMinimizedDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::showMinimized();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showMinimized();
	}
}

void QAbstractPrintDialog_ShowNormalDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::showNormal();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::showNormal();
	}
}

void QAbstractPrintDialog_TabletEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::tabletEvent(static_cast<QTabletEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::tabletEvent(static_cast<QTabletEvent*>(event));
	}
}

void QAbstractPrintDialog_UpdateDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::update();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::update();
	}
}

void QAbstractPrintDialog_UpdateMicroFocusDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::updateMicroFocus();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::updateMicroFocus();
	}
}

void QAbstractPrintDialog_WheelEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::wheelEvent(static_cast<QWheelEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::wheelEvent(static_cast<QWheelEvent*>(event));
	}
}

void* QAbstractPrintDialog_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::paintEngine();
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::paintEngine();
	}
}

void* QAbstractPrintDialog_InputMethodQueryDefault(void* ptr, long long query)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QPrintDialog*>(ptr)->QPrintDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	} else {
		return new QVariant(static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	}
}

char QAbstractPrintDialog_HasHeightForWidthDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::hasHeightForWidth();
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::hasHeightForWidth();
	}
}

int QAbstractPrintDialog_HeightForWidthDefault(void* ptr, int w)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::heightForWidth(w);
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::heightForWidth(w);
	}
}

int QAbstractPrintDialog_MetricDefault(void* ptr, long long m)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPrintDialog*>(ptr)->QPrintDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else {
		return static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	}
}

void QAbstractPrintDialog_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QAbstractPrintDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractPrintDialog_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::customEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractPrintDialog_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::deleteLater();
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::deleteLater();
	}
}

void QAbstractPrintDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractPrintDialog_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QAbstractPrintDialog_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QPrintDialog*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPrintDialog*>(ptr)->QPrintDialog::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QAbstractPrintDialog*>(ptr)->QAbstractPrintDialog::metaObject());
	}
}

class MyQPageSetupDialog: public QPageSetupDialog
{
public:
	MyQPageSetupDialog(QPrinter *printer, QWidget *parent = Q_NULLPTR) : QPageSetupDialog(printer, parent) {QPageSetupDialog_QPageSetupDialog_QRegisterMetaType();};
	MyQPageSetupDialog(QWidget *parent = Q_NULLPTR) : QPageSetupDialog(parent) {QPageSetupDialog_QPageSetupDialog_QRegisterMetaType();};
	int exec() { return callbackQPageSetupDialog_Exec(this); };
	void done(int result) { callbackQPageSetupDialog_Done(this, result); };
	void setVisible(bool visible) { callbackQPageSetupDialog_SetVisible(this, visible); };
	void accept() { callbackQPageSetupDialog_Accept(this); };
	void Signal_Accepted() { callbackQPageSetupDialog_Accepted(this); };
	void closeEvent(QCloseEvent * e) { callbackQPageSetupDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQPageSetupDialog_ContextMenuEvent(this, e); };
	void Signal_Finished(int result) { callbackQPageSetupDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQPageSetupDialog_KeyPressEvent(this, e); };
	
	void reject() { callbackQPageSetupDialog_Reject(this); };
	void Signal_Rejected() { callbackQPageSetupDialog_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQPageSetupDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQPageSetupDialog_ShowEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPageSetupDialog_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPageSetupDialog_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQPageSetupDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPageSetupDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPageSetupDialog_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPageSetupDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQPageSetupDialog_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQPageSetupDialog_ChangeEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQPageSetupDialog_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPageSetupDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPageSetupDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPageSetupDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPageSetupDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPageSetupDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPageSetupDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPageSetupDialog_FocusOutEvent(this, event); };
	void hide() { callbackQPageSetupDialog_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQPageSetupDialog_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPageSetupDialog_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPageSetupDialog_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPageSetupDialog_LeaveEvent(this, event); };
	void lower() { callbackQPageSetupDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPageSetupDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPageSetupDialog_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQPageSetupDialog_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQPageSetupDialog_PaintEvent(this, event); };
	void raise() { callbackQPageSetupDialog_Raise(this); };
	void repaint() { callbackQPageSetupDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQPageSetupDialog_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQPageSetupDialog_SetEnabled(this, vbo); };
	void setFocus() { callbackQPageSetupDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPageSetupDialog_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPageSetupDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPageSetupDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPageSetupDialog_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQPageSetupDialog_Show(this); };
	void showFullScreen() { callbackQPageSetupDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQPageSetupDialog_ShowMaximized(this); };
	void showMinimized() { callbackQPageSetupDialog_ShowMinimized(this); };
	void showNormal() { callbackQPageSetupDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPageSetupDialog_TabletEvent(this, event); };
	void update() { callbackQPageSetupDialog_Update(this); };
	void updateMicroFocus() { callbackQPageSetupDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPageSetupDialog_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQPageSetupDialog_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtPrintSupport_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQPageSetupDialog_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPageSetupDialog_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPageSetupDialog_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQPageSetupDialog_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQPageSetupDialog_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPageSetupDialog_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void childEvent(QChildEvent * event) { callbackQPageSetupDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPageSetupDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPageSetupDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQPageSetupDialog_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQPageSetupDialog_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPageSetupDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPrintSupport_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQPageSetupDialog_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQPageSetupDialog_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPageSetupDialog_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPageSetupDialog*)

int QPageSetupDialog_QPageSetupDialog_QRegisterMetaType(){qRegisterMetaType<QPageSetupDialog*>(); return qRegisterMetaType<MyQPageSetupDialog*>();}

void* QPageSetupDialog_NewQPageSetupDialog(void* printer, void* parent)
{
		return new MyQPageSetupDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent));
}

void* QPageSetupDialog_NewQPageSetupDialog2(void* parent)
{
		return new MyQPageSetupDialog(static_cast<QWidget*>(parent));
}

void* QPageSetupDialog_Printer(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->printer();
}

int QPageSetupDialog_Exec(void* ptr)
{
	return static_cast<QPageSetupDialog*>(ptr)->exec();
}

int QPageSetupDialog_ExecDefault(void* ptr)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::exec();
}

void QPageSetupDialog_Done(void* ptr, int result)
{
	static_cast<QPageSetupDialog*>(ptr)->done(result);
}

void QPageSetupDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::done(result);
}

void QPageSetupDialog_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setVisible(visible != 0);
}

void QPageSetupDialog_DestroyQPageSetupDialog(void* ptr)
{
	static_cast<QPageSetupDialog*>(ptr)->~QPageSetupDialog();
}

void* QPageSetupDialog___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPageSetupDialog___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPageSetupDialog___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPageSetupDialog___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPageSetupDialog___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPageSetupDialog___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPageSetupDialog___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPageSetupDialog___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPageSetupDialog___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPageSetupDialog___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QPageSetupDialog___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QPageSetupDialog___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QPageSetupDialog___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPageSetupDialog___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPageSetupDialog___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPageSetupDialog___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPageSetupDialog___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPageSetupDialog___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPageSetupDialog___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPageSetupDialog___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPageSetupDialog___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPageSetupDialog___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QPageSetupDialog___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPageSetupDialog___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QPageSetupDialog_EventFilterDefault(void* ptr, void* o, void* e)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
}

void QPageSetupDialog_AcceptDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::accept();
}

void QPageSetupDialog_CloseEventDefault(void* ptr, void* e)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QPageSetupDialog_ContextMenuEventDefault(void* ptr, void* e)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPageSetupDialog_KeyPressEventDefault(void* ptr, void* e)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}



void QPageSetupDialog_RejectDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::reject();
}

void QPageSetupDialog_ResizeEventDefault(void* ptr, void* vqr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPageSetupDialog_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QPageSetupDialog_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPageSetupDialog_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QPageSetupDialog_CloseDefault(void* ptr)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::close();
}

char QPageSetupDialog_EventDefault(void* ptr, void* event)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::event(static_cast<QEvent*>(event));
}

char QPageSetupDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusNextPrevChild(next != 0);
}

char QPageSetupDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPageSetupDialog_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QPageSetupDialog_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::changeEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPageSetupDialog_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPageSetupDialog_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPageSetupDialog_DropEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QPageSetupDialog_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::enterEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPageSetupDialog_HideDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hide();
}

void QPageSetupDialog_HideEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QPageSetupDialog_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPageSetupDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPageSetupDialog_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::leaveEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_LowerDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::lower();
}

void QPageSetupDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPageSetupDialog_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void QPageSetupDialog_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPageSetupDialog_RaiseDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::raise();
}

void QPageSetupDialog_RepaintDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::repaint();
}

void QPageSetupDialog_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setDisabled(disable != 0);
}

void QPageSetupDialog_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setEnabled(vbo != 0);
}

void QPageSetupDialog_SetFocus2Default(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setFocus();
}

void QPageSetupDialog_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setHidden(hidden != 0);
}

void QPageSetupDialog_SetStyleSheetDefault(void* ptr, struct QtPrintSupport_PackedString styleSheet)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QPageSetupDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setWindowModified(vbo != 0);
}

void QPageSetupDialog_SetWindowTitleDefault(void* ptr, struct QtPrintSupport_PackedString vqs)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QPageSetupDialog_ShowDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::show();
}

void QPageSetupDialog_ShowFullScreenDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showFullScreen();
}

void QPageSetupDialog_ShowMaximizedDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showMaximized();
}

void QPageSetupDialog_ShowMinimizedDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showMinimized();
}

void QPageSetupDialog_ShowNormalDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::showNormal();
}

void QPageSetupDialog_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPageSetupDialog_UpdateDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::update();
}

void QPageSetupDialog_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::updateMicroFocus();
}

void QPageSetupDialog_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QPageSetupDialog_PaintEngineDefault(void* ptr)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::paintEngine();
}

void* QPageSetupDialog_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QPageSetupDialog_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::hasHeightForWidth();
}

int QPageSetupDialog_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::heightForWidth(w);
}

int QPageSetupDialog_MetricDefault(void* ptr, long long m)
{
		return static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QPageSetupDialog_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QPageSetupDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPageSetupDialog_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::customEvent(static_cast<QEvent*>(event));
}

void QPageSetupDialog_DeleteLaterDefault(void* ptr)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::deleteLater();
}

void QPageSetupDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPageSetupDialog_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QPageSetupDialog_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QPageSetupDialog*>(ptr)->QPageSetupDialog::metaObject());
}

class MyQPrintDialog: public QPrintDialog
{
public:
	MyQPrintDialog(QPrinter *printer, QWidget *parent = Q_NULLPTR) : QPrintDialog(printer, parent) {QPrintDialog_QPrintDialog_QRegisterMetaType();};
	MyQPrintDialog(QWidget *parent = Q_NULLPTR) : QPrintDialog(parent) {QPrintDialog_QPrintDialog_QRegisterMetaType();};
	void done(int result) { callbackQPrintDialog_Done(this, result); };
	int exec() { return callbackQPrintDialog_Exec(this); };
	void Signal_Accepted(QPrinter * printer) { callbackQPrintDialog_Accepted(this, printer); };
	void setVisible(bool visible) { callbackQAbstractPrintDialog_SetVisible(this, visible); };
	void accept() { callbackQAbstractPrintDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQAbstractPrintDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQAbstractPrintDialog_ContextMenuEvent(this, e); };
	void Signal_Finished(int result) { callbackQAbstractPrintDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQAbstractPrintDialog_KeyPressEvent(this, e); };
	void reject() { callbackQAbstractPrintDialog_Reject(this); };
	void Signal_Rejected() { callbackQAbstractPrintDialog_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQAbstractPrintDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQAbstractPrintDialog_ShowEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQAbstractPrintDialog_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQAbstractPrintDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQAbstractPrintDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQAbstractPrintDialog_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQAbstractPrintDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQAbstractPrintDialog_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQAbstractPrintDialog_ChangeEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQAbstractPrintDialog_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQAbstractPrintDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQAbstractPrintDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQAbstractPrintDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQAbstractPrintDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQAbstractPrintDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQAbstractPrintDialog_FocusOutEvent(this, event); };
	void hide() { callbackQAbstractPrintDialog_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQAbstractPrintDialog_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQAbstractPrintDialog_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQAbstractPrintDialog_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQAbstractPrintDialog_LeaveEvent(this, event); };
	void lower() { callbackQAbstractPrintDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQAbstractPrintDialog_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQAbstractPrintDialog_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQAbstractPrintDialog_PaintEvent(this, event); };
	void raise() { callbackQAbstractPrintDialog_Raise(this); };
	void repaint() { callbackQAbstractPrintDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQAbstractPrintDialog_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQAbstractPrintDialog_SetEnabled(this, vbo); };
	void setFocus() { callbackQAbstractPrintDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQAbstractPrintDialog_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQAbstractPrintDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQAbstractPrintDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQAbstractPrintDialog_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQAbstractPrintDialog_Show(this); };
	void showFullScreen() { callbackQAbstractPrintDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQAbstractPrintDialog_ShowMaximized(this); };
	void showMinimized() { callbackQAbstractPrintDialog_ShowMinimized(this); };
	void showNormal() { callbackQAbstractPrintDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQAbstractPrintDialog_TabletEvent(this, event); };
	void update() { callbackQAbstractPrintDialog_Update(this); };
	void updateMicroFocus() { callbackQAbstractPrintDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQAbstractPrintDialog_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQAbstractPrintDialog_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtPrintSupport_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQAbstractPrintDialog_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQAbstractPrintDialog_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQAbstractPrintDialog_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQAbstractPrintDialog_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQAbstractPrintDialog_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQAbstractPrintDialog_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void childEvent(QChildEvent * event) { callbackQAbstractPrintDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractPrintDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractPrintDialog_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractPrintDialog_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractPrintDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPrintSupport_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractPrintDialog_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractPrintDialog_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractPrintDialog_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPrintDialog*)

int QPrintDialog_QPrintDialog_QRegisterMetaType(){qRegisterMetaType<QPrintDialog*>(); return qRegisterMetaType<MyQPrintDialog*>();}

void QPrintDialog_Done(void* ptr, int result)
{
	static_cast<QPrintDialog*>(ptr)->done(result);
}

void QPrintDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QPrintDialog*>(ptr)->QPrintDialog::done(result);
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

void QPrintDialog_Open(void* ptr, void* receiver, char* member)
{
	static_cast<QPrintDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QPrintDialog_SetOption(void* ptr, long long option, char on)
{
	static_cast<QPrintDialog*>(ptr)->setOption(static_cast<QAbstractPrintDialog::PrintDialogOption>(option), on != 0);
}

void QPrintDialog_SetOptions(void* ptr, long long options)
{
	static_cast<QPrintDialog*>(ptr)->setOptions(static_cast<QAbstractPrintDialog::PrintDialogOption>(options));
}

void QPrintDialog_DestroyQPrintDialog(void* ptr)
{
	static_cast<QPrintDialog*>(ptr)->~QPrintDialog();
}

long long QPrintDialog_Options(void* ptr)
{
	return static_cast<QPrintDialog*>(ptr)->options();
}

char QPrintDialog_TestOption(void* ptr, long long option)
{
	return static_cast<QPrintDialog*>(ptr)->testOption(static_cast<QAbstractPrintDialog::PrintDialogOption>(option));
}

class MyQPrintEngine: public QPrintEngine
{
public:
	bool abort() { return callbackQPrintEngine_Abort(this) != 0; };
	bool newPage() { return callbackQPrintEngine_NewPage(this) != 0; };
	void setProperty(QPrintEngine::PrintEnginePropertyKey key, const QVariant & value) { callbackQPrintEngine_SetProperty(this, key, const_cast<QVariant*>(&value)); };
	 ~MyQPrintEngine() { callbackQPrintEngine_DestroyQPrintEngine(this); };
	QPrinter::PrinterState printerState() const { return static_cast<QPrinter::PrinterState>(callbackQPrintEngine_PrinterState(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant property(QPrintEngine::PrintEnginePropertyKey key) const { return *static_cast<QVariant*>(callbackQPrintEngine_Property(const_cast<void*>(static_cast<const void*>(this)), key)); };
	int metric(QPaintDevice::PaintDeviceMetric id) const { return callbackQPrintEngine_Metric(const_cast<void*>(static_cast<const void*>(this)), id); };
};

char QPrintEngine_Abort(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->abort();
}

char QPrintEngine_NewPage(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->newPage();
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
	Q_UNUSED(ptr);

}

long long QPrintEngine_PrinterState(void* ptr)
{
	return static_cast<QPrintEngine*>(ptr)->printerState();
}

void* QPrintEngine_Property(void* ptr, long long key)
{
	return new QVariant(static_cast<QPrintEngine*>(ptr)->property(static_cast<QPrintEngine::PrintEnginePropertyKey>(key)));
}

int QPrintEngine_Metric(void* ptr, long long id)
{
	return static_cast<QPrintEngine*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(id));
}

class MyQPrintPreviewDialog: public QPrintPreviewDialog
{
public:
	MyQPrintPreviewDialog(QPrinter *printer, QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QPrintPreviewDialog(printer, parent, flags) {QPrintPreviewDialog_QPrintPreviewDialog_QRegisterMetaType();};
	MyQPrintPreviewDialog(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QPrintPreviewDialog(parent, flags) {QPrintPreviewDialog_QPrintPreviewDialog_QRegisterMetaType();};
	void done(int result) { callbackQPrintPreviewDialog_Done(this, result); };
	void Signal_PaintRequested(QPrinter * printer) { callbackQPrintPreviewDialog_PaintRequested(this, printer); };
	void setVisible(bool visible) { callbackQPrintPreviewDialog_SetVisible(this, visible); };
	int exec() { return callbackQPrintPreviewDialog_Exec(this); };
	void accept() { callbackQPrintPreviewDialog_Accept(this); };
	void Signal_Accepted() { callbackQPrintPreviewDialog_Accepted(this); };
	void closeEvent(QCloseEvent * e) { callbackQPrintPreviewDialog_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQPrintPreviewDialog_ContextMenuEvent(this, e); };
	void Signal_Finished(int result) { callbackQPrintPreviewDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQPrintPreviewDialog_KeyPressEvent(this, e); };
	
	void reject() { callbackQPrintPreviewDialog_Reject(this); };
	void Signal_Rejected() { callbackQPrintPreviewDialog_Rejected(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQPrintPreviewDialog_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQPrintPreviewDialog_ShowEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewDialog_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewDialog_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQPrintPreviewDialog_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPrintPreviewDialog_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPrintPreviewDialog_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPrintPreviewDialog_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQPrintPreviewDialog_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQPrintPreviewDialog_ChangeEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQPrintPreviewDialog_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPrintPreviewDialog_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPrintPreviewDialog_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPrintPreviewDialog_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPrintPreviewDialog_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPrintPreviewDialog_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPrintPreviewDialog_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPrintPreviewDialog_FocusOutEvent(this, event); };
	void hide() { callbackQPrintPreviewDialog_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQPrintPreviewDialog_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPrintPreviewDialog_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPrintPreviewDialog_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPrintPreviewDialog_LeaveEvent(this, event); };
	void lower() { callbackQPrintPreviewDialog_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPrintPreviewDialog_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQPrintPreviewDialog_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQPrintPreviewDialog_PaintEvent(this, event); };
	void raise() { callbackQPrintPreviewDialog_Raise(this); };
	void repaint() { callbackQPrintPreviewDialog_Repaint(this); };
	void setDisabled(bool disable) { callbackQPrintPreviewDialog_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQPrintPreviewDialog_SetEnabled(this, vbo); };
	void setFocus() { callbackQPrintPreviewDialog_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPrintPreviewDialog_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPrintPreviewDialog_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPrintPreviewDialog_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPrintPreviewDialog_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQPrintPreviewDialog_Show(this); };
	void showFullScreen() { callbackQPrintPreviewDialog_ShowFullScreen(this); };
	void showMaximized() { callbackQPrintPreviewDialog_ShowMaximized(this); };
	void showMinimized() { callbackQPrintPreviewDialog_ShowMinimized(this); };
	void showNormal() { callbackQPrintPreviewDialog_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPrintPreviewDialog_TabletEvent(this, event); };
	void update() { callbackQPrintPreviewDialog_Update(this); };
	void updateMicroFocus() { callbackQPrintPreviewDialog_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPrintPreviewDialog_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQPrintPreviewDialog_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtPrintSupport_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQPrintPreviewDialog_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrintPreviewDialog_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPrintPreviewDialog_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQPrintPreviewDialog_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQPrintPreviewDialog_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPrintPreviewDialog_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void childEvent(QChildEvent * event) { callbackQPrintPreviewDialog_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPrintPreviewDialog_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPrintPreviewDialog_CustomEvent(this, event); };
	void deleteLater() { callbackQPrintPreviewDialog_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQPrintPreviewDialog_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPrintPreviewDialog_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPrintSupport_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQPrintPreviewDialog_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQPrintPreviewDialog_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPrintPreviewDialog_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPrintPreviewDialog*)

int QPrintPreviewDialog_QPrintPreviewDialog_QRegisterMetaType(){qRegisterMetaType<QPrintPreviewDialog*>(); return qRegisterMetaType<MyQPrintPreviewDialog*>();}

void* QPrintPreviewDialog_NewQPrintPreviewDialog(void* printer, void* parent, long long flags)
{
		return new MyQPrintPreviewDialog(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QPrintPreviewDialog_NewQPrintPreviewDialog2(void* parent, long long flags)
{
		return new MyQPrintPreviewDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QPrintPreviewDialog_Printer(void* ptr)
{
	return static_cast<QPrintPreviewDialog*>(ptr)->printer();
}

void QPrintPreviewDialog_Done(void* ptr, int result)
{
	static_cast<QPrintPreviewDialog*>(ptr)->done(result);
}

void QPrintPreviewDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::done(result);
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

void QPrintPreviewDialog_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setVisible(visible != 0);
}

void QPrintPreviewDialog_DestroyQPrintPreviewDialog(void* ptr)
{
	static_cast<QPrintPreviewDialog*>(ptr)->~QPrintPreviewDialog();
}

void* QPrintPreviewDialog___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewDialog___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewDialog___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewDialog___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewDialog___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewDialog___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewDialog___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewDialog___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewDialog___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewDialog___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QPrintPreviewDialog___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QPrintPreviewDialog___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QPrintPreviewDialog___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewDialog___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewDialog___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewDialog___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewDialog___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewDialog___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewDialog___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewDialog___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewDialog___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewDialog___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QPrintPreviewDialog___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewDialog___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QPrintPreviewDialog_EventFilterDefault(void* ptr, void* o, void* e)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::eventFilter(static_cast<QObject*>(o), static_cast<QEvent*>(e));
}

int QPrintPreviewDialog_ExecDefault(void* ptr)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::exec();
}

void QPrintPreviewDialog_AcceptDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::accept();
}

void QPrintPreviewDialog_CloseEventDefault(void* ptr, void* e)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::closeEvent(static_cast<QCloseEvent*>(e));
}

void QPrintPreviewDialog_ContextMenuEventDefault(void* ptr, void* e)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QPrintPreviewDialog_KeyPressEventDefault(void* ptr, void* e)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::keyPressEvent(static_cast<QKeyEvent*>(e));
}



void QPrintPreviewDialog_RejectDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::reject();
}

void QPrintPreviewDialog_ResizeEventDefault(void* ptr, void* vqr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::resizeEvent(static_cast<QResizeEvent*>(vqr));
}

void QPrintPreviewDialog_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showEvent(static_cast<QShowEvent*>(event));
}

void* QPrintPreviewDialog_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewDialog_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QPrintPreviewDialog_CloseDefault(void* ptr)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::close();
}

char QPrintPreviewDialog_EventDefault(void* ptr, void* event)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::event(static_cast<QEvent*>(event));
}

char QPrintPreviewDialog_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusNextPrevChild(next != 0);
}

char QPrintPreviewDialog_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPrintPreviewDialog_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewDialog_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::changeEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewDialog_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewDialog_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewDialog_DropEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewDialog_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewDialog_HideDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hide();
}

void QPrintPreviewDialog_HideEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewDialog_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPrintPreviewDialog_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewDialog_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::leaveEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_LowerDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::lower();
}

void QPrintPreviewDialog_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewDialog_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::moveEvent(static_cast<QMoveEvent*>(event));
}

void QPrintPreviewDialog_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewDialog_RaiseDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::raise();
}

void QPrintPreviewDialog_RepaintDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::repaint();
}

void QPrintPreviewDialog_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setDisabled(disable != 0);
}

void QPrintPreviewDialog_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setEnabled(vbo != 0);
}

void QPrintPreviewDialog_SetFocus2Default(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setFocus();
}

void QPrintPreviewDialog_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setHidden(hidden != 0);
}

void QPrintPreviewDialog_SetStyleSheetDefault(void* ptr, struct QtPrintSupport_PackedString styleSheet)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QPrintPreviewDialog_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setWindowModified(vbo != 0);
}

void QPrintPreviewDialog_SetWindowTitleDefault(void* ptr, struct QtPrintSupport_PackedString vqs)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QPrintPreviewDialog_ShowDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::show();
}

void QPrintPreviewDialog_ShowFullScreenDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showFullScreen();
}

void QPrintPreviewDialog_ShowMaximizedDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showMaximized();
}

void QPrintPreviewDialog_ShowMinimizedDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showMinimized();
}

void QPrintPreviewDialog_ShowNormalDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::showNormal();
}

void QPrintPreviewDialog_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewDialog_UpdateDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::update();
}

void QPrintPreviewDialog_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::updateMicroFocus();
}

void QPrintPreviewDialog_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QPrintPreviewDialog_PaintEngineDefault(void* ptr)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::paintEngine();
}

void* QPrintPreviewDialog_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QPrintPreviewDialog_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::hasHeightForWidth();
}

int QPrintPreviewDialog_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::heightForWidth(w);
}

int QPrintPreviewDialog_MetricDefault(void* ptr, long long m)
{
		return static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QPrintPreviewDialog_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewDialog_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewDialog_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewDialog_DeleteLaterDefault(void* ptr)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::deleteLater();
}

void QPrintPreviewDialog_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewDialog_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QPrintPreviewDialog_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QPrintPreviewDialog*>(ptr)->QPrintPreviewDialog::metaObject());
}

class MyQPrintPreviewWidget: public QPrintPreviewWidget
{
public:
	MyQPrintPreviewWidget(QPrinter *printer, QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QPrintPreviewWidget(printer, parent, flags) {QPrintPreviewWidget_QPrintPreviewWidget_QRegisterMetaType();};
	MyQPrintPreviewWidget(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QPrintPreviewWidget(parent, flags) {QPrintPreviewWidget_QPrintPreviewWidget_QRegisterMetaType();};
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
	bool close() { return callbackQPrintPreviewWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQPrintPreviewWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQPrintPreviewWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPrintPreviewWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQPrintPreviewWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQPrintPreviewWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQPrintPreviewWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQPrintPreviewWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQPrintPreviewWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQPrintPreviewWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQPrintPreviewWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQPrintPreviewWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQPrintPreviewWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQPrintPreviewWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQPrintPreviewWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQPrintPreviewWidget_FocusOutEvent(this, event); };
	void hide() { callbackQPrintPreviewWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQPrintPreviewWidget_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQPrintPreviewWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQPrintPreviewWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQPrintPreviewWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQPrintPreviewWidget_LeaveEvent(this, event); };
	void lower() { callbackQPrintPreviewWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQPrintPreviewWidget_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQPrintPreviewWidget_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQPrintPreviewWidget_PaintEvent(this, event); };
	void raise() { callbackQPrintPreviewWidget_Raise(this); };
	void repaint() { callbackQPrintPreviewWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQPrintPreviewWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQPrintPreviewWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQPrintPreviewWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQPrintPreviewWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQPrintPreviewWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtPrintSupport_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQPrintPreviewWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowModified(bool vbo) { callbackQPrintPreviewWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtPrintSupport_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQPrintPreviewWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQPrintPreviewWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQPrintPreviewWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQPrintPreviewWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQPrintPreviewWidget_ShowMaximized(this); };
	void showMinimized() { callbackQPrintPreviewWidget_ShowMinimized(this); };
	void showNormal() { callbackQPrintPreviewWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQPrintPreviewWidget_TabletEvent(this, event); };
	void update() { callbackQPrintPreviewWidget_Update(this); };
	void updateMicroFocus() { callbackQPrintPreviewWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQPrintPreviewWidget_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQPrintPreviewWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtPrintSupport_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQPrintPreviewWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrintPreviewWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQPrintPreviewWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQPrintPreviewWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQPrintPreviewWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQPrintPreviewWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQPrintPreviewWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPrintPreviewWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQPrintPreviewWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPrintPreviewWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPrintPreviewWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQPrintPreviewWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQPrintPreviewWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPrintPreviewWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPrintSupport_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQPrintPreviewWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQPrintPreviewWidget_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPrintPreviewWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPrintPreviewWidget*)

int QPrintPreviewWidget_QPrintPreviewWidget_QRegisterMetaType(){qRegisterMetaType<QPrintPreviewWidget*>(); return qRegisterMetaType<MyQPrintPreviewWidget*>();}

void* QPrintPreviewWidget_NewQPrintPreviewWidget(void* printer, void* parent, long long flags)
{
		return new MyQPrintPreviewWidget(static_cast<QPrinter*>(printer), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QPrintPreviewWidget_NewQPrintPreviewWidget2(void* parent, long long flags)
{
		return new MyQPrintPreviewWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
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

void QPrintPreviewWidget_ZoomIn(void* ptr, double factor)
{
	QMetaObject::invokeMethod(static_cast<QPrintPreviewWidget*>(ptr), "zoomIn", Q_ARG(qreal, factor));
}

void QPrintPreviewWidget_ZoomInDefault(void* ptr, double factor)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::zoomIn(factor);
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

long long QPrintPreviewWidget_Orientation(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->orientation();
}

long long QPrintPreviewWidget_ViewMode(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->viewMode();
}

long long QPrintPreviewWidget_ZoomMode(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->zoomMode();
}

int QPrintPreviewWidget_CurrentPage(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->currentPage();
}

int QPrintPreviewWidget_PageCount(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->pageCount();
}

double QPrintPreviewWidget_ZoomFactor(void* ptr)
{
	return static_cast<QPrintPreviewWidget*>(ptr)->zoomFactor();
}

void* QPrintPreviewWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QPrintPreviewWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QPrintPreviewWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QPrintPreviewWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QPrintPreviewWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QPrintPreviewWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QPrintPreviewWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QPrintPreviewWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QPrintPreviewWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QPrintPreviewWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPrintPreviewWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QPrintPreviewWidget_CloseDefault(void* ptr)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::close();
}

char QPrintPreviewWidget_EventDefault(void* ptr, void* event)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::event(static_cast<QEvent*>(event));
}

char QPrintPreviewWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusNextPrevChild(next != 0);
}

char QPrintPreviewWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QPrintPreviewWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QPrintPreviewWidget_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::changeEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QPrintPreviewWidget_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QPrintPreviewWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QPrintPreviewWidget_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QPrintPreviewWidget_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QPrintPreviewWidget_DropEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QPrintPreviewWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::enterEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QPrintPreviewWidget_HideDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hide();
}

void QPrintPreviewWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QPrintPreviewWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QPrintPreviewWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QPrintPreviewWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_LowerDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::lower();
}

void QPrintPreviewWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QPrintPreviewWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QPrintPreviewWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPrintPreviewWidget_RaiseDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::raise();
}

void QPrintPreviewWidget_RepaintDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::repaint();
}

void QPrintPreviewWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QPrintPreviewWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setDisabled(disable != 0);
}

void QPrintPreviewWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setEnabled(vbo != 0);
}

void QPrintPreviewWidget_SetFocus2Default(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setFocus();
}

void QPrintPreviewWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setHidden(hidden != 0);
}

void QPrintPreviewWidget_SetStyleSheetDefault(void* ptr, struct QtPrintSupport_PackedString styleSheet)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QPrintPreviewWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setWindowModified(vbo != 0);
}

void QPrintPreviewWidget_SetWindowTitleDefault(void* ptr, struct QtPrintSupport_PackedString vqs)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QPrintPreviewWidget_ShowDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::show();
}

void QPrintPreviewWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QPrintPreviewWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showFullScreen();
}

void QPrintPreviewWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showMaximized();
}

void QPrintPreviewWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showMinimized();
}

void QPrintPreviewWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::showNormal();
}

void QPrintPreviewWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QPrintPreviewWidget_UpdateDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::update();
}

void QPrintPreviewWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::updateMicroFocus();
}

void QPrintPreviewWidget_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QPrintPreviewWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::paintEngine();
}

void* QPrintPreviewWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPrintPreviewWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QPrintPreviewWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::hasHeightForWidth();
}

int QPrintPreviewWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::heightForWidth(w);
}

int QPrintPreviewWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QPrintPreviewWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QPrintPreviewWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::customEvent(static_cast<QEvent*>(event));
}

void QPrintPreviewWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::deleteLater();
}

void QPrintPreviewWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPrintPreviewWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QPrintPreviewWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QPrintPreviewWidget*>(ptr)->QPrintPreviewWidget::metaObject());
}

class MyQPrinter: public QPrinter
{
public:
	MyQPrinter(PrinterMode mode = ScreenResolution) : QPrinter(mode) {};
	MyQPrinter(const QPrinterInfo &printer, PrinterMode mode = ScreenResolution) : QPrinter(printer, mode) {};
	bool newPage() { return callbackQPrinter_NewPage(this) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrinter_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	void setPageSize(QPagedPaintDevice::PageSize size) { callbackQPrinter_SetPageSize2(this, size); };
	void setPageSizeMM(const QSizeF & size) { callbackQPrinter_SetPageSizeMM(this, const_cast<QSizeF*>(&size)); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPrinter_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
};

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

char QPrinter_NewPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->newPage();
}

char QPrinter_NewPageDefault(void* ptr)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::newPage();
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

void QPrinter_SetCreator(void* ptr, struct QtPrintSupport_PackedString creator)
{
	static_cast<QPrinter*>(ptr)->setCreator(QString::fromUtf8(creator.data, creator.len));
}

void QPrinter_SetDocName(void* ptr, struct QtPrintSupport_PackedString name)
{
	static_cast<QPrinter*>(ptr)->setDocName(QString::fromUtf8(name.data, name.len));
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

void QPrinter_SetOutputFileName(void* ptr, struct QtPrintSupport_PackedString fileName)
{
	static_cast<QPrinter*>(ptr)->setOutputFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QPrinter_SetOutputFormat(void* ptr, long long format)
{
	static_cast<QPrinter*>(ptr)->setOutputFormat(static_cast<QPrinter::OutputFormat>(format));
}

void QPrinter_SetPageOrder(void* ptr, long long pageOrder)
{
	static_cast<QPrinter*>(ptr)->setPageOrder(static_cast<QPrinter::PageOrder>(pageOrder));
}

void QPrinter_SetPaperSource(void* ptr, long long source)
{
	static_cast<QPrinter*>(ptr)->setPaperSource(static_cast<QPrinter::PaperSource>(source));
}

void QPrinter_SetPrintProgram(void* ptr, struct QtPrintSupport_PackedString printProg)
{
	static_cast<QPrinter*>(ptr)->setPrintProgram(QString::fromUtf8(printProg.data, printProg.len));
}

void QPrinter_SetPrintRange(void* ptr, long long ran)
{
	static_cast<QPrinter*>(ptr)->setPrintRange(static_cast<QPrinter::PrintRange>(ran));
}

void QPrinter_SetPrinterName(void* ptr, struct QtPrintSupport_PackedString name)
{
	static_cast<QPrinter*>(ptr)->setPrinterName(QString::fromUtf8(name.data, name.len));
}

void QPrinter_SetPrinterSelectionOption(void* ptr, struct QtPrintSupport_PackedString option)
{
	static_cast<QPrinter*>(ptr)->setPrinterSelectionOption(QString::fromUtf8(option.data, option.len));
}

void QPrinter_SetResolution(void* ptr, int dpi)
{
	static_cast<QPrinter*>(ptr)->setResolution(dpi);
}

void QPrinter_DestroyQPrinter(void* ptr)
{
	static_cast<QPrinter*>(ptr)->~QPrinter();
}

long long QPrinter_ColorMode(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->colorMode();
}

long long QPrinter_Duplex(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->duplex();
}

long long QPrinter_OutputFormat(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->outputFormat();
}

long long QPrinter_PageOrder(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->pageOrder();
}

long long QPrinter_PaperSource(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->paperSource();
}

long long QPrinter_PrintRange(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printRange();
}

long long QPrinter_PrinterState(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printerState();
}

struct QtPrintSupport_PackedList QPrinter_SupportedResolutions(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QPrinter*>(ptr)->supportedResolutions()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

void* QPrinter_PaintEngine(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->paintEngine();
}

void* QPrinter_PaintEngineDefault(void* ptr)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::paintEngine();
}

void* QPrinter_PrintEngine(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->printEngine();
}

void* QPrinter_PageRect(void* ptr, long long unit)
{
	return ({ QRectF tmpValue = static_cast<QPrinter*>(ptr)->pageRect(static_cast<QPrinter::Unit>(unit)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPrinter_PaperRect(void* ptr, long long unit)
{
	return ({ QRectF tmpValue = static_cast<QPrinter*>(ptr)->paperRect(static_cast<QPrinter::Unit>(unit)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtPrintSupport_PackedString QPrinter_Creator(void* ptr)
{
	return ({ QByteArray ta24181 = static_cast<QPrinter*>(ptr)->creator().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(ta24181.prepend("WHITESPACE").constData()+10), ta24181.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_DocName(void* ptr)
{
	return ({ QByteArray t72c43e = static_cast<QPrinter*>(ptr)->docName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t72c43e.prepend("WHITESPACE").constData()+10), t72c43e.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_OutputFileName(void* ptr)
{
	return ({ QByteArray tde340f = static_cast<QPrinter*>(ptr)->outputFileName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tde340f.prepend("WHITESPACE").constData()+10), tde340f.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_PrintProgram(void* ptr)
{
	return ({ QByteArray tb89767 = static_cast<QPrinter*>(ptr)->printProgram().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tb89767.prepend("WHITESPACE").constData()+10), tb89767.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_PrinterName(void* ptr)
{
	return ({ QByteArray ta7ec2f = static_cast<QPrinter*>(ptr)->printerName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(ta7ec2f.prepend("WHITESPACE").constData()+10), ta7ec2f.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinter_PrinterSelectionOption(void* ptr)
{
	return ({ QByteArray te1c3f5 = static_cast<QPrinter*>(ptr)->printerSelectionOption().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(te1c3f5.prepend("WHITESPACE").constData()+10), te1c3f5.size()-10 }; });
}

char QPrinter_CollateCopies(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->collateCopies();
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

char QPrinter_SupportsMultipleCopies(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->supportsMultipleCopies();
}

int QPrinter_CopyCount(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->copyCount();
}

int QPrinter_FromPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->fromPage();
}

int QPrinter_Resolution(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->resolution();
}

int QPrinter_ToPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->toPage();
}

int QPrinter___supportedResolutions_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QPrinter___supportedResolutions_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QPrinter___supportedResolutions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void QPrinter_SetPageSize2Default(void* ptr, long long size)
{
		static_cast<QPrinter*>(ptr)->QPrinter::setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPrinter_SetPageSizeMMDefault(void* ptr, void* size)
{
		static_cast<QPrinter*>(ptr)->QPrinter::setPageSizeMM(*static_cast<QSizeF*>(size));
}

int QPrinter_MetricDefault(void* ptr, long long metric)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

struct QtPrintSupport_PackedList QPrinterInfo_QPrinterInfo_AvailablePrinters()
{
	return ({ QList<QPrinterInfo>* tmpValue = new QList<QPrinterInfo>(QPrinterInfo::availablePrinters()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

void* QPrinterInfo_QPrinterInfo_DefaultPrinter()
{
	return new QPrinterInfo(QPrinterInfo::defaultPrinter());
}

void* QPrinterInfo_QPrinterInfo_PrinterInfo(struct QtPrintSupport_PackedString printerName)
{
	return new QPrinterInfo(QPrinterInfo::printerInfo(QString::fromUtf8(printerName.data, printerName.len)));
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

struct QtPrintSupport_PackedString QPrinterInfo_QPrinterInfo_DefaultPrinterName()
{
	return ({ QByteArray t8ce1a9 = QPrinterInfo::defaultPrinterName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t8ce1a9.prepend("WHITESPACE").constData()+10), t8ce1a9.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_QPrinterInfo_AvailablePrinterNames()
{
	return ({ QByteArray tb6ad9e = QPrinterInfo::availablePrinterNames().join("|").toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(tb6ad9e.prepend("WHITESPACE").constData()+10), tb6ad9e.size()-10 }; });
}

void QPrinterInfo_DestroyQPrinterInfo(void* ptr)
{
	static_cast<QPrinterInfo*>(ptr)->~QPrinterInfo();
}

struct QtPrintSupport_PackedList QPrinterInfo_SupportedPageSizes(void* ptr)
{
	return ({ QList<QPageSize>* tmpValue = new QList<QPageSize>(static_cast<QPrinterInfo*>(ptr)->supportedPageSizes()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPrintSupport_PackedList QPrinterInfo_SupportedDuplexModes(void* ptr)
{
	return ({ QList<QPrinter::DuplexMode>* tmpValue = new QList<QPrinter::DuplexMode>(static_cast<QPrinterInfo*>(ptr)->supportedDuplexModes()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPrintSupport_PackedList QPrinterInfo_SupportedResolutions(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QPrinterInfo*>(ptr)->supportedResolutions()); QtPrintSupport_PackedList { tmpValue, tmpValue->size() }; });
}

void* QPrinterInfo_DefaultPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->defaultPageSize());
}

void* QPrinterInfo_MaximumPhysicalPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->maximumPhysicalPageSize());
}

void* QPrinterInfo_MinimumPhysicalPageSize(void* ptr)
{
	return new QPageSize(static_cast<QPrinterInfo*>(ptr)->minimumPhysicalPageSize());
}

long long QPrinterInfo_DefaultDuplexMode(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->defaultDuplexMode();
}

long long QPrinterInfo_State(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->state();
}

struct QtPrintSupport_PackedString QPrinterInfo_Description(void* ptr)
{
	return ({ QByteArray t0d7900 = static_cast<QPrinterInfo*>(ptr)->description().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t0d7900.prepend("WHITESPACE").constData()+10), t0d7900.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_Location(void* ptr)
{
	return ({ QByteArray t360fc3 = static_cast<QPrinterInfo*>(ptr)->location().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t360fc3.prepend("WHITESPACE").constData()+10), t360fc3.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_MakeAndModel(void* ptr)
{
	return ({ QByteArray t372e5c = static_cast<QPrinterInfo*>(ptr)->makeAndModel().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(t372e5c.prepend("WHITESPACE").constData()+10), t372e5c.size()-10 }; });
}

struct QtPrintSupport_PackedString QPrinterInfo_PrinterName(void* ptr)
{
	return ({ QByteArray te478b0 = static_cast<QPrinterInfo*>(ptr)->printerName().toUtf8(); QtPrintSupport_PackedString { const_cast<char*>(te478b0.prepend("WHITESPACE").constData()+10), te478b0.size()-10 }; });
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

char QPrinterInfo_SupportsCustomPageSizes(void* ptr)
{
	return static_cast<QPrinterInfo*>(ptr)->supportsCustomPageSizes();
}

void* QPrinterInfo___availablePrinters_atList(void* ptr, int i)
{
	return new QPrinterInfo(static_cast<QList<QPrinterInfo>*>(ptr)->at(i));
}

void QPrinterInfo___availablePrinters_setList(void* ptr, void* i)
{
	static_cast<QList<QPrinterInfo>*>(ptr)->append(*static_cast<QPrinterInfo*>(i));
}

void* QPrinterInfo___availablePrinters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinterInfo>;
}

void* QPrinterInfo___supportedPageSizes_atList(void* ptr, int i)
{
	return new QPageSize(static_cast<QList<QPageSize>*>(ptr)->at(i));
}

void QPrinterInfo___supportedPageSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QPageSize>*>(ptr)->append(*static_cast<QPageSize*>(i));
}

void* QPrinterInfo___supportedPageSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPageSize>;
}

long long QPrinterInfo___supportedDuplexModes_atList(void* ptr, int i)
{
	return static_cast<QList<QPrinter::DuplexMode>*>(ptr)->at(i);
}

void QPrinterInfo___supportedDuplexModes_setList(void* ptr, long long i)
{
	static_cast<QList<QPrinter::DuplexMode>*>(ptr)->append(static_cast<QPrinter::DuplexMode>(i));
}

void* QPrinterInfo___supportedDuplexModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinter::DuplexMode>;
}

void* QPrinterInfo___supportedPaperSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinter::PaperSize>;
}

int QPrinterInfo___supportedResolutions_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QPrinterInfo___supportedResolutions_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QPrinterInfo___supportedResolutions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

