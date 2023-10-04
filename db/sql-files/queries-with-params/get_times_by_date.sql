-- ATTENTION: the double percents (%%) is to mean a single percent symbol,
-- but it's double because if it's not "fmt.Sprintf" will try to interpret the
-- percent symbol ant then the behavior is unpredictable

-- Notes:
--    the first and the second ? will be replaced by an arbitrary date
--    %%d will be replaced by amount of elements per page
--    third ? will be replaced by the page * amount of elements per page
SELECT * FROM "time"
    WHERE ("repeat" & 32 = 32)
    OR (strftime('%%Y-%%m-%%d', "start") = strftime('%%Y-%%m-%%d', ?))
    OR ("repeat" = (1 << (strftime('%%w', ?) - 1)))
    OR (("repeat" & 64) = 64
        AND (strftime('%%w', "start") = strftime('%%w', ?)))
LIMIT %d OFFSET ?;
