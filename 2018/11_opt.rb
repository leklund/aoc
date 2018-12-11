require 'pry'
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

  # using the summed area table I can search the full range of square
  # sizes MUCH faster than using the sub matrix sums
  def max_power(range = nil)
    max = max_x = max_y = max_w = 0
    range ||= 2..cells.size
    sat = to_sat
    range.each do |width|
      (0..299 - width).each do |x|
        (0...299 - width).each do |y|
          power = summed_area(sat, x, x + width - 1, y, y + width - 1)
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

  # create a Summed area table
  def to_sat
    matrix = to_matrix
    rows = matrix.row_count
    cols = matrix.column_count
    sat = Matrix.zero(rows, cols).to_a

    rows.times do |i|
      cols.times do |j|
        x = i > 0 ? i : 0
        y = j > 0 ? j : 0
        a = matrix[x, y]

        b = sat.fetch(i - 1, []).fetch(j, 0)
        c = sat[i].fetch(j - 1, 0)

        d = i.zero? || j.zero? ? 0 : sat[i - 1][j - 1]

        sat[i][j] = a + b + c - d
      end
    end

    Matrix[*sat]
  end

  ##
  # Calculate summed area
  def summed_area(sat, x0, x1, y0, y1)
    sat ||= to_sat
    x0 = x0.zero? ? 0 : x0 - 1
    y0 = y0.zero? ? 0 : y0 - 1
    sat[x1, y1] + sat[x0, y0] - sat[x1, y0] - sat[x0, y1]
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

def test_area(serial, expected, range = nil)
  battery = build_battery(serial)

  actual = battery.max_power(range)

  abort "❌ expected #{actual} to eq #{expected}" unless actual == expected

  puts '✅'
end

test_battery(57, [122, 79], -5)
test_battery(39, [217, 196], 0)
test_battery(71, [101, 153], 4)
# width 3
test_area(18, [3, 33, 45], 3..3)
test_area(42, [3, 21, 61], 3..3)
# max width
test_area(18, [16, 90, 269])
test_area(42, [12, 232, 251])

input = 7139
battery = build_battery(input)
res = battery.max_power(3..3)
puts "w: #{res[0]}, x: #{res[1]}, y: #{res[2]}"

res = battery.max_power
puts "w: #{res[0]}, x: #{res[1]}, y: #{res[2]}"
