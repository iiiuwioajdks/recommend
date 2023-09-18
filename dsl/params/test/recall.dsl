rule "test_set_group"
begin
set_group(1, "篮球")
set_group(2, "游戏")
set_group(3, "游泳")
set_group(4, "篮球")
set_group(5, "计算机")
end

rule "recall_i2i"
begin
recall_i2i(rc)
end

rule "recall_by_tag"
begin
recall_by_tag(rc)
end

rule "merge"
begin
println("run recall merge")
end