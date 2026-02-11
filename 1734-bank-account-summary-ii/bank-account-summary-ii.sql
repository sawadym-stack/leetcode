# Write your MySQL query statement below
SELECT name,
SUM(t.amount) AS balance
FROM Transactions t
JOIN Users u
ON u.account=t.account
GROUP BY u.account,u.name
HAVING SUM(t.amount)>10000;