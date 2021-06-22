## 基于mysql
### 1.注释的三种方式 

    #方式1

    --方式2 

    /*
        方式三
    */

---

### 2.创建表 

     CREATE TABLE mytable (
    # int 类型，不为空，自增
    id INT NOT NULL AUTO_INCREMENT,
    # int 类型，不可为空，默认值为 1，不为空
    col1 INT NOT NULL DEFAULT 1,
    # 变长字符串类型，最长为 45 个字符，可以为空
    col2 VARCHAR(45) NULL,
    # 日期类型，可为空
    col3 DATE NULL,
    # 设置主键为 id
    PRIMARY KEY (id));

---

### 3.修改表

    添加列

        ALTER TABLE mytable
        ADD col CHAR(20);

    删除列
    
        ALTER TABLE mytable
        DROP COLUMN col;

    删除表
    
     DROP TABLE mytable;

---

### 4.插入数据

    普通插入
    
        INSERT INTO mytable(col1, col2)
        VALUES(val1, val2);

    将一个表的内容插入到一个新表
    
        CREATE TABLE newtable AS
        SELECT * FROM mytable;

---

### 5.更新

    UPDATE mytable
    SET col = val
    WHERE id = 1;

---

### 6.删除

    DELETE FROM mytable
    WHERE id = 1;

---

### 7.查询
    
    distinct 
        
        #作用于所有列,即当所有列相同时才算相同 
        SELECT DISTINCT col1, col2
        FROM mytable;

    limit

        # limit 有两个参数:第一个参数为起始行,从0开始;第二个参数为返回的总行数
        SELECT *
        FROM mytable
        LIMIT 5;

        SELECT *
        FROM mytable
        LIMIT 0, 5;

---

### 8.排序

    asc:升序
    desc:降序 

    SELECT *
    FROM mytable
    ORDER BY col1 DESC, col2 ASC;

---

### 9.过滤(where子句)

    SELECT *
    FROM mytable
    WHERE col IS NULL;
    
---

### 10.通配符

    % 匹配 >=0 个任意字符；

    _ 匹配 ==1 个任意字符；

    [ ] 可以匹配集合内的字符;例如 [ab] 将匹配字符 a 或者 b。用脱字符 ^ 可以对其进行否定，也就是不匹配集合内的字符。

    SELECT *
    FROM mytable
    WHERE col LIKE '[^AB]%'; -- 不以 A 和 B 开头的任意文本

---

### 13.分组(group by子句)

    where过滤行,having过滤分组,行过滤优先于分组过滤

    group by子句出现在where子句之后,order by子句之前 

    除了汇总字段外,select语句的每一个字段都必须在group by子句中出现

    例:
        SELECT col, COUNT(*) AS num
        FROM mytable
        WHERE col > 2
        GROUP BY col
        HAVING num >= 2;

---

### 14.子查询(in,not in)

---

### 15.连接

    包含内连接、外连接和自然连接等。

    inner join:两个表的交集

    left join/right join:

    full join:全集

---

### 16.组合查询(union)

    使用 UNION 来组合两个查询，如果第一个查询返回 M 行，第二个查询返回 N 行，那么组合查询的结果一般为 M+N 行。

    每个查询必须包含相同的列、表达式和聚集函数。
    
    默认会去除相同行，如果需要保留相同行，使用 UNION ALL。

    SELECT col
    FROM mytable
    WHERE col = 1
    UNION
    SELECT col
    FROM mytable
    WHERE col =2;

---

### 17.视图

    把视图当作是对sql语句的封装
    
    格式:
        create view viewName as 
            sql语句;

---

### 18.存储过程

    看作一系列语句的批处理.
    