-- ("repeat" & 64 AND ((unixepoch("start") - unixepoch(?)) / (86400)) == 7)
-- statement checks if date is 7 days after "start"
-- Notes:
--    86400 is the amount of seconds in 7 days
--    ? will be replaced by an arbitrary date
--    %d will be replaced by amount of elements per page
--    second ? will be replaced by the page * amount of elements per page
SELECT * FROM "time"
WHERE "repeat" & 32 == 32
    OR ("repeat" & 64 AND ((unixepoch("start") - unixepoch(?)) / (86400)) == 7)
LIMIT %d OFFSET ?;
