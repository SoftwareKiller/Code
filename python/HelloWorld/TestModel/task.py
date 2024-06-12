from apscheduler.schedulers.background import BackgroundScheduler
from django_apscheduler.jobstores import DjangoJobStore, register_job

def my_task_function():
    # 这里写你的任务逻辑
    print("定时任务执行了！")

# 实例化调度器
scheduler = BackgroundScheduler()
# 使用DjangoJobStore作为作业存储
scheduler.add_jobstore(DjangoJobStore(), 'default')
# 添加任务
register_job(scheduler=scheduler, func=my_task_function, trigger='interval', seconds=10)