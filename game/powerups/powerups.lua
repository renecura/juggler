local M = {}

M.data = {
	{id = "speedup_1", title = "Speed Up", minimum_level = 0, power = 1, rarity = 1},
	{id = "extra_ball_1", title = "Extra Ball", minimum_level = 0, power = 1, rarity = 1},
	{id = "multiplier_1", title = "Multi", minimum_level = 0, power = 1, rarity = 1}
}

M.by_id = {}
for _, p in ipairs(M.data) do M.by_id[p.id] = p end

function M.picktwo()
	local pool = {}
	for i, p in ipairs(M.data) do pool[i] = p end
	local a = table.remove(pool, math.random(#pool))
	local b = table.remove(pool, math.random(#pool))
	return a.id, b.id
end

return M