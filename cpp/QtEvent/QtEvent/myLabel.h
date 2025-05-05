#pragma once
#pragma execution_character_set("utf-8")

#include <QLabel>

class myLabel : public QLabel
{
	Q_OBJECT

public:
	myLabel(QWidget* parent = nullptr);
	~myLabel();

	//// 鼠标进入事件
	//void enterEvent(QEvent* event);

	//// 鼠标的离开
	//void leaveEvent(QEvent* event);


	// 鼠标按下
	virtual void mousePressEvent(QMouseEvent* ev);

	// 鼠标释放
	virtual void mouseReleaseEvent(QMouseEvent* ev);

	// 鼠标移动
	virtual void mouseMoveEvent(QMouseEvent* ev);

signals:

public slots:
};
