-- Write your PostgreSQL query statement below
SELECT e.name AS Employee
FROM Employee e
JOIN Employee m
ON e.managerid = m.id
where e.salary>m.salary;