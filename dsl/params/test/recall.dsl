rule "test_set_group"
begin
set_group(6, "后端")
set_group(7, "前端")
set_group(8, "前端")
set_group(9, "网络")
set_group(10, "推荐系统")
set_group(10, "后端")
end

rule "recall_i2i"
begin
recall_i2i(rc)
end

rule "recall_love_tag"
begin
recall_love_tag(rc)
end

rule "merge"
begin
recall_merge(rc)
end