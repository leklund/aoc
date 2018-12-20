# I started by building out a tree structure
# where I would traverse the whole tree to find all
# the possible paths but then I thought about walking
# the path and was saying to myslef you go east then north
# then north and then we pop back to where we were and go
# west then north etc.

# It's also lucky that this even works. If the input had branches that
# were not dead ends this would never work. A simple input
# like '^E(N|S)E$' wouldn't work with this stack-based solution

# map the cardinal directions in to real/imaginary space
# A trick I learned from a previous day when
# reading solutions
def imaginary_compass
  {
    'W' => Complex(0, -1),
    'E' => Complex(0, 1),
    'N' => Complex(1, 0),
    'S' => Complex(-1, 0)
  }
end

def walk(input)
  locations = []
  current_location = Complex(0)
  doors = {
    current_location => 0
  }
  input.chars.each do |c|
    case c
    when /[EWNS]/
      new_location = current_location + imaginary_compass[c]
      doors[new_location] ||= doors[current_location] + 1

      current_location = new_location
    when '('
      locations.push current_location
    when '|'
      current_location = locations.last
    when ')'
      current_location = locations.pop
    end
  end
  doors
end

def max_dist(doors)
  doors.values.max
end

def count_rooms(doors, num_doors)
  doors.count { |_loc, dist| dist >= num_doors }
end

def test
  res = walk('^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$')
  puts "expected: 23, actual: #{max_dist(res)}"

  res = walk('^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$')
  puts "expected: 31, actual: #{max_dist(res)}"

  binding.pry
end

test

input = File.read('20.input').chomp

res = walk(input)

puts max_dist(res)
puts count_rooms(res, 1000)
