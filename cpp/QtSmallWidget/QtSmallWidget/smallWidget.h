#pragma once
#include<QWidget>
namespace Ui {
	class SmallWidget;
}

class SmallWidget : public QWidget
{
	Q_OBJECT
public:
	SmallWidget(QWidget* parent = nullptr);
	~SmallWidget();
private:
	Ui::SmallWidget ui;
};
