local M = {}

function M.reset()
	M.level = 1
	M.score = 0
	M.score_target = 3

	M.balls = 1
	M.speed = 2.0
	M.multi = 1.0
	
	M.powerups = {}      -- id -> { level = n }
end

function M.add_powerup(id, upgraded)
	local p = M.powerups[id] or { level = 0 }
	p.level = p.level + (upgraded and 2 or 1)
	M.powerups[id] = p
end

M.reset()

return M