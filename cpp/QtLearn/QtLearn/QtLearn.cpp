#include "QtLearn.h"
#include <QPushButton>

QtLearn::QtLearn(QWidget* parent)
	: QWidget(parent)
{
	//resize(500, 200);
	//m_lableInfo = new QLabel(tr("<h1>Hello Widget!</h1>"), this);
	//m_lableInfo->setGeometry(20, 20, 300, 40);
	////ui.setupUi(this);
	QPushButton* btn = new QPushButton(parent);
	//btn->show(); // show 以顶层的方式弹出窗口控件
	btn->setParent(this);
	btn->setText("First Button");

	QPushButton* btn2 = new QPushButton("Second Button", this);

	btn2->move(100, 100);
	btn2->resize(50, 50);

	resize(600, 400);

	setFixedSize(600, 400);

	setWindowTitle("This is a title");

	// 
	//connect(btn, &QPushButton::clicked, this, &QWidget::close);

	this->t = new Teacher();
	this->st = new Student();

	void(Teacher:: * teacherHungryNoParam)() = &Teacher::hungry;
	void(Student:: * studentTreatNoParam)() = &Student::treat;
	void(Teacher:: * teacherHungry)(QString) = &Teacher::hungry;
	void(Student:: * studentTreat)(QString) = &Student::treat;

	connect(t, teacherHungryNoParam, st, studentTreatNoParam);
	connect(t, teacherHungry, st, studentTreat);

	afterClass();
	connect(btn, &QPushButton::clicked, t, teacherHungryNoParam);
}

QtLearn::~QtLearn()
{
	delete m_lableInfo;
	m_lableInfo = nullptr;
}

void QtLearn::afterClass()
{
	emit t->hungry();
	emit t->hungry("牛肉凉面");
}

