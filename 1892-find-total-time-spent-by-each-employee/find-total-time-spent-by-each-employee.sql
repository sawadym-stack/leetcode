-- Write your PostgreSQL query statement below
SELECT event_day AS day,
       emp_id,
       SUM(OUT_TIME-IN_TIME) AS total_time
       FROM Employees
       GROUP BY event_day,emp_id;