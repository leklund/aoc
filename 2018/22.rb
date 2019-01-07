require 'pry'
require 'matrix'

MOD = 20183
DEPTH = 9465
TARGET = Complex(13, 704)
CLIMB = 1
NONE = 0
TORCH = 2
ROCKY = 0
WET = 1
NARROW = 2
TOOLS = [NONE, CLIMB, TORCH]

@gim = {}
@elm = {}

def gi(coord)
  x = coord.real
  y = coord.imag
  return @gim[coord] if @gim[coord]

  return @gim[coord] = 0 if x == 0 && y == 0
  return @gim[coord] = 0 if TARGET == coord
  return @gim[coord] = x * 16807 if y == 0
  return @gim[coord] = y * 48271 if x == 0

  @gim[coord] = el(Complex(x - 1, y)) * el(Complex(x, y - 1))
end

def el(coord)
  return @elm[coord] if @elm[coord]

  @elm[coord] = (gi(coord) + DEPTH) % 20183
end

def type(coord)
  el(coord) % 3
end

# tools: 0 none, 1 climbing gear, 2 torch
# types: 0 rocky 1 wet, 2 narrow
# valid:
#   rocky (0)  - climbing (1) / torch (2)
#   wet (1)    - climbing (1) / none (0)
#   narrow (2) - torch (2) / none (0)

def valid_tool?(dest, tool)
  dest_type = type(dest)
  case dest_type
  when ROCKY
    [CLIMB, TORCH].include? tool
  when WET
    [CLIMB, NONE].include? tool
  when NARROW
    [TORCH, NONE].include? tool
  end
end

# array of moves
# [[dest_coords, tool, cost]]
def moves(coord_tool)
  coord = coord_tool.first
  tool = coord_tool.last

  valid_moves = []

  dirs = [Complex(1, 0), Complex(-1, 0), Complex(0, 1), Complex(0, -1)]

  # move
  dirs.each do |dir|
    dest = coord + dir

    next if dest.real < 0 || dest.imag < 0
    next unless valid_tool?(dest, tool)

    valid_moves.push([[dest, tool], 1])
  end

  # switch tool
  TOOLS.each do |dest_tool|
    next if dest_tool == tool
    next unless valid_tool?(coord, dest_tool)

    valid_moves.push([[coord, dest_tool], 7])
  end

  valid_moves
end


# PART ONE
sum = 0
(0..TARGET.imag).each do |y|
  (0..TARGET.real).each do |x|
    coord = Complex(x, y)
    modu = type(coord)
    sum += modu
    case modu
    when 0
      print '.'
    when 1
      print '='
    when 2
      print '|'
    end
  end
  print "\n"
end
print "\n\n"

puts sum


###
# Part Two
# Implement an A* search with a PriorityQueue
require 'priority_queue'

# tools
# 0: none
# 1: climbing gear
# 2: torch

# a star search heurstic
def heuristic(current:, target:)
  d = (current[0].real - target[0].real).abs + (current[0].imag - target[0].imag).abs
  d += 7 if d <= 8 && current[1] != target[1]

  d
end

# origin and target are arrays [coord, tool]
def astar_search(origin:, target:)
  costs = {}
  seen = {}
  path = {}
  initial_estimate = heuristic(current: origin, target: target)
  costs[origin] = 0

  pq = PriorityQueue.new

  pq.push(origin, initial_estimate)

  loop do
    current = pq.delete_min_return_key

    # made it
    return path if current == target

    seen[current] = true

    # get moves and estimate costs
    moves(current).each do |move|
      move_pair = move[0]
      cost = move[1]

      next if seen[move_pair]

      new_cost = costs[current] + cost

      # skip if more costly
      next if pq.has_key?(move_pair) && !costs[move_pair].nil? && new_cost >= costs[move_pair]

      path[move_pair] = current
      costs[move_pair] = new_cost
      estimate = new_cost + heuristic(current: move_pair, target: target)

      if pq.has_key?(move_pair)
        pq.change_priority(move_pair, estimate)
      else
        pq.push(move_pair, estimate)
      end
    end
  end
end

origin = Complex(0,0)
out = astar_search(origin: [origin, TORCH], target: [TARGET, TORCH])

# walk back
t = [TARGET, TORCH]
path = [t]
while out.key?(t)
  t = out[t]
  path.push(t)
end

# calculate the time for each step
prev = nil
t = 0

path.each do |loc|
  if prev.nil?
    prev = loc
    next
  end

  if prev[1] != loc[1]
    t += 7
  else
    t += 1
  end
  prev = loc
end

puts "PART TWO: #{t}"
