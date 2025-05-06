#include "QtPainterEvent.h"
#include <QPainter> // 画家类

QtPainterEvent::QtPainterEvent(QWidget* parent)
	: QMainWindow(parent)
{
	ui.setupUi(this);
}

QtPainterEvent::~QtPainterEvent()
{
}

void QtPainterEvent::paintEvent(QPaintEvent*)
{
	//QPainter painter(this);

	//QPen pen(QColor(255, 0, 0));
	//pen.setWidth(3);
	//pen.setStyle(Qt::DotLine);
	//// 设置画笔
	//painter.setPen(pen);

	//QBrush brush(Qt::cyan);
	//brush.setStyle(Qt::Dense7Pattern);

	//// 设置画刷
	//painter.setBrush(brush);

	//// 画线
	//painter.drawLine(QPoint(0, 0), QPoint(100, 100));

	//// 画圆
	//painter.drawEllipse(QPoint(100, 100), 50, 50);

	//// 画矩形
	//painter.drawRect(QRect(20, 20, 100, 100));

	//// 画字体
	//painter.drawText(QRect(10, 200, 100, 50), "别失恋了，来学习");

	QPainter painter(this);
	//painter.drawEllipse(QPoint(100, 500), 500, 500);
	//// 设置抗锯齿能力   效率较低
	//painter.setRenderHint(QPainter::HighQualityAntialiasing);
	//painter.drawEllipse(QPoint(200, 500), 500, 500);


	//painter.drawRect(QRect(20, 20, 50, 50));

	//// 转换坐标系
	//painter.translate(100, 0);
	//painter.drawRect(QRect(20, 20, 50, 50));

	//// 保存当前坐标系
	//painter.save();
	//painter.translate(100, 0);
	//// 还原坐标系
	//painter.restore();
	//painter.drawRect(QRect(20, 20, 50, 50));


	painter.drawPixmap(20, 0, QPixmap(":/igm.png"));
}
