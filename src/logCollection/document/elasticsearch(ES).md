# elasticsearch API介绍

1. 查看健康情况

   ```bash
   curl -X GET http://192.168.56.1:9200/_cat/health?v
   ```

2. 创建索引

   ```bash
   curl -X PUT http://192.168.56.1:9200/web
   ```

3. 查询当前es集群中所有的indices

   ```bash
   curl -X GET http://192.168.56.1:9200/_cat/indices?v
   ```

4. 删除索引

   ```bash
   curl -X DELETE http://192.168.56.1:9200/web
   ```

5. 往user索引里面添加person

   ```bash
   # 创建user索引
   curl -X PUT http://192.168.56.1:9200/user
   
   # 添加person
   curl -H "Content-type:application/json" -X POST http://192.168.56.1:9200/user/_doc/person1 -d '
   {
   "name":"zhaoyun",
   "age":38,
   "married":false
   }'
   
   curl -H "Content-type:application/json" -X POST http://192.168.56.1:9200/user/_doc/person2 -d '
   {
   "name":"kkite",
   "age":22,
   "married":false
   }'
   
   curl -H "Content-type:application/json" -X POST http://192.168.56.1:9200/user/_doc/person3 -d '
   {
   "name":"daisy",
   "age":3,
   "married":false
   }'
   ```

6. 查询user索引

   ```bash
   curl -X GET http://192.168.56.1:9200/user/_search
   ```

7. 条件查询

   ```bash
   curl -H "Content-type:application/json" -X GET http://192.168.56.1:9200/user/_search -d '
   {
   "query":{
   	"match":{"name":"daisy"}
   	}
   }'
   ```

   

