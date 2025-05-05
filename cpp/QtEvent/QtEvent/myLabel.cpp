#include "myLabel.h"
#include <QDebug>
#include <QMouseEvent>

myLabel::myLabel(QWidget* parent)
	: QLabel(parent)
{
	// 设置鼠标追踪
	setMouseTracking(true);
}

myLabel::~myLabel()
{
}

//// 鼠标进入事件
//void myLabel::enterEvent(QEvent* event)
//{
//	qDebug() << "鼠标进入了";
//}
//
//// 鼠标的离开
//void myLabel::leaveEvent(QEvent* event)
//{
//	qDebug() << "鼠标离开了";
//}

// 鼠标按下
void myLabel::mousePressEvent(QMouseEvent* ev)
{
	if (ev->button() == Qt::LeftButton) {
		QString str = QString("鼠标左键按下了 x = %1 y = %2, global X %3 global Y %4").arg(ev->x()).arg(ev->y()).arg(ev->globalX()).arg(ev->globalY());
		qDebug() << str;
	}
}

// 鼠标释放
void myLabel::mouseReleaseEvent(QMouseEvent* ev)
{
	qDebug() << "鼠标释放了";
}

// 鼠标移动
void myLabel::mouseMoveEvent(QMouseEvent* ev)
{
	//if (ev->buttons() & Qt::LeftButton)
	//{
	QString str = QString("鼠标左键移动了 x = %1 y = %2, global X %3 global Y %4").arg(ev->x()).arg(ev->y()).arg(ev->globalX()).arg(ev->globalY());
	qDebug() << str;
	//}
}
