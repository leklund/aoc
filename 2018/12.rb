PAD = Array.new(4, 0)

##
# When I got to part two and the number of steps went from 20 to 5e9 I realized there
# was no way my simple loop would return in a reasonable time. I hoped that like Conway's
# life that it would reach a stable configuration at some point so I started by printing each step
# and it did become stable very fast. Only 193 steps in the case of my input. So then
# it was just matter of tracking the delta and multiplying by the reamining steps to get
# the final count.

def run(arr, rules, steps = 20)
  initial_size = arr.size
  current_count = 0
  deltas = Array.new(10, rand)
  completed_steps = steps.times do |step|
    arr = pad(arr)

    arr = arr.each_with_object([]).with_index do |(_e, a), i|
      # skip padded elements
      if i == 0 || i == 1 || 
         i == arr.size - 1  || i == arr.size - 2
        a << 0
        next
      end

      state = arr.slice(i - 2, 5).join.to_i(2)

      a << rules.fetch(state, 0)
    end

    count = 0
    offset = (arr.size - initial_size) / 2
    arr.each_with_index { |e, i| count += (i - offset) if e == 1 }
    deltas.shift && deltas.push(count - current_count)
    current_count = count

    break step if deltas.uniq.size == 1
  end

  offset = (arr.size - initial_size) / 2
  count = 0
  arr.each_with_index { |e, i| count += (i - offset) if e == 1 }

  if completed_steps != steps # we reached a stable state before all steps ran
    count += (deltas.first * (steps - completed_steps - 1))
  end

  [arr, count]
end

def pad(arr)
  # This might add some extra 0s to one side but meh
  return arr if arr[0, 4] == PAD && arr[-4, 4] == PAD

  PAD + arr + PAD
end

def parse(lines)
  initial_state = lines.shift
  arr = initial_state.match(/: (.*)/)[1].tr('#','1').tr('.','0').chars.map(&:to_i)

  lines.shift # blank

  rules = lines.each_with_object({}) do |line, h|
    key = line[0, 5].tr('#', '1').tr('.', '0').to_i(2)
    h[key] = line[-1].tr('#', '1').tr('.', '0').to_i
  end

  [arr, rules]
end

def test
  input =<<~INPUT
initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
  INPUT

  arr, rules = parse(input.split("\n"))
  final_state, count = run(arr, rules)

  expected = 325
  abort "❌ expected #{count} to eq #{expected}" unless count == expected

  puts '✅'
end

test

data = DATA.read.chomp

arr, rules = parse(data.split("\n"))
final_state, count = run(arr, rules)
puts final_state.join
puts count

_, count = run(arr, rules, 50000000000)
puts count

__END__
initial state: ##.#..########..##..#..##.....##..###.####.###.##.###...###.##..#.##...#.#.#...###..###.###.#.#

####. => #
##.#. => .
.##.# => .
..##. => .
..... => .
.#.#. => #
.###. => .
.#.## => .
#.#.# => .
.#... => #
#..#. => #
....# => .
###.. => .
##..# => #
#..## => #
..#.. => .
##### => .
.#### => #
#.##. => #
#.### => #
...#. => .
###.# => .
#.#.. => #
##... => #
...## => #
.#..# => .
#.... => .
#...# => .
.##.. => #
..### => .
##.## => .
..#.# => #
