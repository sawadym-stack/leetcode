-- Write your PostgreSQL query statement below
SELECT  name FROM customer
WHERE referee_id IS null OR REFEREE_ID !=2;