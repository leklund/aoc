class Cave
  attr_accessor :map, :units

  def initialize(m)
    @map = m.clone
    @units = []
  end

  def width
    @width ||= map.index("\n") + 1
  end

  def adjacent(position)
    [
      position - width, # Up
      position - 1,     # Left
      position + 1,     # Right
      position + width  # Down
    ]
  end

  def adjacent_open(position)
    adjacent(position).select { |i| map[i] == '.' }
  end

  def elf_count
    units.select(&:elf).size
  end

  def turn
    # each unit
    # move && attack
    units.sort_by(&:to_i).each do |unit|
      next if unit.hp <= 0

      return true if units.map(&:to_s).uniq.size == 1

      unit.move && unit.attack
    end

    false
  end
end

class Unit
  attr_accessor :position, :hp, :power, :elf, :cave

  def initialize(i, hp, power, elf, cave)
    @position = i # index in the array
    @hp = hp
    @power = power
    @elf = elf # bool
    @cave = cave
  end

  def width
    cave.width
  end

  def adjacent
    [
      position - width, # Up
      position - 1,     # Left
      position + 1,     # Right
      position + width  # Down
    ]
  end

  def already_in_range?
    !active_target.nil?
  end

  def active_target
    cave.units.sort_by(&:to_i).select { |u| u.elf != elf && adjacent.include?(u.position) }&.sort_by { |target| [target.hp, target.to_i] }&.first
  end

  def attack
    return unless active_target
    active_target.hp -= power
    if active_target.hp <= 0
      cave.map[active_target.to_i] = '.'
      cave.units.delete(active_target)
    end
  end

  def move
    return true if already_in_range?

    #array of units of opposite type
    targets = cave.units.reject { |u| elf == u.elf }

    # array of in_range index positions
    in_range = targets.map { |target| target.open_steps }.flatten.uniq

    # array of possible next steps [next_step, distance, target_i]
    next_steps = in_range.map { |target| next_step(target) }

    return false if next_steps.empty?

    next_posistion = next_steps.compact.sort.first&.last
    return true if next_posistion.nil?

    cave.map[position] = '.'
    cave.map[next_posistion] = to_s
    self.position = next_posistion
    true
  end

  def open?
    !open_steps.empty?
  end

  def open_steps
    adjacent.select { |i| cave.map[i] == '.' }
  end

  def next_step(target)
    # walk back from target to self and return
    #   [distance, target, next_step]
    # returned in this order to make sorting easier
    possible_steps = open_steps
    d = 1
    return [d, target, target] if possible_steps.include? target

    visited = [target]
    step_back = target
    d += 1
    path = cave.adjacent_open(step_back).reject { |i| visited.include?(i) }
    branch = []

    loop do
      # walk it back now
      if path.empty?
        # dead end
        break if branch.empty?

        d += 1

        path = branch
        branch = []
      end

      step_back = path.shift

      return [d, target, step_back] if possible_steps.include?(step_back)

      visited << step_back

      branch |= cave.adjacent_open(step_back) - visited
    end

    nil
  end

  def to_i
    position
  end

  def to_s
    elf ? 'E' : 'G'
  end
end

HP = 200
GPOWER = 3

def run(input)
  (4..1000).each do |elfpower|
    cave = Cave.new(input)

    cave.map.chars.each_with_index do |c, i|
      power = c == 'E' ? elfpower : GPOWER
      cave.units << Unit.new(i, HP, power, c == 'E', cave) if c.match?(/G|E/)
    end

    required_elves = cave.elf_count

    i = 0
    done = false
    until done do
      done = cave.turn
      done = true if cave.elf_count != required_elves
      i += 1 unless done
    end
    if cave.elf_count == required_elves
      return cave.units.map(&:hp).sum * i
    end
  end
end

def sample_inputs
  expected = [4988, 31284, 3478, 6474, 1140]
  samples =<<~INPUT
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
-----
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
-----
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
-----
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
-----
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
INPUT
  samples.split("-----\n").each_with_index do |sample, i|
    actual = run(sample)
    puts "#{actual} - #{expected[i]}"
  end
end

sample_inputs

input = DATA.read

puts run(input)

__END__
################################
####################.###########
###################..##..#######
###################.###..#######
###################.###.########
##################G.###.########
##..############.#..##..########
#...#####.####.....##...########
#G.....##..###.#.GG#...G.....G##
#...G....G.G##.....#.#.......###
##..............G............###
#......GG.....G............#####
#####.........#####.......E..###
#####.....G..#######.......#####
#####.......#########......###.#
######......#########..........#
#####.....G.#########..........#
##.....#...G#########......#...#
###.#....G..#########G.........#
##..###......#######E..........#
##..###...G...#####E..E.....E..#
###..##...............E.#.E....#
###..##..##E...###......##..####
##..G...###....###......########
##......###E....##....##########
###...#####..E..###...##########
##....####......####.###########
####..#####.....##...###########
############..####....##########
############.#####....##########
##########...#####.#.###########
################################
