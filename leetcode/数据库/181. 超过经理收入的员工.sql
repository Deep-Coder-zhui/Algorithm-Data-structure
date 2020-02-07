# Write your MySQL query statement below
# 两张表的连接
# 274ms + 26.33%
select E1.Name as Employee
from Employee as E1, Employee as E2 
where E1.ManagerId = E2.Id and E1.Salary > E2.Salary
# 267ms + 34.54%，先过滤掉Employee表中ManagerId为空的行得到新的表tmp。。。
select tmp.Name as Employee from (select * from Employee where Employee.ManagerId <> "")tmp, Employee 
where tmp.ManagerId = Employee.Id and tmp.Salary > Employee.Salary

# left join(左联接) on 返回包括左表中的所有和右表中联结字段相等的记录。
# 题目描述里的sql架构，里边的建表语句没有把id创建为主键，所有此处用distinct可提速。
select 
    Employee.Name as Employee
from Employee left join (select distinct Id,Salary from Employee) m  
    on Employee.ManagerId  = m.Id 
where Employee.Salary  > m.Salary 
