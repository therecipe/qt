

#define protected public
#define private public

#include "moc.h"
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
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLabel>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMovie>
#include <QObject>
#include <QPaintEvent>
#include <QPicture>
#include <QPixmap>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

class DropArea: public QLabel
{
Q_OBJECT
public:
	DropArea(QWidget *parent, Qt::WindowFlags f) : QLabel(parent, f) {};
	void Signal_Changed(QMimeData* mimeData) { callbackDropArea_Changed(this, this->objectName().toUtf8().data(), mimeData); };
	void setPixmap(const QPixmap & vqp) { callbackDropArea_SetPixmap(this, this->objectName().toUtf8().data(), new QPixmap(vqp)); };
	void setText(const QString & vqs) { callbackDropArea_SetText(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void changeEvent(QEvent * ev) { callbackDropArea_ChangeEvent(this, this->objectName().toUtf8().data(), ev); };
	void clear() { callbackDropArea_Clear(this, this->objectName().toUtf8().data()); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackDropArea_ContextMenuEvent(this, this->objectName().toUtf8().data(), ev); };
	void focusInEvent(QFocusEvent * ev) { callbackDropArea_FocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	bool focusNextPrevChild(bool next) { return callbackDropArea_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	void focusOutEvent(QFocusEvent * ev) { callbackDropArea_FocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	int heightForWidth(int w) const { return callbackDropArea_HeightForWidth(const_cast<DropArea*>(this), this->objectName().toUtf8().data(), w); };
	void keyPressEvent(QKeyEvent * ev) { callbackDropArea_KeyPressEvent(this, this->objectName().toUtf8().data(), ev); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackDropArea_MinimumSizeHint(const_cast<DropArea*>(this), this->objectName().toUtf8().data())); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackDropArea_MouseMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackDropArea_MousePressEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackDropArea_MouseReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void paintEvent(QPaintEvent * vqp) { callbackDropArea_PaintEvent(this, this->objectName().toUtf8().data(), vqp); };
	void setMovie(QMovie * movie) { callbackDropArea_SetMovie(this, this->objectName().toUtf8().data(), movie); };
	
	void setNum(int num) { callbackDropArea_SetNum(this, this->objectName().toUtf8().data(), num); };
	void setPicture(const QPicture & picture) { callbackDropArea_SetPicture(this, this->objectName().toUtf8().data(), new QPicture(picture)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackDropArea_SizeHint(const_cast<DropArea*>(this), this->objectName().toUtf8().data())); };
	void actionEvent(QActionEvent * event) { callbackDropArea_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackDropArea_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackDropArea_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackDropArea_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackDropArea_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackDropArea_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackDropArea_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackDropArea_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackDropArea_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackDropArea_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackDropArea_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackDropArea_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackDropArea_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackDropArea_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackDropArea_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackDropArea_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackDropArea_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	bool hasHeightForWidth() const { return callbackDropArea_HasHeightForWidth(const_cast<DropArea*>(this), this->objectName().toUtf8().data()) != 0; };
	void hide() { callbackDropArea_Hide(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackDropArea_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackDropArea_InputMethodQuery(const_cast<DropArea*>(this), this->objectName().toUtf8().data(), query)); };
	void keyReleaseEvent(QKeyEvent * event) { callbackDropArea_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackDropArea_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackDropArea_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackDropArea_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackDropArea_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackDropArea_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackDropArea_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackDropArea_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackDropArea_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackDropArea_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackDropArea_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackDropArea_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackDropArea_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackDropArea_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackDropArea_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackDropArea_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackDropArea_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackDropArea_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackDropArea_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackDropArea_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackDropArea_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackDropArea_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackDropArea_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackDropArea_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDropArea_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackDropArea_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	
signals:
	void changed(QMimeData* mimeData);
public slots:
};

void DropArea_ConnectChanged(void* ptr)
{
	QObject::connect(static_cast<DropArea*>(ptr), static_cast<void (DropArea::*)(QMimeData*)>(&DropArea::changed), static_cast<DropArea*>(ptr), static_cast<void (DropArea::*)(QMimeData*)>(&DropArea::Signal_Changed));
}

void DropArea_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<DropArea*>(ptr), static_cast<void (DropArea::*)(QMimeData*)>(&DropArea::changed), static_cast<DropArea*>(ptr), static_cast<void (DropArea::*)(QMimeData*)>(&DropArea::Signal_Changed));
}

void DropArea_Changed(void* ptr, void* mimeData)
{
	static_cast<DropArea*>(ptr)->changed(static_cast<QMimeData*>(mimeData));
}

void* DropArea_NewDropArea(void* parent, int f)
{
	return new DropArea(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void DropArea_DestroyDropArea(void* ptr)
{
	static_cast<DropArea*>(ptr)->~DropArea();
}

void DropArea_SetPixmap(void* ptr, void* vqp)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setPixmap", Q_ARG(QPixmap, *static_cast<QPixmap*>(vqp)));
}

void DropArea_SetPixmapDefault(void* ptr, void* vqp)
{
	static_cast<DropArea*>(ptr)->QLabel::setPixmap(*static_cast<QPixmap*>(vqp));
}

void DropArea_SetText(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setText", Q_ARG(QString, QString(vqs)));
}

void DropArea_SetTextDefault(void* ptr, char* vqs)
{
	static_cast<DropArea*>(ptr)->QLabel::setText(QString(vqs));
}

void DropArea_ChangeEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->changeEvent(static_cast<QEvent*>(ev));
}

void DropArea_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::changeEvent(static_cast<QEvent*>(ev));
}

void DropArea_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "clear");
}

void DropArea_ClearDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::clear();
}

void DropArea_ContextMenuEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void DropArea_ContextMenuEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void DropArea_FocusInEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void DropArea_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::focusInEvent(static_cast<QFocusEvent*>(ev));
}

int DropArea_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<DropArea*>(ptr)->focusNextPrevChild(next != 0);
}

int DropArea_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<DropArea*>(ptr)->QLabel::focusNextPrevChild(next != 0);
}

void DropArea_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void DropArea_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

int DropArea_HeightForWidth(void* ptr, int w)
{
	return static_cast<DropArea*>(ptr)->heightForWidth(w);
}

int DropArea_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<DropArea*>(ptr)->QLabel::heightForWidth(w);
}

void DropArea_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void DropArea_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void* DropArea_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropArea*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<DropArea*>(ptr)->minimumSizeHint()).height());
}

void* DropArea_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropArea*>(ptr)->QLabel::minimumSizeHint()).width(), static_cast<QSize>(static_cast<DropArea*>(ptr)->QLabel::minimumSizeHint()).height());
}

void DropArea_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_MousePressEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<DropArea*>(ptr)->QLabel::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void DropArea_PaintEvent(void* ptr, void* vqp)
{
	static_cast<DropArea*>(ptr)->paintEvent(static_cast<QPaintEvent*>(vqp));
}

void DropArea_PaintEventDefault(void* ptr, void* vqp)
{
	static_cast<DropArea*>(ptr)->QLabel::paintEvent(static_cast<QPaintEvent*>(vqp));
}

void DropArea_SetMovie(void* ptr, void* movie)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setMovie", Q_ARG(QMovie*, static_cast<QMovie*>(movie)));
}

void DropArea_SetMovieDefault(void* ptr, void* movie)
{
	static_cast<DropArea*>(ptr)->QLabel::setMovie(static_cast<QMovie*>(movie));
}





void DropArea_SetNum(void* ptr, int num)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setNum", Q_ARG(int, num));
}

void DropArea_SetNumDefault(void* ptr, int num)
{
	static_cast<DropArea*>(ptr)->QLabel::setNum(num);
}

void DropArea_SetPicture(void* ptr, void* picture)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setPicture", Q_ARG(QPicture, *static_cast<QPicture*>(picture)));
}

void DropArea_SetPictureDefault(void* ptr, void* picture)
{
	static_cast<DropArea*>(ptr)->QLabel::setPicture(*static_cast<QPicture*>(picture));
}

void* DropArea_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropArea*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<DropArea*>(ptr)->sizeHint()).height());
}

void* DropArea_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropArea*>(ptr)->QLabel::sizeHint()).width(), static_cast<QSize>(static_cast<DropArea*>(ptr)->QLabel::sizeHint()).height());
}

void DropArea_ActionEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void DropArea_ActionEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::actionEvent(static_cast<QActionEvent*>(event));
}

void DropArea_DragEnterEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void DropArea_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void DropArea_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void DropArea_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void DropArea_DragMoveEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void DropArea_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void DropArea_DropEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void DropArea_DropEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::dropEvent(static_cast<QDropEvent*>(event));
}

void DropArea_EnterEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void DropArea_EnterEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::enterEvent(static_cast<QEvent*>(event));
}

void DropArea_HideEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void DropArea_HideEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::hideEvent(static_cast<QHideEvent*>(event));
}

void DropArea_LeaveEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void DropArea_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::leaveEvent(static_cast<QEvent*>(event));
}

void DropArea_MoveEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void DropArea_MoveEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::moveEvent(static_cast<QMoveEvent*>(event));
}

void DropArea_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void DropArea_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<DropArea*>(ptr)->QLabel::setEnabled(vbo != 0);
}

void DropArea_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void DropArea_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<DropArea*>(ptr)->QLabel::setStyleSheet(QString(styleSheet));
}

void DropArea_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void DropArea_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<DropArea*>(ptr)->QLabel::setVisible(visible != 0);
}

void DropArea_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void DropArea_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<DropArea*>(ptr)->QLabel::setWindowModified(vbo != 0);
}

void DropArea_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void DropArea_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<DropArea*>(ptr)->QLabel::setWindowTitle(QString(vqs));
}

void DropArea_ShowEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void DropArea_ShowEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::showEvent(static_cast<QShowEvent*>(event));
}

int DropArea_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int DropArea_CloseDefault(void* ptr)
{
	return static_cast<DropArea*>(ptr)->QLabel::close();
}

void DropArea_CloseEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void DropArea_CloseEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::closeEvent(static_cast<QCloseEvent*>(event));
}

int DropArea_HasHeightForWidth(void* ptr)
{
	return static_cast<DropArea*>(ptr)->hasHeightForWidth();
}

int DropArea_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<DropArea*>(ptr)->QLabel::hasHeightForWidth();
}

void DropArea_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "hide");
}

void DropArea_HideDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::hide();
}

void DropArea_InputMethodEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void DropArea_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* DropArea_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<DropArea*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* DropArea_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<DropArea*>(ptr)->QLabel::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void DropArea_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void DropArea_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void DropArea_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "lower");
}

void DropArea_LowerDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::lower();
}

void DropArea_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void DropArea_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

int DropArea_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<DropArea*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int DropArea_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<DropArea*>(ptr)->QLabel::nativeEvent(QByteArray(eventType), message, &result);
}

void DropArea_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "raise");
}

void DropArea_RaiseDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::raise();
}

void DropArea_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "repaint");
}

void DropArea_RepaintDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::repaint();
}

void DropArea_ResizeEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void DropArea_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::resizeEvent(static_cast<QResizeEvent*>(event));
}

void DropArea_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void DropArea_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<DropArea*>(ptr)->QLabel::setDisabled(disable != 0);
}

void DropArea_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setFocus");
}

void DropArea_SetFocus2Default(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::setFocus();
}

void DropArea_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void DropArea_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<DropArea*>(ptr)->QLabel::setHidden(hidden != 0);
}

void DropArea_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "show");
}

void DropArea_ShowDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::show();
}

void DropArea_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "showFullScreen");
}

void DropArea_ShowFullScreenDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::showFullScreen();
}

void DropArea_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "showMaximized");
}

void DropArea_ShowMaximizedDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::showMaximized();
}

void DropArea_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "showMinimized");
}

void DropArea_ShowMinimizedDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::showMinimized();
}

void DropArea_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "showNormal");
}

void DropArea_ShowNormalDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::showNormal();
}

void DropArea_TabletEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void DropArea_TabletEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::tabletEvent(static_cast<QTabletEvent*>(event));
}

void DropArea_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "update");
}

void DropArea_UpdateDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::update();
}

void DropArea_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "updateMicroFocus");
}

void DropArea_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::updateMicroFocus();
}

void DropArea_WheelEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void DropArea_WheelEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::wheelEvent(static_cast<QWheelEvent*>(event));
}

void DropArea_TimerEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void DropArea_TimerEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::timerEvent(static_cast<QTimerEvent*>(event));
}

void DropArea_ChildEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void DropArea_ChildEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::childEvent(static_cast<QChildEvent*>(event));
}

void DropArea_ConnectNotify(void* ptr, void* sign)
{
	static_cast<DropArea*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropArea_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DropArea*>(ptr)->QLabel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropArea_CustomEvent(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void DropArea_CustomEventDefault(void* ptr, void* event)
{
	static_cast<DropArea*>(ptr)->QLabel::customEvent(static_cast<QEvent*>(event));
}

void DropArea_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropArea*>(ptr), "deleteLater");
}

void DropArea_DeleteLaterDefault(void* ptr)
{
	static_cast<DropArea*>(ptr)->QLabel::deleteLater();
}

void DropArea_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<DropArea*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropArea_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DropArea*>(ptr)->QLabel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int DropArea_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<DropArea*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int DropArea_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<DropArea*>(ptr)->QLabel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}





class DropSiteWindow: public QWidget
{
Q_OBJECT
public:
	DropSiteWindow(QWidget *parent, Qt::WindowFlags f) : QWidget(parent, f) {};
	void actionEvent(QActionEvent * event) { callbackDropSiteWindow_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackDropSiteWindow_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackDropSiteWindow_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackDropSiteWindow_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackDropSiteWindow_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackDropSiteWindow_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackDropSiteWindow_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackDropSiteWindow_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackDropSiteWindow_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackDropSiteWindow_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackDropSiteWindow_MinimumSizeHint(const_cast<DropSiteWindow*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackDropSiteWindow_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackDropSiteWindow_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackDropSiteWindow_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackDropSiteWindow_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackDropSiteWindow_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackDropSiteWindow_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackDropSiteWindow_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackDropSiteWindow_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackDropSiteWindow_SizeHint(const_cast<DropSiteWindow*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackDropSiteWindow_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackDropSiteWindow_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackDropSiteWindow_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackDropSiteWindow_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackDropSiteWindow_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackDropSiteWindow_HasHeightForWidth(const_cast<DropSiteWindow*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackDropSiteWindow_HeightForWidth(const_cast<DropSiteWindow*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackDropSiteWindow_Hide(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackDropSiteWindow_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackDropSiteWindow_InputMethodQuery(const_cast<DropSiteWindow*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackDropSiteWindow_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackDropSiteWindow_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackDropSiteWindow_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackDropSiteWindow_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackDropSiteWindow_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackDropSiteWindow_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackDropSiteWindow_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackDropSiteWindow_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackDropSiteWindow_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackDropSiteWindow_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackDropSiteWindow_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackDropSiteWindow_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackDropSiteWindow_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackDropSiteWindow_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackDropSiteWindow_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackDropSiteWindow_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackDropSiteWindow_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackDropSiteWindow_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackDropSiteWindow_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackDropSiteWindow_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackDropSiteWindow_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackDropSiteWindow_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackDropSiteWindow_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackDropSiteWindow_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackDropSiteWindow_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackDropSiteWindow_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackDropSiteWindow_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackDropSiteWindow_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDropSiteWindow_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackDropSiteWindow_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	
signals:
public slots:
	void updateFormatsTable(QMimeData* mimeData) { callbackDropSiteWindow_UpdateFormatsTable(this, this->objectName().toUtf8().data(), mimeData); };
};

void DropSiteWindow_UpdateFormatsTable(void* ptr, void* mimeData)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "updateFormatsTable", Q_ARG(QMimeData*, static_cast<QMimeData*>(mimeData)));
}

void* DropSiteWindow_NewDropSiteWindow(void* parent, int f)
{
	return new DropSiteWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void DropSiteWindow_DestroyDropSiteWindow(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->~DropSiteWindow();
}

void DropSiteWindow_ActionEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void DropSiteWindow_ActionEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void DropSiteWindow_DragEnterEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void DropSiteWindow_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void DropSiteWindow_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void DropSiteWindow_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void DropSiteWindow_DragMoveEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void DropSiteWindow_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void DropSiteWindow_DropEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void DropSiteWindow_DropEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void DropSiteWindow_EnterEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_EnterEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::enterEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_FocusInEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void DropSiteWindow_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void DropSiteWindow_FocusOutEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void DropSiteWindow_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void DropSiteWindow_HideEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void DropSiteWindow_HideEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void DropSiteWindow_LeaveEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* DropSiteWindow_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->minimumSizeHint()).height());
}

void* DropSiteWindow_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->QWidget::minimumSizeHint()).width(), static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->QWidget::minimumSizeHint()).height());
}

void DropSiteWindow_MoveEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void DropSiteWindow_MoveEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void DropSiteWindow_PaintEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void DropSiteWindow_PaintEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void DropSiteWindow_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void DropSiteWindow_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setEnabled(vbo != 0);
}

void DropSiteWindow_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void DropSiteWindow_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setStyleSheet(QString(styleSheet));
}

void DropSiteWindow_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void DropSiteWindow_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setVisible(visible != 0);
}

void DropSiteWindow_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void DropSiteWindow_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setWindowModified(vbo != 0);
}

void DropSiteWindow_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void DropSiteWindow_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setWindowTitle(QString(vqs));
}

void DropSiteWindow_ShowEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void DropSiteWindow_ShowEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* DropSiteWindow_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->sizeHint()).height());
}

void* DropSiteWindow_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->QWidget::sizeHint()).width(), static_cast<QSize>(static_cast<DropSiteWindow*>(ptr)->QWidget::sizeHint()).height());
}

void DropSiteWindow_ChangeEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::changeEvent(static_cast<QEvent*>(event));
}

int DropSiteWindow_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int DropSiteWindow_CloseDefault(void* ptr)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::close();
}

void DropSiteWindow_CloseEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void DropSiteWindow_CloseEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void DropSiteWindow_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void DropSiteWindow_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int DropSiteWindow_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<DropSiteWindow*>(ptr)->focusNextPrevChild(next != 0);
}

int DropSiteWindow_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::focusNextPrevChild(next != 0);
}

int DropSiteWindow_HasHeightForWidth(void* ptr)
{
	return static_cast<DropSiteWindow*>(ptr)->hasHeightForWidth();
}

int DropSiteWindow_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::hasHeightForWidth();
}

int DropSiteWindow_HeightForWidth(void* ptr, int w)
{
	return static_cast<DropSiteWindow*>(ptr)->heightForWidth(w);
}

int DropSiteWindow_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::heightForWidth(w);
}

void DropSiteWindow_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "hide");
}

void DropSiteWindow_HideDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::hide();
}

void DropSiteWindow_InputMethodEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void DropSiteWindow_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* DropSiteWindow_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<DropSiteWindow*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* DropSiteWindow_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<DropSiteWindow*>(ptr)->QWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void DropSiteWindow_KeyPressEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void DropSiteWindow_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void DropSiteWindow_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void DropSiteWindow_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void DropSiteWindow_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "lower");
}

void DropSiteWindow_LowerDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::lower();
}

void DropSiteWindow_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MousePressEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void DropSiteWindow_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int DropSiteWindow_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<DropSiteWindow*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int DropSiteWindow_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void DropSiteWindow_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "raise");
}

void DropSiteWindow_RaiseDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::raise();
}

void DropSiteWindow_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "repaint");
}

void DropSiteWindow_RepaintDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::repaint();
}

void DropSiteWindow_ResizeEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void DropSiteWindow_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void DropSiteWindow_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void DropSiteWindow_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setDisabled(disable != 0);
}

void DropSiteWindow_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setFocus");
}

void DropSiteWindow_SetFocus2Default(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setFocus();
}

void DropSiteWindow_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void DropSiteWindow_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::setHidden(hidden != 0);
}

void DropSiteWindow_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "show");
}

void DropSiteWindow_ShowDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::show();
}

void DropSiteWindow_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "showFullScreen");
}

void DropSiteWindow_ShowFullScreenDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::showFullScreen();
}

void DropSiteWindow_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "showMaximized");
}

void DropSiteWindow_ShowMaximizedDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::showMaximized();
}

void DropSiteWindow_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "showMinimized");
}

void DropSiteWindow_ShowMinimizedDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::showMinimized();
}

void DropSiteWindow_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "showNormal");
}

void DropSiteWindow_ShowNormalDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::showNormal();
}

void DropSiteWindow_TabletEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void DropSiteWindow_TabletEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void DropSiteWindow_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "update");
}

void DropSiteWindow_UpdateDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::update();
}

void DropSiteWindow_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "updateMicroFocus");
}

void DropSiteWindow_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::updateMicroFocus();
}

void DropSiteWindow_WheelEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void DropSiteWindow_WheelEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void DropSiteWindow_TimerEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void DropSiteWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void DropSiteWindow_ChildEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void DropSiteWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::childEvent(static_cast<QChildEvent*>(event));
}

void DropSiteWindow_ConnectNotify(void* ptr, void* sign)
{
	static_cast<DropSiteWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropSiteWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropSiteWindow_CustomEvent(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::customEvent(static_cast<QEvent*>(event));
}

void DropSiteWindow_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<DropSiteWindow*>(ptr), "deleteLater");
}

void DropSiteWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::deleteLater();
}

void DropSiteWindow_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<DropSiteWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void DropSiteWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<DropSiteWindow*>(ptr)->QWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int DropSiteWindow_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<DropSiteWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int DropSiteWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<DropSiteWindow*>(ptr)->QWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}





#include "moc_moc.h"
