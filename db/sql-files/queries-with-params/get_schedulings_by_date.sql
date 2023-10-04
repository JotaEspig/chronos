SELECT * FROM "scheduling"
WHERE strftime('%%Y-%%m-%%d', "start") = strftime('%%Y-%%m-%%d', ?)
LIMIT %d OFFSET ?;
