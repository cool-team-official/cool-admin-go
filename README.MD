# CoolAdminGo
CoolAdmin的Go版本后端

开发中，未完成

## TODO

- [x] 基础CURD结构搭建
- [x] model层快速生成
- [x] service层快速生成
- [x] controller层快速生成
- [] 中间件错误状态及错误码与node版本统一
- [] 全局错误码统一
- [] controller层接口输入与node版本统一
- [] service层逻辑补全
- [] 其他项目需求
```bash
gentool \
-dsn "root:123456@(127.0.0.1:3306)/cool?charset=utf8mb4&parseTime=True&loc=Local" \
-db "mysql" \
-outPath "./temp/gentemp" \
-onlyModel \
-fieldNullable \
-fieldSignable \
-fieldWithIndexTag true\
-fieldWithTypeTag 
```