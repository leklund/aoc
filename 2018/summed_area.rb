require 'matrix'

class Matrix
  def summed_area_table
    return @sat if @sat
    rows = row_count
    cols = column_count

    sat = Matrix.zero(rows, cols).to_a

    rows.times do |i|
      cols.times do |j|
        x = i > 0 ? i : 0
        y = j > 0 ? j : 0
        a = self[x, y]

        b = sat.fetch(i - 1, []).fetch(j, 0)
        c = sat[i].fetch(j - 1, 0)

        d = i.zero? || j.zero? ? 0 : sat[i - 1][j - 1]

        sat[i][j] = a + b + c - d
      end
    end

    @sat = Matrix[*sat]
  end

  def summed_area(x0, y0, x1, y1)
    sat = summed_area_table
    x0 = x0.zero? ? 0 : x0 - 1
    y0 = y0.zero? ? 0 : y0 - 1
    sat[x1, y1] + sat[x0, y0] - sat[x1, y0] - sat[x0, y1]
  end
end
