---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by saeipi.
--- DateTime: 2023/7/21 2:20 PM
---
---

local product_key = KEYS[1]
local total_num_key = ARGV[1]
local sold_num_key = ARGV[2]

local values = redis.call('HMGET', product_key, total_num_key, sold_num_key)
if #values ~= 2 then
    return 0
end

local total_num = tonumber(values[1])
local sold_num = tonumber(values[2])
if total_num > sold_num then
    return redis.call('HINCRBY', product_key, sold_num_key, 1)
end
return 0