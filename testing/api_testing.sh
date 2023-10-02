#/bin/bash

# Test the /query?date= endpoint
echo "------------TEST1: query?date=2015-12-31 -----------"
curl -X GET "weather:8080/query?date=2015-12-31" >test1.out

# Test the /query?weather= endpoint
echo "------------TEST2: query?weather=rain -----------"
curl -X GET "weather:8080/query?weather=rain" >test2.out

# Test the /query?limit= endpoints
echo "------------TEST3 query?limit=2 -----------"
curl -X GET "weather:8080/query?limit=2" >test3.out

# Test the /query?weather= and /query?limit= endpoints
echo "------------TEST4 query?weather=sun&limit=2 -----------"
curl -X GET "weather:8080/query?weather=sun&limit=2" >test4.out

# Test the / endpoints
echo "------------TEST5 / -----------"
curl -X GET "weather:8080/" >test5.out
