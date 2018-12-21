require 'pry'
require 'scanf'

# I started with pen and paper and found the loop and relialized the exit condition
# would be when intruction 28 was called with r[4] == r[0]. So Part one is just to print
# the value of r[4] and break when instruction 28 is called.
# Part 2 was confusing because of the language. But if there is a max bound then the list of values
# that causes the program to halt must loop at some point. It was easy enough to reverse
# engineer the loop and find the last value before the r[4] value looped. It would
# have taken me a lot fewer guesses if I hadn't had silly bugs like checking every value against
# the list and not just halting values. That one caused me to check several different values.
# Finally I just started a brute force run and compared it against my generated list and once they
# diverged I facepalmed because I knew where the bug was immediately.

class Oper
  attr_reader :register, :instruction

  INSTRUCTIONS = %i(addr addi mulr muli banr bani borr bori setr seti gtir gtri gtrr eqir eqri eqrr)

  def initialize(instruction, register)
    @instruction = instruction # operation a b c
    @register = register # Array.new(6)
  end

  def self.method_missing(meth, *args)
    o = new(*args)
    o.send(meth)
    o.register
  end

  def code
    instruction[0]
  end

  def a
    instruction[1]
  end

  def b
    instruction[2]
  end

  def c
    instruction[3]
  end

  def addr
    register[c] = register[a] + register[b]
  end

  def addi
    register[c] = register[a] + b
  end

  def mulr
    register[c] = register[a] * register[b]
  end

  def muli
    register[c] = register[a] * b
  end

  def banr
    register[c] = register[a] & register[b]
  end

  def bani
    register[c] = register[a] & b
  end

  def borr
    register[c] = register[a] | register[b]
  end

  def bori
    register[c] = register[a] | b
  end

  def setr
    register[c] = register[a]
  end

  def seti
    register[c] = a
  end

  def gtir
    register[c] = a > register[b] ? 1 : 0
  end

  def gtri
    register[c] = register[a] > b ? 1 : 0
  end

  def gtrr
    register[c] = register[a] > register[b] ? 1 : 0
  end

  def eqir
    register[c] = a == register[b] ? 1 : 0
  end

  def eqri
    register[c] = register[a] == b ? 1 : 0
  end

  def eqrr
    register[c] = register[a] == register[b] ? 1 : 0
  end
end

class Computer
  # HEY KIDS I'M A COMPUTER
  attr_reader :instructions, :pointer_index
  attr_accessor :registers

  def initialize(instruction_set)
    @registers = Array.new(6, 0)
    @instructions = parse(instruction_set)
  end

  def parse(instruction_set)
    instructions = instruction_set.chomp.split("\n")
    @pointer_index = instructions.shift.scan(/\d/).first.to_i

    instructions.map do |i|
      a = i.scanf('%4c %d %d %d')
      a[0] = a[0].to_sym
      a
    end
  end

  def increment
    registers[pointer_index] += 1
  end

  def pointer
    registers[pointer_index]
  end

  def out_of_range
    pointer > instructions.size - 1 || pointer < 0
  end

  def run
    halt_codes = {}
    i = 1
      seen = []
    loop do
      instr = instructions[pointer]

      if pointer == 28
        # Part one solution
        puts "#{registers[4]}"
        break
      end
      send(instr.first, instr, registers)

      increment

      break if out_of_range
      i += 1
    end
  end

  def to_s
    #s = "Pointer index: #{pointer_index}\n"
    #s += "| 0 | 1 | 2 | 3 | 4 | 5 |\n"
    s = "| #{registers.join(" | ")} |"
  end

  def method_missing(meth, *args)
    Oper.send(meth, *args)
  end
end

# reverse engineered the loop
def find_loop
  puts '--------------'
  c0 = 16777215
  c1 = 65899
  c2 = 65536
  c3 = 255
  c4 = 6152285

  # intiial values
  r4 = 8004575
  r2 = c2
  nr2 = r2

  seen = []

  loop do
    pr4 = r4
    r2 = nr2

    # only want to record the values that will make the program halt
    halt = r2 / 256 == 0

    if halt
      puts "LOOP DETECTED with r4 == #{r4}" if seen.index(r4)
      break if seen.index(r4)
      seen << r4
    end

    n = halt ? c4 : r4
    nr2 = halt ? (pr4 | c2) : r2 / 256
    r4 = ((((nr2 & c3) + n) & c0) * c1) & c0
  end

  seen.last
end

def test
  input =<<~INPUT
#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5
  INPUT

  computer = Computer.new(input)
  computer.run
  puts computer.to_s
end

test
puts "------------"
data = DATA.read

computer = Computer.new(data)

computer.run

puts find_loop

__END__
#ip 3
seti 123 0 4
bani 4 456 4
eqri 4 72 4
addr 4 3 3
seti 0 0 3
seti 0 9 4
bori 4 65536 2
seti 6152285 4 4
bani 2 255 1
addr 4 1 4
bani 4 16777215 4
muli 4 65899 4
bani 4 16777215 4
gtir 256 2 1
addr 1 3 3
addi 3 1 3
seti 27 4 3
seti 0 3 1
addi 1 1 5
muli 5 256 5
gtrr 5 2 5
addr 5 3 3
addi 3 1 3
seti 25 9 3
addi 1 1 1
seti 17 4 3
setr 1 9 2
seti 7 4 3
eqrr 4 0 1
addr 1 3 3
seti 5 6 3
