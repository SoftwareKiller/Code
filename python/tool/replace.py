import re
from datetime import datetime

def replace_variable(text, key, replacement):
    pattern = re.compile(r'^(.*' + key + '.*)\{dt\}(.*)$', re.MULTILINE)
    return pattern.sub(r'\1' + replacement + r'\2', text)

str = "from_unixtime(unix_timestamp(close_accounting_date,'yyyyMMdd'),'yyyy-MM-dd')= '{dt}')loan_account_tab"

close_acc_date = datetime.now().strftime("%Y-%m-%d")
update_txt = replace_variable(str, "close_accounting_date", close_acc_date)
print(update_txt)