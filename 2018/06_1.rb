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
end

def run(input)
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

def test
  input = <<~INPUT
    1, 1
    1, 6
    8, 3
    3, 4
    5, 5
    8, 9
  INPUT

  actual = run(input.split("\n"))
  expected = 17

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts "✅"
end

test

data = DATA.read.chomp

out = run(data.split("\n"))
puts "answer: #{out}"

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
