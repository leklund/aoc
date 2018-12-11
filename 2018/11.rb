require 'matrix'

class Battery
  attr_accessor :cells

  def initialize
    self.cells = Matrix.zero(300).to_a
  end

  def add(x, y, seed)
    rack_id = x + 10
    val = (rack_id * y) + seed
    val *= rack_id
    val = val.to_s[-3].to_i - 5

    cells[x][y] = val
  end

  def [](x,y)
    cells[x][y]
  end

  def max_power(range = 3..3)
    max = max_x = max_y = max_w = 0

    range.each do |width|
      cell_matrix = to_matrix
      (0..299 - width).each do |x|
        (0...299 - width).each do |y|
          power = cell_matrix.minor(x..(x + width - 1), y..(y + width - 1)).sum
          if power > max
            max = power
            max_x = x
            max_y = y
            max_w = width
          end
        end
      end
    end
    [max_w, max_x, max_y]
  end

  def to_matrix
    Matrix[*cells]
  end
end

def build_battery(seed)
  b = Battery.new
  300.times do |x|
    300.times do |y|
      b.add(x, y, seed)
    end
  end
  b
end

def test_battery(serial, cell, power)
  battery = build_battery(serial)

  actual = battery[*cell]
  expected = power

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts '✅'
end

def test_area(serial, expected, range = 3..3)
  battery = build_battery(serial)

  actual = battery.max_power(range)

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts '✅'
end

test_battery(57, [122, 79], -5)
test_battery(39, [217, 196], 0)
test_battery(71, [101, 153], 4)
# width 3
test_area(18, [3, 33, 45])
test_area(42, [3, 21, 61])
# max width
test_area(18, [16, 90, 269], 3..20)
test_area(42, [12, 232, 251], 3..20)

input = 7139
battery = build_battery(input)
res = battery.max_power
puts "w: #{res[0]}, x: #{res[1]}, y: #{res[2]}"
res = battery.max_power(range = 3..20)
puts "w: #{res[0]}, x: #{res[1]}, y: #{res[2]}"


