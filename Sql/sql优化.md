# sql 优化

**原则：**

1. 减少数据访问：设置合理字段类型， 通过索引访问
2. 返回更少的数据：只返回需要的字段
3. 减少交互次数： 批量操作
4. 减少CPU开销： 尽量减少全表查询



**总结**

1. 最大化利用索引
2. 尽可能的避免全表扫描
3. 减少无效数据的查询



### 避免不走索引场景

1. 避免在字段开头使用模糊查询
2. 避免使用in ,not in 可使用（between, exists）替代
3. 避免使用or 可使用（union）
4. 避免在where条件等号左侧进行计算
5. 当数据量大的时候， 避免使用where 1=1 ，这样回进行全表扫描



### select 优化

1. 避免出现select * 

2. 调整where条件的连接顺序（把过滤数据大的放前面）

   



