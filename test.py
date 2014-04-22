import json
#a = {"id": 1, "name": "a", "content": {"address": "add", "zip": "z"}}
# b = {"name": "a"}
# c = json.loads(b)
# json.loads({"uuid":"f4609dcc2109ecc786d545697ebfcc6954992ddc","language":"en","device_uuid":"9f4c28da3a6785237003dfb7b35981b7a68a7055","country":"us","app":"a5e7529567987ddf386b070-ed35c5b8-df42-11e2-8ba2-005cf8cbabd8","nth":44,"device_new":false,"platform":"iPhone","carrier":"Verizon Wireless","app_ver":"1.0.3","at":"2014-03-09 23:17:19","session_uuid":"66ca17ae4865d9499238e81f1595c7df77ea2e90","library_ver":"iOSa_2.21.1","install_uuid":"07a9047c5369327444ff876dc9f0e8c332a38a48","model":"iPhone 5S","custom":{"NumberObjects":"7","ProcessingTimeMS":"1"},"type":"e","os_ver":"7.0.6","name":"SHTDecodeResults"})
# print c["at"]

a = json.dumps({"id": 1, "name": "a"})
print 'a (json)'
print a
b = json.loads(a)
print 'b (dictionary)'
print b
# c = json.loads({"uuid":"f4609dcc2109ecc786d545697ebfcc6954992ddc","language":"en","device_uuid":"9f4c28da3a6785237003dfb7b35981b7a68a7055","country":"us","app":"a5e7529567987ddf386b070-ed35c5b8-df42-11e2-8ba2-005cf8cbabd8","nth":44,"platform":"iPhone","carrier":"Verizon Wireless","app_ver":"1.0.3","at":"2014-03-09 23:17:19","session_uuid":"66ca17ae4865d9499238e81f1595c7df77ea2e90","library_ver":"iOSa_2.21.1","install_uuid":"07a9047c5369327444ff876dc9f0e8c332a38a48","model":"iPhone 5S","custom":{"NumberObjects":"7","ProcessingTimeMS":"1"},"type":"e","os_ver":"7.0.6","name":"SHTDecodeResults"})
c = json.loads({"type":"e"})
print 'c (dictionary)'
print c
