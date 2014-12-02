#include "qlabel.h"
#include <QLabel>
#include "cgoexport.h"

//MyQLabel
class MyQLabel: public QLabel {
Q_OBJECT
public:
void Signal_LinkActivated() { callbackQLabel(this, QString("linkActivated").toUtf8().data()); };
void Signal_LinkHovered() { callbackQLabel(this, QString("linkHovered").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_SetText(QString text);

};
#include "qlabel.moc"


//Public Functions
QtObjectPtr QLabel_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QLabel(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

QtObjectPtr QLabel_New_String_QWidget_WindowType(char* text, QtObjectPtr parent, int f)
{
	return new QLabel(QString(text), ((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QLabel_Destroy(QtObjectPtr ptr)
{
	((QLabel*)(ptr))->~QLabel();
}

int QLabel_Alignment(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->alignment();
}

QtObjectPtr QLabel_Buddy(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->buddy();
}

int QLabel_HasScaledContents(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->hasScaledContents();
}

int QLabel_HasSelectedText(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->hasSelectedText();
}

int QLabel_Indent(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->indent();
}

int QLabel_Margin(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->margin();
}

int QLabel_OpenExternalLinks(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->openExternalLinks();
}

char* QLabel_SelectedText(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->selectedText().toUtf8().data();
}

int QLabel_SelectionStart(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->selectionStart();
}

void QLabel_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment)
{
	((QLabel*)(ptr))->setAlignment(((Qt::AlignmentFlag)(alignment)));
}

void QLabel_SetBuddy_QWidget(QtObjectPtr ptr, QtObjectPtr buddy)
{
	((QLabel*)(ptr))->setBuddy(((QWidget*)(buddy)));
}

void QLabel_SetIndent_Int(QtObjectPtr ptr, int indent)
{
	((QLabel*)(ptr))->setIndent(indent);
}

void QLabel_SetMargin_Int(QtObjectPtr ptr, int margin)
{
	((QLabel*)(ptr))->setMargin(margin);
}

void QLabel_SetOpenExternalLinks_Bool(QtObjectPtr ptr, int open)
{
	((QLabel*)(ptr))->setOpenExternalLinks(open != 0);
}

void QLabel_SetScaledContents_Bool(QtObjectPtr ptr, int scaledContents)
{
	((QLabel*)(ptr))->setScaledContents(scaledContents != 0);
}

void QLabel_SetSelection_Int_Int(QtObjectPtr ptr, int start, int length)
{
	((QLabel*)(ptr))->setSelection(start, length);
}

void QLabel_SetTextFormat_TextFormat(QtObjectPtr ptr, int textFormat)
{
	((QLabel*)(ptr))->setTextFormat(((Qt::TextFormat)(textFormat)));
}

void QLabel_SetTextInteractionFlags_TextInteractionFlag(QtObjectPtr ptr, int flags)
{
	((QLabel*)(ptr))->setTextInteractionFlags(((Qt::TextInteractionFlag)(flags)));
}

void QLabel_SetWordWrap_Bool(QtObjectPtr ptr, int on)
{
	((QLabel*)(ptr))->setWordWrap(on != 0);
}

char* QLabel_Text(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->text().toUtf8().data();
}

int QLabel_TextFormat(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->textFormat();
}

int QLabel_TextInteractionFlags(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->textInteractionFlags();
}

int QLabel_WordWrap(QtObjectPtr ptr)
{
	return ((QLabel*)(ptr))->wordWrap();
}

//Public Slots
void QLabel_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQLabel*)(ptr)), &MyQLabel::Slot_Clear, ((QLabel*)(ptr)), &QLabel::clear, Qt::QueuedConnection);
}

void QLabel_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLabel*)(ptr)), &MyQLabel::Slot_Clear, ((QLabel*)(ptr)), &QLabel::clear);
}

void QLabel_Clear(QtObjectPtr ptr)
{
	((MyQLabel*)(ptr))->Slot_Clear();
}

void QLabel_ConnectSlotSetText(QtObjectPtr ptr)
{
	QObject::connect(((MyQLabel*)(ptr)), &MyQLabel::Slot_SetText, ((QLabel*)(ptr)), &QLabel::setText, Qt::QueuedConnection);
}

void QLabel_DisconnectSlotSetText(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQLabel*)(ptr)), &MyQLabel::Slot_SetText, ((QLabel*)(ptr)), &QLabel::setText);
}

void QLabel_SetText_String(QtObjectPtr ptr, char* text)
{
	((MyQLabel*)(ptr))->Slot_SetText(QString(text));
}

//Signals
void QLabel_ConnectSignalLinkActivated(QtObjectPtr ptr)
{
	QObject::connect(((QLabel*)(ptr)), &QLabel::linkActivated, ((MyQLabel*)(ptr)), &MyQLabel::Signal_LinkActivated, Qt::QueuedConnection);
}

void QLabel_DisconnectSignalLinkActivated(QtObjectPtr ptr)
{
	QObject::disconnect(((QLabel*)(ptr)), &QLabel::linkActivated, ((MyQLabel*)(ptr)), &MyQLabel::Signal_LinkActivated);
}

void QLabel_ConnectSignalLinkHovered(QtObjectPtr ptr)
{
	QObject::connect(((QLabel*)(ptr)), &QLabel::linkHovered, ((MyQLabel*)(ptr)), &MyQLabel::Signal_LinkHovered, Qt::QueuedConnection);
}

void QLabel_DisconnectSignalLinkHovered(QtObjectPtr ptr)
{
	QObject::disconnect(((QLabel*)(ptr)), &QLabel::linkHovered, ((MyQLabel*)(ptr)), &MyQLabel::Signal_LinkHovered);
}

