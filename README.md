# sqlstruct

# 缘由 
从https://github.com/go-programming-tour-book/tour继承过来,但是根据实际开发中的需求，支持了以下特性

# 功能
### 不但支持表，而且支持整库, 通过不指定--table就是整库 
### 支持可空字段的映射 Bool Int Time Decimal Float String
### 支持导出到文件,可以通过--file 和--package 选项来指定

# 其他
### 可空字段的另外一种解决方案[sql_null](https://github.com/15125505/doc/blob/master/go/sql_null.md)，其实在开发中，尽量不要去设置字段可空，但是对于遗留系统，我们也往往不能控制这种情况发生

