require 'matrix'

class Matrix
  def summed_area_table
    return @sat if @sat
    rows = row_count
    cols = column_count

    sat = Matrix.zero(rows, cols).to_a

    rows.times do |x|
      cols.times do |y|
        a = self[x, y]

        b = sat.fetch(x - 1, []).fetch(y, 0)
        c = sat[x].fetch(y - 1, 0)

        d = x.zero? || y.zero? ? 0 : sat[x - 1][y - 1]

        sat[x][y] = a + b + c - d
      end
    end

    @sat = Matrix[*sat]
  end

  def summed_area(x0, y0, x1, y1)
    x0 = x0.zero? ? 0 : x0 - 1
    y0 = y0.zero? ? 0 : y0 - 1
    summed_area_table[x1, y1] +
      summed_area_table[x0, y0] -
      summed_area_table[x1, y0] -
      summed_area_table[x0, y1]
  end
end
