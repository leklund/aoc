require 'pry'
require 'scanf'

# collection of points
class Map
  attr_accessor :points, :cursor

  SPRING = [500, 0]

  def initialize
    @points = {}
    # spring
    points[[500, 0]] = '+'
    @cursor = [SPRING]
  end

  def add(p, range, axis, char = '#')
    range.each do |n|
      coords = axis == :x ? [p, n] : [n, p]
      points[coords] = char
    end
  end

  def xy
    points.keys
  end

  def ymax
    @ymax ||= xy.map(&:last).max
  end

  def ymin
    @ymin ||= xy.map(&:last).sort[1]
  end

  def print_s
    xmin, xmax = points.keys.map { |k| k.first }.minmax
    ymin, ymax = points.keys.map { |k| k.last }.minmax

    (ymin..ymax).each  do |y|
      (xmin..xmax).each do |x|
        print points.fetch([x, y], '.')
      end
      print "\n"
    end
  end

  def flow
    return '_DONE_' if cursor.empty?

    curr = cursor.last
    if curr.last + 1 > ymax
      cursor.pop
      return
    end

    down = [curr.first, curr.last + 1]
    below = points[down]

    if below.nil?
      points[down] = '|'
      cursor << down
    elsif ['#', '~'].include?(below)
      # possible well. check bounds
      check_stream = below == '~'
      left_bound, right_bound, overflow = bounds(curr, check_stream)
      curry = curr.last

      char = overflow.any? ? '|' : '~'
      add(curry, (left_bound + 1)..(right_bound - 1), :y, char)

      # step back one step and add "springs" for the overflow points
      cursor.pop
      cursor.push([left_bound + 1, curry]) if overflow[0]
      cursor.push([right_bound - 1, curry]) if overflow[1]
    elsif below == '|'
      # check the row below
      if points[[curr.first, curr.last + 2]] != '|' && points[[curr.first + 1, curr.last + 2]] && points[[curr.first - 1, curr.last + 2]]
        left_bound, right_bound, overflow = bounds(curr)

        if overflow.none?
          add(curr.last, (left_bound + 1)..(right_bound - 1), :y, '~')
        end
      end
      # move back up the stream
      cursor.pop
    end
  end

  def bounds(coord, check_stream = true)
    x, y = coord
    row = xy.map { |p| p.first if p.last == y }.compact.sort
    nextrow = xy.map { |p| p.first if p.last == y + 1 }.compact.sort
    right_bound = left_bound = nil

    overflow = [false, false]
    if check_stream
      right_bound = row.find { |i| i > x && points[[i, y]] != '|' }
      left_bound = row.reverse.find { |i| i < x && points[[i, y]] != '|' }
    else
      right_bound = row.find { |i| i > x }
      left_bound = row.reverse.find { |i| i < x }
    end

    right_bound = nextrow.max + 1 if right_bound.nil?
    left_bound = nextrow.min - 1 if left_bound.nil?

    x.upto(right_bound) do |xx|
      below = points[[xx, y + 1]]
      point = points[[xx, y]]

      if below.nil? || (below == '|' && (point == '|' || point.nil?)) # overflow
        right_bound = xx + 1
        overflow[1] = true
        break
      elsif below == '#' || below == '~'
        right_bound = xx if points[[xx, y]] == '#'
      end
    end

    x.downto(left_bound) do |xx|
      below = points[[xx, y + 1]]
      point = points[[xx, y]]
  
      if below.nil? || (below == '|' && (point == '|' || point.nil?)) # overflow
        left_bound = xx - 1
        overflow[0] = true
        break
      elsif below == '#' || below == '~'
        left_bound = xx if points[[xx, y]] == '#'
      end
    end

    [left_bound, right_bound, overflow]
  end

  def unbounded?(coord)
    bounds(coord?).all(&:nil?)
  end

  def water_count(chars = ['|', '~'])
    points.select { |k, _v| k.last >= ymin && k.last <= ymax }.to_h.values.select { |v| chars.include?(v) }.size
  end
end

def run(map)
  done = false

  map.ymin
  map.ymax

  until done
    res = map.flow
    done = res == '_DONE_'
  end

  map.print_s
  puts "FLOWING: #{map.water_count}"
  puts "RETAINED: #{map.water_count(['~'])}"

end

def make_map(input)
  m = Map.new

  input.chomp.split("\n").each do |line|
    axis, pos, _xy, r1, r2 = line.scanf('%c=%d, %c=%d..%d')
    m.add(pos, r1..r2, axis.to_sym)
  end

  m
end

def test
  input =<<~INPUT
x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
  INPUT

  m = make_map(input)
  run(m)

  input2 = <<~INPUT
x=490, y=3..13
x=510, y=3..13
y=13, x=490..510
x=497, y=5..8
x=503, y=5..8
y=8, x=497..503
  INPUT

  m = make_map(input2)
  run(m)

  input3 = <<~INPUT
x=495, y=3..6
x=500, y=4..6
y=6, x=495..500
  INPUT
  m = make_map(input3)
  run(m)

  m = make_map(input3)
  m.cursor.unshift [495,0]
  run(m)
end

test

input = File.read('17.input')

m = make_map(input)

run(m)
