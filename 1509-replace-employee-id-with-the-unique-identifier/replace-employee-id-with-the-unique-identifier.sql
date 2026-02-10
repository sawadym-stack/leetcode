# Write your MySQL query statement below
SELECT unique_id,name 
FROM EmployeeUNI en
RIGHT JOIN Employees e
ON en.id=e.id;