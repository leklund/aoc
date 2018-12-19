require 'pry'
require 'scanf'
require 'prime'

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
    i = 1
    loop do
      instr = instructions[pointer]

      send(instr.first, instr, registers)

      increment

      break if out_of_range
      break if i > 10000000
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
puts computer.to_s
n = computer.registers[5]

puts n + n.prime_division.flatten.uniq.sum

# After much confusion I looked for hints and directed to prime divisors.
# This isn't so much a code problem as a puzzle, but okay.
# I break after 1 million cycles and calculate the sum of the prime divisors of the last register
# I could break after a much smaller amount but this allos me to see the answer in part
# one without summing the divisors.
# Needed a few more hints before I realized that for n, both 1 and n are factors.
# I suppose I should have reverse engineered the loop and figured out what it was doing but
# I really had no clue on this one.

computer = Computer.new(data)
computer.registers[0] = 1
computer.run
puts computer.to_s
n = computer.registers[5]
puts n + n.prime_division.flatten.uniq.sum


__END__
#ip 3
addi 3 16 3
seti 1 8 1
seti 1 3 4
mulr 1 4 2
eqrr 2 5 2
addr 2 3 3
addi 3 1 3
addr 1 0 0
addi 4 1 4
gtrr 4 5 2
addr 3 2 3
seti 2 6 3
addi 1 1 1
gtrr 1 5 2
addr 2 3 3
seti 1 5 3
mulr 3 3 3
addi 5 2 5
mulr 5 5 5
mulr 3 5 5
muli 5 11 5
addi 2 5 2
mulr 2 3 2
addi 2 21 2
addr 5 2 5
addr 3 0 3
seti 0 4 3
setr 3 1 2
mulr 2 3 2
addr 3 2 2
mulr 3 2 2
muli 2 14 2
mulr 2 3 2
addr 5 2 5
seti 0 3 0
seti 0 6 3
