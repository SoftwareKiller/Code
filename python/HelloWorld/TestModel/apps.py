from django.apps import AppConfig


class TestmodelConfig(AppConfig):
    default_auto_field = 'django.db.models.BigAutoField'
    name = 'TestModel'
    def ready(self):
        # 启动定时任务调度器
        from .task import scheduler
        print("Starting scheduler...")
        scheduler.start()
