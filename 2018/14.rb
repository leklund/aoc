class Elf
  attr_accessor :score, :pos

  def initialize(pos, score)
    @pos = pos
    @score = score
  end

  def move(arr)
    len = arr.size
    d = (score + 1 + pos) % len
    self.pos = d
    self.score = arr[d]
  end
end

def run(input, steps)
  elves = input.each_with_object([]).with_index do |(e, a), i|
    a << Elf.new(i, e)
  end

  while input.size < (steps + 10)
    step(input, elves[0], elves[1])
  end

  input
end

def run2(input, pattern)
  elves = input.each_with_object([]).with_index do |(e, a), i|
    a << Elf.new(i, e)
  end

  mm = []
  seen_starting_idx = nil

  while true
    adds = step(input, elves[0], elves[1])

    # check each added element to see if it the start of the pattern we are looking for.
    adds.each_with_index do |el, i|
      if el == pattern[mm.size]
        mm << el
        seen_starting_idx = input.size - 1 if seen_starting_idx.nil?
        return seen_starting_idx if pattern == mm
      elsif i == 1 && el == pattern[0]
        mm = [el]
        seen_starting_idx = input.size - 1
      else
        mm = []
        seen_starting_idx = nil
      end
    end
  end

  input.size - pattern.size
end

def step(arr, elf1, elf2)
  new_digits = (elf1.score + elf2.score).digits.reverse

  arr.concat(new_digits)

  elf1.move(arr)
  elf2.move(arr)
  new_digits
end

seed = [3, 7]

o = run(seed, 2018)

puts o[9, 10].join == '5158916779'
puts o[5, 10].join == '0124515891'
puts o[18, 10].join == '9251071085'
puts o[2018, 10].join == '5941429882'

o = run([3, 7], 704321)
puts o[704321, 10].join

puts "PART 2"

puts run2([3, 7], 51589.digits.reverse)
puts run2([3, 7], '01245'.chars.map(&:to_i))
puts run2([3, 7], 92510.digits.reverse)
puts run2([3, 7], 59414.digits.reverse)


puts run2([3,7], 704321.digits.reverse)
