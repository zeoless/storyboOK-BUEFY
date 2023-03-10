
SELECT abs(-17.4);
SELECT cbrt(27.0);
SELECT ceil(-42.8);
SELECT ceiling(-95.3);
SELECT degrees(0.5);
SELECT div(9,4);
SELECT exp(1.0);
SELECT floor(-42.8);
SELECT ln(2.0);
SELECT log(100.0);
SELECT log(2.0, 64.0);
SELECT mod(9,4);
SELECT pi();
SELECT power(9.0, 3.0);
SELECT radians(45.0);
SELECT round(42.4);
SELECT round(42.4382, 2);
SELECT scale(8.41);
SELECT sign(-8.4);
SELECT sqrt(2.0);
SELECT trunc(42.8);
SELECT trunc(42.4382, 2);
SELECT width_bucket(5.35, 0.024, 10.06, 5);
SELECT width_bucket(now(), array['yesterday', 'today', 'tomorrow']::timestamptz[]);
create schema if not exists sqlc;