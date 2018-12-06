class Point
  attr_accessor :x, :y, :area, :unbounded

  def initialize(x:, y:)
    @x = x
    @y = y
    @area = 0
    @unbounded = false
  end

  def coords
    [x,y]
  end

  def dist(xx, yy)
    (x - xx).abs + (y - yy).abs
  end
end

class Graph
  attr_accessor :corners, :points
  attr_reader :xmin, :xmax, :ymin, :ymax

  def initialize
    @points = {}
  end

  def add(point)
    points[point.coords] = point
  end

  def [](coords)
    points[coords]
  end

  # brute force nearest neighbor
  def nn(x,y)
    measured = points.each_with_object(Hash.new { |h, k| h[k] = [] }) do |(coords, point), h|
      next if point.coords == [x, y]

      dist = point.dist(x, y)

      h[dist] << coords
    end

    measured.min_by { |k, _v| k }[1]
  end

  # set x/y min/max
  def set_bounds
    @xmin, @xmax = points.values.minmax_by(&:x).map(&:x)
    @ymin, @ymax = points.values.minmax_by(&:y).map(&:y)

    points.values.each { |p| p.unbounded = true if [xmin, xmax].include?(p.x) || [ymin, ymax].include?(p.y) }
  end

  def bounded
    points.select { |_k, v| v.unbounded == false }
  end

  def distance_to_all_points(x:, y:, max_dist:)
    points.map { |_x, point|
      dist = point.dist(x, y)
      return max_dist if dist >= max_dist

      dist
    }.compact.sum
  end
end

def run_part_one(input)
  graph = Graph.new
  input.each do |p|
    x, y = p.split(', ')
    graph.add(Point.new(x: Integer(x), y: Integer(y)))
  end

  graph.set_bounds

  # brute force the whole graph from the corners
  (graph.xmin..graph.xmax).each do |x|
    (graph.ymin..graph.ymax).each do |y|
      edge = [graph.xmin, graph.xmax].include?(x) || [graph.ymin, graph.ymax].include?(y)
      next if graph[[x, y]]

      nn = graph.nn(x, y)

      next if nn.size != 1

      # if we are on an edge and we get back on neighor, mark it unbounded
      nn.each { |coords| graph[coords].unbounded = true } if edge
      graph[nn.first].area += 1
    end
  end

  point = graph.bounded.values.max_by(&:area)

  point.area + 1
end

def run_part_two(input, distance)
  graph = Graph.new
  input.each do |p|
    x, y = p.split(', ')
    graph.add(Point.new(x: Integer(x), y: Integer(y)))
  end

  graph.set_bounds

  num_points = graph.points.size
  area = 0

  # Brute force every point inside the corners of the graph.
  # Depending on the input distance this is not guaranteed to return
  # the correct answer.
  (graph.xmin..graph.xmax).each do |x|
    (graph.ymin..graph.ymax).each do |y|
      area += 1 if graph.distance_to_all_points(x: x, y: y, max_dist: distance) < distance
    end
  end

  area
end

def test_part_one
  input = <<~INPUT
    1, 1
    1, 6
    8, 3
    3, 4
    5, 5
    8, 9
  INPUT

  actual = run_part_one(input.split("\n"))
  expected = 17

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts "✅"
end

def test_part_two
  input = <<~INPUT
    1, 1
    1, 6
    8, 3
    3, 4
    5, 5
    8, 9
  INPUT

  actual = run_part_two(input.split("\n"), 32)
  expected = 16

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts "✅"
end

test_part_one

data = DATA.read.chomp

out = run_part_one(data.split("\n"))
puts "part one: #{out}"

test_part_two

out = run_part_two(data.split("\n"), 10_000)
puts "part two: #{out}"

__END__
315, 342
59, 106
44, 207
52, 81
139, 207
93, 135
152, 187
271, 47
223, 342
50, 255
332, 68
322, 64
250, 72
165, 209
129, 350
139, 118
282, 129
311, 264
216, 246
134, 42
66, 151
263, 199
222, 169
236, 212
320, 178
202, 288
273, 190
83, 153
88, 156
284, 305
131, 90
152, 88
358, 346
272, 248
317, 122
166, 179
301, 307
156, 128
261, 290
268, 312
89, 53
324, 173
353, 177
91, 69
303, 164
40, 221
146, 344
61, 314
319, 224
98, 143
